// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: radio.proto

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

type Radio int32

const (
	Radio_R0 Radio = 0
	Radio_R1 Radio = 1
)

// Enum value maps for Radio.
var (
	Radio_name = map[int32]string{
		0: "R0",
		1: "R1",
	}
	Radio_value = map[string]int32{
		"R0": 0,
		"R1": 1,
	}
)

func (x Radio) Enum() *Radio {
	p := new(Radio)
	*p = x
	return p
}

func (x Radio) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Radio) Descriptor() protoreflect.EnumDescriptor {
	return file_radio_proto_enumTypes[0].Descriptor()
}

func (Radio) Type() protoreflect.EnumType {
	return &file_radio_proto_enumTypes[0]
}

func (x Radio) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Radio.Descriptor instead.
func (Radio) EnumDescriptor() ([]byte, []int) {
	return file_radio_proto_rawDescGZIP(), []int{0}
}

type Spreading int32

const (
	Spreading_SF_UNDEFINED Spreading = 0
	Spreading_SF7          Spreading = 1
	Spreading_SF8          Spreading = 2
	Spreading_SF9          Spreading = 3
	Spreading_SF10         Spreading = 4
	Spreading_SF11         Spreading = 5
	Spreading_SF12         Spreading = 6
)

// Enum value maps for Spreading.
var (
	Spreading_name = map[int32]string{
		0: "SF_UNDEFINED",
		1: "SF7",
		2: "SF8",
		3: "SF9",
		4: "SF10",
		5: "SF11",
		6: "SF12",
	}
	Spreading_value = map[string]int32{
		"SF_UNDEFINED": 0,
		"SF7":          1,
		"SF8":          2,
		"SF9":          3,
		"SF10":         4,
		"SF11":         5,
		"SF12":         6,
	}
)

func (x Spreading) Enum() *Spreading {
	p := new(Spreading)
	*p = x
	return p
}

func (x Spreading) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Spreading) Descriptor() protoreflect.EnumDescriptor {
	return file_radio_proto_enumTypes[1].Descriptor()
}

func (Spreading) Type() protoreflect.EnumType {
	return &file_radio_proto_enumTypes[1]
}

func (x Spreading) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Spreading.Descriptor instead.
func (Spreading) EnumDescriptor() ([]byte, []int) {
	return file_radio_proto_rawDescGZIP(), []int{1}
}

type Bandwidth int32

const (
	Bandwidth_BW_UNDEFINED Bandwidth = 0
	Bandwidth_BW7_8kHz     Bandwidth = 1
	Bandwidth_BW15_6kHz    Bandwidth = 2
	Bandwidth_BW31_2kHz    Bandwidth = 3
	Bandwidth_BW62_5kHz    Bandwidth = 4
	Bandwidth_BW125kHz     Bandwidth = 5
	Bandwidth_BW250kHz     Bandwidth = 6
	Bandwidth_BW500kHz     Bandwidth = 7
)

// Enum value maps for Bandwidth.
var (
	Bandwidth_name = map[int32]string{
		0: "BW_UNDEFINED",
		1: "BW7_8kHz",
		2: "BW15_6kHz",
		3: "BW31_2kHz",
		4: "BW62_5kHz",
		5: "BW125kHz",
		6: "BW250kHz",
		7: "BW500kHz",
	}
	Bandwidth_value = map[string]int32{
		"BW_UNDEFINED": 0,
		"BW7_8kHz":     1,
		"BW15_6kHz":    2,
		"BW31_2kHz":    3,
		"BW62_5kHz":    4,
		"BW125kHz":     5,
		"BW250kHz":     6,
		"BW500kHz":     7,
	}
)

func (x Bandwidth) Enum() *Bandwidth {
	p := new(Bandwidth)
	*p = x
	return p
}

func (x Bandwidth) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Bandwidth) Descriptor() protoreflect.EnumDescriptor {
	return file_radio_proto_enumTypes[2].Descriptor()
}

func (Bandwidth) Type() protoreflect.EnumType {
	return &file_radio_proto_enumTypes[2]
}

