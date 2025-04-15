package core

import (
	"context"
	"errors"
	"fmt"
	"go-eth-auction/logger"
	"go-eth-auction/shared"
	"math/big"

	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type RepoInterface interface {
	Insert(a *AuctionModel) error
	Delete(a *AuctionModel) error
	Get(name string) *AuctionModel
	AddProduct(auctionName string, p *ProductModel) error
	AuthorizeBidder(auctionName string, b *BidderModel) error
	Bid(auctionName string, p *Bid) error
	BlockchainView(auctionName string) ([]byte, error)
	Backup() ([]byte, error)
	Restore(b []byte) error
}

type Repository struct {
	lock       *sync.RWMutex
	storage    map[string]*AuctionModel
	blockchain *BlockChain
	workers    *shared.Worker

	count int
}

const (
	JobSetAuctionTimes string = "setAuctionTimings"
	JobBidAuthorize    string = "bidAuthorize"
	JobAddProduct      string = "addProduct"
	JobBid             string = "bid"
)

type Job struct {
	name       string
	params     map[string]any
	rollbackFn func()
}

func NewRepository(blockchain *BlockChain) *Repository {
	r := &Repository{
		lock:       &sync.RWMutex{},
		storage:    make(map[string]*AuctionModel),
		blockchain: blockchain,
		workers:    shared.NewWorkerManager(10, 1, nil),
		count:      0,
	}

	r.workers.SetWorker(func(workerIndex int, job interface{}) {
		v, _ := job.(*Job)
		if v.name == JobBidAuthorize {
			r.callAuthorize(v.params["auctionName"].(string),
				v.params["bidder"].(*BidderModel), v.rollbackFn)
		} else if v.name == JobAddProduct {
			r.callAddProduct(v.params["auctionName"].(string),
				v.params["product"].(*ProductModel), v.rollbackFn)
		} else if v.name == JobBid {
			r.callBid(v.params["auctionName"].(string),
				v.params["pCode"].(uint), v.params["bid"].(*big.Int), v.rollbackFn)
		} else if v.name == JobSetAuctionTimes {
			r.callSetTimings(v.params["auctionName"].(string),
				v.params["startDate"].(*big.Int), v.params["endDate"].(*big.Int), v.rollbackFn)
		}
	})

	go r.workers.Start()

	return r
}

func (ad *Repository) callSetTimings(auctionName string, startDate, endDate *big.Int, rollbackFn func()) error {

	chainid, err := ad.blockchain.Backend.ChainID(context.Background())
	if err != nil {
		return err
	}
	tx, err := ad.blockchain.AuctionContract(auctionName).SetAuctionTiming(&bind.TransactOpts{
		From: *ad.blockchain.owners[auctionName],
		Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(t, types.NewLondonSigner(chainid), ad.blockchain.keys[auctionName])
		},
	}, startDate, endDate)
	if err != nil {
		logger.Get().Append(auctionName, JobSetAuctionTimes,
			fmt.Sprintf("contract call failed: %s", err.Error()))

		if rollbackFn != nil {
			rollbackFn()
		}
		return err
	}

	logger.Get().Append(auctionName, JobSetAuctionTimes, "contract call successful on SetAuctionTimings")

	fmt.Println("SetAuctionTiming on smart contract, successfully: " + tx.ChainId().String())

	return nil
}

func (ad *Repository) Insert(a *AuctionModel) error {
	ad.lock.Lock()
	defer ad.lock.Unlock()

	// this piece of code is not efficient, we need
	// a separate endpoint for setting the times
	ad.workers.Enqueue(&Job{name: JobSetAuctionTimes,
		params: map[string]any{"auctionName": a.AuctionName, "startDate": big.NewInt(a.StartDate),
			"endDate": big.NewInt(a.EndDate)},

		// @todo we need to set a rollback fn to set back the previous
		// start and end dates for the auction in case of failure
		rollbackFn: nil,
	})
	if _, ok := ad.storage[a.AuctionName]; !ok {
		ad.storage[a.AuctionName] = a
		ad.count++
		logger.Get().Append(a.AuctionName, "auction", "created an auction")
		return nil
	}
	return errors.New("duplicate auction")
}

