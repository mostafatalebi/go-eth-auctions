// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AuctionMetaData contains all meta data concerning the Auction contract.
var AuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumActivationType\",\"name\":\"_startType\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrAuctionCannotBeStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAuctionClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAuctionIsManual\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAuctionIsTemporal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAuctionNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAuctionNotStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAuctionNotYetClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAuctionStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrBadProductCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrBidCannotBeLowerThanPrevious\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrBidTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrBidderNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrDuplicateBidder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrDurTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrForbidden\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrOutOfBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrProductNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrTooManyProducts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrWithdrawFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"productCode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedBuyers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toBeBuyer\",\"type\":\"address\"}],\"name\":\"authorize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"productCode\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"productCode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"bidAs\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currentBids\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"productCode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"put\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"productCode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"}],\"name\":\"getCurrentBids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"productCode\",\"type\":\"uint256\"}],\"name\":\"getHighestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMyBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWinners\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liveProductsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"productCode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startingBidPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRemove\",\"type\":\"bool\"}],\"name\":\"product\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"productsKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"productsMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"setAuctionTiming\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toBeBuyer\",\"type\":\"address\"}],\"name\":\"unauthorize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"withdrawAs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionMetaData.ABI instead.
var AuctionABI = AuctionMetaData.ABI

// Auction is an auto generated Go binding around an Ethereum contract.
type Auction struct {
	AuctionCaller     // Read-only binding to the contract
	AuctionTransactor // Write-only binding to the contract
	AuctionFilterer   // Log filterer for contract events
}

// AuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionSession struct {
	Contract     *Auction          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionCallerSession struct {
	Contract *AuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionTransactorSession struct {
	Contract     *AuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionRaw struct {
	Contract *Auction // Generic contract binding to access the raw methods on
}

// AuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionCallerRaw struct {
	Contract *AuctionCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionTransactorRaw struct {
	Contract *AuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuction creates a new instance of Auction, bound to a specific deployed contract.
func NewAuction(address common.Address, backend bind.ContractBackend) (*Auction, error) {
	contract, err := bindAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// NewAuctionCaller creates a new read-only instance of Auction, bound to a specific deployed contract.
func NewAuctionCaller(address common.Address, caller bind.ContractCaller) (*AuctionCaller, error) {
	contract, err := bindAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionCaller{contract: contract}, nil
}

// NewAuctionTransactor creates a new write-only instance of Auction, bound to a specific deployed contract.
func NewAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionTransactor, error) {
	contract, err := bindAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionTransactor{contract: contract}, nil
}

// NewAuctionFilterer creates a new log filterer instance of Auction, bound to a specific deployed contract.
func NewAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionFilterer, error) {
	contract, err := bindAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionFilterer{contract: contract}, nil
}

// bindAuction binds a generic wrapper to an already deployed contract.
func bindAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.AuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transact(opts, method, params...)
}

// AllowedBuyers is a free data retrieval call binding the contract method 0xb9b48595.
//
// Solidity: function allowedBuyers(address ) view returns(bool)
func (_Auction *AuctionCaller) AllowedBuyers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "allowedBuyers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedBuyers is a free data retrieval call binding the contract method 0xb9b48595.
//
// Solidity: function allowedBuyers(address ) view returns(bool)
func (_Auction *AuctionSession) AllowedBuyers(arg0 common.Address) (bool, error) {
	return _Auction.Contract.AllowedBuyers(&_Auction.CallOpts, arg0)
}

// AllowedBuyers is a free data retrieval call binding the contract method 0xb9b48595.
//
// Solidity: function allowedBuyers(address ) view returns(bool)
func (_Auction *AuctionCallerSession) AllowedBuyers(arg0 common.Address) (bool, error) {
	return _Auction.Contract.AllowedBuyers(&_Auction.CallOpts, arg0)
}

// CurrentBids is a free data retrieval call binding the contract method 0x69b2643d.
//
// Solidity: function currentBids(uint256 , address ) view returns(address buyer, uint256 productCode, uint256 amount, bool put)
func (_Auction *AuctionCaller) CurrentBids(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Buyer       common.Address
	ProductCode *big.Int
	Amount      *big.Int
	Put         bool
}, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "currentBids", arg0, arg1)

	outstruct := new(struct {
		Buyer       common.Address
		ProductCode *big.Int
		Amount      *big.Int
		Put         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Buyer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ProductCode = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Put = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// CurrentBids is a free data retrieval call binding the contract method 0x69b2643d.
//
// Solidity: function currentBids(uint256 , address ) view returns(address buyer, uint256 productCode, uint256 amount, bool put)
func (_Auction *AuctionSession) CurrentBids(arg0 *big.Int, arg1 common.Address) (struct {
	Buyer       common.Address
	ProductCode *big.Int
	Amount      *big.Int
	Put         bool
}, error) {
	return _Auction.Contract.CurrentBids(&_Auction.CallOpts, arg0, arg1)
}

// CurrentBids is a free data retrieval call binding the contract method 0x69b2643d.
//
// Solidity: function currentBids(uint256 , address ) view returns(address buyer, uint256 productCode, uint256 amount, bool put)
func (_Auction *AuctionCallerSession) CurrentBids(arg0 *big.Int, arg1 common.Address) (struct {
	Buyer       common.Address
	ProductCode *big.Int
	Amount      *big.Int
	Put         bool
}, error) {
	return _Auction.Contract.CurrentBids(&_Auction.CallOpts, arg0, arg1)
}

// GetCurrentBids is a free data retrieval call binding the contract method 0x3c7a1e7d.
//
// Solidity: function getCurrentBids(uint256 productCode, address bidderAddress) view returns(uint256)
func (_Auction *AuctionCaller) GetCurrentBids(opts *bind.CallOpts, productCode *big.Int, bidderAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "getCurrentBids", productCode, bidderAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBids is a free data retrieval call binding the contract method 0x3c7a1e7d.
//
// Solidity: function getCurrentBids(uint256 productCode, address bidderAddress) view returns(uint256)
func (_Auction *AuctionSession) GetCurrentBids(productCode *big.Int, bidderAddress common.Address) (*big.Int, error) {
	return _Auction.Contract.GetCurrentBids(&_Auction.CallOpts, productCode, bidderAddress)
}

// GetCurrentBids is a free data retrieval call binding the contract method 0x3c7a1e7d.
//
// Solidity: function getCurrentBids(uint256 productCode, address bidderAddress) view returns(uint256)
func (_Auction *AuctionCallerSession) GetCurrentBids(productCode *big.Int, bidderAddress common.Address) (*big.Int, error) {
	return _Auction.Contract.GetCurrentBids(&_Auction.CallOpts, productCode, bidderAddress)
}

// GetHighestBid is a free data retrieval call binding the contract method 0x8f288644.
//
// Solidity: function getHighestBid(uint256 productCode) view returns(uint256)
func (_Auction *AuctionCaller) GetHighestBid(opts *bind.CallOpts, productCode *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "getHighestBid", productCode)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHighestBid is a free data retrieval call binding the contract method 0x8f288644.
//
// Solidity: function getHighestBid(uint256 productCode) view returns(uint256)
func (_Auction *AuctionSession) GetHighestBid(productCode *big.Int) (*big.Int, error) {
	return _Auction.Contract.GetHighestBid(&_Auction.CallOpts, productCode)
}

// GetHighestBid is a free data retrieval call binding the contract method 0x8f288644.
//
// Solidity: function getHighestBid(uint256 productCode) view returns(uint256)
func (_Auction *AuctionCallerSession) GetHighestBid(productCode *big.Int) (*big.Int, error) {
	return _Auction.Contract.GetHighestBid(&_Auction.CallOpts, productCode)
}

// GetMyBalance is a free data retrieval call binding the contract method 0x4c738909.
//
// Solidity: function getMyBalance() view returns(uint256)
func (_Auction *AuctionCaller) GetMyBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "getMyBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMyBalance is a free data retrieval call binding the contract method 0x4c738909.
//
// Solidity: function getMyBalance() view returns(uint256)
func (_Auction *AuctionSession) GetMyBalance() (*big.Int, error) {
	return _Auction.Contract.GetMyBalance(&_Auction.CallOpts)
}

// GetMyBalance is a free data retrieval call binding the contract method 0x4c738909.
//
// Solidity: function getMyBalance() view returns(uint256)
func (_Auction *AuctionCallerSession) GetMyBalance() (*big.Int, error) {
	return _Auction.Contract.GetMyBalance(&_Auction.CallOpts)
}

// GetWinners is a free data retrieval call binding the contract method 0xdf15c37e.
//
// Solidity: function getWinners() view returns(string)
func (_Auction *AuctionCaller) GetWinners(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "getWinners")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetWinners is a free data retrieval call binding the contract method 0xdf15c37e.
//
// Solidity: function getWinners() view returns(string)
func (_Auction *AuctionSession) GetWinners() (string, error) {
	return _Auction.Contract.GetWinners(&_Auction.CallOpts)
}

// GetWinners is a free data retrieval call binding the contract method 0xdf15c37e.
//
// Solidity: function getWinners() view returns(string)
func (_Auction *AuctionCallerSession) GetWinners() (string, error) {
	return _Auction.Contract.GetWinners(&_Auction.CallOpts)
}

// LiveProductsCount is a free data retrieval call binding the contract method 0x825ae97e.
//
// Solidity: function liveProductsCount() view returns(uint256)
func (_Auction *AuctionCaller) LiveProductsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "liveProductsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiveProductsCount is a free data retrieval call binding the contract method 0x825ae97e.
//
// Solidity: function liveProductsCount() view returns(uint256)
func (_Auction *AuctionSession) LiveProductsCount() (*big.Int, error) {
	return _Auction.Contract.LiveProductsCount(&_Auction.CallOpts)
}

// LiveProductsCount is a free data retrieval call binding the contract method 0x825ae97e.
//
// Solidity: function liveProductsCount() view returns(uint256)
func (_Auction *AuctionCallerSession) LiveProductsCount() (*big.Int, error) {
	return _Auction.Contract.LiveProductsCount(&_Auction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionSession) Owner() (common.Address, error) {
	return _Auction.Contract.Owner(&_Auction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionCallerSession) Owner() (common.Address, error) {
	return _Auction.Contract.Owner(&_Auction.CallOpts)
}

// ProductsKeys is a free data retrieval call binding the contract method 0x87e06e5e.
//
// Solidity: function productsKeys(uint256 ) view returns(uint256)
func (_Auction *AuctionCaller) ProductsKeys(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "productsKeys", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProductsKeys is a free data retrieval call binding the contract method 0x87e06e5e.
//
// Solidity: function productsKeys(uint256 ) view returns(uint256)
func (_Auction *AuctionSession) ProductsKeys(arg0 *big.Int) (*big.Int, error) {
	return _Auction.Contract.ProductsKeys(&_Auction.CallOpts, arg0)
}

// ProductsKeys is a free data retrieval call binding the contract method 0x87e06e5e.
//
// Solidity: function productsKeys(uint256 ) view returns(uint256)
func (_Auction *AuctionCallerSession) ProductsKeys(arg0 *big.Int) (*big.Int, error) {
	return _Auction.Contract.ProductsKeys(&_Auction.CallOpts, arg0)
}

// ProductsMap is a free data retrieval call binding the contract method 0x77dec24c.
//
// Solidity: function productsMap(uint256 ) view returns(uint256 code, uint256 startingPrice, bool exists)
func (_Auction *AuctionCaller) ProductsMap(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Code          *big.Int
	StartingPrice *big.Int
	Exists        bool
}, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "productsMap", arg0)

	outstruct := new(struct {
		Code          *big.Int
		StartingPrice *big.Int
		Exists        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Code = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartingPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// ProductsMap is a free data retrieval call binding the contract method 0x77dec24c.
//
// Solidity: function productsMap(uint256 ) view returns(uint256 code, uint256 startingPrice, bool exists)
func (_Auction *AuctionSession) ProductsMap(arg0 *big.Int) (struct {
	Code          *big.Int
	StartingPrice *big.Int
	Exists        bool
}, error) {
	return _Auction.Contract.ProductsMap(&_Auction.CallOpts, arg0)
}

// ProductsMap is a free data retrieval call binding the contract method 0x77dec24c.
//
// Solidity: function productsMap(uint256 ) view returns(uint256 code, uint256 startingPrice, bool exists)
func (_Auction *AuctionCallerSession) ProductsMap(arg0 *big.Int) (struct {
	Code          *big.Int
	StartingPrice *big.Int
	Exists        bool
}, error) {
	return _Auction.Contract.ProductsMap(&_Auction.CallOpts, arg0)
}

// Authorize is a paid mutator transaction binding the contract method 0xb6a5d7de.
//
// Solidity: function authorize(address toBeBuyer) returns()
func (_Auction *AuctionTransactor) Authorize(opts *bind.TransactOpts, toBeBuyer common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "authorize", toBeBuyer)
}

// Authorize is a paid mutator transaction binding the contract method 0xb6a5d7de.
//
// Solidity: function authorize(address toBeBuyer) returns()
func (_Auction *AuctionSession) Authorize(toBeBuyer common.Address) (*types.Transaction, error) {
	return _Auction.Contract.Authorize(&_Auction.TransactOpts, toBeBuyer)
}

// Authorize is a paid mutator transaction binding the contract method 0xb6a5d7de.
//
// Solidity: function authorize(address toBeBuyer) returns()
func (_Auction *AuctionTransactorSession) Authorize(toBeBuyer common.Address) (*types.Transaction, error) {
	return _Auction.Contract.Authorize(&_Auction.TransactOpts, toBeBuyer)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 productCode) payable returns()
func (_Auction *AuctionTransactor) Bid(opts *bind.TransactOpts, productCode *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "bid", productCode)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 productCode) payable returns()
func (_Auction *AuctionSession) Bid(productCode *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts, productCode)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 productCode) payable returns()
func (_Auction *AuctionTransactorSession) Bid(productCode *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts, productCode)
}

// BidAs is a paid mutator transaction binding the contract method 0x0d1f55e3.
//
// Solidity: function bidAs(uint256 productCode, address bidder) payable returns()
func (_Auction *AuctionTransactor) BidAs(opts *bind.TransactOpts, productCode *big.Int, bidder common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "bidAs", productCode, bidder)
}

// BidAs is a paid mutator transaction binding the contract method 0x0d1f55e3.
//
// Solidity: function bidAs(uint256 productCode, address bidder) payable returns()
func (_Auction *AuctionSession) BidAs(productCode *big.Int, bidder common.Address) (*types.Transaction, error) {
	return _Auction.Contract.BidAs(&_Auction.TransactOpts, productCode, bidder)
}

// BidAs is a paid mutator transaction binding the contract method 0x0d1f55e3.
//
// Solidity: function bidAs(uint256 productCode, address bidder) payable returns()
func (_Auction *AuctionTransactorSession) BidAs(productCode *big.Int, bidder common.Address) (*types.Transaction, error) {
	return _Auction.Contract.BidAs(&_Auction.TransactOpts, productCode, bidder)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x378252f2.
//
// Solidity: function closeAuction() returns()
func (_Auction *AuctionTransactor) CloseAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "closeAuction")
}

// CloseAuction is a paid mutator transaction binding the contract method 0x378252f2.
//
// Solidity: function closeAuction() returns()
func (_Auction *AuctionSession) CloseAuction() (*types.Transaction, error) {
	return _Auction.Contract.CloseAuction(&_Auction.TransactOpts)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x378252f2.
//
// Solidity: function closeAuction() returns()
func (_Auction *AuctionTransactorSession) CloseAuction() (*types.Transaction, error) {
	return _Auction.Contract.CloseAuction(&_Auction.TransactOpts)
}

// Product is a paid mutator transaction binding the contract method 0x0471b94b.
//
// Solidity: function product(uint256 productCode, uint256 startingBidPrice, bool isRemove) returns()
func (_Auction *AuctionTransactor) Product(opts *bind.TransactOpts, productCode *big.Int, startingBidPrice *big.Int, isRemove bool) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "product", productCode, startingBidPrice, isRemove)
}