func (x Bandwidth) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Bandwidth.Descriptor instead.
func (Bandwidth) EnumDescriptor() ([]byte, []int) {
	return file_radio_proto_rawDescGZIP(), []int{2}
}

type Coderate int32

const (
	Coderate_CR_UNDEFINED Coderate = 0
	Coderate_CR4_5        Coderate = 1
	Coderate_CR4_6        Coderate = 2
	Coderate_CR4_7        Coderate = 3
	Coderate_CR4_8        Coderate = 4
)

// Enum value maps for Coderate.
var (
	Coderate_name = map[int32]string{
		0: "CR_UNDEFINED",
		1: "CR4_5",
		2: "CR4_6",
		3: "CR4_7",
		4: "CR4_8",
	}
	Coderate_value = map[string]int32{
		"CR_UNDEFINED": 0,
		"CR4_5":        1,
		"CR4_6":        2,
		"CR4_7":        3,
		"CR4_8":        4,
	}
)

func (x Coderate) Enum() *Coderate {
	p := new(Coderate)
	*p = x
	return p
}

func (x Coderate) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Coderate) Descriptor() protoreflect.EnumDescriptor {
	return file_radio_proto_enumTypes[3].Descriptor()
}

func (Coderate) Type() protoreflect.EnumType {
	return &file_radio_proto_enumTypes[3]
}

func (x Coderate) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Coderate.Descriptor instead.
func (Coderate) EnumDescriptor() ([]byte, []int) {
	return file_radio_proto_rawDescGZIP(), []int{3}
}

type RadioReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Kind:
	//
	//	*RadioReq_Tx
	Kind isRadioReq_Kind `protobuf_oneof:"kind"`
}

func (x *RadioReq) Reset() {
	*x = RadioReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_radio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioReq) ProtoMessage() {}

func (x *RadioReq) ProtoReflect() protoreflect.Message {
	mi := &file_radio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioReq.ProtoReflect.Descriptor instead.
func (*RadioReq) Descriptor() ([]byte, []int) {
	return file_radio_proto_rawDescGZIP(), []int{0}
}

func (x *RadioReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (m *RadioReq) GetKind() isRadioReq_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *RadioReq) GetTx() *RadioTxReq {
	if x, ok := x.GetKind().(*RadioReq_Tx); ok {
		return x.Tx
	}
	return nil
}

type isRadioReq_Kind interface {
	isRadioReq_Kind()
}

type RadioReq_Tx struct {
	Tx *RadioTxReq `protobuf:"bytes,2,opt,name=tx,proto3,oneof"`
}

func (*RadioReq_Tx) isRadioReq_Kind() {}

type RadioResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Kind:
	//
	//	*RadioResp_Tx
	//	*RadioResp_RxPacket
	//	*RadioResp_ParseErr
	Kind isRadioResp_Kind `protobuf_oneof:"kind"`
}

func (x *RadioResp) Reset() {
	*x = RadioResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_radio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioResp) ProtoMessage() {}

func (x *RadioResp) ProtoReflect() protoreflect.Message {
	mi := &file_radio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioResp.ProtoReflect.Descriptor instead.
func (*RadioResp) Descriptor() ([]byte, []int) {
	return file_radio_proto_rawDescGZIP(), []int{1}
}

func (x *RadioResp) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (m *RadioResp) GetKind() isRadioResp_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *RadioResp) GetTx() *RadioTxResp {
	if x, ok := x.GetKind().(*RadioResp_Tx); ok {
		return x.Tx
	}
	return nil
}

func (x *RadioResp) GetRxPacket() *RadioRxPacket {
	if x, ok := x.GetKind().(*RadioResp_RxPacket); ok {
		return x.RxPacket
	}
	return nil
}

func (x *RadioResp) GetParseErr() []byte {
	if x, ok := x.GetKind().(*RadioResp_ParseErr); ok {
		return x.ParseErr
	}
	return nil
}

type isRadioResp_Kind interface {
	isRadioResp_Kind()
}

