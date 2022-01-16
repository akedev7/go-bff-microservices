// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: quote.proto

package quotepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// QuoteServiceClient is the client API for QuoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuoteServiceClient interface {
	GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...grpc.CallOption) (*GetQuoteReponse, error)
}

type quoteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuoteServiceClient(cc grpc.ClientConnInterface) QuoteServiceClient {
	return &quoteServiceClient{cc}
}

func (c *quoteServiceClient) GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...grpc.CallOption) (*GetQuoteReponse, error) {
	out := new(GetQuoteReponse)
	err := c.cc.Invoke(ctx, "/quotepb.QuoteService/GetQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuoteServiceServer is the server API for QuoteService service.
// All implementations must embed UnimplementedQuoteServiceServer
// for forward compatibility
type QuoteServiceServer interface {
	GetQuote(context.Context, *GetQuoteRequest) (*GetQuoteReponse, error)
	mustEmbedUnimplementedQuoteServiceServer()
}

// UnimplementedQuoteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuoteServiceServer struct {
}

func (UnimplementedQuoteServiceServer) GetQuote(context.Context, *GetQuoteRequest) (*GetQuoteReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuote not implemented")
}
func (UnimplementedQuoteServiceServer) mustEmbedUnimplementedQuoteServiceServer() {}

// UnsafeQuoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuoteServiceServer will
// result in compilation errors.
type UnsafeQuoteServiceServer interface {
	mustEmbedUnimplementedQuoteServiceServer()
}

func RegisterQuoteServiceServer(s grpc.ServiceRegistrar, srv QuoteServiceServer) {
	s.RegisterService(&QuoteService_ServiceDesc, srv)
}

func _QuoteService_GetQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).GetQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotepb.QuoteService/GetQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).GetQuote(ctx, req.(*GetQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuoteService_ServiceDesc is the grpc.ServiceDesc for QuoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "quotepb.QuoteService",
	HandlerType: (*QuoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuote",
			Handler:    _QuoteService_GetQuote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quote.proto",
}