// Product is a paid mutator transaction binding the contract method 0x0471b94b.
//
// Solidity: function product(uint256 productCode, uint256 startingBidPrice, bool isRemove) returns()
func (_Auction *AuctionSession) Product(productCode *big.Int, startingBidPrice *big.Int, isRemove bool) (*types.Transaction, error) {
	return _Auction.Contract.Product(&_Auction.TransactOpts, productCode, startingBidPrice, isRemove)
}

// Product is a paid mutator transaction binding the contract method 0x0471b94b.
//
// Solidity: function product(uint256 productCode, uint256 startingBidPrice, bool isRemove) returns()
func (_Auction *AuctionTransactorSession) Product(productCode *big.Int, startingBidPrice *big.Int, isRemove bool) (*types.Transaction, error) {
	return _Auction.Contract.Product(&_Auction.TransactOpts, productCode, startingBidPrice, isRemove)
}

// SetAuctionTiming is a paid mutator transaction binding the contract method 0x5b0ba29b.
//
// Solidity: function setAuctionTiming(uint256 start, uint256 end) returns()
func (_Auction *AuctionTransactor) SetAuctionTiming(opts *bind.TransactOpts, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "setAuctionTiming", start, end)
}

// SetAuctionTiming is a paid mutator transaction binding the contract method 0x5b0ba29b.
//
// Solidity: function setAuctionTiming(uint256 start, uint256 end) returns()
func (_Auction *AuctionSession) SetAuctionTiming(start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.SetAuctionTiming(&_Auction.TransactOpts, start, end)
}