type RadioResp_Tx struct {
	Tx *RadioTxResp `protobuf:"bytes,2,opt,name=tx,proto3,oneof"`
}

type RadioResp_RxPacket struct {
	RxPacket *RadioRxPacket `protobuf:"bytes,3,opt,name=rx_packet,json=rxPacket,proto3,oneof"`
}

type RadioResp_ParseErr struct {
	ParseErr []byte `protobuf:"bytes,4,opt,name=parse_err,json=parseErr,proto3,oneof"`
}

func (*RadioResp_Tx) isRadioResp_Kind() {}

func (*RadioResp_RxPacket) isRadioResp_Kind() {}

func (*RadioResp_ParseErr) isRadioResp_Kind() {}

type RadioTxReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Center frequency to transmit on.
	Freq uint32 `protobuf:"varint,1,opt,name=freq,proto3" json:"freq,omitempty"`
	// Which radio to transmit on.
	Radio Radio `protobuf:"varint,2,opt,name=radio,proto3,enum=helium.radio.Radio" json:"radio,omitempty"`
	// TX power (in dBm).
	Power int32 `protobuf:"varint,3,opt,name=power,proto3" json:"power,omitempty"`
	// Modulation bandwidth.
	Bandwidth Bandwidth `protobuf:"varint,4,opt,name=bandwidth,proto3,enum=helium.radio.Bandwidth" json:"bandwidth,omitempty"`
	// Spreading factor to use with this packet.
	Spreading Spreading `protobuf:"varint,5,opt,name=spreading,proto3,enum=helium.radio.Spreading" json:"spreading,omitempty"`
	// Error-correcting-code of the packet.
	Coderate Coderate `protobuf:"varint,6,opt,name=coderate,proto3,enum=helium.radio.Coderate" json:"coderate,omitempty"`
	// Invert signal polarity for orthogonal downlinks.
	InvertPolarity bool `protobuf:"varint,7,opt,name=invert_polarity,json=invertPolarity,proto3" json:"invert_polarity,omitempty"`
	// Do not send a CRC in the packet.
	OmitCrc bool `protobuf:"varint,8,opt,name=omit_crc,json=omitCrc,proto3" json:"omit_crc,omitempty"`
	// Enable implicit header mode.
	ImplicitHeader bool `protobuf:"varint,9,opt,name=implicit_header,json=implicitHeader,proto3" json:"implicit_header,omitempty"`
	// Arbitrary user-defined payload to transmit.
	Payload []byte `protobuf:"bytes,10,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *RadioTxReq) Reset() {
	*x = RadioTxReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_radio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioTxReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioTxReq) ProtoMessage() {}

func (x *RadioTxReq) ProtoReflect() protoreflect.Message {
	mi := &file_radio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioTxReq.ProtoReflect.Descriptor instead.
func (*RadioTxReq) Descriptor() ([]byte, []int) {
	return file_radio_proto_rawDescGZIP(), []int{2}
}

func (x *RadioTxReq) GetFreq() uint32 {
	if x != nil {
		return x.Freq
	}
	return 0
}

func (x *RadioTxReq) GetRadio() Radio {
	if x != nil {
		return x.Radio
	}
	return Radio_R0
}

func (x *RadioTxReq) GetPower() int32 {
	if x != nil {
		return x.Power
	}
	return 0
}

func (x *RadioTxReq) GetBandwidth() Bandwidth {
	if x != nil {
		return x.Bandwidth
	}
	return Bandwidth_BW_UNDEFINED
}

func (x *RadioTxReq) GetSpreading() Spreading {
	if x != nil {
		return x.Spreading
	}
	return Spreading_SF_UNDEFINED
}

func (x *RadioTxReq) GetCoderate() Coderate {
	if x != nil {
		return x.Coderate
	}
	return Coderate_CR_UNDEFINED
}

func (x *RadioTxReq) GetInvertPolarity() bool {
	if x != nil {
		return x.InvertPolarity
	}
	return false
}

func (x *RadioTxReq) GetOmitCrc() bool {
	if x != nil {
		return x.OmitCrc
	}
	return false
}

func (x *RadioTxReq) GetImplicitHeader() bool {
	if x != nil {
		return x.ImplicitHeader
	}
	return false
}

func (x *RadioTxReq) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type RadioRxPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Center frequency of the channel this packet was received on.
	Freq uint32 `protobuf:"varint,1,opt,name=freq,proto3" json:"freq,omitempty"`
	// Channel this packet packet was received on.
	IfChain uint32 `protobuf:"varint,2,opt,name=if_chain,json=ifChain,proto3" json:"if_chain,omitempty"`
	// Status of CRC check.
	CrcCheck bool `protobuf:"varint,3,opt,name=crc_check,json=crcCheck,proto3" json:"crc_check,omitempty"`
	// 1uS-resolution timestamp derived from concentrator's internal counter.
	Timestamp uint64 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// RF chain this packet was received on.
	Radio Radio `protobuf:"varint,5,opt,name=radio,proto3,enum=helium.radio.Radio" json:"radio,omitempty"`
	// Modulation bandwidth.
	Bandwidth Bandwidth `protobuf:"varint,6,opt,name=bandwidth,proto3,enum=helium.radio.Bandwidth" json:"bandwidth,omitempty"`
	// Spreading factor of this packet.
	Spreading Spreading `protobuf:"varint,7,opt,name=spreading,proto3,enum=helium.radio.Spreading" json:"spreading,omitempty"`
	// Error Correcting Code rate of this packet.
	Coderate Coderate `protobuf:"varint,8,opt,name=coderate,proto3,enum=helium.radio.Coderate" json:"coderate,omitempty"`
	// Average packet RSSI in dB.
	Rssi float32 `protobuf:"fixed32,9,opt,name=rssi,proto3" json:"rssi,omitempty"`
	// Average packet SNR, in dB.
	Snr float32 `protobuf:"fixed32,10,opt,name=snr,proto3" json:"snr,omitempty"`
	// This packet's payload.
	Payload []byte `protobuf:"bytes,11,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *RadioRxPacket) Reset() {
	*x = RadioRxPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_radio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioRxPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioRxPacket) ProtoMessage() {}

