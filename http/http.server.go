package http

import (
	"fmt"
	"net"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const (
	RouteGreeting               = "/greeting"
	RouteAuctionCreate          = "/auction/create"
	RouteAuctionView            = "/auction/view" // uses query string auctionName={}
	RouteAuctionProductAdd      = "/auction/product/add"
	RouteAuctionBidderAuthorize = "/auction/bidder/authorize"
	RouteAuctionBid             = "/auction/bid"
	RouteAuctionLogs            = "/auction/logs"
	RouteBlockchainView         = "/blockchain/view"
)

type Server struct {
	port     uint
	engine   *fasthttp.Server
	handlers *HttpHandlers
}

func NewServer(port uint) (*Server, error) {
	var handlers, err = NewHttpHandlers()
	if err != nil {
		return nil, err
	}
	return &Server{
		port:     port,
		handlers: handlers,
		engine:   &fasthttp.Server{},
	}, nil
}

func (s *Server) Start() error {
	var lis, err = net.Listen("tcp4", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	var route = router.New()

	route.GET(RouteGreeting, s.handlers.Greeting)
	route.GET(RouteAuctionView, s.handlers.AuctionView)
	route.GET(RouteAuctionLogs, s.handlers.AuctionLogs)
	route.POST(RouteAuctionCreate, s.handlers.AuctionCreate)
	route.POST(RouteAuctionProductAdd, s.handlers.AuctionProductAdd)
	route.POST(RouteAuctionBidderAuthorize, s.handlers.AuctionAuthorizeBidder)
	route.POST(RouteAuctionBid, s.handlers.Bid)
	route.GET(RouteBlockchainView, s.handlers.BlockchainAuctionView)

	s.engine.Handler = route.Handler
	return s.engine.Serve(lis)
}