// SetAuctionTiming is a paid mutator transaction binding the contract method 0x5b0ba29b.
//
// Solidity: function setAuctionTiming(uint256 start, uint256 end) returns()
func (_Auction *AuctionTransactorSession) SetAuctionTiming(start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.SetAuctionTiming(&_Auction.TransactOpts, start, end)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_Auction *AuctionTransactor) StartAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "startAuction")
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_Auction *AuctionSession) StartAuction() (*types.Transaction, error) {
	return _Auction.Contract.StartAuction(&_Auction.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_Auction *AuctionTransactorSession) StartAuction() (*types.Transaction, error) {
	return _Auction.Contract.StartAuction(&_Auction.TransactOpts)
}

// Unauthorize is a paid mutator transaction binding the contract method 0xf0b37c04.
//
// Solidity: function unauthorize(address toBeBuyer) returns()
func (_Auction *AuctionTransactor) Unauthorize(opts *bind.TransactOpts, toBeBuyer common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "unauthorize", toBeBuyer)
}

// Unauthorize is a paid mutator transaction binding the contract method 0xf0b37c04.
//
// Solidity: function unauthorize(address toBeBuyer) returns()
func (_Auction *AuctionSession) Unauthorize(toBeBuyer common.Address) (*types.Transaction, error) {
	return _Auction.Contract.Unauthorize(&_Auction.TransactOpts, toBeBuyer)
}

