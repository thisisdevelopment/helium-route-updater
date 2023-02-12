// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: service/transaction.proto

package transaction

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

type TxnStatus int32

const (
	TxnStatus_pending   TxnStatus = 0
	TxnStatus_not_found TxnStatus = 1
)

// Enum value maps for TxnStatus.
var (
	TxnStatus_name = map[int32]string{
		0: "pending",
		1: "not_found",
	}
	TxnStatus_value = map[string]int32{
		"pending":   0,
		"not_found": 1,
	}
)

func (x TxnStatus) Enum() *TxnStatus {
	p := new(TxnStatus)
	*p = x
	return p
}

func (x TxnStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TxnStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_service_transaction_proto_enumTypes[0].Descriptor()
}

func (TxnStatus) Type() protoreflect.EnumType {
	return &file_service_transaction_proto_enumTypes[0]
}

func (x TxnStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TxnStatus.Descriptor instead.
func (TxnStatus) EnumDescriptor() ([]byte, []int) {
	return file_service_transaction_proto_rawDescGZIP(), []int{0}
}

type Acceptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height   uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	QueuePos uint32 `protobuf:"varint,2,opt,name=queue_pos,json=queuePos,proto3" json:"queue_pos,omitempty"`
	QueueLen uint32 `protobuf:"varint,3,opt,name=queue_len,json=queueLen,proto3" json:"queue_len,omitempty"`
	PubKey   []byte `protobuf:"bytes,4,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (x *Acceptor) Reset() {
	*x = Acceptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Acceptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acceptor) ProtoMessage() {}

func (x *Acceptor) ProtoReflect() protoreflect.Message {
	mi := &file_service_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acceptor.ProtoReflect.Descriptor instead.
func (*Acceptor) Descriptor() ([]byte, []int) {
	return file_service_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Acceptor) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Acceptor) GetQueuePos() uint32 {
	if x != nil {
		return x.QueuePos
	}
	return 0
}

func (x *Acceptor) GetQueueLen() uint32 {
	if x != nil {
		return x.QueueLen
	}
	return 0
}

func (x *Acceptor) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

type Rejector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Reason []byte `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	PubKey []byte `protobuf:"bytes,3,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (x *Rejector) Reset() {
	*x = Rejector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rejector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rejector) ProtoMessage() {}

func (x *Rejector) ProtoReflect() protoreflect.Message {
	mi := &file_service_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rejector.ProtoReflect.Descriptor instead.
func (*Rejector) Descriptor() ([]byte, []int) {
	return file_service_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *Rejector) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Rejector) GetReason() []byte {
	if x != nil {
		return x.Reason
	}
	return nil
}

func (x *Rejector) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

type TxnSubmitReqV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txn *helium.BlockchainTxn `protobuf:"bytes,1,opt,name=txn,proto3" json:"txn,omitempty"`
	Key []byte                `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *TxnSubmitReqV1) Reset() {
	*x = TxnSubmitReqV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxnSubmitReqV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxnSubmitReqV1) ProtoMessage() {}