func (ad *Repository) Delete(a *AuctionModel) error {
	ad.lock.Lock()
	defer ad.lock.Unlock()
	if _, ok := ad.storage[a.AuctionName]; !ok {
		delete(ad.storage, a.AuctionName)
		logger.Get().Append(a.AuctionName, "auction", "deleted an auction")
		ad.count--
	}
	return errors.New("no auction found")
}

func (ad *Repository) Get(name string) *AuctionModel {
	ad.lock.RLock()
	defer ad.lock.RUnlock()
	if v, ok := ad.storage[name]; ok {
		return v.Copy()
	}
	return nil
}

func (ad *Repository) BlockchainView(auctionName string) ([]byte, error) {
	ad.lock.RLock()
	defer ad.lock.RUnlock()

	a, err := ad.blockchain.AuctionContract(auctionName).GetWinners(&bind.CallOpts{
		From: *ad.blockchain.owners[auctionName],
	})

	if err != nil {
		return nil, err
	}

	return []byte(a), err
}

func (ad *Repository) AddProduct(auctionName string, p *ProductModel) error {
	ad.lock.Lock()
	defer ad.lock.Unlock()
	if _, ok := ad.storage[auctionName]; ok {
		if ad.storage[auctionName].Products == nil {
			ad.storage[auctionName].Products = make(map[uint]*ProductModel)
		}
		if v, okk := ad.storage[auctionName].Products[p.Code]; okk {
			var oldStartingPrice = v.StartingPrice
			var code = p.Code
			ad.workers.Enqueue(&Job{name: JobAddProduct,
				params: map[string]any{"auctionName": auctionName, "product": p},
				rollbackFn: func() {
					ad.lock.Lock()
					defer ad.lock.Unlock()
					ad.storage[auctionName].Products[code].StartingPrice = oldStartingPrice
				},
			})
			v.StartingPrice = p.StartingPrice
		} else {
			ad.workers.Enqueue(&Job{name: JobAddProduct,
				params: map[string]any{"auctionName": auctionName, "product": p},
				rollbackFn: func() {
					ad.lock.Lock()
					defer ad.lock.Unlock()
					delete(ad.storage[auctionName].Products, p.Code)
				},
			})
			logger.Get().Append(auctionName, "product", "added a product")
			ad.storage[auctionName].Products[p.Code] = p.Copy()
		}
		return nil
	}
	return errors.New("auction not found")
}

func (ad *Repository) callAddProduct(auctionName string, p *ProductModel, rollbackFn func()) error {

	chainid, err := ad.blockchain.Backend.ChainID(context.Background())
	if err != nil {
		return err
	}
	tx, err := ad.blockchain.AuctionContract(auctionName).Product(&bind.TransactOpts{
		From: *ad.blockchain.owners[auctionName],
		Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(t, types.NewLondonSigner(chainid), ad.blockchain.keys[auctionName])
		},
	}, big.NewInt(int64(p.Code)), p.StartingPrice, false)
	if err != nil {
		logger.Get().Append(auctionName, JobAddProduct,
			fmt.Sprintf("contract call failed: %s", err.Error()))

		return err
	}

	logger.Get().Append(auctionName, JobAddProduct, "contract call successful")

	fmt.Println("bidder authorized on smart contract, successfully: " + tx.ChainId().String())

	return nil
}

