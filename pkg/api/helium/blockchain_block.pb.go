// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: blockchain_block.proto

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

type BlockchainBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Block:
	//	*BlockchainBlock_V1
	Block isBlockchainBlock_Block `protobuf_oneof:"block"`
}

func (x *BlockchainBlock) Reset() {
	*x = BlockchainBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainBlock) ProtoMessage() {}

func (x *BlockchainBlock) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainBlock.ProtoReflect.Descriptor instead.
func (*BlockchainBlock) Descriptor() ([]byte, []int) {
	return file_blockchain_block_proto_rawDescGZIP(), []int{0}
}

func (m *BlockchainBlock) GetBlock() isBlockchainBlock_Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (x *BlockchainBlock) GetV1() *BlockchainBlockV1 {
	if x, ok := x.GetBlock().(*BlockchainBlock_V1); ok {
		return x.V1
	}
	return nil
}

type isBlockchainBlock_Block interface {
	isBlockchainBlock_Block()
}

type BlockchainBlock_V1 struct {
	V1 *BlockchainBlockV1 `protobuf:"bytes,1,opt,name=v1,proto3,oneof"`
}

func (*BlockchainBlock_V1) isBlockchainBlock_Block() {}

var File_blockchain_block_proto protoreflect.FileDescriptor

var file_blockchain_block_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x26, 0x0a, 0x02, 0x76, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x76, 0x31, 0x48, 0x00, 0x52, 0x02, 0x76, 0x31, 0x42,
	0x07, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_block_proto_rawDescOnce sync.Once
	file_blockchain_block_proto_rawDescData = file_blockchain_block_proto_rawDesc
)

func file_blockchain_block_proto_rawDescGZIP() []byte {
	file_blockchain_block_proto_rawDescOnce.Do(func() {
		file_blockchain_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_block_proto_rawDescData)
	})
	return file_blockchain_block_proto_rawDescData
}

var file_blockchain_block_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_blockchain_block_proto_goTypes = []interface{}{
	(*BlockchainBlock)(nil),   // 0: blockchain_block
	(*BlockchainBlockV1)(nil), // 1: blockchain_block_v1
}
var file_blockchain_block_proto_depIdxs = []int32{
	1, // 0: blockchain_block.v1:type_name -> blockchain_block_v1
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blockchain_block_proto_init() }
func file_blockchain_block_proto_init() {
	if File_blockchain_block_proto != nil {
		return
	}
	file_blockchain_block_v1_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_blockchain_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainBlock); i {
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
	file_blockchain_block_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BlockchainBlock_V1)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blockchain_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_block_proto_goTypes,
		DependencyIndexes: file_blockchain_block_proto_depIdxs,
		MessageInfos:      file_blockchain_block_proto_msgTypes,
	}.Build()
	File_blockchain_block_proto = out.File
	file_blockchain_block_proto_rawDesc = nil
	file_blockchain_block_proto_goTypes = nil
	file_blockchain_block_proto_depIdxs = nil
}
