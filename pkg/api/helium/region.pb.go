// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: region.proto

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

type Region int32

const (
	Region_US915 Region = 0
	Region_EU868 Region = 1
	Region_EU433 Region = 2
	Region_CN470 Region = 3
	// Deprecated: Do not use.
	Region_CN779     Region = 4
	Region_AU915     Region = 5
	Region_AS923_1   Region = 6
	Region_KR920     Region = 7
	Region_IN865     Region = 8
	Region_AS923_2   Region = 9
	Region_AS923_3   Region = 10
	Region_AS923_4   Region = 11
	Region_AS923_1B  Region = 12
	Region_CD900_1A  Region = 13
	Region_RU864     Region = 14
	Region_EU868_A   Region = 15
	Region_EU868_B   Region = 16
	Region_EU868_C   Region = 17
	Region_EU868_D   Region = 18
	Region_EU868_E   Region = 19
	Region_EU868_F   Region = 20
	Region_AU915_SB1 Region = 21
	Region_AU915_SB2 Region = 22
	Region_AS923_1A  Region = 23
	Region_AS923_1C  Region = 24
	Region_AS923_1D  Region = 25
	Region_AS923_1E  Region = 26
	Region_AS923_1F  Region = 27
)

// Enum value maps for Region.
var (
	Region_name = map[int32]string{
		0:  "US915",
		1:  "EU868",
		2:  "EU433",
		3:  "CN470",
		4:  "CN779",
		5:  "AU915",
		6:  "AS923_1",
		7:  "KR920",
		8:  "IN865",
		9:  "AS923_2",
		10: "AS923_3",
		11: "AS923_4",
		12: "AS923_1B",
		13: "CD900_1A",
		14: "RU864",
		15: "EU868_A",
		16: "EU868_B",
		17: "EU868_C",
		18: "EU868_D",
		19: "EU868_E",
		20: "EU868_F",
		21: "AU915_SB1",
		22: "AU915_SB2",
		23: "AS923_1A",
		24: "AS923_1C",
		25: "AS923_1D",
		26: "AS923_1E",
		27: "AS923_1F",
	}
	Region_value = map[string]int32{
		"US915":     0,
		"EU868":     1,
		"EU433":     2,
		"CN470":     3,
		"CN779":     4,
		"AU915":     5,
		"AS923_1":   6,
		"KR920":     7,
		"IN865":     8,
		"AS923_2":   9,
		"AS923_3":   10,
		"AS923_4":   11,
		"AS923_1B":  12,
		"CD900_1A":  13,
		"RU864":     14,
		"EU868_A":   15,
		"EU868_B":   16,
		"EU868_C":   17,
		"EU868_D":   18,
		"EU868_E":   19,
		"EU868_F":   20,
		"AU915_SB1": 21,
		"AU915_SB2": 22,
		"AS923_1A":  23,
		"AS923_1C":  24,
		"AS923_1D":  25,
		"AS923_1E":  26,
		"AS923_1F":  27,
	}
)

func (x Region) Enum() *Region {
	p := new(Region)
	*p = x
	return p
}

func (x Region) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Region) Descriptor() protoreflect.EnumDescriptor {
	return file_region_proto_enumTypes[0].Descriptor()
}

func (Region) Type() protoreflect.EnumType {
	return &file_region_proto_enumTypes[0]
}

func (x Region) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Region.Descriptor instead.
func (Region) EnumDescriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{0}
}

var File_region_proto protoreflect.FileDescriptor

var file_region_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2a, 0xf1, 0x02, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x53, 0x39, 0x31, 0x35, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x55, 0x38, 0x36, 0x38, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x55, 0x34, 0x33, 0x33,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4e, 0x34, 0x37, 0x30, 0x10, 0x03, 0x12, 0x0d, 0x0a,
	0x05, 0x43, 0x4e, 0x37, 0x37, 0x39, 0x10, 0x04, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x41, 0x55, 0x39, 0x31, 0x35, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x53, 0x39, 0x32, 0x33,
	0x5f, 0x31, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x4b, 0x52, 0x39, 0x32, 0x30, 0x10, 0x07, 0x12,
	0x09, 0x0a, 0x05, 0x49, 0x4e, 0x38, 0x36, 0x35, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x53,
	0x39, 0x32, 0x33, 0x5f, 0x32, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x53, 0x39, 0x32, 0x33,
	0x5f, 0x33, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x53, 0x39, 0x32, 0x33, 0x5f, 0x34, 0x10,
	0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x53, 0x39, 0x32, 0x33, 0x5f, 0x31, 0x42, 0x10, 0x0c, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x44, 0x39, 0x30, 0x30, 0x5f, 0x31, 0x41, 0x10, 0x0d, 0x12, 0x09, 0x0a,
	0x05, 0x52, 0x55, 0x38, 0x36, 0x34, 0x10, 0x0e, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x55, 0x38, 0x36,
	0x38, 0x5f, 0x41, 0x10, 0x0f, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x55, 0x38, 0x36, 0x38, 0x5f, 0x42,
	0x10, 0x10, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x55, 0x38, 0x36, 0x38, 0x5f, 0x43, 0x10, 0x11, 0x12,
	0x0b, 0x0a, 0x07, 0x45, 0x55, 0x38, 0x36, 0x38, 0x5f, 0x44, 0x10, 0x12, 0x12, 0x0b, 0x0a, 0x07,
	0x45, 0x55, 0x38, 0x36, 0x38, 0x5f, 0x45, 0x10, 0x13, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x55, 0x38,
	0x36, 0x38, 0x5f, 0x46, 0x10, 0x14, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x55, 0x39, 0x31, 0x35, 0x5f,
	0x53, 0x42, 0x31, 0x10, 0x15, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x55, 0x39, 0x31, 0x35, 0x5f, 0x53,
	0x42, 0x32, 0x10, 0x16, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x53, 0x39, 0x32, 0x33, 0x5f, 0x31, 0x41,
	0x10, 0x17, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x53, 0x39, 0x32, 0x33, 0x5f, 0x31, 0x43, 0x10, 0x18,
	0x12, 0x0c, 0x0a, 0x08, 0x41, 0x53, 0x39, 0x32, 0x33, 0x5f, 0x31, 0x44, 0x10, 0x19, 0x12, 0x0c,
	0x0a, 0x08, 0x41, 0x53, 0x39, 0x32, 0x33, 0x5f, 0x31, 0x45, 0x10, 0x1a, 0x12, 0x0c, 0x0a, 0x08,
	0x41, 0x53, 0x39, 0x32, 0x33, 0x5f, 0x31, 0x46, 0x10, 0x1b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_region_proto_rawDescOnce sync.Once
	file_region_proto_rawDescData = file_region_proto_rawDesc
)

func file_region_proto_rawDescGZIP() []byte {
	file_region_proto_rawDescOnce.Do(func() {
		file_region_proto_rawDescData = protoimpl.X.CompressGZIP(file_region_proto_rawDescData)
	})
	return file_region_proto_rawDescData
}

var file_region_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_region_proto_goTypes = []interface{}{
	(Region)(0), // 0: helium.region
}
var file_region_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_region_proto_init() }
func file_region_proto_init() {
	if File_region_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_region_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_region_proto_goTypes,
		DependencyIndexes: file_region_proto_depIdxs,
		EnumInfos:         file_region_proto_enumTypes,
	}.Build()
	File_region_proto = out.File
	file_region_proto_rawDesc = nil
	file_region_proto_goTypes = nil
	file_region_proto_depIdxs = nil
}
