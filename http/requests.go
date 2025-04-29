package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-eth-auction/config"
	"go-eth-auction/core"
	"go-eth-auction/utils"
	"math/big"
	"time"
)

var minimumAllowedStartingPrice = big.NewInt(1)

type ErrorResp struct {
	Error string `json:"error"`
}

func ValidateIncomingAuctionCreateReq(b []byte) (*core.AuctionModel, error) {
	// this is for demonstration purposes, for high speed
	// servers, the JSON method must be different as golang's
	// default json pkg is not efficient
	var auctionCreate *core.AuctionModel
	var err = json.Unmarshal(b, &auctionCreate)

	if err != nil {
		return nil, err
	}
	var now = time.Now()
	err = nil
	if auctionCreate.AuctionName == "" {
		err = errors.New("validation error: auction's name is required")
	}

	if auctionCreate.StartDate < now.Unix() {
		err = errors.New("validation error: startDate cannot be older than now")
	}

	if auctionCreate.EndDate < now.Unix() {
		err = errors.New("validation error: endDate cannot be older than now")
	}

	if auctionCreate.EndDate-auctionCreate.StartDate < int64(config.ConfigObj.AuctionMinimumDuration.Seconds()) {
		err = fmt.Errorf("validation error: auction duration cannot be shorter than %v minutes", config.ConfigObj.AuctionMinimumDuration.Minutes())
	}

	if err != nil {
		return nil, err
	}

	auctionCreate.Products = nil
	auctionCreate.Bidders = nil
	return auctionCreate, nil
}

func ValidateIncomingAuctionProductAddReq(b []byte) (*core.ProductModel, error) {
	// this is for demonstration purposes, for high speed
	// servers, the JSON method must be different as golang's
	// default json pkg is not efficient
	var product *core.ProductModel
	var err = json.Unmarshal(b, &product)

	if err != nil {
		return nil, err
	}
	err = nil

	if product.Code < 1 {
		err = errors.New("validation error: product code must be a positive integer")
	}

	if product.StartingPrice.Cmp(minimumAllowedStartingPrice) == -1 {
		err = errors.New("validation error: starting price should be a positive integer indiciating amount of gwei")
	}

	if err != nil {
		return nil, err
	}

	return product, nil
}

func ValidateIncomingAuctionAuthorizeBidderReq(b []byte) (*core.BidderModel, error) {
	var model *core.BidderModel
	var err = json.Unmarshal(b, &model)

	if err != nil {
		return nil, err
	}
	err = nil

	if model.Address == nil || utils.ValidateAddress(*model.Address) == false {
		err = errors.New("validation error: address is not properly formatted")
	}

	if model.Name == "" {
		err = errors.New("validation error: bidder's name is required")
	}

	if err != nil {
		return nil, err
	}

	return model, nil
}

func ValidateIncomingAddressReq(b []byte) (*core.AddressRequestModel, error) {
	var model *core.AddressRequestModel
	var err = json.Unmarshal(b, &model)

	if err != nil {
		return nil, err
	}
	err = nil

	if model.Address == nil || utils.ValidateAddress(*model.Address) == false {
		err = errors.New("validation error: address is not properly formatted")
	}

	if err != nil {
		return nil, err
	}

	return model, nil
}

func RespondWithError(msg string) []byte {
	var e = &ErrorResp{
		Error: msg,
	}

	var d, _ = json.Marshal(e)
	return d
}

func RespondWithData(data any) ([]byte, error) {
	return json.Marshal(data)
}

func ValidateIncomingBidReq(b []byte) (*core.Bid, error) {
	var model *core.Bid
	var err = json.Unmarshal(b, &model)

	if err != nil {
		return nil, err
	}
	err = nil

	if model.AuctionName == "" {
		err = errors.New("validation error: address is not properly formatted")
	}

	if model.Address == nil || utils.ValidateAddress(*model.Address) == false {
		err = errors.New("validation error: address is not properly formatted")
	}

	if model.Amount == nil {
		err = errors.New("validation error: bid amount is required")
	}

	if model.ProductCode == 0 {
		err = errors.New("validation error: product cannot be zero")
	}

	if err != nil {
		return nil, err
	}

	return model, nil
}

func ValidateWithdrawReq(b []byte) (*core.WithdrawRequestModel, error) {
	var model *core.WithdrawRequestModel

	var err = json.Unmarshal(b, &model)

	if err != nil {
		return nil, err
	}

	return model, nil
}
