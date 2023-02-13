// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: blockchain_ledger_entry_v2.proto

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

type BlockchainLedgerEntryV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Top level nonce since entry-specific nonces don't really matter
	Nonce         uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	HntBalance    uint64 `protobuf:"varint,2,opt,name=hnt_balance,json=hntBalance,proto3" json:"hnt_balance,omitempty"`
	HstBalance    uint64 `protobuf:"varint,3,opt,name=hst_balance,json=hstBalance,proto3" json:"hst_balance,omitempty"`
	MobileBalance uint64 `protobuf:"varint,4,opt,name=mobile_balance,json=mobileBalance,proto3" json:"mobile_balance,omitempty"`
	IotBalance    uint64 `protobuf:"varint,5,opt,name=iot_balance,json=iotBalance,proto3" json:"iot_balance,omitempty"`
}

func (x *BlockchainLedgerEntryV2) Reset() {
	*x = BlockchainLedgerEntryV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_ledger_entry_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainLedgerEntryV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainLedgerEntryV2) ProtoMessage() {}

func (x *BlockchainLedgerEntryV2) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_ledger_entry_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainLedgerEntryV2.ProtoReflect.Descriptor instead.
func (*BlockchainLedgerEntryV2) Descriptor() ([]byte, []int) {
	return file_blockchain_ledger_entry_v2_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainLedgerEntryV2) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *BlockchainLedgerEntryV2) GetHntBalance() uint64 {
	if x != nil {
		return x.HntBalance
	}
	return 0
}

func (x *BlockchainLedgerEntryV2) GetHstBalance() uint64 {
	if x != nil {
		return x.HstBalance
	}
	return 0
}

func (x *BlockchainLedgerEntryV2) GetMobileBalance() uint64 {
	if x != nil {
		return x.MobileBalance
	}
	return 0
}

func (x *BlockchainLedgerEntryV2) GetIotBalance() uint64 {
	if x != nil {
		return x.IotBalance
	}
	return 0
}

var File_blockchain_ledger_entry_v2_proto protoreflect.FileDescriptor

var file_blockchain_ledger_entry_v2_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x22, 0xbc, 0x01, 0x0a, 0x1a, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x76, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x68, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x68, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x68, 0x73, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6f, 0x74, 0x5f,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x69,
	0x6f, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_blockchain_ledger_entry_v2_proto_rawDescOnce sync.Once
	file_blockchain_ledger_entry_v2_proto_rawDescData = file_blockchain_ledger_entry_v2_proto_rawDesc
)

func file_blockchain_ledger_entry_v2_proto_rawDescGZIP() []byte {
	file_blockchain_ledger_entry_v2_proto_rawDescOnce.Do(func() {
		file_blockchain_ledger_entry_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_ledger_entry_v2_proto_rawDescData)
	})
	return file_blockchain_ledger_entry_v2_proto_rawDescData
}

var file_blockchain_ledger_entry_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_blockchain_ledger_entry_v2_proto_goTypes = []interface{}{
	(*BlockchainLedgerEntryV2)(nil), // 0: helium.blockchain_ledger_entry_v2
}
var file_blockchain_ledger_entry_v2_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blockchain_ledger_entry_v2_proto_init() }
func file_blockchain_ledger_entry_v2_proto_init() {
	if File_blockchain_ledger_entry_v2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_ledger_entry_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainLedgerEntryV2); i {
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
			RawDescriptor: file_blockchain_ledger_entry_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_ledger_entry_v2_proto_goTypes,
		DependencyIndexes: file_blockchain_ledger_entry_v2_proto_depIdxs,
		MessageInfos:      file_blockchain_ledger_entry_v2_proto_msgTypes,
	}.Build()
	File_blockchain_ledger_entry_v2_proto = out.File
	file_blockchain_ledger_entry_v2_proto_rawDesc = nil
	file_blockchain_ledger_entry_v2_proto_goTypes = nil
	file_blockchain_ledger_entry_v2_proto_depIdxs = nil
}
