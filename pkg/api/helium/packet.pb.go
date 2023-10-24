// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: packet.proto

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

type PacketPacketType int32

const (
	Packet_longfi  PacketPacketType = 0
	Packet_lorawan PacketPacketType = 1
)

// Enum value maps for PacketPacketType.
var (
	PacketPacketType_name = map[int32]string{
		0: "longfi",
		1: "lorawan",
	}
	PacketPacketType_value = map[string]int32{
		"longfi":  0,
		"lorawan": 1,
	}
)

func (x PacketPacketType) Enum() *PacketPacketType {
	p := new(PacketPacketType)
	*p = x
	return p
}

func (x PacketPacketType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PacketPacketType) Descriptor() protoreflect.EnumDescriptor {
	return file_packet_proto_enumTypes[0].Descriptor()
}

func (PacketPacketType) Type() protoreflect.EnumType {
	return &file_packet_proto_enumTypes[0]
}

func (x PacketPacketType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PacketPacketType.Descriptor instead.
func (PacketPacketType) EnumDescriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{3, 0}
}

type Eui struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deveui uint64 `protobuf:"varint,1,opt,name=deveui,proto3" json:"deveui,omitempty"`
	Appeui uint64 `protobuf:"varint,2,opt,name=appeui,proto3" json:"appeui,omitempty"`
}

func (x *Eui) Reset() {
	*x = Eui{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Eui) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Eui) ProtoMessage() {}

func (x *Eui) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Eui.ProtoReflect.Descriptor instead.
func (*Eui) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{0}
}

func (x *Eui) GetDeveui() uint64 {
	if x != nil {
		return x.Deveui
	}
	return 0
}

func (x *Eui) GetAppeui() uint64 {
	if x != nil {
		return x.Appeui
	}
	return 0
}

type RoutingInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*RoutingInformation_Devaddr
	//	*RoutingInformation_Eui
	Data isRoutingInformation_Data `protobuf_oneof:"data"`
}

func (x *RoutingInformation) Reset() {
	*x = RoutingInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutingInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingInformation) ProtoMessage() {}

func (x *RoutingInformation) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingInformation.ProtoReflect.Descriptor instead.
func (*RoutingInformation) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{1}
}

func (m *RoutingInformation) GetData() isRoutingInformation_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *RoutingInformation) GetDevaddr() uint32 {
	if x, ok := x.GetData().(*RoutingInformation_Devaddr); ok {
		return x.Devaddr
	}
	return 0
}

func (x *RoutingInformation) GetEui() *Eui {
	if x, ok := x.GetData().(*RoutingInformation_Eui); ok {
		return x.Eui
	}
	return nil
}

type isRoutingInformation_Data interface {
	isRoutingInformation_Data()
}

type RoutingInformation_Devaddr struct {
	Devaddr uint32 `protobuf:"varint,1,opt,name=devaddr,proto3,oneof"`
}

type RoutingInformation_Eui struct {
	Eui *Eui `protobuf:"bytes,2,opt,name=eui,proto3,oneof"`
}

func (*RoutingInformation_Devaddr) isRoutingInformation_Data() {}

func (*RoutingInformation_Eui) isRoutingInformation_Data() {}

type Window struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp uint64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Frequency float32 `protobuf:"fixed32,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Datarate  string  `protobuf:"bytes,3,opt,name=datarate,proto3" json:"datarate,omitempty"`
}

func (x *Window) Reset() {
	*x = Window{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Window) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Window) ProtoMessage() {}

func (x *Window) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Window.ProtoReflect.Descriptor instead.
func (*Window) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{2}
}

func (x *Window) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Window) GetFrequency() float32 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *Window) GetDatarate() string {
	if x != nil {
		return x.Datarate
	}
	return ""
}

type Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oui            uint32              `protobuf:"varint,1,opt,name=oui,proto3" json:"oui,omitempty"`
	Type           PacketPacketType    `protobuf:"varint,2,opt,name=type,proto3,enum=helium.PacketPacketType" json:"type,omitempty"`
	Payload        []byte              `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp      uint64              `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SignalStrength float32             `protobuf:"fixed32,5,opt,name=signal_strength,json=signalStrength,proto3" json:"signal_strength,omitempty"`
	Frequency      float32             `protobuf:"fixed32,6,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Datarate       string              `protobuf:"bytes,7,opt,name=datarate,proto3" json:"datarate,omitempty"`
	Snr            float32             `protobuf:"fixed32,8,opt,name=snr,proto3" json:"snr,omitempty"`
	Routing        *RoutingInformation `protobuf:"bytes,9,opt,name=routing,proto3" json:"routing,omitempty"`
	Rx2Window      *Window             `protobuf:"bytes,10,opt,name=rx2_window,json=rx2Window,proto3" json:"rx2_window,omitempty"`
}

