// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: blockchain_region_param_v1.proto

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

type RegionSpreading int32

const (
	RegionSpreading_SF_INVALID RegionSpreading = 0
	RegionSpreading_SF7        RegionSpreading = 1
	RegionSpreading_SF8        RegionSpreading = 2
	RegionSpreading_SF9        RegionSpreading = 3
	RegionSpreading_SF10       RegionSpreading = 4
	RegionSpreading_SF11       RegionSpreading = 5
	RegionSpreading_SF12       RegionSpreading = 6
)

// Enum value maps for RegionSpreading.
var (
	RegionSpreading_name = map[int32]string{
		0: "SF_INVALID",
		1: "SF7",
		2: "SF8",
		3: "SF9",
		4: "SF10",
		5: "SF11",
		6: "SF12",
	}
	RegionSpreading_value = map[string]int32{
		"SF_INVALID": 0,
		"SF7":        1,
		"SF8":        2,
		"SF9":        3,
		"SF10":       4,
		"SF11":       5,
		"SF12":       6,
	}
)

func (x RegionSpreading) Enum() *RegionSpreading {
	p := new(RegionSpreading)
	*p = x
	return p
}

func (x RegionSpreading) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegionSpreading) Descriptor() protoreflect.EnumDescriptor {
	return file_blockchain_region_param_v1_proto_enumTypes[0].Descriptor()
}

func (RegionSpreading) Type() protoreflect.EnumType {
	return &file_blockchain_region_param_v1_proto_enumTypes[0]
}

func (x RegionSpreading) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegionSpreading.Descriptor instead.
func (RegionSpreading) EnumDescriptor() ([]byte, []int) {
	return file_blockchain_region_param_v1_proto_rawDescGZIP(), []int{0}
}

type BlockchainRegionParamsV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionParams []*BlockchainRegionParamV1 `protobuf:"bytes,1,rep,name=region_params,json=regionParams,proto3" json:"region_params,omitempty"`
}

func (x *BlockchainRegionParamsV1) Reset() {
	*x = BlockchainRegionParamsV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_region_param_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainRegionParamsV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainRegionParamsV1) ProtoMessage() {}

func (x *BlockchainRegionParamsV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_region_param_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainRegionParamsV1.ProtoReflect.Descriptor instead.
func (*BlockchainRegionParamsV1) Descriptor() ([]byte, []int) {
	return file_blockchain_region_param_v1_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainRegionParamsV1) GetRegionParams() []*BlockchainRegionParamV1 {
	if x != nil {
		return x.RegionParams
	}
	return nil
}

type TaggedSpreading struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionSpreading RegionSpreading `protobuf:"varint,1,opt,name=region_spreading,json=regionSpreading,proto3,enum=helium.RegionSpreading" json:"region_spreading,omitempty"`
	MaxPacketSize   uint32          `protobuf:"varint,2,opt,name=max_packet_size,json=maxPacketSize,proto3" json:"max_packet_size,omitempty"`
}

func (x *TaggedSpreading) Reset() {
	*x = TaggedSpreading{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_region_param_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaggedSpreading) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaggedSpreading) ProtoMessage() {}

func (x *TaggedSpreading) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_region_param_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaggedSpreading.ProtoReflect.Descriptor instead.
func (*TaggedSpreading) Descriptor() ([]byte, []int) {
	return file_blockchain_region_param_v1_proto_rawDescGZIP(), []int{1}
}

func (x *TaggedSpreading) GetRegionSpreading() RegionSpreading {
	if x != nil {
		return x.RegionSpreading
	}
	return RegionSpreading_SF_INVALID
}

func (x *TaggedSpreading) GetMaxPacketSize() uint32 {
	if x != nil {
		return x.MaxPacketSize
	}
	return 0
}

type BlockchainRegionSpreadingV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaggedSpreading []*TaggedSpreading `protobuf:"bytes,1,rep,name=tagged_spreading,json=taggedSpreading,proto3" json:"tagged_spreading,omitempty"`
}

func (x *BlockchainRegionSpreadingV1) Reset() {
	*x = BlockchainRegionSpreadingV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_region_param_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainRegionSpreadingV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainRegionSpreadingV1) ProtoMessage() {}

