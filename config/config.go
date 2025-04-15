package config

import (
	"crypto/ecdsa"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ContractAuction = "auction"
)

type ContractConfigEntry struct {
	Address      common.Address
	Owner        common.Address
	Name         string
	Key          *ecdsa.PrivateKey
	ContractType string
}

type config struct {
	AuctionMinimumDuration      time.Duration
	MaxAllowedNumberOfProducts  int
	EthereumNodeAddress         string
	AuctionSmartContractAddress common.Address
	AuctionSmartContractSk      string

	Contracts []*ContractConfigEntry
}

var ConfigObj *config

// for a production level code-base, secrets must be managed
// from a credentials management system (such as Vault)
func init() {
	var contracts = make([]*ContractConfigEntry, 0)

	// auction contract
	auctionPk, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		panic(err.Error())
	}

	// for a production level project, all this should be configured to be
	// generated from a panel or API or whatever
	contracts = append(contracts, &ContractConfigEntry{
		Address:      common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"),
		Owner:        common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		Name:         "camera-auction-texas",
		ContractType: ContractAuction,
		Key:          auctionPk,
	})

	ConfigObj = &config{
		AuctionMinimumDuration:     time.Minute * 30,
		MaxAllowedNumberOfProducts: 5,
		EthereumNodeAddress:        "http://127.0.0.1:8545",
		Contracts:                  contracts,
	}
}
