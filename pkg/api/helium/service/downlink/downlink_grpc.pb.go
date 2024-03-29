// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: service/downlink.proto

package downlink

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

// HttpRoamingClient is the client API for HttpRoaming service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HttpRoamingClient interface {
	Stream(ctx context.Context, in *HttpRoamingRegisterV1, opts ...grpc.CallOption) (HttpRoaming_StreamClient, error)
}

type httpRoamingClient struct {
	cc grpc.ClientConnInterface
}

func NewHttpRoamingClient(cc grpc.ClientConnInterface) HttpRoamingClient {
	return &httpRoamingClient{cc}
}

func (c *httpRoamingClient) Stream(ctx context.Context, in *HttpRoamingRegisterV1, opts ...grpc.CallOption) (HttpRoaming_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &HttpRoaming_ServiceDesc.Streams[0], "/helium.downlink.http_roaming/stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &httpRoamingStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HttpRoaming_StreamClient interface {
	Recv() (*HttpRoamingDownlinkV1, error)
	grpc.ClientStream
}

type httpRoamingStreamClient struct {
	grpc.ClientStream
}

func (x *httpRoamingStreamClient) Recv() (*HttpRoamingDownlinkV1, error) {
	m := new(HttpRoamingDownlinkV1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HttpRoamingServer is the server API for HttpRoaming service.
// All implementations must embed UnimplementedHttpRoamingServer
// for forward compatibility
type HttpRoamingServer interface {
	Stream(*HttpRoamingRegisterV1, HttpRoaming_StreamServer) error
	mustEmbedUnimplementedHttpRoamingServer()
}

// UnimplementedHttpRoamingServer must be embedded to have forward compatible implementations.
type UnimplementedHttpRoamingServer struct {
}

func (UnimplementedHttpRoamingServer) Stream(*HttpRoamingRegisterV1, HttpRoaming_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedHttpRoamingServer) mustEmbedUnimplementedHttpRoamingServer() {}

// UnsafeHttpRoamingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HttpRoamingServer will
// result in compilation errors.
type UnsafeHttpRoamingServer interface {
	mustEmbedUnimplementedHttpRoamingServer()
}

func RegisterHttpRoamingServer(s grpc.ServiceRegistrar, srv HttpRoamingServer) {
	s.RegisterService(&HttpRoaming_ServiceDesc, srv)
}

func _HttpRoaming_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HttpRoamingRegisterV1)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HttpRoamingServer).Stream(m, &httpRoamingStreamServer{stream})
}

type HttpRoaming_StreamServer interface {
	Send(*HttpRoamingDownlinkV1) error
	grpc.ServerStream
}

type httpRoamingStreamServer struct {
	grpc.ServerStream
}

func (x *httpRoamingStreamServer) Send(m *HttpRoamingDownlinkV1) error {
	return x.ServerStream.SendMsg(m)
}

// HttpRoaming_ServiceDesc is the grpc.ServiceDesc for HttpRoaming service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HttpRoaming_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helium.downlink.http_roaming",
	HandlerType: (*HttpRoamingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "stream",
			Handler:       _HttpRoaming_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service/downlink.proto",
}
