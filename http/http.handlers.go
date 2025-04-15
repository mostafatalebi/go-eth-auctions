package http

import (
	"go-eth-auction/core"
	"go-eth-auction/logger"
	"strings"

	"github.com/valyala/fasthttp"
)

type HttpHandlers struct {
	core *core.Container
}

func NewHttpHandlers() (*HttpHandlers, error) {
	var a, err = core.NewContainer()
	if err != nil {
		return nil, err
	}
	return &HttpHandlers{
		core: a,
	}, nil
}

func (hh *HttpHandlers) Greeting(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetBody([]byte("Hello from Ethereum-Based Go Multi-Product Auction Server!"))
}

func (hh *HttpHandlers) AuctionCreate(ctx *fasthttp.RequestCtx) {
	var body = ctx.Request.Body()

	var auc, err = ValidateIncomingAuctionCreateReq(body)
	ctx.Response.Header.Set("Content-Type", "application/json")
	if err != nil {
		ctx.Response.SetStatusCode(400)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	err = hh.core.Repo.Insert(auc)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.Response.SetStatusCode(404)
			return
		} else if strings.Contains(err.Error(), "duplicate") {
			ctx.Response.SetStatusCode(409)
			return
		}
		ctx.Response.SetStatusCode(500)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	data, err := RespondWithData(auc)

	if err != nil {
		ctx.Response.SetStatusCode(500)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	ctx.Response.SetStatusCode(200)
	ctx.Response.SetBody(data)
}

func (hh *HttpHandlers) AuctionProductAdd(ctx *fasthttp.RequestCtx) {
	var body = ctx.Request.Body()

	var model, err = ValidateIncomingAuctionProductAddReq(body)
	ctx.Response.Header.Set("Content-Type", "application/json")
	if err != nil {
		ctx.Response.SetStatusCode(400)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	err = hh.core.Repo.AddProduct(model.AuctionName, model)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.Response.SetStatusCode(404)
			return
		}
		ctx.Response.SetStatusCode(500)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	data, err := RespondWithData(model)

	if err != nil {
		ctx.Response.SetStatusCode(500)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	ctx.Response.SetStatusCode(200)
	ctx.Response.SetBody(data)
}

func (hh *HttpHandlers) AuctionAuthorizeBidder(ctx *fasthttp.RequestCtx) {
	var body = ctx.Request.Body()

	var model, err = ValidateIncomingAuctionAuthorizeBidderReq(body)
	ctx.Response.Header.Set("Content-Type", "application/json")
	if err != nil {
		ctx.Response.SetStatusCode(400)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	err = hh.core.Repo.AuthorizeBidder(model.AuctionName, model)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.Response.SetStatusCode(404)
			return
		}
		ctx.Response.SetStatusCode(500)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	data, err := RespondWithData(model)

	if err != nil {
		ctx.Response.SetStatusCode(500)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	ctx.Response.SetStatusCode(200)
	ctx.Response.SetBody(data)
}

func (hh *HttpHandlers) AuctionLogs(ctx *fasthttp.RequestCtx) {
	var auctionName = string(ctx.QueryArgs().Peek("auctionName"))
	ctx.Response.Header.Set("Content-Type", "application/json")
	if len(auctionName) == 0 {
		ctx.Response.SetStatusCode(400)
		ctx.Response.SetBody(RespondWithError("auctionName query param is required"))
		return
	}

	var logs, err = logger.Get().ListAsJson(auctionName)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.Response.SetStatusCode(404)
			return
		}
		ctx.Response.SetStatusCode(500)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	ctx.Response.SetStatusCode(200)
	ctx.Response.SetBody(logs)
}

func (hh *HttpHandlers) AuctionView(ctx *fasthttp.RequestCtx) {

	var auctionName = string(ctx.QueryArgs().Peek("auctionName"))
	ctx.Response.Header.Set("Content-Type", "application/json")
	if len(auctionName) == 0 {
		ctx.Response.SetStatusCode(400)
		ctx.Response.SetBody(RespondWithError("auctionName query param is required"))
		return
	}

	var model = hh.core.Repo.Get(auctionName)

	if model == nil {
		ctx.Response.SetStatusCode(404)
		ctx.Response.SetBody(RespondWithError("auction not found"))
		return
	}

	data, err := RespondWithData(model)

	if err != nil {
		ctx.Response.SetStatusCode(500)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	ctx.Response.SetStatusCode(200)
	ctx.Response.SetBody(data)
}

func (hh *HttpHandlers) BlockchainAuctionView(ctx *fasthttp.RequestCtx) {

	var auctionName = string(ctx.QueryArgs().Peek("auctionName"))
	ctx.Response.Header.Set("Content-Type", "application/json")
	if len(auctionName) == 0 {
		ctx.Response.SetStatusCode(400)
		ctx.Response.SetBody(RespondWithError("auctionName query param is required"))
		return
	}

	var data, err = hh.core.Repo.BlockchainView(auctionName)

	if err != nil {
		ctx.Response.SetStatusCode(500)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	ctx.Response.SetStatusCode(200)
	ctx.Response.SetBody(data)
}

func (hh *HttpHandlers) Bid(ctx *fasthttp.RequestCtx) {
	var body = ctx.Request.Body()

	var model, err = ValidateIncomingBidReq(body)
	ctx.Response.Header.Set("Content-Type", "application/json")
	if err != nil {

		ctx.Response.SetStatusCode(400)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	err = hh.core.Repo.Bid(model.AuctionName, model)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.Response.SetStatusCode(404)
			return
		}
		ctx.Response.SetStatusCode(500)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	data, err := RespondWithData(model)

	if err != nil {
		ctx.Response.SetStatusCode(500)
		ctx.Response.SetBody(RespondWithError(err.Error()))
		return
	}

	ctx.Response.SetStatusCode(200)
	ctx.Response.SetBody(data)
}
