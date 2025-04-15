package core

import (
	"go-eth-auction/config"
	"go-eth-auction/contract"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Container struct {
	Repo RepoInterface
}

func NewContainer() (*Container, error) {
	a := &Container{}
	var backend, err = ethclient.Dial(config.ConfigObj.EthereumNodeAddress)
	if err != nil {
		return nil, err
	}

	blockchain, err := NewBlockChain(backend)
	if err != nil {
		return nil, err
	}

	if config.ConfigObj.Contracts != nil {
		for _, v := range config.ConfigObj.Contracts {
			if v.ContractType == config.ContractAuction {
				contractObj, err := contract.NewAuction(v.Address, backend)
				if err != nil {
					return nil, err
				}
				blockchain.registerContract(v.Name, contractObj, v.Key, v.Owner)
			}
		}
	}

	a.Repo = NewRepository(blockchain)
	return a, nil
}
