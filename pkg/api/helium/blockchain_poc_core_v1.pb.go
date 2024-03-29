// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: blockchain_poc_core_v1.proto

package helium

import (
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

type Origin int32

const (
	Origin_origin_p2p   Origin = 0 //TODO: needed fixing; orig: p2p = 0;
	Origin_origin_radio Origin = 1 //TODO: needed fixing; orig: radio = 1
)

// Enum value maps for Origin.
var (
	Origin_name = map[int32]string{
		0: "origin_p2p",
		1: "origin_radio",
	}
	Origin_value = map[string]int32{
		"origin_p2p":   0,
		"origin_radio": 1,
	}
)

func (x Origin) Enum() *Origin {
	p := new(Origin)
	*p = x
	return p
}

func (x Origin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Origin) Descriptor() protoreflect.EnumDescriptor {
	return file_blockchain_poc_core_v1_proto_enumTypes[0].Descriptor()
}

func (Origin) Type() protoreflect.EnumType {
	return &file_blockchain_poc_core_v1_proto_enumTypes[0]
}

func (x Origin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Origin.Descriptor instead.
func (Origin) EnumDescriptor() ([]byte, []int) {
	return file_blockchain_poc_core_v1_proto_rawDescGZIP(), []int{0}
}

type BlockchainPocReceiptV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gateway   []byte  `protobuf:"bytes,1,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Timestamp uint64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signal    int32   `protobuf:"zigzag32,3,opt,name=signal,proto3" json:"signal,omitempty"`
	Data      []byte  `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Origin    Origin  `protobuf:"varint,5,opt,name=origin,proto3,enum=helium.Origin" json:"origin,omitempty"`
	Signature []byte  `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	Snr       float32 `protobuf:"fixed32,7,opt,name=snr,proto3" json:"snr,omitempty"`
	Frequency float32 `protobuf:"fixed32,8,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Channel   int32   `protobuf:"varint,9,opt,name=channel,proto3" json:"channel,omitempty"`
	Datarate  string  `protobuf:"bytes,10,opt,name=datarate,proto3" json:"datarate,omitempty"`
	AddrHash  []byte  `protobuf:"bytes,11,opt,name=addr_hash,json=addrHash,proto3" json:"addr_hash,omitempty"`
	// Transmit power at which this packet was transmitted
	// It is x10, for example: 270 = 27db, 36 = 3.6db etc
	TxPower int32 `protobuf:"varint,12,opt,name=tx_power,json=txPower,proto3" json:"tx_power,omitempty"`
	// fixed point reward shares added by off-chain poc verifier,
	// propose 2 decimal places
	RewardShares uint32 `protobuf:"varint,13,opt,name=reward_shares,json=rewardShares,proto3" json:"reward_shares,omitempty"`
}

func (x *BlockchainPocReceiptV1) Reset() {
	*x = BlockchainPocReceiptV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_poc_core_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainPocReceiptV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainPocReceiptV1) ProtoMessage() {}

func (x *BlockchainPocReceiptV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_poc_core_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainPocReceiptV1.ProtoReflect.Descriptor instead.
func (*BlockchainPocReceiptV1) Descriptor() ([]byte, []int) {
	return file_blockchain_poc_core_v1_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainPocReceiptV1) GetGateway() []byte {
	if x != nil {
		return x.Gateway
	}
	return nil
}

func (x *BlockchainPocReceiptV1) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BlockchainPocReceiptV1) GetSignal() int32 {
	if x != nil {
		return x.Signal
	}
	return 0
}

func (x *BlockchainPocReceiptV1) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *BlockchainPocReceiptV1) GetOrigin() Origin {
	if x != nil {
		return x.Origin
	}
	return Origin_origin_p2p
}

func (x *BlockchainPocReceiptV1) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *BlockchainPocReceiptV1) GetSnr() float32 {
	if x != nil {
		return x.Snr
	}
	return 0
}

func (x *BlockchainPocReceiptV1) GetFrequency() float32 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *BlockchainPocReceiptV1) GetChannel() int32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

func (x *BlockchainPocReceiptV1) GetDatarate() string {
	if x != nil {
		return x.Datarate
	}
	return ""
}

func (x *BlockchainPocReceiptV1) GetAddrHash() []byte {
	if x != nil {
		return x.AddrHash
	}
	return nil
}

func (x *BlockchainPocReceiptV1) GetTxPower() int32 {
	if x != nil {
		return x.TxPower
	}
	return 0
}

func (x *BlockchainPocReceiptV1) GetRewardShares() uint32 {
	if x != nil {
		return x.RewardShares
	}
	return 0
}

