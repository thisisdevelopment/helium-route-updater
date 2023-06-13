// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: service/multi_buy.proto

package multi_buy

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

type MultiBuyIncReqV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *MultiBuyIncReqV1) Reset() {
	*x = MultiBuyIncReqV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_multi_buy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiBuyIncReqV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiBuyIncReqV1) ProtoMessage() {}

func (x *MultiBuyIncReqV1) ProtoReflect() protoreflect.Message {
	mi := &file_service_multi_buy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiBuyIncReqV1.ProtoReflect.Descriptor instead.
func (*MultiBuyIncReqV1) Descriptor() ([]byte, []int) {
	return file_service_multi_buy_proto_rawDescGZIP(), []int{0}
}

func (x *MultiBuyIncReqV1) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type MultiBuyIncResV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *MultiBuyIncResV1) Reset() {
	*x = MultiBuyIncResV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_multi_buy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiBuyIncResV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiBuyIncResV1) ProtoMessage() {}

func (x *MultiBuyIncResV1) ProtoReflect() protoreflect.Message {
	mi := &file_service_multi_buy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiBuyIncResV1.ProtoReflect.Descriptor instead.
func (*MultiBuyIncResV1) Descriptor() ([]byte, []int) {
	return file_service_multi_buy_proto_rawDescGZIP(), []int{1}
}

func (x *MultiBuyIncResV1) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_service_multi_buy_proto protoreflect.FileDescriptor

var file_service_multi_buy_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f,
	0x62, 0x75, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x68, 0x65, 0x6c, 0x69, 0x75,
	0x6d, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x62, 0x75, 0x79, 0x22, 0x28, 0x0a, 0x14, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x69, 0x6e, 0x63, 0x5f, 0x72, 0x65, 0x71,
	0x5f, 0x76, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x2c, 0x0a, 0x14, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x62,
	0x75, 0x79, 0x5f, 0x69, 0x6e, 0x63, 0x5f, 0x72, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x32, 0x62, 0x0a, 0x09, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x62, 0x75, 0x79,
	0x12, 0x55, 0x0a, 0x03, 0x69, 0x6e, 0x63, 0x12, 0x26, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d,
	0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x62, 0x75, 0x79, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x5f, 0x62, 0x75, 0x79, 0x5f, 0x69, 0x6e, 0x63, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x76, 0x31, 0x1a,
	0x26, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x62,
	0x75, 0x79, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x69, 0x6e, 0x63,
	0x5f, 0x72, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_multi_buy_proto_rawDescOnce sync.Once
	file_service_multi_buy_proto_rawDescData = file_service_multi_buy_proto_rawDesc
)

func file_service_multi_buy_proto_rawDescGZIP() []byte {
	file_service_multi_buy_proto_rawDescOnce.Do(func() {
		file_service_multi_buy_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_multi_buy_proto_rawDescData)
	})
	return file_service_multi_buy_proto_rawDescData
}

var file_service_multi_buy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_multi_buy_proto_goTypes = []interface{}{
	(*MultiBuyIncReqV1)(nil), // 0: helium.multi_buy.multi_buy_inc_req_v1
	(*MultiBuyIncResV1)(nil), // 1: helium.multi_buy.multi_buy_inc_res_v1
}
var file_service_multi_buy_proto_depIdxs = []int32{
	0, // 0: helium.multi_buy.multi_buy.inc:input_type -> helium.multi_buy.multi_buy_inc_req_v1
	1, // 1: helium.multi_buy.multi_buy.inc:output_type -> helium.multi_buy.multi_buy_inc_res_v1
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_multi_buy_proto_init() }
func file_service_multi_buy_proto_init() {
	if File_service_multi_buy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_multi_buy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiBuyIncReqV1); i {
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
		file_service_multi_buy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiBuyIncResV1); i {
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
			RawDescriptor: file_service_multi_buy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_multi_buy_proto_goTypes,
		DependencyIndexes: file_service_multi_buy_proto_depIdxs,
		MessageInfos:      file_service_multi_buy_proto_msgTypes,
	}.Build()
	File_service_multi_buy_proto = out.File
	file_service_multi_buy_proto_rawDesc = nil
	file_service_multi_buy_proto_goTypes = nil
	file_service_multi_buy_proto_depIdxs = nil
}
