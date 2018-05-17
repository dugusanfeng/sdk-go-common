// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet/services.proto

package wallet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TxSrvc service

type TxSrvcClient interface {
	// ProcessTx handles the UTXO transaction request, TxIN -> TxOUT
	ProcessTx(ctx context.Context, in *TX, opts ...grpc.CallOption) (*ProcessTxResponse, error)
	QueryTx(ctx context.Context, in *QueryTxRequest, opts ...grpc.CallOption) (*TX, error)
	QueryUTXO(ctx context.Context, in *QueryUTXORequest, opts ...grpc.CallOption) (*QueryUTXOResponse, error)
}

type txSrvcClient struct {
	cc *grpc.ClientConn
}

func NewTxSrvcClient(cc *grpc.ClientConn) TxSrvcClient {
	return &txSrvcClient{cc}
}

func (c *txSrvcClient) ProcessTx(ctx context.Context, in *TX, opts ...grpc.CallOption) (*ProcessTxResponse, error) {
	out := new(ProcessTxResponse)
	err := grpc.Invoke(ctx, "/wallet.TxSrvc/ProcessTx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txSrvcClient) QueryTx(ctx context.Context, in *QueryTxRequest, opts ...grpc.CallOption) (*TX, error) {
	out := new(TX)
	err := grpc.Invoke(ctx, "/wallet.TxSrvc/QueryTx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txSrvcClient) QueryUTXO(ctx context.Context, in *QueryUTXORequest, opts ...grpc.CallOption) (*QueryUTXOResponse, error) {
	out := new(QueryUTXOResponse)
	err := grpc.Invoke(ctx, "/wallet.TxSrvc/QueryUTXO", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TxSrvc service

type TxSrvcServer interface {
	// ProcessTx handles the UTXO transaction request, TxIN -> TxOUT
	ProcessTx(context.Context, *TX) (*ProcessTxResponse, error)
	QueryTx(context.Context, *QueryTxRequest) (*TX, error)
	QueryUTXO(context.Context, *QueryUTXORequest) (*QueryUTXOResponse, error)
}

func RegisterTxSrvcServer(s *grpc.Server, srv TxSrvcServer) {
	s.RegisterService(&_TxSrvc_serviceDesc, srv)
}

func _TxSrvc_ProcessTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TX)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxSrvcServer).ProcessTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.TxSrvc/ProcessTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxSrvcServer).ProcessTx(ctx, req.(*TX))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxSrvc_QueryTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxSrvcServer).QueryTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.TxSrvc/QueryTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxSrvcServer).QueryTx(ctx, req.(*QueryTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxSrvc_QueryUTXO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUTXORequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxSrvcServer).QueryUTXO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.TxSrvc/QueryUTXO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxSrvcServer).QueryUTXO(ctx, req.(*QueryUTXORequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TxSrvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wallet.TxSrvc",
	HandlerType: (*TxSrvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessTx",
			Handler:    _TxSrvc_ProcessTx_Handler,
		},
		{
			MethodName: "QueryTx",
			Handler:    _TxSrvc_QueryTx_Handler,
		},
		{
			MethodName: "QueryUTXO",
			Handler:    _TxSrvc_QueryUTXO_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet/services.proto",
}

func init() { proto.RegisterFile("wallet/services.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4f, 0xcc, 0xc9,
	0x49, 0x2d, 0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x83, 0x08, 0x4b, 0xf1, 0x43, 0xa5, 0x4b, 0x2a, 0x20, 0x12, 0x46, 0x9b, 0x19,
	0xb9, 0xd8, 0x42, 0x2a, 0x82, 0x8b, 0xca, 0x92, 0x85, 0x4c, 0xb8, 0x38, 0x03, 0x8a, 0xf2, 0x93,
	0x53, 0x8b, 0x8b, 0x43, 0x2a, 0x84, 0xb8, 0xf4, 0x20, 0x2a, 0xf5, 0x42, 0x22, 0xa4, 0x24, 0x61,
	0x6c, 0xb8, 0x74, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x12, 0x83, 0x90, 0x3e, 0x17,
	0x7b, 0x60, 0x69, 0x6a, 0x51, 0x65, 0x48, 0x85, 0x90, 0x18, 0x4c, 0x1d, 0x54, 0x20, 0x28, 0xb5,
	0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x0a, 0xc9, 0x2c, 0x25, 0x06, 0x21, 0x27, 0x2e, 0x4e, 0xb0, 0x7c,
	0x68, 0x48, 0x84, 0xbf, 0x90, 0x04, 0x8a, 0x16, 0x90, 0x10, 0x4c, 0x93, 0x24, 0x16, 0x19, 0x98,
	0xa5, 0x4e, 0xc6, 0x51, 0x86, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x89, 0x45, 0x15, 0x89, 0x79, 0xc9, 0x19, 0x89, 0x99, 0x79, 0xfa, 0xc5, 0x29, 0xd9, 0xba, 0xe9,
	0xf9, 0xba, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0xfa, 0x60, 0x2f, 0x16, 0xeb, 0x43, 0x0c, 0x4a,
	0x62, 0x03, 0x73, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x95, 0x8f, 0x97, 0x4c, 0x23, 0x01,
	0x00, 0x00,
}