func (x *BlockchainRegionSpreadingV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_region_param_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainRegionSpreadingV1.ProtoReflect.Descriptor instead.
func (*BlockchainRegionSpreadingV1) Descriptor() ([]byte, []int) {
	return file_blockchain_region_param_v1_proto_rawDescGZIP(), []int{2}
}

func (x *BlockchainRegionSpreadingV1) GetTaggedSpreading() []*TaggedSpreading {
	if x != nil {
		return x.TaggedSpreading
	}
	return nil
}

type BlockchainRegionParamV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// in hertz
	ChannelFrequency uint64 `protobuf:"varint,1,opt,name=channel_frequency,json=channelFrequency,proto3" json:"channel_frequency,omitempty"`
	// in hertz
	Bandwidth uint32 `protobuf:"varint,2,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	// in dBi x 10
	MaxEirp uint32 `protobuf:"varint,3,opt,name=max_eirp,json=maxEirp,proto3" json:"max_eirp,omitempty"`
	// list of atoms
	Spreading *BlockchainRegionSpreadingV1 `protobuf:"bytes,4,opt,name=spreading,proto3" json:"spreading,omitempty"`
}

func (x *BlockchainRegionParamV1) Reset() {
	*x = BlockchainRegionParamV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_region_param_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainRegionParamV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainRegionParamV1) ProtoMessage() {}

func (x *BlockchainRegionParamV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_region_param_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainRegionParamV1.ProtoReflect.Descriptor instead.
func (*BlockchainRegionParamV1) Descriptor() ([]byte, []int) {
	return file_blockchain_region_param_v1_proto_rawDescGZIP(), []int{3}
}

func (x *BlockchainRegionParamV1) GetChannelFrequency() uint64 {
	if x != nil {
		return x.ChannelFrequency
	}
	return 0
}

func (x *BlockchainRegionParamV1) GetBandwidth() uint32 {
	if x != nil {
		return x.Bandwidth
	}
	return 0
}

func (x *BlockchainRegionParamV1) GetMaxEirp() uint32 {
	if x != nil {
		return x.MaxEirp
	}
	return 0
}

func (x *BlockchainRegionParamV1) GetSpreading() *BlockchainRegionSpreadingV1 {
	if x != nil {
		return x.Spreading
	}
	return nil
}

var File_blockchain_region_param_v1_proto protoreflect.FileDescriptor

var file_blockchain_region_param_v1_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x22, 0x66, 0x0a, 0x1b, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x12, 0x47, 0x0a, 0x0d, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x5f, 0x76, 0x31, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x22, 0x7e, 0x0a, 0x10, 0x74, 0x61, 0x67, 0x67, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x72,
	0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x42, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x53, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61,
	0x78, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0x65, 0x0a, 0x1e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x76, 0x31, 0x12, 0x43, 0x0a, 0x10, 0x74, 0x61, 0x67, 0x67, 0x65, 0x64, 0x5f, 0x73,
	0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x74, 0x61, 0x67, 0x67, 0x65, 0x64, 0x5f, 0x73,
	0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x64,
	0x53, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x22, 0xc8, 0x01, 0x0a, 0x1a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x76, 0x31, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x10, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x46, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x69, 0x72, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x45, 0x69, 0x72, 0x70, 0x12, 0x44,
	0x0a, 0x09, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x72,
	0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x52, 0x09, 0x73, 0x70, 0x72, 0x65, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x2a, 0x5a, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x70,
	0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x46, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x46, 0x37, 0x10, 0x01,
	0x12, 0x07, 0x0a, 0x03, 0x53, 0x46, 0x38, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x46, 0x39,
	0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x46, 0x31, 0x30, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04,
	0x53, 0x46, 0x31, 0x31, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x46, 0x31, 0x32, 0x10, 0x06,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_region_param_v1_proto_rawDescOnce sync.Once
	file_blockchain_region_param_v1_proto_rawDescData = file_blockchain_region_param_v1_proto_rawDesc
)

func file_blockchain_region_param_v1_proto_rawDescGZIP() []byte {
	file_blockchain_region_param_v1_proto_rawDescOnce.Do(func() {
		file_blockchain_region_param_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_region_param_v1_proto_rawDescData)
	})
	return file_blockchain_region_param_v1_proto_rawDescData
}

var file_blockchain_region_param_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_blockchain_region_param_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_blockchain_region_param_v1_proto_goTypes = []interface{}{
	(RegionSpreading)(0),                // 0: helium.RegionSpreading
	(*BlockchainRegionParamsV1)(nil),    // 1: helium.blockchain_region_params_v1
	(*TaggedSpreading)(nil),             // 2: helium.tagged_spreading
	(*BlockchainRegionSpreadingV1)(nil), // 3: helium.blockchain_region_spreading_v1
	(*BlockchainRegionParamV1)(nil),     // 4: helium.blockchain_region_param_v1
}
var file_blockchain_region_param_v1_proto_depIdxs = []int32{
	4, // 0: helium.blockchain_region_params_v1.region_params:type_name -> helium.blockchain_region_param_v1
	0, // 1: helium.tagged_spreading.region_spreading:type_name -> helium.RegionSpreading
	2, // 2: helium.blockchain_region_spreading_v1.tagged_spreading:type_name -> helium.tagged_spreading
	3, // 3: helium.blockchain_region_param_v1.spreading:type_name -> helium.blockchain_region_spreading_v1
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_blockchain_region_param_v1_proto_init() }
func file_blockchain_region_param_v1_proto_init() {
	if File_blockchain_region_param_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_region_param_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainRegionParamsV1); i {
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
		file_blockchain_region_param_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaggedSpreading); i {
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
		file_blockchain_region_param_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainRegionSpreadingV1); i {
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
		file_blockchain_region_param_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainRegionParamV1); i {
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
			RawDescriptor: file_blockchain_region_param_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_region_param_v1_proto_goTypes,
		DependencyIndexes: file_blockchain_region_param_v1_proto_depIdxs,
		EnumInfos:         file_blockchain_region_param_v1_proto_enumTypes,
		MessageInfos:      file_blockchain_region_param_v1_proto_msgTypes,
	}.Build()
	File_blockchain_region_param_v1_proto = out.File
	file_blockchain_region_param_v1_proto_rawDesc = nil
	file_blockchain_region_param_v1_proto_goTypes = nil
	file_blockchain_region_param_v1_proto_depIdxs = nil
}
