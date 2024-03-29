// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: blockchain_snapshot_handler.proto

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

type BlockchainSnapshotReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Hash   []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *BlockchainSnapshotReq) Reset() {
	*x = BlockchainSnapshotReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_snapshot_handler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainSnapshotReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainSnapshotReq) ProtoMessage() {}

func (x *BlockchainSnapshotReq) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_snapshot_handler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainSnapshotReq.ProtoReflect.Descriptor instead.
func (*BlockchainSnapshotReq) Descriptor() ([]byte, []int) {
	return file_blockchain_snapshot_handler_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainSnapshotReq) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockchainSnapshotReq) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type BlockchainSnapshotResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextHeight int64  `protobuf:"zigzag64,1,opt,name=next_height,json=nextHeight,proto3" json:"next_height,omitempty"`
	Snapshot   []byte `protobuf:"bytes,2,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
}

func (x *BlockchainSnapshotResp) Reset() {
	*x = BlockchainSnapshotResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_snapshot_handler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainSnapshotResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainSnapshotResp) ProtoMessage() {}

func (x *BlockchainSnapshotResp) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_snapshot_handler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainSnapshotResp.ProtoReflect.Descriptor instead.
func (*BlockchainSnapshotResp) Descriptor() ([]byte, []int) {
	return file_blockchain_snapshot_handler_proto_rawDescGZIP(), []int{1}
}

func (x *BlockchainSnapshotResp) GetNextHeight() int64 {
	if x != nil {
		return x.NextHeight
	}
	return 0
}

func (x *BlockchainSnapshotResp) GetSnapshot() []byte {
	if x != nil {
		return x.Snapshot
	}
	return nil
}

var File_blockchain_snapshot_handler_proto protoreflect.FileDescriptor

var file_blockchain_snapshot_handler_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x17, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x57, 0x0a, 0x18, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0a, 0x6e, 0x65, 0x78,
	0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_snapshot_handler_proto_rawDescOnce sync.Once
	file_blockchain_snapshot_handler_proto_rawDescData = file_blockchain_snapshot_handler_proto_rawDesc
)

func file_blockchain_snapshot_handler_proto_rawDescGZIP() []byte {
	file_blockchain_snapshot_handler_proto_rawDescOnce.Do(func() {
		file_blockchain_snapshot_handler_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_snapshot_handler_proto_rawDescData)
	})
	return file_blockchain_snapshot_handler_proto_rawDescData
}

var file_blockchain_snapshot_handler_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_blockchain_snapshot_handler_proto_goTypes = []interface{}{
	(*BlockchainSnapshotReq)(nil),  // 0: blockchain_snapshot_req
	(*BlockchainSnapshotResp)(nil), // 1: blockchain_snapshot_resp
}
var file_blockchain_snapshot_handler_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blockchain_snapshot_handler_proto_init() }
func file_blockchain_snapshot_handler_proto_init() {
	if File_blockchain_snapshot_handler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_snapshot_handler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainSnapshotReq); i {
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
		file_blockchain_snapshot_handler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainSnapshotResp); i {
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
			RawDescriptor: file_blockchain_snapshot_handler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_snapshot_handler_proto_goTypes,
		DependencyIndexes: file_blockchain_snapshot_handler_proto_depIdxs,
		MessageInfos:      file_blockchain_snapshot_handler_proto_msgTypes,
	}.Build()
	File_blockchain_snapshot_handler_proto = out.File
	file_blockchain_snapshot_handler_proto_rawDesc = nil
	file_blockchain_snapshot_handler_proto_goTypes = nil
	file_blockchain_snapshot_handler_proto_depIdxs = nil
}
