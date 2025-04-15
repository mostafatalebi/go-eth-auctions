package core

import (
	"crypto/ecdsa"
	"go-eth-auction/contract"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AbstractContract interface{}

type BlockChain struct {
	Backend *ethclient.Client

	// map[auctionName]auctionAbi
	Contracts map[string]AbstractContract
	keys      map[string]*ecdsa.PrivateKey
	owners    map[string]*common.Address
}

func NewBlockChain(b *ethclient.Client) (*BlockChain, error) {
	bc := &BlockChain{
		Backend: b,
	}
	return bc, nil
}

func (bc *BlockChain) registerContract(contractName string, contract AbstractContract, pk *ecdsa.PrivateKey, owner common.Address) error {
	if bc.Contracts == nil {
		bc.Contracts = make(map[string]AbstractContract)
	}
	if bc.keys == nil {
		bc.keys = make(map[string]*ecdsa.PrivateKey)
	}
	if bc.owners == nil {
		bc.owners = make(map[string]*common.Address)
	}
	bc.Contracts[contractName] = contract
	bc.keys[contractName] = pk
	bc.owners[contractName] = &owner

	return nil
}

// helper
// for accessing auction contract
func (bc *BlockChain) AuctionContract(name string) *contract.Auction {
	var v, ok = bc.Contracts[name].(*contract.Auction)
	if ok {
		return v
	}
	return nil
}