func (x *RadioRxPacket) ProtoReflect() protoreflect.Message {
	mi := &file_radio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioRxPacket.ProtoReflect.Descriptor instead.
func (*RadioRxPacket) Descriptor() ([]byte, []int) {
	return file_radio_proto_rawDescGZIP(), []int{3}
}

func (x *RadioRxPacket) GetFreq() uint32 {
	if x != nil {
		return x.Freq
	}
	return 0
}

func (x *RadioRxPacket) GetIfChain() uint32 {
	if x != nil {
		return x.IfChain
	}
	return 0
}

func (x *RadioRxPacket) GetCrcCheck() bool {
	if x != nil {
		return x.CrcCheck
	}
	return false
}

func (x *RadioRxPacket) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *RadioRxPacket) GetRadio() Radio {
	if x != nil {
		return x.Radio
	}
	return Radio_R0
}

func (x *RadioRxPacket) GetBandwidth() Bandwidth {
	if x != nil {
		return x.Bandwidth
	}
	return Bandwidth_BW_UNDEFINED
}

func (x *RadioRxPacket) GetSpreading() Spreading {
	if x != nil {
		return x.Spreading
	}
	return Spreading_SF_UNDEFINED
}

func (x *RadioRxPacket) GetCoderate() Coderate {
	if x != nil {
		return x.Coderate
	}
	return Coderate_CR_UNDEFINED
}

func (x *RadioRxPacket) GetRssi() float32 {
	if x != nil {
		return x.Rssi
	}
	return 0
}

func (x *RadioRxPacket) GetSnr() float32 {
	if x != nil {
		return x.Snr
	}
	return 0
}

func (x *RadioRxPacket) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type RadioTxResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RadioTxResp) Reset() {
	*x = RadioTxResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_radio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioTxResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioTxResp) ProtoMessage() {}

