// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: data_rate.proto

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

type DataRate int32

const (
	DataRate_SF12BW125     DataRate = 0
	DataRate_SF11BW125     DataRate = 1
	DataRate_SF10BW125     DataRate = 2
	DataRate_SF9BW125      DataRate = 3
	DataRate_SF8BW125      DataRate = 4
	DataRate_SF7BW125      DataRate = 5
	DataRate_SF12BW250     DataRate = 6
	DataRate_SF11BW250     DataRate = 7
	DataRate_SF10BW250     DataRate = 8
	DataRate_SF9BW250      DataRate = 9
	DataRate_SF8BW250      DataRate = 10
	DataRate_SF7BW250      DataRate = 11
	DataRate_SF12BW500     DataRate = 12
	DataRate_SF11BW500     DataRate = 13
	DataRate_SF10BW500     DataRate = 14
	DataRate_SF9BW500      DataRate = 15
	DataRate_SF8BW500      DataRate = 16
	DataRate_SF7BW500      DataRate = 17
	DataRate_LRFHSS1BW137  DataRate = 18
	DataRate_LRFHSS2BW137  DataRate = 19
	DataRate_LRFHSS1BW336  DataRate = 20
	DataRate_LRFHSS2BW336  DataRate = 21
	DataRate_LRFHSS1BW1523 DataRate = 22
	DataRate_LRFHSS2BW1523 DataRate = 23
	DataRate_FSK50         DataRate = 24
)

// Enum value maps for DataRate.
var (
	DataRate_name = map[int32]string{
		0:  "SF12BW125",
		1:  "SF11BW125",
		2:  "SF10BW125",
		3:  "SF9BW125",
		4:  "SF8BW125",
		5:  "SF7BW125",
		6:  "SF12BW250",
		7:  "SF11BW250",
		8:  "SF10BW250",
		9:  "SF9BW250",
		10: "SF8BW250",
		11: "SF7BW250",
		12: "SF12BW500",
		13: "SF11BW500",
		14: "SF10BW500",
		15: "SF9BW500",
		16: "SF8BW500",
		17: "SF7BW500",
		18: "LRFHSS1BW137",
		19: "LRFHSS2BW137",
		20: "LRFHSS1BW336",
		21: "LRFHSS2BW336",
		22: "LRFHSS1BW1523",
		23: "LRFHSS2BW1523",
		24: "FSK50",
	}
	DataRate_value = map[string]int32{
		"SF12BW125":     0,
		"SF11BW125":     1,
		"SF10BW125":     2,
		"SF9BW125":      3,
		"SF8BW125":      4,
		"SF7BW125":      5,
		"SF12BW250":     6,
		"SF11BW250":     7,
		"SF10BW250":     8,
		"SF9BW250":      9,
		"SF8BW250":      10,
		"SF7BW250":      11,
		"SF12BW500":     12,
		"SF11BW500":     13,
		"SF10BW500":     14,
		"SF9BW500":      15,
		"SF8BW500":      16,
		"SF7BW500":      17,
		"LRFHSS1BW137":  18,
		"LRFHSS2BW137":  19,
		"LRFHSS1BW336":  20,
		"LRFHSS2BW336":  21,
		"LRFHSS1BW1523": 22,
		"LRFHSS2BW1523": 23,
		"FSK50":         24,
	}
)

func (x DataRate) Enum() *DataRate {
	p := new(DataRate)
	*p = x
	return p
}

func (x DataRate) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataRate) Descriptor() protoreflect.EnumDescriptor {
	return file_data_rate_proto_enumTypes[0].Descriptor()
}

func (DataRate) Type() protoreflect.EnumType {
	return &file_data_rate_proto_enumTypes[0]
}

func (x DataRate) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataRate.Descriptor instead.
func (DataRate) EnumDescriptor() ([]byte, []int) {
	return file_data_rate_proto_rawDescGZIP(), []int{0}
}

var File_data_rate_proto protoreflect.FileDescriptor

