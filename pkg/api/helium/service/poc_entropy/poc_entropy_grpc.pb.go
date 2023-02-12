// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service/poc_entropy.proto

package poc_entropy

import (
	context "context"
	helium "github.com/thisisdevelopment/helium-route-updater/pkg/api/helium"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PocEntropyClient is the client API for PocEntropy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PocEntropyClient interface {
	Entropy(ctx context.Context, in *EntropyReqV1, opts ...grpc.CallOption) (*helium.EntropyReportV1, error)
}

type pocEntropyClient struct {
	cc grpc.ClientConnInterface
}

func NewPocEntropyClient(cc grpc.ClientConnInterface) PocEntropyClient {
	return &pocEntropyClient{cc}
}

func (c *pocEntropyClient) Entropy(ctx context.Context, in *EntropyReqV1, opts ...grpc.CallOption) (*helium.EntropyReportV1, error) {
	out := new(helium.EntropyReportV1)
	err := c.cc.Invoke(ctx, "/helium.poc_entropy.poc_entropy/entropy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PocEntropyServer is the server API for PocEntropy service.
// All implementations must embed UnimplementedPocEntropyServer
// for forward compatibility
type PocEntropyServer interface {
	Entropy(context.Context, *EntropyReqV1) (*helium.EntropyReportV1, error)
	mustEmbedUnimplementedPocEntropyServer()
}

// UnimplementedPocEntropyServer must be embedded to have forward compatible implementations.
type UnimplementedPocEntropyServer struct {
}

func (UnimplementedPocEntropyServer) Entropy(context.Context, *EntropyReqV1) (*helium.EntropyReportV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Entropy not implemented")
}
func (UnimplementedPocEntropyServer) mustEmbedUnimplementedPocEntropyServer() {}

// UnsafePocEntropyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PocEntropyServer will
// result in compilation errors.
type UnsafePocEntropyServer interface {
	mustEmbedUnimplementedPocEntropyServer()
}

func RegisterPocEntropyServer(s grpc.ServiceRegistrar, srv PocEntropyServer) {
	s.RegisterService(&PocEntropy_ServiceDesc, srv)
}

func _PocEntropy_Entropy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntropyReqV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PocEntropyServer).Entropy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helium.poc_entropy.poc_entropy/entropy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PocEntropyServer).Entropy(ctx, req.(*EntropyReqV1))
	}
	return interceptor(ctx, in, info, handler)
}

// PocEntropy_ServiceDesc is the grpc.ServiceDesc for PocEntropy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PocEntropy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helium.poc_entropy.poc_entropy",
	HandlerType: (*PocEntropyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "entropy",
			Handler:    _PocEntropy_Entropy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/poc_entropy.proto",
}