func (x *Packet) Reset() {
	*x = Packet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_packet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_packet_proto_rawDescGZIP(), []int{3}
}

func (x *Packet) GetOui() uint32 {
	if x != nil {
		return x.Oui
	}
	return 0
}

func (x *Packet) GetType() PacketPacketType {
	if x != nil {
		return x.Type
	}
	return Packet_longfi
}

func (x *Packet) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Packet) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Packet) GetSignalStrength() float32 {
	if x != nil {
		return x.SignalStrength
	}
	return 0
}

func (x *Packet) GetFrequency() float32 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *Packet) GetDatarate() string {
	if x != nil {
		return x.Datarate
	}
	return ""
}

func (x *Packet) GetSnr() float32 {
	if x != nil {
		return x.Snr
	}
	return 0
}

func (x *Packet) GetRouting() *RoutingInformation {
	if x != nil {
		return x.Routing
	}
	return nil
}

func (x *Packet) GetRx2Window() *Window {
	if x != nil {
		return x.Rx2Window
	}
	return nil
}

var File_packet_proto protoreflect.FileDescriptor

var file_packet_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x22, 0x35, 0x0a, 0x03, 0x65, 0x75, 0x69, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x65, 0x76, 0x65, 0x75, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x64,
	0x65, 0x76, 0x65, 0x75, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x65, 0x75, 0x69, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x70, 0x70, 0x65, 0x75, 0x69, 0x22, 0x5a, 0x0a,
	0x13, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x76, 0x61, 0x64, 0x64, 0x72,
	0x12, 0x1f, 0x0a, 0x03, 0x65, 0x75, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x65, 0x75, 0x69, 0x48, 0x00, 0x52, 0x03, 0x65, 0x75,
	0x69, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x60, 0x0a, 0x06, 0x77, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x72, 0x61, 0x74, 0x65, 0x22, 0x85, 0x03, 0x0a, 0x06,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x75, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x6f, 0x75, 0x69, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x27, 0x0a, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x66, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x72,
	0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6e, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x73, 0x6e, 0x72, 0x12, 0x35, 0x0a, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2d, 0x0a, 0x0a,
	0x72, 0x78, 0x32, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x52, 0x09, 0x72, 0x78, 0x32, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x22, 0x26, 0x0a, 0x0b, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x6c, 0x6f,
	0x6e, 0x67, 0x66, 0x69, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61,
	0x6e, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packet_proto_rawDescOnce sync.Once
	file_packet_proto_rawDescData = file_packet_proto_rawDesc
)

func file_packet_proto_rawDescGZIP() []byte {
	file_packet_proto_rawDescOnce.Do(func() {
		file_packet_proto_rawDescData = protoimpl.X.CompressGZIP(file_packet_proto_rawDescData)
	})
	return file_packet_proto_rawDescData
}

var file_packet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_packet_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_packet_proto_goTypes = []interface{}{
	(PacketPacketType)(0),      // 0: helium.packet.packet_type
	(*Eui)(nil),                // 1: helium.eui
	(*RoutingInformation)(nil), // 2: helium.routing_information
	(*Window)(nil),             // 3: helium.window
	(*Packet)(nil),             // 4: helium.packet
}
var file_packet_proto_depIdxs = []int32{
	1, // 0: helium.routing_information.eui:type_name -> helium.eui
	0, // 1: helium.packet.type:type_name -> helium.packet.packet_type
	2, // 2: helium.packet.routing:type_name -> helium.routing_information
	3, // 3: helium.packet.rx2_window:type_name -> helium.window
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_packet_proto_init() }
func file_packet_proto_init() {
	if File_packet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_packet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Eui); i {
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
		file_packet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutingInformation); i {
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
		file_packet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Window); i {
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
		file_packet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Packet); i {
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
	file_packet_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*RoutingInformation_Devaddr)(nil),
		(*RoutingInformation_Eui)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_packet_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packet_proto_goTypes,
		DependencyIndexes: file_packet_proto_depIdxs,
		EnumInfos:         file_packet_proto_enumTypes,
		MessageInfos:      file_packet_proto_msgTypes,
	}.Build()
	File_packet_proto = out.File
	file_packet_proto_rawDesc = nil
	file_packet_proto_goTypes = nil
	file_packet_proto_depIdxs = nil
}