func (x *TxnSubmitReqV1) ProtoReflect() protoreflect.Message {
	mi := &file_service_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxnSubmitReqV1.ProtoReflect.Descriptor instead.
func (*TxnSubmitReqV1) Descriptor() ([]byte, []int) {
	return file_service_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *TxnSubmitReqV1) GetTxn() *helium.BlockchainTxn {
	if x != nil {
		return x.Txn
	}
	return nil
}

func (x *TxnSubmitReqV1) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type TxnSubmitRespV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        []byte                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	RecvHeight uint64                 `protobuf:"varint,2,opt,name=recv_height,json=recvHeight,proto3" json:"recv_height,omitempty"`
	Validator  *helium.RoutingAddress `protobuf:"bytes,3,opt,name=validator,proto3" json:"validator,omitempty"`
	Signature  []byte                 `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *TxnSubmitRespV1) Reset() {
	*x = TxnSubmitRespV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxnSubmitRespV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxnSubmitRespV1) ProtoMessage() {}

func (x *TxnSubmitRespV1) ProtoReflect() protoreflect.Message {
	mi := &file_service_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxnSubmitRespV1.ProtoReflect.Descriptor instead.
func (*TxnSubmitRespV1) Descriptor() ([]byte, []int) {
	return file_service_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *TxnSubmitRespV1) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *TxnSubmitRespV1) GetRecvHeight() uint64 {
	if x != nil {
		return x.RecvHeight
	}
	return 0
}

func (x *TxnSubmitRespV1) GetValidator() *helium.RoutingAddress {
	if x != nil {
		return x.Validator
	}
	return nil
}

func (x *TxnSubmitRespV1) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type TxnQueryReqV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *TxnQueryReqV1) Reset() {
	*x = TxnQueryReqV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxnQueryReqV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxnQueryReqV1) ProtoMessage() {}

func (x *TxnQueryReqV1) ProtoReflect() protoreflect.Message {
	mi := &file_service_transaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxnQueryReqV1.ProtoReflect.Descriptor instead.
func (*TxnQueryReqV1) Descriptor() ([]byte, []int) {
	return file_service_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *TxnQueryReqV1) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type TxnQueryRespV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     TxnStatus   `protobuf:"varint,1,opt,name=status,proto3,enum=helium.transaction.TxnStatus" json:"status,omitempty"`
	Key        []byte      `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Acceptors  []*Acceptor `protobuf:"bytes,3,rep,name=acceptors,proto3" json:"acceptors,omitempty"`
	Rejectors  []*Rejector `protobuf:"bytes,4,rep,name=rejectors,proto3" json:"rejectors,omitempty"`
	RecvHeight uint64      `protobuf:"varint,5,opt,name=recv_height,json=recvHeight,proto3" json:"recv_height,omitempty"`
	Height     uint64      `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	Signature  []byte      `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *TxnQueryRespV1) Reset() {
	*x = TxnQueryRespV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_transaction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxnQueryRespV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxnQueryRespV1) ProtoMessage() {}

func (x *TxnQueryRespV1) ProtoReflect() protoreflect.Message {
	mi := &file_service_transaction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxnQueryRespV1.ProtoReflect.Descriptor instead.
func (*TxnQueryRespV1) Descriptor() ([]byte, []int) {
	return file_service_transaction_proto_rawDescGZIP(), []int{5}
}

func (x *TxnQueryRespV1) GetStatus() TxnStatus {
	if x != nil {
		return x.Status
	}
	return TxnStatus_pending
}

func (x *TxnQueryRespV1) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *TxnQueryRespV1) GetAcceptors() []*Acceptor {
	if x != nil {
		return x.Acceptors
	}
	return nil
}

func (x *TxnQueryRespV1) GetRejectors() []*Rejector {
	if x != nil {
		return x.Rejectors
	}
	return nil
}

func (x *TxnQueryRespV1) GetRecvHeight() uint64 {
	if x != nil {
		return x.RecvHeight
	}
	return 0
}

func (x *TxnQueryRespV1) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *TxnQueryRespV1) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_service_transaction_proto protoreflect.FileDescriptor

