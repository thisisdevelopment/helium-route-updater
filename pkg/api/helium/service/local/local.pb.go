// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: service/local.proto

package local

import (
	helium "github.com/thisisdevelopment/helium-route-updater/pkg/api/helium"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PubkeyRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address           []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	OnboardingAddress []byte `protobuf:"bytes,2,opt,name=onboarding_address,json=onboardingAddress,proto3" json:"onboarding_address,omitempty"`
}

func (x *PubkeyRes) Reset() {
	*x = PubkeyRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_local_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubkeyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubkeyRes) ProtoMessage() {}

func (x *PubkeyRes) ProtoReflect() protoreflect.Message {
	mi := &file_service_local_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubkeyRes.ProtoReflect.Descriptor instead.
func (*PubkeyRes) Descriptor() ([]byte, []int) {
	return file_service_local_proto_rawDescGZIP(), []int{0}
}

func (x *PubkeyRes) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *PubkeyRes) GetOnboardingAddress() []byte {
	if x != nil {
		return x.OnboardingAddress
	}
	return nil
}

type PubkeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PubkeyReq) Reset() {
	*x = PubkeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_local_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubkeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubkeyReq) ProtoMessage() {}

func (x *PubkeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_local_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubkeyReq.ProtoReflect.Descriptor instead.
func (*PubkeyReq) Descriptor() ([]byte, []int) {
	return file_service_local_proto_rawDescGZIP(), []int{1}
}

type KeyedUri struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Uri     string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *KeyedUri) Reset() {
	*x = KeyedUri{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_local_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyedUri) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyedUri) ProtoMessage() {}

func (x *KeyedUri) ProtoReflect() protoreflect.Message {
	mi := &file_service_local_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyedUri.ProtoReflect.Descriptor instead.
func (*KeyedUri) Descriptor() ([]byte, []int) {
	return file_service_local_proto_rawDescGZIP(), []int{2}
}

func (x *KeyedUri) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *KeyedUri) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type RegionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegionReq) Reset() {
	*x = RegionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_local_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionReq) ProtoMessage() {}

func (x *RegionReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_local_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionReq.ProtoReflect.Descriptor instead.
func (*RegionReq) Descriptor() ([]byte, []int) {
	return file_service_local_proto_rawDescGZIP(), []int{3}
}

type RegionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region int32 `protobuf:"varint,1,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *RegionRes) Reset() {
	*x = RegionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_local_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionRes) ProtoMessage() {}

func (x *RegionRes) ProtoReflect() protoreflect.Message {
	mi := &file_service_local_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionRes.ProtoReflect.Descriptor instead.
func (*RegionRes) Descriptor() ([]byte, []int) {
	return file_service_local_proto_rawDescGZIP(), []int{4}
}

func (x *RegionRes) GetRegion() int32 {
	if x != nil {
		return x.Region
	}
	return 0
}

type RouterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RouterReq) Reset() {
	*x = RouterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_local_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouterReq) ProtoMessage() {}

func (x *RouterReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_local_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouterReq.ProtoReflect.Descriptor instead.
func (*RouterReq) Descriptor() ([]byte, []int) {
	return file_service_local_proto_rawDescGZIP(), []int{5}
}

type RouterRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri       string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Connected bool   `protobuf:"varint,2,opt,name=connected,proto3" json:"connected,omitempty"`
}

func (x *RouterRes) Reset() {
	*x = RouterRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_local_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouterRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouterRes) ProtoMessage() {}

func (x *RouterRes) ProtoReflect() protoreflect.Message {
	mi := &file_service_local_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouterRes.ProtoReflect.Descriptor instead.
func (*RouterRes) Descriptor() ([]byte, []int) {
	return file_service_local_proto_rawDescGZIP(), []int{6}
}

func (x *RouterRes) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *RouterRes) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

type AddGatewayReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner       []byte                    `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Payer       []byte                    `protobuf:"bytes,2,opt,name=payer,proto3" json:"payer,omitempty"`
	StakingMode helium.GatewayStakingMode `protobuf:"varint,3,opt,name=staking_mode,json=stakingMode,proto3,enum=helium.GatewayStakingMode" json:"staking_mode,omitempty"`
}

func (x *AddGatewayReq) Reset() {
	*x = AddGatewayReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_local_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGatewayReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGatewayReq) ProtoMessage() {}

func (x *AddGatewayReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_local_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGatewayReq.ProtoReflect.Descriptor instead.
func (*AddGatewayReq) Descriptor() ([]byte, []int) {
	return file_service_local_proto_rawDescGZIP(), []int{7}
}

func (x *AddGatewayReq) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *AddGatewayReq) GetPayer() []byte {
	if x != nil {
		return x.Payer
	}
	return nil
}

func (x *AddGatewayReq) GetStakingMode() helium.GatewayStakingMode {
	if x != nil {
		return x.StakingMode
	}
	return helium.GatewayStakingMode(0)
}

type AddGatewayRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddGatewayTxn []byte `protobuf:"bytes,1,opt,name=add_gateway_txn,json=addGatewayTxn,proto3" json:"add_gateway_txn,omitempty"`
}

func (x *AddGatewayRes) Reset() {
	*x = AddGatewayRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_local_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGatewayRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGatewayRes) ProtoMessage() {}

func (x *AddGatewayRes) ProtoReflect() protoreflect.Message {
	mi := &file_service_local_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGatewayRes.ProtoReflect.Descriptor instead.
func (*AddGatewayRes) Descriptor() ([]byte, []int) {
	return file_service_local_proto_rawDescGZIP(), []int{8}
}

func (x *AddGatewayRes) GetAddGatewayTxn() []byte {
	if x != nil {
		return x.AddGatewayTxn
	}
	return nil
}

var File_service_local_proto protoreflect.FileDescriptor

var file_service_local_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x1a, 0x1a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x55, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x6f, 0x6e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x11, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79,
	0x5f, 0x72, 0x65, 0x71, 0x22, 0x37, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x65, 0x64, 0x5f, 0x75, 0x72,
	0x69, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x22, 0x0c, 0x0a,
	0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x22, 0x24, 0x0a, 0x0a, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x22, 0x0c, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x22,
	0x3c, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x7e, 0x0a,
	0x0f, 0x61, 0x64, 0x64, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x52, 0x0b, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x39, 0x0a,
	0x0f, 0x61, 0x64, 0x64, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f,
	0x74, 0x78, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x54, 0x78, 0x6e, 0x32, 0x8c, 0x02, 0x0a, 0x03, 0x61, 0x70, 0x69,
	0x12, 0x3c, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x2e, 0x68, 0x65, 0x6c,
	0x69, 0x75, 0x6d, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79,
	0x5f, 0x72, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x2e, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x3c,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75,
	0x6d, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x71, 0x1a, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x06,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71,
	0x1a, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x0b, 0x61, 0x64,
	0x64, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x1d, 0x2e, 0x68, 0x65, 0x6c, 0x69,
	0x75, 0x6d, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x61, 0x64, 0x64, 0x5f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75,
	0x6d, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x61, 0x64, 0x64, 0x5f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_local_proto_rawDescOnce sync.Once
	file_service_local_proto_rawDescData = file_service_local_proto_rawDesc
)

func file_service_local_proto_rawDescGZIP() []byte {
	file_service_local_proto_rawDescOnce.Do(func() {
		file_service_local_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_local_proto_rawDescData)
	})
	return file_service_local_proto_rawDescData
}

var file_service_local_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_local_proto_goTypes = []interface{}{
	(*PubkeyRes)(nil),              // 0: helium.local.pubkey_res
	(*PubkeyReq)(nil),              // 1: helium.local.pubkey_req
	(*KeyedUri)(nil),               // 2: helium.local.keyed_uri
	(*RegionReq)(nil),              // 3: helium.local.region_req
	(*RegionRes)(nil),              // 4: helium.local.region_res
	(*RouterReq)(nil),              // 5: helium.local.router_req
	(*RouterRes)(nil),              // 6: helium.local.router_res
	(*AddGatewayReq)(nil),          // 7: helium.local.add_gateway_req
	(*AddGatewayRes)(nil),          // 8: helium.local.add_gateway_res
	(helium.GatewayStakingMode)(0), // 9: helium.gateway_staking_mode
}
var file_service_local_proto_depIdxs = []int32{
	9, // 0: helium.local.add_gateway_req.staking_mode:type_name -> helium.gateway_staking_mode
	1, // 1: helium.local.api.pubkey:input_type -> helium.local.pubkey_req
	3, // 2: helium.local.api.region:input_type -> helium.local.region_req
	5, // 3: helium.local.api.router:input_type -> helium.local.router_req
	7, // 4: helium.local.api.add_gateway:input_type -> helium.local.add_gateway_req
	0, // 5: helium.local.api.pubkey:output_type -> helium.local.pubkey_res
	4, // 6: helium.local.api.region:output_type -> helium.local.region_res
	6, // 7: helium.local.api.router:output_type -> helium.local.router_res
	8, // 8: helium.local.api.add_gateway:output_type -> helium.local.add_gateway_res
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_local_proto_init() }
func file_service_local_proto_init() {
	if File_service_local_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_local_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubkeyRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_local_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubkeyReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_local_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyedUri); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_local_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_local_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_local_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouterReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_local_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouterRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_local_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGatewayReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_local_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGatewayRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_local_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_local_proto_goTypes,
		DependencyIndexes: file_service_local_proto_depIdxs,
		MessageInfos:      file_service_local_proto_msgTypes,
	}.Build()
	File_service_local_proto = out.File
	file_service_local_proto_rawDesc = nil
	file_service_local_proto_goTypes = nil
	file_service_local_proto_depIdxs = nil
}
