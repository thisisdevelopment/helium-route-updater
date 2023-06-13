// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service/local.proto

package local

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

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	Pubkey(ctx context.Context, in *PubkeyReq, opts ...grpc.CallOption) (*PubkeyRes, error)
	Region(ctx context.Context, in *RegionReq, opts ...grpc.CallOption) (*RegionRes, error)
	Router(ctx context.Context, in *RouterReq, opts ...grpc.CallOption) (*RouterRes, error)
	AddGateway(ctx context.Context, in *AddGatewayReq, opts ...grpc.CallOption) (*AddGatewayRes, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) Pubkey(ctx context.Context, in *PubkeyReq, opts ...grpc.CallOption) (*PubkeyRes, error) {
	out := new(PubkeyRes)
	err := c.cc.Invoke(ctx, "/helium.local.api/pubkey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Region(ctx context.Context, in *RegionReq, opts ...grpc.CallOption) (*RegionRes, error) {
	out := new(RegionRes)
	err := c.cc.Invoke(ctx, "/helium.local.api/region", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Router(ctx context.Context, in *RouterReq, opts ...grpc.CallOption) (*RouterRes, error) {
	out := new(RouterRes)
	err := c.cc.Invoke(ctx, "/helium.local.api/router", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) AddGateway(ctx context.Context, in *AddGatewayReq, opts ...grpc.CallOption) (*AddGatewayRes, error) {
	out := new(AddGatewayRes)
	err := c.cc.Invoke(ctx, "/helium.local.api/add_gateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	Pubkey(context.Context, *PubkeyReq) (*PubkeyRes, error)
	Region(context.Context, *RegionReq) (*RegionRes, error)
	Router(context.Context, *RouterReq) (*RouterRes, error)
	AddGateway(context.Context, *AddGatewayReq) (*AddGatewayRes, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) Pubkey(context.Context, *PubkeyReq) (*PubkeyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pubkey not implemented")
}
func (UnimplementedApiServer) Region(context.Context, *RegionReq) (*RegionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Region not implemented")
}
func (UnimplementedApiServer) Router(context.Context, *RouterReq) (*RouterRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Router not implemented")
}
func (UnimplementedApiServer) AddGateway(context.Context, *AddGatewayReq) (*AddGatewayRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGateway not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_Pubkey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubkeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Pubkey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.local.api/pubkey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Pubkey(ctx, req.(*PubkeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Region_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Region(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.local.api/region",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Region(ctx, req.(*RegionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Router_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Router(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.local.api/router",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Router(ctx, req.(*RouterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_AddGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGatewayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).AddGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.local.api/add_gateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).AddGateway(ctx, req.(*AddGatewayReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helium.local.api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "pubkey",
			Handler:    _Api_Pubkey_Handler,
		},
		{
			MethodName: "region",
			Handler:    _Api_Region_Handler,
		},
		{
			MethodName: "router",
			Handler:    _Api_Router_Handler,
		},
		{
			MethodName: "add_gateway",
			Handler:    _Api_AddGateway_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/local.proto",
}