func (x *RadioTxResp) ProtoReflect() protoreflect.Message {
	mi := &file_radio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioTxResp.ProtoReflect.Descriptor instead.
func (*RadioTxResp) Descriptor() ([]byte, []int) {
	return file_radio_proto_rawDescGZIP(), []int{4}
}

func (x *RadioTxResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_radio_proto protoreflect.FileDescriptor

var file_radio_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x68,
	0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x22, 0x4e, 0x0a, 0x08, 0x52,
	0x61, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x72, 0x61, 0x64,
	0x69, 0x6f, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x54, 0x78, 0x52, 0x65, 0x71, 0x48, 0x00, 0x52,
	0x02, 0x74, 0x78, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0xab, 0x01, 0x0a, 0x09,
	0x52, 0x61, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x02, 0x74, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x72,
	0x61, 0x64, 0x69, 0x6f, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x48, 0x00, 0x52, 0x02, 0x74, 0x78, 0x12, 0x3a, 0x0a, 0x09, 0x72, 0x78, 0x5f, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x69,
	0x75, 0x6d, 0x2e, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x52, 0x78,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x00, 0x52, 0x08, 0x72, 0x78, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x73, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x72, 0x73, 0x65, 0x45, 0x72,
	0x72, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x8a, 0x03, 0x0a, 0x0a, 0x52, 0x61,
	0x64, 0x69, 0x6f, 0x54, 0x78, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x65, 0x71,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x72, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x05,
	0x72, 0x61, 0x64, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x68, 0x65,
	0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f,
	0x52, 0x05, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x35, 0x0a,
	0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x2e,
	0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x52, 0x09, 0x62, 0x61, 0x6e, 0x64, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x12, 0x35, 0x0a, 0x09, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d,
	0x2e, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x2e, 0x53, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x09, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x32, 0x0a, 0x08, 0x63,
	0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x50, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6d, 0x69, 0x74,
	0x5f, 0x63, 0x72, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6f, 0x6d, 0x69, 0x74,
	0x43, 0x72, 0x63, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x6d,
	0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x86, 0x03, 0x0a, 0x0d, 0x52, 0x61, 0x64, 0x69, 0x6f,
	0x52, 0x78, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x65, 0x71,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x72, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x66, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x69, 0x66, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x63, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x72, 0x63, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x29, 0x0a, 0x05, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x72, 0x61, 0x64, 0x69, 0x6f,
	0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x52, 0x05, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x12, 0x35, 0x0a,
	0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x2e,
	0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x52, 0x09, 0x62, 0x61, 0x6e, 0x64, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x12, 0x35, 0x0a, 0x09, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d,
	0x2e, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x2e, 0x53, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x09, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x32, 0x0a, 0x08, 0x63,
	0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x73, 0x73, 0x69, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72,
	0x73, 0x73, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6e, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x73, 0x6e, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x27, 0x0a, 0x0b, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2a, 0x17, 0x0a, 0x05, 0x52, 0x61, 0x64, 0x69,
	0x6f, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x30, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x31, 0x10,
	0x01, 0x2a, 0x56, 0x0a, 0x09, 0x53, 0x70, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x46, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x53, 0x46, 0x37, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x46, 0x38,
	0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x46, 0x39, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x53,
	0x46, 0x31, 0x30, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x46, 0x31, 0x31, 0x10, 0x05, 0x12,
	0x08, 0x0a, 0x04, 0x53, 0x46, 0x31, 0x32, 0x10, 0x06, 0x2a, 0x82, 0x01, 0x0a, 0x09, 0x42, 0x61,
	0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x57, 0x5f, 0x55, 0x4e,
	0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x57, 0x37,
	0x5f, 0x38, 0x6b, 0x48, 0x7a, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x57, 0x31, 0x35, 0x5f,
	0x36, 0x6b, 0x48, 0x7a, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x57, 0x33, 0x31, 0x5f, 0x32,
	0x6b, 0x48, 0x7a, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x57, 0x36, 0x32, 0x5f, 0x35, 0x6b,
	0x48, 0x7a, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x57, 0x31, 0x32, 0x35, 0x6b, 0x48, 0x7a,
	0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x57, 0x32, 0x35, 0x30, 0x6b, 0x48, 0x7a, 0x10, 0x06,
	0x12, 0x0c, 0x0a, 0x08, 0x42, 0x57, 0x35, 0x30, 0x30, 0x6b, 0x48, 0x7a, 0x10, 0x07, 0x2a, 0x48,
	0x0a, 0x08, 0x43, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x52,
	0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x43, 0x52, 0x34, 0x5f, 0x35, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x52, 0x34, 0x5f, 0x36,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x52, 0x34, 0x5f, 0x37, 0x10, 0x03, 0x12, 0x09, 0x0a,
	0x05, 0x43, 0x52, 0x34, 0x5f, 0x38, 0x10, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_radio_proto_rawDescOnce sync.Once
	file_radio_proto_rawDescData = file_radio_proto_rawDesc
)

func file_radio_proto_rawDescGZIP() []byte {
	file_radio_proto_rawDescOnce.Do(func() {
		file_radio_proto_rawDescData = protoimpl.X.CompressGZIP(file_radio_proto_rawDescData)
	})
	return file_radio_proto_rawDescData
}

var file_radio_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_radio_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_radio_proto_goTypes = []interface{}{
	(Radio)(0),            // 0: helium.radio.Radio
	(Spreading)(0),        // 1: helium.radio.Spreading
	(Bandwidth)(0),        // 2: helium.radio.Bandwidth
	(Coderate)(0),         // 3: helium.radio.Coderate
	(*RadioReq)(nil),      // 4: helium.radio.RadioReq
	(*RadioResp)(nil),     // 5: helium.radio.RadioResp
	(*RadioTxReq)(nil),    // 6: helium.radio.RadioTxReq
	(*RadioRxPacket)(nil), // 7: helium.radio.RadioRxPacket
	(*RadioTxResp)(nil),   // 8: helium.radio.RadioTxResp
}
var file_radio_proto_depIdxs = []int32{
	6,  // 0: helium.radio.RadioReq.tx:type_name -> helium.radio.RadioTxReq
	8,  // 1: helium.radio.RadioResp.tx:type_name -> helium.radio.RadioTxResp
	7,  // 2: helium.radio.RadioResp.rx_packet:type_name -> helium.radio.RadioRxPacket
	0,  // 3: helium.radio.RadioTxReq.radio:type_name -> helium.radio.Radio
	2,  // 4: helium.radio.RadioTxReq.bandwidth:type_name -> helium.radio.Bandwidth
	1,  // 5: helium.radio.RadioTxReq.spreading:type_name -> helium.radio.Spreading
	3,  // 6: helium.radio.RadioTxReq.coderate:type_name -> helium.radio.Coderate
	0,  // 7: helium.radio.RadioRxPacket.radio:type_name -> helium.radio.Radio
	2,  // 8: helium.radio.RadioRxPacket.bandwidth:type_name -> helium.radio.Bandwidth
	1,  // 9: helium.radio.RadioRxPacket.spreading:type_name -> helium.radio.Spreading
	3,  // 10: helium.radio.RadioRxPacket.coderate:type_name -> helium.radio.Coderate
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_radio_proto_init() }
func file_radio_proto_init() {
	if File_radio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_radio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioReq); i {
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
		file_radio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioResp); i {
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
		file_radio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioTxReq); i {
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
		file_radio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioRxPacket); i {
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
		file_radio_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioTxResp); i {
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
	file_radio_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RadioReq_Tx)(nil),
	}
	file_radio_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*RadioResp_Tx)(nil),
		(*RadioResp_RxPacket)(nil),
		(*RadioResp_ParseErr)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_radio_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_radio_proto_goTypes,
		DependencyIndexes: file_radio_proto_depIdxs,
		EnumInfos:         file_radio_proto_enumTypes,
		MessageInfos:      file_radio_proto_msgTypes,
	}.Build()
	File_radio_proto = out.File
	file_radio_proto_rawDesc = nil
	file_radio_proto_goTypes = nil
	file_radio_proto_depIdxs = nil
}