var file_data_rate_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2a, 0x89, 0x03, 0x0a, 0x09, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x46, 0x31, 0x32, 0x42,
	0x57, 0x31, 0x32, 0x35, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x46, 0x31, 0x31, 0x42, 0x57,
	0x31, 0x32, 0x35, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x46, 0x31, 0x30, 0x42, 0x57, 0x31,
	0x32, 0x35, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x46, 0x39, 0x42, 0x57, 0x31, 0x32, 0x35,
	0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x46, 0x38, 0x42, 0x57, 0x31, 0x32, 0x35, 0x10, 0x04,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x46, 0x37, 0x42, 0x57, 0x31, 0x32, 0x35, 0x10, 0x05, 0x12, 0x0d,
	0x0a, 0x09, 0x53, 0x46, 0x31, 0x32, 0x42, 0x57, 0x32, 0x35, 0x30, 0x10, 0x06, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x46, 0x31, 0x31, 0x42, 0x57, 0x32, 0x35, 0x30, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x46, 0x31, 0x30, 0x42, 0x57, 0x32, 0x35, 0x30, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x46, 0x39, 0x42, 0x57, 0x32, 0x35, 0x30, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x46, 0x38,
	0x42, 0x57, 0x32, 0x35, 0x30, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x46, 0x37, 0x42, 0x57,
	0x32, 0x35, 0x30, 0x10, 0x0b, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x46, 0x31, 0x32, 0x42, 0x57, 0x35,
	0x30, 0x30, 0x10, 0x0c, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x46, 0x31, 0x31, 0x42, 0x57, 0x35, 0x30,
	0x30, 0x10, 0x0d, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x46, 0x31, 0x30, 0x42, 0x57, 0x35, 0x30, 0x30,
	0x10, 0x0e, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x46, 0x39, 0x42, 0x57, 0x35, 0x30, 0x30, 0x10, 0x0f,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x46, 0x38, 0x42, 0x57, 0x35, 0x30, 0x30, 0x10, 0x10, 0x12, 0x0c,
	0x0a, 0x08, 0x53, 0x46, 0x37, 0x42, 0x57, 0x35, 0x30, 0x30, 0x10, 0x11, 0x12, 0x10, 0x0a, 0x0c,
	0x4c, 0x52, 0x46, 0x48, 0x53, 0x53, 0x31, 0x42, 0x57, 0x31, 0x33, 0x37, 0x10, 0x12, 0x12, 0x10,
	0x0a, 0x0c, 0x4c, 0x52, 0x46, 0x48, 0x53, 0x53, 0x32, 0x42, 0x57, 0x31, 0x33, 0x37, 0x10, 0x13,
	0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x52, 0x46, 0x48, 0x53, 0x53, 0x31, 0x42, 0x57, 0x33, 0x33, 0x36,
	0x10, 0x14, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x52, 0x46, 0x48, 0x53, 0x53, 0x32, 0x42, 0x57, 0x33,
	0x33, 0x36, 0x10, 0x15, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x52, 0x46, 0x48, 0x53, 0x53, 0x31, 0x42,
	0x57, 0x31, 0x35, 0x32, 0x33, 0x10, 0x16, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x52, 0x46, 0x48, 0x53,
	0x53, 0x32, 0x42, 0x57, 0x31, 0x35, 0x32, 0x33, 0x10, 0x17, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x53,
	0x4b, 0x35, 0x30, 0x10, 0x18, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_rate_proto_rawDescOnce sync.Once
	file_data_rate_proto_rawDescData = file_data_rate_proto_rawDesc
)

func file_data_rate_proto_rawDescGZIP() []byte {
	file_data_rate_proto_rawDescOnce.Do(func() {
		file_data_rate_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_rate_proto_rawDescData)
	})
	return file_data_rate_proto_rawDescData
}

var file_data_rate_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_data_rate_proto_goTypes = []interface{}{
	(DataRate)(0), // 0: helium.data_rate
}
var file_data_rate_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_data_rate_proto_init() }
func file_data_rate_proto_init() {
	if File_data_rate_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_rate_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_rate_proto_goTypes,
		DependencyIndexes: file_data_rate_proto_depIdxs,
		EnumInfos:         file_data_rate_proto_enumTypes,
	}.Build()
	File_data_rate_proto = out.File
	file_data_rate_proto_rawDesc = nil
	file_data_rate_proto_goTypes = nil
	file_data_rate_proto_depIdxs = nil
}
