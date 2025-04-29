package core

import (
	"errors"
	"go-eth-auction/shared"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type AuctionModel struct {
	AuctionName string                  `json:"name"`
	StartDate   int64                   `json:"startDate"`
	EndDate     int64                   `json:"endDate"`
	Products    map[uint]*ProductModel  `json:"products"`
	Bidders     map[string]*BidderModel `json:"bidders"`
}

func (am *AuctionModel) Copy() *AuctionModel {
	copied := &AuctionModel{
		AuctionName: am.AuctionName,
		StartDate:   am.StartDate,
		EndDate:     am.EndDate,
		Bidders:     make(map[string]*BidderModel),
		Products:    make(map[uint]*ProductModel),
	}

	for k, v := range am.Products {
		copied.Products[k] = v.Copy()
	}

	for k, v := range am.Bidders {
		copied.Bidders[k] = v.Copy()
	}

	return copied
}

type Bid struct {
	AuctionName string          `json:"auctionName"`
	Address     *common.Address `json:"address"`
	Amount      *big.Int        `json:"amount"`
	ProductCode uint            `json:"productCode"`
}

func (b *Bid) Copy() *Bid {
	return &Bid{
		AuctionName: b.AuctionName,
		Address:     b.Address,
		Amount:      b.Amount,
		ProductCode: b.ProductCode,
	}
}

type ProductModel struct {
	// We need to put auction name here too, for a more advanced system
	// we would have done something different
	AuctionName string `json:"auctionName"`

	Code          uint     `json:"code"`
	StartingPrice *big.Int `json:"startingPrice"`
	Bids          map[string]*Bid
	HighestBid    *Bid
}

// forceOverride by default, this function returns error
// if the new bid's amount is less than previous amount for
// the same bidder address. Setting this flag to true, ignores
// this condition
// it returns three values: prevBid and prevHighest are used
// by the caller of this function for possible rollbacking (in case
// contract call fails)
func (p *ProductModel) InsertBids(b *Bid, forceOverride bool) (previousBid *big.Int,
	previousHighest *Bid, err error) {
	if b != nil {
		if p.Bids == nil {
			p.Bids = make(map[string]*Bid, 0)
		}
		if v, ok := p.Bids[b.Address.Hex()]; ok && !forceOverride {
			if b.Amount.Cmp(v.Amount) == -1 {
				err = errors.New("bid amount lower than previous bid, skipped")
			}
			previousBid = v.Amount
		}
		p.Bids[b.Address.Hex()] = b
		if p.HighestBid == nil {
			previousHighest = nil
			p.HighestBid = b.Copy()
		} else if b.Amount.Cmp(p.HighestBid.Amount) == 1 {
			previousHighest = p.HighestBid.Copy()
			p.HighestBid = b.Copy()
		}
		return
	}
	err = errors.New("bid *Bid cannot be nil")
	return
}

func (ap *ProductModel) Copy() *ProductModel {
	pm := &ProductModel{
		AuctionName:   ap.AuctionName,
		Code:          ap.Code,
		StartingPrice: ap.StartingPrice,
		Bids:          make(map[string]*Bid),
	}

	if len(ap.Bids) > 0 {
		for k, v := range ap.Bids {
			pm.Bids[k] = v.Copy()
		}
	}

	return pm
}

type AddressRequestModel struct {
	Address *common.Address `json:"address"`
}

type WithdrawRequestModel struct {
	Address     *common.Address `json:"address"`
	AcutionName string          `json:"auctionName`
}

type BidderModel struct {
	AuctionName string          `json:"auctionName"`
	Name        string          `json:"name"`
	Address     *common.Address `json:"address"`
	Status      shared.Status   `json:"status"`
}

func NewBidderModel() *BidderModel {
	return &BidderModel{}
}

func (bm *BidderModel) SetStatus(status string) {
	bm.Status = status
}

func (bm *BidderModel) Copy() *BidderModel {
	copied := &BidderModel{
		AuctionName: bm.AuctionName,
		Name:        bm.Name,
		Address:     bm.Address,
		Status:      shared.StatusPending,
	}

	return copied
}
