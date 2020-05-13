// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/trace-common.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SpanTypeV1 int32

const (
	SpanTypeV1_Entry SpanTypeV1 = 0
	SpanTypeV1_Exit  SpanTypeV1 = 1
	SpanTypeV1_Local SpanTypeV1 = 2
)

var SpanTypeV1_name = map[int32]string{
	0: "Entry",
	1: "Exit",
	2: "Local",
}

var SpanTypeV1_value = map[string]int32{
	"Entry": 0,
	"Exit":  1,
	"Local": 2,
}

func (x SpanTypeV1) String() string {
	return proto.EnumName(SpanTypeV1_name, int32(x))
}

func (SpanTypeV1) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{0}
}

type RefTypeV1 int32

const (
	RefTypeV1_CrossProcess RefTypeV1 = 0
	RefTypeV1_CrossThread  RefTypeV1 = 1
)

var RefTypeV1_name = map[int32]string{
	0: "CrossProcess",
	1: "CrossThread",
}

var RefTypeV1_value = map[string]int32{
	"CrossProcess": 0,
	"CrossThread":  1,
}

func (x RefTypeV1) String() string {
	return proto.EnumName(RefTypeV1_name, int32(x))
}

func (RefTypeV1) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{1}
}

type SpanLayerV1 int32

const (
	SpanLayerV1_Unknown      SpanLayerV1 = 0
	SpanLayerV1_Database     SpanLayerV1 = 1
	SpanLayerV1_RPCFramework SpanLayerV1 = 2
	SpanLayerV1_Http         SpanLayerV1 = 3
	SpanLayerV1_MQ           SpanLayerV1 = 4
	SpanLayerV1_Cache        SpanLayerV1 = 5
)

var SpanLayerV1_name = map[int32]string{
	0: "Unknown",
	1: "Database",
	2: "RPCFramework",
	3: "Http",
	4: "MQ",
	5: "Cache",
}

var SpanLayerV1_value = map[string]int32{
	"Unknown":      0,
	"Database":     1,
	"RPCFramework": 2,
	"Http":         3,
	"MQ":           4,
	"Cache":        5,
}

func (x SpanLayerV1) String() string {
	return proto.EnumName(SpanLayerV1_name, int32(x))
}

func (SpanLayerV1) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{2}
}

type UpstreamSegment struct {
	GlobalTraceIds       []*UniqueId `protobuf:"bytes,1,rep,name=globalTraceIds,proto3" json:"globalTraceIds,omitempty"`
	Segment              []byte      `protobuf:"bytes,2,opt,name=segment,proto3" json:"segment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpstreamSegment) Reset()         { *m = UpstreamSegment{} }
func (m *UpstreamSegment) String() string { return proto.CompactTextString(m) }
func (*UpstreamSegment) ProtoMessage()    {}
func (*UpstreamSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{0}
}

func (m *UpstreamSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSegment.Unmarshal(m, b)
}
func (m *UpstreamSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSegment.Marshal(b, m, deterministic)
}
func (m *UpstreamSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSegment.Merge(m, src)
}
func (m *UpstreamSegment) XXX_Size() int {
	return xxx_messageInfo_UpstreamSegment.Size(m)
}
func (m *UpstreamSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSegment.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSegment proto.InternalMessageInfo

func (m *UpstreamSegment) GetGlobalTraceIds() []*UniqueId {
	if m != nil {
		return m.GlobalTraceIds
	}
	return nil
}

func (m *UpstreamSegment) GetSegment() []byte {
	if m != nil {
		return m.Segment
	}
	return nil
}

type UniqueId struct {
	IdParts              []int64  `protobuf:"varint,1,rep,packed,name=idParts,proto3" json:"idParts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UniqueId) Reset()         { *m = UniqueId{} }
func (m *UniqueId) String() string { return proto.CompactTextString(m) }
func (*UniqueId) ProtoMessage()    {}
func (*UniqueId) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{1}
}

func (m *UniqueId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UniqueId.Unmarshal(m, b)
}
func (m *UniqueId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UniqueId.Marshal(b, m, deterministic)
}
func (m *UniqueId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UniqueId.Merge(m, src)
}
func (m *UniqueId) XXX_Size() int {
	return xxx_messageInfo_UniqueId.Size(m)
}
func (m *UniqueId) XXX_DiscardUnknown() {
	xxx_messageInfo_UniqueId.DiscardUnknown(m)
}

var xxx_messageInfo_UniqueId proto.InternalMessageInfo

func (m *UniqueId) GetIdParts() []int64 {
	if m != nil {
		return m.IdParts
	}
	return nil
}

func init() {
	proto.RegisterEnum("SpanTypeV1", SpanTypeV1_name, SpanTypeV1_value)
	proto.RegisterEnum("RefTypeV1", RefTypeV1_name, RefTypeV1_value)
	proto.RegisterEnum("SpanLayerV1", SpanLayerV1_name, SpanLayerV1_value)
	proto.RegisterType((*UpstreamSegment)(nil), "UpstreamSegment")
	proto.RegisterType((*UniqueId)(nil), "UniqueId")
}

