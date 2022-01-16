package client

import (
	"context"
	"errors"

	"github.com/akedev7/go-bff-microservices/quote/quotepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type Quote struct {
	Quote string `json:"quote"`
}

type QuoteClient struct {
}

var (
	_                      = loadLocalEnv()
	quoteGrpcService       = GetEnv("ADVICE_GRPC_SERVICE")
	quoteGrpcServiceClient quotepb.QuoteServiceClient
)

func prepareQuoteGrpcClient(c *context.Context) error {

	conn, err := grpc.DialContext(*c, quoteGrpcService, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock()}...)

	if err != nil {
		quoteGrpcServiceClient = nil
		return errors.New("connection to advice gRPC service failed")
	}

	if quoteGrpcServiceClient != nil {
		conn.Close()
		return nil
	}

	quoteGrpcServiceClient = quotepb.NewQuoteServiceClient(conn)
	return nil
}

func (qc *QuoteClient) GetQuote(id string, c *context.Context) (*Quote, error) {

	if err := prepareQuoteGrpcClient(c); err != nil {
		return nil, err
	}

	res, err := quoteGrpcServiceClient.GetQuote(*c, &quotepb.GetQuoteRequest{Id: id})
	if err != nil {
		return nil, errors.New(status.Convert(err).Message())
	}
	return &Quote{Quote: res.Quote}, nil
}