func (ad *Repository) callAuthorize(auctionName string, b *BidderModel, rollbackFn func()) error {

	chainid, err := ad.blockchain.Backend.ChainID(context.Background())
	if err != nil {
		return err
	}
	tx, err := ad.blockchain.AuctionContract(auctionName).Authorize(&bind.TransactOpts{
		From: *ad.blockchain.owners[auctionName],
		Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(t, types.NewLondonSigner(chainid), ad.blockchain.keys[auctionName])
		},
	}, *b.Address)
	if err != nil {
		logger.Get().Append(auctionName, JobBidAuthorize,
			fmt.Sprintf("contract call failed: %s", err.Error()))

		if rollbackFn != nil {
			rollbackFn()
		}
		return err
	}

	logger.Get().Append(auctionName, JobBidAuthorize, "contract call successful")

	fmt.Println("bidder authorized on smart contract, successfully: " + tx.ChainId().String())

	return nil
}

func (ad *Repository) AuthorizeBidder(auctionName string, b *BidderModel) error {
	ad.lock.Lock()
	defer ad.lock.Unlock()
	if v, ok := ad.storage[auctionName]; ok {
		if _, ok := ad.blockchain.keys[auctionName]; !ok {
			return errors.New("contract keys not found")
		}
		if _, okk := v.Bidders[b.Address.Hex()]; !okk {
			var addr = b.Address.Hex()
			ad.workers.Enqueue(&Job{name: JobBidAuthorize,
				params: map[string]any{"auctionName": auctionName, "bidder": b},
				rollbackFn: func() {
					ad.lock.Lock()
					defer ad.lock.Unlock()
					delete(ad.storage[auctionName].Bidders, addr)
				},
			})
		}

		if ad.storage[auctionName].Bidders == nil {
			ad.storage[auctionName].Bidders = make(map[string]*BidderModel)
		}
		b.SetStatus(shared.StatusPending)
		ad.storage[auctionName].Bidders[b.Address.Hex()] = b.Copy()

		return nil
	}
	return errors.New("auction not found")
}

func (ad *Repository) callBid(auctionName string, pCode uint, bidAmount *big.Int, rollbackFn func()) error {

	chainid, err := ad.blockchain.Backend.ChainID(context.Background())
	if err != nil {
		return err
	}
	tx, err := ad.blockchain.AuctionContract(auctionName).Bid(&bind.TransactOpts{
		From: *ad.blockchain.owners[auctionName],
		Signer: func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(t, types.NewLondonSigner(chainid), ad.blockchain.keys[auctionName])
		},
		Value: bidAmount,
	}, big.NewInt(int64(pCode)))
	if err != nil {
		logger.Get().Append(auctionName, JobBidAuthorize,
			fmt.Sprintf("contract call failed: %s", err.Error()))

		if rollbackFn != nil {
			rollbackFn()
		}
		return err
	}

	logger.Get().Append(auctionName, JobBid, "contract call successful")

	fmt.Println("bid placed on smart contract, successfully: " + tx.ChainId().String())

	return nil
}

func (ad *Repository) Bid(auctionName string, bid *Bid) error {
	ad.lock.RLock()
	defer ad.lock.RUnlock()
	if v, ok := ad.storage[auctionName]; ok {
		if vv, okk := v.Products[bid.ProductCode]; okk {
			prevBid, prevHighest, err := vv.InsertBids(bid, false)
			if err != nil {
				return err
			}

			ad.workers.Enqueue(&Job{name: JobBid,
				params: map[string]any{"auctionName": auctionName, "pCode": bid.ProductCode,
					"bid": bid.Amount},
				rollbackFn: func() {
					ad.lock.Lock()
					defer ad.lock.Unlock()
					if prevBid != nil {
						ad.storage[auctionName].Products[bid.ProductCode].
							Bids[bid.Address.Hex()].Amount = prevBid
					}
					if prevHighest != nil {
						ad.storage[auctionName].Products[bid.ProductCode].
							HighestBid.Amount = prevBid
					}
				},
			})

			return nil
		}
	}
	return errors.New("auction not found")
}

func (ad *Repository) Backup() ([]byte, error) {
	return nil, nil
}

func (ad *Repository) Restore(b []byte) error {
	return nil
}
