// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: proto/multiplication.proto

package proto

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

// MultiplicationClient is the client API for Multiplication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MultiplicationClient interface {
	Multiply(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error)
}

type multiplicationClient struct {
	cc grpc.ClientConnInterface
}

func NewMultiplicationClient(cc grpc.ClientConnInterface) MultiplicationClient {
	return &multiplicationClient{cc}
}

func (c *multiplicationClient) Multiply(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := c.cc.Invoke(ctx, "/proto.Multiplication/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MultiplicationServer is the server API for Multiplication service.
// All implementations must embed UnimplementedMultiplicationServer
// for forward compatibility
type MultiplicationServer interface {
	Multiply(context.Context, *Input) (*Output, error)
	mustEmbedUnimplementedMultiplicationServer()
}

// UnimplementedMultiplicationServer must be embedded to have forward compatible implementations.
type UnimplementedMultiplicationServer struct {
}

func (UnimplementedMultiplicationServer) Multiply(context.Context, *Input) (*Output, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedMultiplicationServer) mustEmbedUnimplementedMultiplicationServer() {}

// UnsafeMultiplicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MultiplicationServer will
// result in compilation errors.
type UnsafeMultiplicationServer interface {
	mustEmbedUnimplementedMultiplicationServer()
}

func RegisterMultiplicationServer(s grpc.ServiceRegistrar, srv MultiplicationServer) {
	s.RegisterService(&Multiplication_ServiceDesc, srv)
}

func _Multiplication_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiplicationServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Multiplication/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiplicationServer).Multiply(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

// Multiplication_ServiceDesc is the grpc.ServiceDesc for Multiplication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Multiplication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Multiplication",
	HandlerType: (*MultiplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Multiply",
			Handler:    _Multiplication_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/multiplication.proto",
}