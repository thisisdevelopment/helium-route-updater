// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: gateway_staking_mode.proto

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

type GatewayStakingMode int32

const (
	GatewayStakingMode_dataonly GatewayStakingMode = 0
	GatewayStakingMode_full     GatewayStakingMode = 1
	GatewayStakingMode_light    GatewayStakingMode = 2
)

// Enum value maps for GatewayStakingMode.
var (
	GatewayStakingMode_name = map[int32]string{
		0: "dataonly",
		1: "full",
		2: "light",
	}
	GatewayStakingMode_value = map[string]int32{
		"dataonly": 0,
		"full":     1,
		"light":    2,
	}
)

func (x GatewayStakingMode) Enum() *GatewayStakingMode {
	p := new(GatewayStakingMode)
	*p = x
	return p
}

func (x GatewayStakingMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GatewayStakingMode) Descriptor() protoreflect.EnumDescriptor {
	return file_gateway_staking_mode_proto_enumTypes[0].Descriptor()
}

func (GatewayStakingMode) Type() protoreflect.EnumType {
	return &file_gateway_staking_mode_proto_enumTypes[0]
}

func (x GatewayStakingMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GatewayStakingMode.Descriptor instead.
func (GatewayStakingMode) EnumDescriptor() ([]byte, []int) {
	return file_gateway_staking_mode_proto_rawDescGZIP(), []int{0}
}

var File_gateway_staking_mode_proto protoreflect.FileDescriptor

var file_gateway_staking_mode_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x65,
	0x6c, 0x69, 0x75, 0x6d, 0x2a, 0x39, 0x0a, 0x14, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x0c, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x6f, 0x6e, 0x6c, 0x79, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x66, 0x75,
	0x6c, 0x6c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x10, 0x02, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_staking_mode_proto_rawDescOnce sync.Once
	file_gateway_staking_mode_proto_rawDescData = file_gateway_staking_mode_proto_rawDesc
)

func file_gateway_staking_mode_proto_rawDescGZIP() []byte {
	file_gateway_staking_mode_proto_rawDescOnce.Do(func() {
		file_gateway_staking_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_staking_mode_proto_rawDescData)
	})
	return file_gateway_staking_mode_proto_rawDescData
}

var file_gateway_staking_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gateway_staking_mode_proto_goTypes = []interface{}{
	(GatewayStakingMode)(0), // 0: helium.gateway_staking_mode
}
var file_gateway_staking_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gateway_staking_mode_proto_init() }
func file_gateway_staking_mode_proto_init() {
	if File_gateway_staking_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_staking_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_staking_mode_proto_goTypes,
		DependencyIndexes: file_gateway_staking_mode_proto_depIdxs,
		EnumInfos:         file_gateway_staking_mode_proto_enumTypes,
	}.Build()
	File_gateway_staking_mode_proto = out.File
	file_gateway_staking_mode_proto_rawDesc = nil
	file_gateway_staking_mode_proto_goTypes = nil
	file_gateway_staking_mode_proto_depIdxs = nil
}
