cd quote/quotepb/
protoc --go_out=. quote.proto
protoc --go-grpc_out=. quote.proto