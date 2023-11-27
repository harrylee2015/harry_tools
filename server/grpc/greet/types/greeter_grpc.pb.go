// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: greeter.proto

package types

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

const (
	Greeter_Greeter_FullMethodName  = "/types.Greeter/Greeter"
	Greeter_GetGreet_FullMethodName = "/types.Greeter/GetGreet"
)

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	Greeter(ctx context.Context, in *GreeterRequest, opts ...grpc.CallOption) (*GreeterResponse, error)
	GetGreet(ctx context.Context, in *GreeterRequest, opts ...grpc.CallOption) (*GreeterResponse, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Greeter(ctx context.Context, in *GreeterRequest, opts ...grpc.CallOption) (*GreeterResponse, error) {
	out := new(GreeterResponse)
	err := c.cc.Invoke(ctx, Greeter_Greeter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetGreet(ctx context.Context, in *GreeterRequest, opts ...grpc.CallOption) (*GreeterResponse, error) {
	out := new(GreeterResponse)
	err := c.cc.Invoke(ctx, Greeter_GetGreet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations should embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	Greeter(context.Context, *GreeterRequest) (*GreeterResponse, error)
	GetGreet(context.Context, *GreeterRequest) (*GreeterResponse, error)
}

// UnimplementedGreeterServer should be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) Greeter(context.Context, *GreeterRequest) (*GreeterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greeter not implemented")
}
func (UnimplementedGreeterServer) GetGreet(context.Context, *GreeterRequest) (*GreeterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGreet not implemented")
}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_Greeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreeterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Greeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_Greeter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Greeter(ctx, req.(*GreeterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetGreet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreeterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetGreet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_GetGreet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetGreet(ctx, req.(*GreeterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "types.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greeter",
			Handler:    _Greeter_Greeter_Handler,
		},
		{
			MethodName: "GetGreet",
			Handler:    _Greeter_GetGreet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeter.proto",
}