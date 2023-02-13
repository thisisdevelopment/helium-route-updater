// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: blockchain_txn_stake_validator_v1.proto

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

type BlockchainTxnStakeValidatorV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address        []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Owner          []byte `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Stake          uint64 `protobuf:"varint,3,opt,name=stake,proto3" json:"stake,omitempty"`
	OwnerSignature []byte `protobuf:"bytes,4,opt,name=owner_signature,json=ownerSignature,proto3" json:"owner_signature,omitempty"`
	Fee            uint64 `protobuf:"varint,5,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (x *BlockchainTxnStakeValidatorV1) Reset() {
	*x = BlockchainTxnStakeValidatorV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_txn_stake_validator_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainTxnStakeValidatorV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainTxnStakeValidatorV1) ProtoMessage() {}

func (x *BlockchainTxnStakeValidatorV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_txn_stake_validator_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainTxnStakeValidatorV1.ProtoReflect.Descriptor instead.
func (*BlockchainTxnStakeValidatorV1) Descriptor() ([]byte, []int) {
	return file_blockchain_txn_stake_validator_v1_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainTxnStakeValidatorV1) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *BlockchainTxnStakeValidatorV1) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *BlockchainTxnStakeValidatorV1) GetStake() uint64 {
	if x != nil {
		return x.Stake
	}
	return 0
}

func (x *BlockchainTxnStakeValidatorV1) GetOwnerSignature() []byte {
	if x != nil {
		return x.OwnerSignature
	}
	return nil
}

func (x *BlockchainTxnStakeValidatorV1) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

var File_blockchain_txn_stake_validator_v1_proto protoreflect.FileDescriptor

var file_blockchain_txn_stake_validator_v1_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x78, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x65, 0x6c, 0x69, 0x75,
	0x6d, 0x22, 0xa4, 0x01, 0x0a, 0x21, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x76, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x6b, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x65, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_txn_stake_validator_v1_proto_rawDescOnce sync.Once
	file_blockchain_txn_stake_validator_v1_proto_rawDescData = file_blockchain_txn_stake_validator_v1_proto_rawDesc
)

func file_blockchain_txn_stake_validator_v1_proto_rawDescGZIP() []byte {
	file_blockchain_txn_stake_validator_v1_proto_rawDescOnce.Do(func() {
		file_blockchain_txn_stake_validator_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_txn_stake_validator_v1_proto_rawDescData)
	})
	return file_blockchain_txn_stake_validator_v1_proto_rawDescData
}

var file_blockchain_txn_stake_validator_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_blockchain_txn_stake_validator_v1_proto_goTypes = []interface{}{
	(*BlockchainTxnStakeValidatorV1)(nil), // 0: helium.blockchain_txn_stake_validator_v1
}
var file_blockchain_txn_stake_validator_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blockchain_txn_stake_validator_v1_proto_init() }
func file_blockchain_txn_stake_validator_v1_proto_init() {
	if File_blockchain_txn_stake_validator_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_txn_stake_validator_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainTxnStakeValidatorV1); i {
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
			RawDescriptor: file_blockchain_txn_stake_validator_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_txn_stake_validator_v1_proto_goTypes,
		DependencyIndexes: file_blockchain_txn_stake_validator_v1_proto_depIdxs,
		MessageInfos:      file_blockchain_txn_stake_validator_v1_proto_msgTypes,
	}.Build()
	File_blockchain_txn_stake_validator_v1_proto = out.File
	file_blockchain_txn_stake_validator_v1_proto_rawDesc = nil
	file_blockchain_txn_stake_validator_v1_proto_goTypes = nil
	file_blockchain_txn_stake_validator_v1_proto_depIdxs = nil
}