// Unauthorize is a paid mutator transaction binding the contract method 0xf0b37c04.
//
// Solidity: function unauthorize(address toBeBuyer) returns()
func (_Auction *AuctionTransactorSession) Unauthorize(toBeBuyer common.Address) (*types.Transaction, error) {
	return _Auction.Contract.Unauthorize(&_Auction.TransactOpts, toBeBuyer)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionSession) Withdraw() (*types.Transaction, error) {
	return _Auction.Contract.Withdraw(&_Auction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Auction.Contract.Withdraw(&_Auction.TransactOpts)
}

// WithdrawAs is a paid mutator transaction binding the contract method 0x954d4fbd.
//
// Solidity: function withdrawAs(address bidder) returns()
func (_Auction *AuctionTransactor) WithdrawAs(opts *bind.TransactOpts, bidder common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "withdrawAs", bidder)
}

// WithdrawAs is a paid mutator transaction binding the contract method 0x954d4fbd.
//
// Solidity: function withdrawAs(address bidder) returns()
func (_Auction *AuctionSession) WithdrawAs(bidder common.Address) (*types.Transaction, error) {
	return _Auction.Contract.WithdrawAs(&_Auction.TransactOpts, bidder)
}

// WithdrawAs is a paid mutator transaction binding the contract method 0x954d4fbd.
//
// Solidity: function withdrawAs(address bidder) returns()
func (_Auction *AuctionTransactorSession) WithdrawAs(bidder common.Address) (*types.Transaction, error) {
	return _Auction.Contract.WithdrawAs(&_Auction.TransactOpts, bidder)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Auction *AuctionTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Auction *AuctionSession) Receive() (*types.Transaction, error) {
	return _Auction.Contract.Receive(&_Auction.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Auction *AuctionTransactorSession) Receive() (*types.Transaction, error) {
	return _Auction.Contract.Receive(&_Auction.TransactOpts)
}

// AuctionBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the Auction contract.
type AuctionBidPlacedIterator struct {
	Event *AuctionBidPlaced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AuctionBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionBidPlaced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AuctionBidPlaced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AuctionBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionBidPlaced represents a BidPlaced event raised by the Auction contract.
type AuctionBidPlaced struct {
	ProductCode *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x0fc69407039e8ba8cdd517e2a4cae9dd19697b0a500ff37a2425c59ebc649b31.
//
// Solidity: event BidPlaced(uint256 productCode, uint256 amount)
func (_Auction *AuctionFilterer) FilterBidPlaced(opts *bind.FilterOpts) (*AuctionBidPlacedIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "BidPlaced")
	if err != nil {
		return nil, err
	}
	return &AuctionBidPlacedIterator{contract: _Auction.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x0fc69407039e8ba8cdd517e2a4cae9dd19697b0a500ff37a2425c59ebc649b31.
//
// Solidity: event BidPlaced(uint256 productCode, uint256 amount)
func (_Auction *AuctionFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *AuctionBidPlaced) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "BidPlaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionBidPlaced)
				if err := _Auction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBidPlaced is a log parse operation binding the contract event 0x0fc69407039e8ba8cdd517e2a4cae9dd19697b0a500ff37a2425c59ebc649b31.
//
// Solidity: event BidPlaced(uint256 productCode, uint256 amount)
func (_Auction *AuctionFilterer) ParseBidPlaced(log types.Log) (*AuctionBidPlaced, error) {
	event := new(AuctionBidPlaced)
	if err := _Auction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