func init() { proto.RegisterFile("common/trace-common.proto", fileDescriptor_f5d7d2e503402948) }

var fileDescriptor_f5d7d2e503402948 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x9b, 0x74, 0xff, 0xb4, 0x27, 0xc5, 0x1d, 0xe6, 0xaa, 0x7a, 0xb5, 0x2c, 0x7b, 0xb1,
	0x14, 0x9d, 0x18, 0x7d, 0x03, 0xeb, 0x8a, 0x0b, 0xab, 0xc4, 0xb4, 0xdd, 0x82, 0x17, 0xc2, 0x69,
	0x72, 0x4c, 0x43, 0x92, 0x99, 0x38, 0x33, 0x35, 0xe6, 0x95, 0x7c, 0x4a, 0x99, 0xa6, 0x41, 0xd8,
	0xcb, 0xef, 0xe3, 0x77, 0x7e, 0x1c, 0xf8, 0xe0, 0x65, 0xaa, 0xea, 0x5a, 0xc9, 0xd0, 0x6a, 0x4c,
	0xe9, 0x4d, 0x1f, 0x44, 0xa3, 0x95, 0x55, 0x37, 0x3f, 0xe0, 0x6a, 0xd3, 0x18, 0xab, 0x09, 0xeb,
	0x15, 0xe5, 0x35, 0x49, 0xcb, 0x23, 0x78, 0x91, 0x57, 0x6a, 0x87, 0xd5, 0xda, 0xe1, 0x0f, 0x99,
	0x99, 0x7b, 0xd7, 0xe3, 0xbb, 0xe0, 0xdd, 0x54, 0x6c, 0x64, 0xf1, 0xeb, 0x40, 0x0f, 0x59, 0xf2,
	0x0c, 0xe0, 0x73, 0xb8, 0x34, 0xfd, 0xf5, 0xdc, 0xbf, 0xf6, 0xee, 0x66, 0xc9, 0x10, 0x6f, 0x6e,
	0x61, 0x32, 0x5c, 0x39, 0xaa, 0xc8, 0x62, 0xd4, 0xb6, 0x37, 0x8e, 0x93, 0x21, 0x2e, 0x5e, 0x03,
	0xac, 0x1a, 0x94, 0xeb, 0xae, 0xa1, 0xa7, 0x88, 0x4f, 0xe1, 0xfc, 0x5e, 0x5a, 0xdd, 0xb1, 0x11,
	0x9f, 0xc0, 0xd9, 0xfd, 0x9f, 0xc2, 0x32, 0xcf, 0x95, 0x8f, 0x2a, 0xc5, 0x8a, 0xf9, 0x0b, 0x01,
	0xd3, 0x84, 0x7e, 0x9e, 0x60, 0x06, 0xb3, 0xa5, 0x56, 0xc6, 0xc4, 0x5a, 0xa5, 0x64, 0x0c, 0x1b,
	0xf1, 0x2b, 0x08, 0x8e, 0xcd, 0x7a, 0xaf, 0x09, 0x33, 0xe6, 0x2d, 0xb6, 0x10, 0x38, 0xfb, 0x23,
	0x76, 0xa4, 0x9f, 0x22, 0x1e, 0xc0, 0xe5, 0x46, 0x96, 0x52, 0xb5, 0x92, 0x8d, 0xf8, 0x0c, 0x26,
	0x1f, 0xd1, 0xe2, 0x0e, 0x0d, 0x31, 0xcf, 0xc9, 0x92, 0x78, 0xf9, 0x49, 0x63, 0x4d, 0xad, 0xd2,
	0x25, 0xf3, 0xdd, 0x03, 0x9f, 0xad, 0x6d, 0xd8, 0x98, 0x5f, 0x80, 0xff, 0xe5, 0x1b, 0x3b, 0x73,
	0x8f, 0x2c, 0x31, 0xdd, 0x13, 0x3b, 0xff, 0xd0, 0xc2, 0x5b, 0xa5, 0x73, 0x81, 0x8d, 0xcb, 0xc2,
	0x94, 0x5d, 0x8b, 0x55, 0x59, 0x48, 0xd7, 0xd4, 0x42, 0x92, 0x75, 0x0e, 0x51, 0xa1, 0xcc, 0x0f,
	0x98, 0x93, 0xc0, 0x9c, 0xa4, 0x8d, 0xbd, 0xef, 0xb7, 0xff, 0xc1, 0xf0, 0x04, 0x85, 0x03, 0x14,
	0x1e, 0xa1, 0xf0, 0x77, 0xf4, 0xd7, 0x7f, 0xb5, 0x2a, 0xbb, 0xed, 0xc9, 0xf7, 0xb5, 0xc7, 0x62,
	0x37, 0x59, 0xaa, 0xaa, 0xdd, 0xc5, 0x71, 0xbc, 0xf7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x38,
	0xd4, 0x7a, 0xc7, 0xd9, 0x01, 0x00, 0x00,
}