// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service/gateway.proto

package gateway

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

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	Version(ctx context.Context, in *GatewayVersionReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error)
	FollowSc(ctx context.Context, opts ...grpc.CallOption) (Gateway_FollowScClient, error)
	Routing(ctx context.Context, in *GatewayRoutingReqV1, opts ...grpc.CallOption) (Gateway_RoutingClient, error)
	StreamPoc(ctx context.Context, in *GatewayPocReqV1, opts ...grpc.CallOption) (Gateway_StreamPocClient, error)
	ConfigUpdate(ctx context.Context, in *GatewayConfigUpdateReqV1, opts ...grpc.CallOption) (Gateway_ConfigUpdateClient, error)
	IsActiveSc(ctx context.Context, in *GatewayScIsActiveReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error)
	IsOverpaidSc(ctx context.Context, in *GatewayScIsOverpaidReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error)
	CloseSc(ctx context.Context, in *GatewayScCloseReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error)
	CheckChallengeTarget(ctx context.Context, in *GatewayPocCheckChallengeTargetReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error)
	SendReport(ctx context.Context, in *GatewayPocReportReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error)
	AddressToPublicUri(ctx context.Context, in *GatewayAddressRoutingDataReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error)
	PocKeyToPublicUri(ctx context.Context, in *GatewayPocKeyRoutingDataReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error)
	Config(ctx context.Context, in *GatewayConfigReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error)
	Validators(ctx context.Context, in *GatewayValidatorsReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error)
	RegionParamsUpdate(ctx context.Context, in *GatewayRegionParamsUpdateReqV1, opts ...grpc.CallOption) (Gateway_RegionParamsUpdateClient, error)
	RegionParams(ctx context.Context, in *GatewayRegionParamsReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) Version(ctx context.Context, in *GatewayVersionReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error) {
	out := new(GatewayRespV1)
	err := c.cc.Invoke(ctx, "/helium.gateway/version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) FollowSc(ctx context.Context, opts ...grpc.CallOption) (Gateway_FollowScClient, error) {
	stream, err := c.cc.NewStream(ctx, &Gateway_ServiceDesc.Streams[0], "/helium.gateway/follow_sc", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayFollowScClient{stream}
	return x, nil
}

type Gateway_FollowScClient interface {
	Send(*GatewayScFollowReqV1) error
	Recv() (*GatewayRespV1, error)
	grpc.ClientStream
}

type gatewayFollowScClient struct {
	grpc.ClientStream
}

func (x *gatewayFollowScClient) Send(m *GatewayScFollowReqV1) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gatewayFollowScClient) Recv() (*GatewayRespV1, error) {
	m := new(GatewayRespV1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gatewayClient) Routing(ctx context.Context, in *GatewayRoutingReqV1, opts ...grpc.CallOption) (Gateway_RoutingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Gateway_ServiceDesc.Streams[1], "/helium.gateway/routing", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayRoutingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Gateway_RoutingClient interface {
	Recv() (*GatewayRespV1, error)
	grpc.ClientStream
}

type gatewayRoutingClient struct {
	grpc.ClientStream
}

func (x *gatewayRoutingClient) Recv() (*GatewayRespV1, error) {
	m := new(GatewayRespV1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gatewayClient) StreamPoc(ctx context.Context, in *GatewayPocReqV1, opts ...grpc.CallOption) (Gateway_StreamPocClient, error) {
	stream, err := c.cc.NewStream(ctx, &Gateway_ServiceDesc.Streams[2], "/helium.gateway/stream_poc", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayStreamPocClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Gateway_StreamPocClient interface {
	Recv() (*GatewayRespV1, error)
	grpc.ClientStream
}

type gatewayStreamPocClient struct {
	grpc.ClientStream
}

func (x *gatewayStreamPocClient) Recv() (*GatewayRespV1, error) {
	m := new(GatewayRespV1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gatewayClient) ConfigUpdate(ctx context.Context, in *GatewayConfigUpdateReqV1, opts ...grpc.CallOption) (Gateway_ConfigUpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Gateway_ServiceDesc.Streams[3], "/helium.gateway/config_update", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayConfigUpdateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Gateway_ConfigUpdateClient interface {
	Recv() (*GatewayRespV1, error)
	grpc.ClientStream
}

type gatewayConfigUpdateClient struct {
	grpc.ClientStream
}

func (x *gatewayConfigUpdateClient) Recv() (*GatewayRespV1, error) {
	m := new(GatewayRespV1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gatewayClient) IsActiveSc(ctx context.Context, in *GatewayScIsActiveReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error) {
	out := new(GatewayRespV1)
	err := c.cc.Invoke(ctx, "/helium.gateway/is_active_sc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) IsOverpaidSc(ctx context.Context, in *GatewayScIsOverpaidReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error) {
	out := new(GatewayRespV1)
	err := c.cc.Invoke(ctx, "/helium.gateway/is_overpaid_sc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CloseSc(ctx context.Context, in *GatewayScCloseReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error) {
	out := new(GatewayRespV1)
	err := c.cc.Invoke(ctx, "/helium.gateway/close_sc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CheckChallengeTarget(ctx context.Context, in *GatewayPocCheckChallengeTargetReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error) {
	out := new(GatewayRespV1)
	err := c.cc.Invoke(ctx, "/helium.gateway/check_challenge_target", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) SendReport(ctx context.Context, in *GatewayPocReportReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error) {
	out := new(GatewayRespV1)
	err := c.cc.Invoke(ctx, "/helium.gateway/send_report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AddressToPublicUri(ctx context.Context, in *GatewayAddressRoutingDataReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error) {
	out := new(GatewayRespV1)
	err := c.cc.Invoke(ctx, "/helium.gateway/address_to_public_uri", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) PocKeyToPublicUri(ctx context.Context, in *GatewayPocKeyRoutingDataReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error) {
	out := new(GatewayRespV1)
	err := c.cc.Invoke(ctx, "/helium.gateway/poc_key_to_public_uri", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Config(ctx context.Context, in *GatewayConfigReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error) {
	out := new(GatewayRespV1)
	err := c.cc.Invoke(ctx, "/helium.gateway/config", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Validators(ctx context.Context, in *GatewayValidatorsReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error) {
	out := new(GatewayRespV1)
	err := c.cc.Invoke(ctx, "/helium.gateway/validators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) RegionParamsUpdate(ctx context.Context, in *GatewayRegionParamsUpdateReqV1, opts ...grpc.CallOption) (Gateway_RegionParamsUpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Gateway_ServiceDesc.Streams[4], "/helium.gateway/region_params_update", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayRegionParamsUpdateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Gateway_RegionParamsUpdateClient interface {
	Recv() (*GatewayRespV1, error)
	grpc.ClientStream
}

type gatewayRegionParamsUpdateClient struct {
	grpc.ClientStream
}

func (x *gatewayRegionParamsUpdateClient) Recv() (*GatewayRespV1, error) {
	m := new(GatewayRespV1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gatewayClient) RegionParams(ctx context.Context, in *GatewayRegionParamsReqV1, opts ...grpc.CallOption) (*GatewayRespV1, error) {
	out := new(GatewayRespV1)
	err := c.cc.Invoke(ctx, "/helium.gateway/region_params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	Version(context.Context, *GatewayVersionReqV1) (*GatewayRespV1, error)
	FollowSc(Gateway_FollowScServer) error
	Routing(*GatewayRoutingReqV1, Gateway_RoutingServer) error
	StreamPoc(*GatewayPocReqV1, Gateway_StreamPocServer) error
	ConfigUpdate(*GatewayConfigUpdateReqV1, Gateway_ConfigUpdateServer) error
	IsActiveSc(context.Context, *GatewayScIsActiveReqV1) (*GatewayRespV1, error)
	IsOverpaidSc(context.Context, *GatewayScIsOverpaidReqV1) (*GatewayRespV1, error)
	CloseSc(context.Context, *GatewayScCloseReqV1) (*GatewayRespV1, error)
	CheckChallengeTarget(context.Context, *GatewayPocCheckChallengeTargetReqV1) (*GatewayRespV1, error)
	SendReport(context.Context, *GatewayPocReportReqV1) (*GatewayRespV1, error)
	AddressToPublicUri(context.Context, *GatewayAddressRoutingDataReqV1) (*GatewayRespV1, error)
	PocKeyToPublicUri(context.Context, *GatewayPocKeyRoutingDataReqV1) (*GatewayRespV1, error)
	Config(context.Context, *GatewayConfigReqV1) (*GatewayRespV1, error)
	Validators(context.Context, *GatewayValidatorsReqV1) (*GatewayRespV1, error)
	RegionParamsUpdate(*GatewayRegionParamsUpdateReqV1, Gateway_RegionParamsUpdateServer) error
	RegionParams(context.Context, *GatewayRegionParamsReqV1) (*GatewayRespV1, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) Version(context.Context, *GatewayVersionReqV1) (*GatewayRespV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedGatewayServer) FollowSc(Gateway_FollowScServer) error {
	return status.Errorf(codes.Unimplemented, "method FollowSc not implemented")
}
func (UnimplementedGatewayServer) Routing(*GatewayRoutingReqV1, Gateway_RoutingServer) error {
	return status.Errorf(codes.Unimplemented, "method Routing not implemented")
}
func (UnimplementedGatewayServer) StreamPoc(*GatewayPocReqV1, Gateway_StreamPocServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPoc not implemented")
}
func (UnimplementedGatewayServer) ConfigUpdate(*GatewayConfigUpdateReqV1, Gateway_ConfigUpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method ConfigUpdate not implemented")
}
func (UnimplementedGatewayServer) IsActiveSc(context.Context, *GatewayScIsActiveReqV1) (*GatewayRespV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsActiveSc not implemented")
}
func (UnimplementedGatewayServer) IsOverpaidSc(context.Context, *GatewayScIsOverpaidReqV1) (*GatewayRespV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsOverpaidSc not implemented")
}
func (UnimplementedGatewayServer) CloseSc(context.Context, *GatewayScCloseReqV1) (*GatewayRespV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseSc not implemented")
}
func (UnimplementedGatewayServer) CheckChallengeTarget(context.Context, *GatewayPocCheckChallengeTargetReqV1) (*GatewayRespV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckChallengeTarget not implemented")
}
func (UnimplementedGatewayServer) SendReport(context.Context, *GatewayPocReportReqV1) (*GatewayRespV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendReport not implemented")
}
func (UnimplementedGatewayServer) AddressToPublicUri(context.Context, *GatewayAddressRoutingDataReqV1) (*GatewayRespV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressToPublicUri not implemented")
}
func (UnimplementedGatewayServer) PocKeyToPublicUri(context.Context, *GatewayPocKeyRoutingDataReqV1) (*GatewayRespV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PocKeyToPublicUri not implemented")
}
func (UnimplementedGatewayServer) Config(context.Context, *GatewayConfigReqV1) (*GatewayRespV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Config not implemented")
}
func (UnimplementedGatewayServer) Validators(context.Context, *GatewayValidatorsReqV1) (*GatewayRespV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validators not implemented")
}
func (UnimplementedGatewayServer) RegionParamsUpdate(*GatewayRegionParamsUpdateReqV1, Gateway_RegionParamsUpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method RegionParamsUpdate not implemented")
}
func (UnimplementedGatewayServer) RegionParams(context.Context, *GatewayRegionParamsReqV1) (*GatewayRespV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegionParams not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayVersionReqV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.gateway/version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Version(ctx, req.(*GatewayVersionReqV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_FollowSc_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GatewayServer).FollowSc(&gatewayFollowScServer{stream})
}

type Gateway_FollowScServer interface {
	Send(*GatewayRespV1) error
	Recv() (*GatewayScFollowReqV1, error)
	grpc.ServerStream
}

type gatewayFollowScServer struct {
	grpc.ServerStream
}

func (x *gatewayFollowScServer) Send(m *GatewayRespV1) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gatewayFollowScServer) Recv() (*GatewayScFollowReqV1, error) {
	m := new(GatewayScFollowReqV1)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Gateway_Routing_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GatewayRoutingReqV1)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayServer).Routing(m, &gatewayRoutingServer{stream})
}

type Gateway_RoutingServer interface {
	Send(*GatewayRespV1) error
	grpc.ServerStream
}

type gatewayRoutingServer struct {
	grpc.ServerStream
}

func (x *gatewayRoutingServer) Send(m *GatewayRespV1) error {
	return x.ServerStream.SendMsg(m)
}

func _Gateway_StreamPoc_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GatewayPocReqV1)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayServer).StreamPoc(m, &gatewayStreamPocServer{stream})
}

type Gateway_StreamPocServer interface {
	Send(*GatewayRespV1) error
	grpc.ServerStream
}

type gatewayStreamPocServer struct {
	grpc.ServerStream
}

func (x *gatewayStreamPocServer) Send(m *GatewayRespV1) error {
	return x.ServerStream.SendMsg(m)
}

func _Gateway_ConfigUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GatewayConfigUpdateReqV1)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayServer).ConfigUpdate(m, &gatewayConfigUpdateServer{stream})
}

type Gateway_ConfigUpdateServer interface {
	Send(*GatewayRespV1) error
	grpc.ServerStream
}

type gatewayConfigUpdateServer struct {
	grpc.ServerStream
}

func (x *gatewayConfigUpdateServer) Send(m *GatewayRespV1) error {
	return x.ServerStream.SendMsg(m)
}

func _Gateway_IsActiveSc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayScIsActiveReqV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).IsActiveSc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.gateway/is_active_sc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).IsActiveSc(ctx, req.(*GatewayScIsActiveReqV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_IsOverpaidSc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayScIsOverpaidReqV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).IsOverpaidSc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.gateway/is_overpaid_sc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).IsOverpaidSc(ctx, req.(*GatewayScIsOverpaidReqV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CloseSc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayScCloseReqV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CloseSc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.gateway/close_sc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CloseSc(ctx, req.(*GatewayScCloseReqV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CheckChallengeTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayPocCheckChallengeTargetReqV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CheckChallengeTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.gateway/check_challenge_target",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CheckChallengeTarget(ctx, req.(*GatewayPocCheckChallengeTargetReqV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_SendReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayPocReportReqV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).SendReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.gateway/send_report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).SendReport(ctx, req.(*GatewayPocReportReqV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AddressToPublicUri_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayAddressRoutingDataReqV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AddressToPublicUri(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.gateway/address_to_public_uri",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AddressToPublicUri(ctx, req.(*GatewayAddressRoutingDataReqV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_PocKeyToPublicUri_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayPocKeyRoutingDataReqV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).PocKeyToPublicUri(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.gateway/poc_key_to_public_uri",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).PocKeyToPublicUri(ctx, req.(*GatewayPocKeyRoutingDataReqV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayConfigReqV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Config(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.gateway/config",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Config(ctx, req.(*GatewayConfigReqV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Validators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayValidatorsReqV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Validators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.gateway/validators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Validators(ctx, req.(*GatewayValidatorsReqV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_RegionParamsUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GatewayRegionParamsUpdateReqV1)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayServer).RegionParamsUpdate(m, &gatewayRegionParamsUpdateServer{stream})
}

type Gateway_RegionParamsUpdateServer interface {
	Send(*GatewayRespV1) error
	grpc.ServerStream
}

type gatewayRegionParamsUpdateServer struct {
	grpc.ServerStream
}

func (x *gatewayRegionParamsUpdateServer) Send(m *GatewayRespV1) error {
	return x.ServerStream.SendMsg(m)
}

func _Gateway_RegionParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayRegionParamsReqV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).RegionParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.gateway/region_params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).RegionParams(ctx, req.(*GatewayRegionParamsReqV1))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helium.gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "version",
			Handler:    _Gateway_Version_Handler,
		},
		{
			MethodName: "is_active_sc",
			Handler:    _Gateway_IsActiveSc_Handler,
		},
		{
			MethodName: "is_overpaid_sc",
			Handler:    _Gateway_IsOverpaidSc_Handler,
		},
		{
			MethodName: "close_sc",
			Handler:    _Gateway_CloseSc_Handler,
		},
		{
			MethodName: "check_challenge_target",
			Handler:    _Gateway_CheckChallengeTarget_Handler,
		},
		{
			MethodName: "send_report",
			Handler:    _Gateway_SendReport_Handler,
		},
		{
			MethodName: "address_to_public_uri",
			Handler:    _Gateway_AddressToPublicUri_Handler,
		},
		{
			MethodName: "poc_key_to_public_uri",
			Handler:    _Gateway_PocKeyToPublicUri_Handler,
		},
		{
			MethodName: "config",
			Handler:    _Gateway_Config_Handler,
		},
		{
			MethodName: "validators",
			Handler:    _Gateway_Validators_Handler,
		},
		{
			MethodName: "region_params",
			Handler:    _Gateway_RegionParams_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "follow_sc",
			Handler:       _Gateway_FollowSc_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "routing",
			Handler:       _Gateway_Routing_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "stream_poc",
			Handler:       _Gateway_StreamPoc_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "config_update",
			Handler:       _Gateway_ConfigUpdate_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "region_params_update",
			Handler:       _Gateway_RegionParamsUpdate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service/gateway.proto",
}
