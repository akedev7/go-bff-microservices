package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/akedev7/go-bff-microservices/quote/quotepb"
	"google.golang.org/grpc"
)

var (
	timeout = time.Second
)

type server struct {
	quotepb.UnimplementedQuoteServiceServer
}

func (*server) GetQuote(ctx context.Context, req *quotepb.GetQuoteRequest) (*quotepb.GetQuoteReponse, error) {

	_, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	log.Println("Called GetAdvice for User Id", req.Id)
	return &quotepb.GetQuoteReponse{
		Quote: req.Id,
	}, nil
}

func main() {
	log.Println("Quote Service")

	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Println("ERROR:", err.Error())
	}

	s := grpc.NewServer()
	quotepb.RegisterQuoteServiceServer(s, &server{})
	log.Printf("Server started at %v", lis.Addr().String())

	err = s.Serve(lis)
	if err != nil {
		log.Println("ERROR:", err.Error())
	}

}