type BlockchainPocWitnessV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gateway    []byte  `protobuf:"bytes,1,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Timestamp  uint64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signal     int32   `protobuf:"zigzag32,3,opt,name=signal,proto3" json:"signal,omitempty"`
	PacketHash []byte  `protobuf:"bytes,4,opt,name=packet_hash,json=packetHash,proto3" json:"packet_hash,omitempty"`
	Signature  []byte  `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Snr        float32 `protobuf:"fixed32,6,opt,name=snr,proto3" json:"snr,omitempty"`
	Frequency  float32 `protobuf:"fixed32,7,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Channel    int32   `protobuf:"varint,8,opt,name=channel,proto3" json:"channel,omitempty"`
	Datarate   string  `protobuf:"bytes,9,opt,name=datarate,proto3" json:"datarate,omitempty"`
	// fixed point reward shares added by off-chain poc verifier,
	// propose 2 decimal places
	RewardShares uint32 `protobuf:"varint,10,opt,name=reward_shares,json=rewardShares,proto3" json:"reward_shares,omitempty"`
}

func (x *BlockchainPocWitnessV1) Reset() {
	*x = BlockchainPocWitnessV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_poc_core_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainPocWitnessV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainPocWitnessV1) ProtoMessage() {}

func (x *BlockchainPocWitnessV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_poc_core_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainPocWitnessV1.ProtoReflect.Descriptor instead.
func (*BlockchainPocWitnessV1) Descriptor() ([]byte, []int) {
	return file_blockchain_poc_core_v1_proto_rawDescGZIP(), []int{1}
}

func (x *BlockchainPocWitnessV1) GetGateway() []byte {
	if x != nil {
		return x.Gateway
	}
	return nil
}

func (x *BlockchainPocWitnessV1) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BlockchainPocWitnessV1) GetSignal() int32 {
	if x != nil {
		return x.Signal
	}
	return 0
}

func (x *BlockchainPocWitnessV1) GetPacketHash() []byte {
	if x != nil {
		return x.PacketHash
	}
	return nil
}

func (x *BlockchainPocWitnessV1) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *BlockchainPocWitnessV1) GetSnr() float32 {
	if x != nil {
		return x.Snr
	}
	return 0
}

func (x *BlockchainPocWitnessV1) GetFrequency() float32 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *BlockchainPocWitnessV1) GetChannel() int32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

func (x *BlockchainPocWitnessV1) GetDatarate() string {
	if x != nil {
		return x.Datarate
	}
	return ""
}

func (x *BlockchainPocWitnessV1) GetRewardShares() uint32 {
	if x != nil {
		return x.RewardShares
	}
	return 0
}

type BlockchainPocResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*BlockchainPocResponseV1_Receipt
	//	*BlockchainPocResponseV1_Witness
	Payload isBlockchainPocResponseV1_Payload `protobuf_oneof:"payload"`
}

func (x *BlockchainPocResponseV1) Reset() {
	*x = BlockchainPocResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_poc_core_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainPocResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainPocResponseV1) ProtoMessage() {}

func (x *BlockchainPocResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_poc_core_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainPocResponseV1.ProtoReflect.Descriptor instead.
func (*BlockchainPocResponseV1) Descriptor() ([]byte, []int) {
	return file_blockchain_poc_core_v1_proto_rawDescGZIP(), []int{2}
}

func (m *BlockchainPocResponseV1) GetPayload() isBlockchainPocResponseV1_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *BlockchainPocResponseV1) GetReceipt() *BlockchainPocReceiptV1 {
	if x, ok := x.GetPayload().(*BlockchainPocResponseV1_Receipt); ok {
		return x.Receipt
	}
	return nil
}

func (x *BlockchainPocResponseV1) GetWitness() *BlockchainPocWitnessV1 {
	if x, ok := x.GetPayload().(*BlockchainPocResponseV1_Witness); ok {
		return x.Witness
	}
	return nil
}

type isBlockchainPocResponseV1_Payload interface {
	isBlockchainPocResponseV1_Payload()
}

type BlockchainPocResponseV1_Receipt struct {
	Receipt *BlockchainPocReceiptV1 `protobuf:"bytes,1,opt,name=receipt,proto3,oneof"`
}

type BlockchainPocResponseV1_Witness struct {
	Witness *BlockchainPocWitnessV1 `protobuf:"bytes,2,opt,name=witness,proto3,oneof"`
}

func (*BlockchainPocResponseV1_Receipt) isBlockchainPocResponseV1_Payload() {}

func (*BlockchainPocResponseV1_Witness) isBlockchainPocResponseV1_Payload() {}

type BlockchainPocPathElementV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Challengee []byte                    `protobuf:"bytes,1,opt,name=challengee,proto3" json:"challengee,omitempty"`
	Receipt    *BlockchainPocReceiptV1   `protobuf:"bytes,2,opt,name=receipt,proto3" json:"receipt,omitempty"`
	Witnesses  []*BlockchainPocWitnessV1 `protobuf:"bytes,3,rep,name=witnesses,proto3" json:"witnesses,omitempty"`
}

func (x *BlockchainPocPathElementV1) Reset() {
	*x = BlockchainPocPathElementV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_poc_core_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainPocPathElementV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainPocPathElementV1) ProtoMessage() {}

func (x *BlockchainPocPathElementV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_poc_core_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainPocPathElementV1.ProtoReflect.Descriptor instead.
func (*BlockchainPocPathElementV1) Descriptor() ([]byte, []int) {
	return file_blockchain_poc_core_v1_proto_rawDescGZIP(), []int{3}
}

func (x *BlockchainPocPathElementV1) GetChallengee() []byte {
	if x != nil {
		return x.Challengee
	}
	return nil
}

func (x *BlockchainPocPathElementV1) GetReceipt() *BlockchainPocReceiptV1 {
	if x != nil {
		return x.Receipt
	}
	return nil
}

func (x *BlockchainPocPathElementV1) GetWitnesses() []*BlockchainPocWitnessV1 {
	if x != nil {
		return x.Witnesses
	}
	return nil
}

var File_blockchain_poc_core_v1_proto protoreflect.FileDescriptor

var file_blockchain_poc_core_v1_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x63,
	0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x22, 0x88, 0x03, 0x0a, 0x19, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x63, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x5f, 0x76, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x6e, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x73, 0x6e, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x72, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x78, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x73, 0x22, 0xb5, 0x02, 0x0a, 0x19, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x70, 0x6f, 0x63, 0x5f, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x76, 0x31, 0x12,
	0x18, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x6e, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x73, 0x6e, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x72, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x1a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x63, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x76, 0x31, 0x12, 0x3d, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x69,
	0x75, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x6f,
	0x63, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x76, 0x31, 0x48, 0x00, 0x52, 0x07,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x3d, 0x0a, 0x07, 0x77, 0x69, 0x74, 0x6e, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75,
	0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x63,
	0x5f, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x76, 0x31, 0x48, 0x00, 0x52, 0x07, 0x77,
	0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0xbe, 0x01, 0x0a, 0x1e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x70, 0x6f, 0x63, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x76, 0x31, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x63, 0x5f, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x76, 0x31, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x12, 0x3f, 0x0a, 0x09, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x63, 0x5f, 0x77, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x5f, 0x76, 0x31, 0x52, 0x09, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x2a, 0x2a, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x0a,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x32, 0x70, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x10, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_poc_core_v1_proto_rawDescOnce sync.Once
	file_blockchain_poc_core_v1_proto_rawDescData = file_blockchain_poc_core_v1_proto_rawDesc
)

func file_blockchain_poc_core_v1_proto_rawDescGZIP() []byte {
	file_blockchain_poc_core_v1_proto_rawDescOnce.Do(func() {
		file_blockchain_poc_core_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_poc_core_v1_proto_rawDescData)
	})
	return file_blockchain_poc_core_v1_proto_rawDescData
}

var file_blockchain_poc_core_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_blockchain_poc_core_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_blockchain_poc_core_v1_proto_goTypes = []interface{}{
	(Origin)(0),                        // 0: helium.origin
	(*BlockchainPocReceiptV1)(nil),     // 1: helium.blockchain_poc_receipt_v1
	(*BlockchainPocWitnessV1)(nil),     // 2: helium.blockchain_poc_witness_v1
	(*BlockchainPocResponseV1)(nil),    // 3: helium.blockchain_poc_response_v1
	(*BlockchainPocPathElementV1)(nil), // 4: helium.blockchain_poc_path_element_v1
}
var file_blockchain_poc_core_v1_proto_depIdxs = []int32{
	0, // 0: helium.blockchain_poc_receipt_v1.origin:type_name -> helium.origin
	1, // 1: helium.blockchain_poc_response_v1.receipt:type_name -> helium.blockchain_poc_receipt_v1
	2, // 2: helium.blockchain_poc_response_v1.witness:type_name -> helium.blockchain_poc_witness_v1
	1, // 3: helium.blockchain_poc_path_element_v1.receipt:type_name -> helium.blockchain_poc_receipt_v1
	2, // 4: helium.blockchain_poc_path_element_v1.witnesses:type_name -> helium.blockchain_poc_witness_v1
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_blockchain_poc_core_v1_proto_init() }
func file_blockchain_poc_core_v1_proto_init() {
	if File_blockchain_poc_core_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_poc_core_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainPocReceiptV1); i {
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
		file_blockchain_poc_core_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainPocWitnessV1); i {
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
		file_blockchain_poc_core_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainPocResponseV1); i {
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
		file_blockchain_poc_core_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainPocPathElementV1); i {
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
	file_blockchain_poc_core_v1_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*BlockchainPocResponseV1_Receipt)(nil),
		(*BlockchainPocResponseV1_Witness)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blockchain_poc_core_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_poc_core_v1_proto_goTypes,
		DependencyIndexes: file_blockchain_poc_core_v1_proto_depIdxs,
		EnumInfos:         file_blockchain_poc_core_v1_proto_enumTypes,
		MessageInfos:      file_blockchain_poc_core_v1_proto_msgTypes,
	}.Build()
	File_blockchain_poc_core_v1_proto = out.File
	file_blockchain_poc_core_v1_proto_rawDesc = nil
	file_blockchain_poc_core_v1_proto_goTypes = nil
	file_blockchain_poc_core_v1_proto_depIdxs = nil
}