var file_service_transaction_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x68, 0x65, 0x6c,
	0x69, 0x75, 0x6d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x78,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x50, 0x6f, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x4c, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x22, 0x53,
	0x0a, 0x08, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x75,
	0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62,
	0x4b, 0x65, 0x79, 0x22, 0x4f, 0x0a, 0x11, 0x74, 0x78, 0x6e, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x76, 0x31, 0x12, 0x28, 0x0a, 0x03, 0x74, 0x78, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x78, 0x6e, 0x52, 0x03, 0x74,
	0x78, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x9c, 0x01, 0x0a, 0x12, 0x74, 0x78, 0x6e, 0x5f, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x76, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x63, 0x76, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x76, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x35,
	0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x22, 0x24, 0x0a, 0x10, 0x74, 0x78, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x72, 0x65, 0x71, 0x5f, 0x76, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xac, 0x02, 0x0a, 0x11, 0x74, 0x78,
	0x6e, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x76, 0x31, 0x12,
	0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x78, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x09, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x68,
	0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75,
	0x6d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x76, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x76, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2a, 0x28, 0x0a, 0x0a, 0x74, 0x78, 0x6e, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x10, 0x01, 0x32, 0xbc, 0x01, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x25, 0x2e, 0x68,
	0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x74, 0x78, 0x6e, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x71,
	0x5f, 0x76, 0x31, 0x1a, 0x26, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x78, 0x6e, 0x5f, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x76, 0x31, 0x12, 0x54, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x24, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x78, 0x6e, 0x5f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x76, 0x31, 0x1a, 0x25, 0x2e, 0x68, 0x65, 0x6c,
	0x69, 0x75, 0x6d, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x74, 0x78, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_transaction_proto_rawDescOnce sync.Once
	file_service_transaction_proto_rawDescData = file_service_transaction_proto_rawDesc
)

func file_service_transaction_proto_rawDescGZIP() []byte {
	file_service_transaction_proto_rawDescOnce.Do(func() {
		file_service_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_transaction_proto_rawDescData)
	})
	return file_service_transaction_proto_rawDescData
}

var file_service_transaction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_transaction_proto_goTypes = []interface{}{
	(TxnStatus)(0),                // 0: helium.transaction.txn_status
	(*Acceptor)(nil),              // 1: helium.transaction.acceptor
	(*Rejector)(nil),              // 2: helium.transaction.rejector
	(*TxnSubmitReqV1)(nil),        // 3: helium.transaction.txn_submit_req_v1
	(*TxnSubmitRespV1)(nil),       // 4: helium.transaction.txn_submit_resp_v1
	(*TxnQueryReqV1)(nil),         // 5: helium.transaction.txn_query_req_v1
	(*TxnQueryRespV1)(nil),        // 6: helium.transaction.txn_query_resp_v1
	(*helium.BlockchainTxn)(nil),  // 7: helium.blockchain_txn
	(*helium.RoutingAddress)(nil), // 8: helium.routing_address
}
var file_service_transaction_proto_depIdxs = []int32{
	7, // 0: helium.transaction.txn_submit_req_v1.txn:type_name -> helium.blockchain_txn
	8, // 1: helium.transaction.txn_submit_resp_v1.validator:type_name -> helium.routing_address
	0, // 2: helium.transaction.txn_query_resp_v1.status:type_name -> helium.transaction.txn_status
	1, // 3: helium.transaction.txn_query_resp_v1.acceptors:type_name -> helium.transaction.acceptor
	2, // 4: helium.transaction.txn_query_resp_v1.rejectors:type_name -> helium.transaction.rejector
	3, // 5: helium.transaction.transaction.submit:input_type -> helium.transaction.txn_submit_req_v1
	5, // 6: helium.transaction.transaction.query:input_type -> helium.transaction.txn_query_req_v1
	4, // 7: helium.transaction.transaction.submit:output_type -> helium.transaction.txn_submit_resp_v1
	6, // 8: helium.transaction.transaction.query:output_type -> helium.transaction.txn_query_resp_v1
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_service_transaction_proto_init() }
func file_service_transaction_proto_init() {
	if File_service_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Acceptor); i {
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
		file_service_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rejector); i {
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
		file_service_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxnSubmitReqV1); i {
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
		file_service_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxnSubmitRespV1); i {
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
		file_service_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxnQueryReqV1); i {
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
		file_service_transaction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxnQueryRespV1); i {
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
			RawDescriptor: file_service_transaction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_transaction_proto_goTypes,
		DependencyIndexes: file_service_transaction_proto_depIdxs,
		EnumInfos:         file_service_transaction_proto_enumTypes,
		MessageInfos:      file_service_transaction_proto_msgTypes,
	}.Build()
	File_service_transaction_proto = out.File
	file_service_transaction_proto_rawDesc = nil
	file_service_transaction_proto_goTypes = nil
	file_service_transaction_proto_depIdxs = nil
}
