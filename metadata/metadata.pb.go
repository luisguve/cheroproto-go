// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata.proto

// package metadata defines message types for describing content

package metadata

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// ContentStatus defines the different categories that fits content at a
// given time, trying to describe the quality of the content.
type ContentStatus int32

const (
	ContentStatus_NEW ContentStatus = 0
	ContentStatus_REL ContentStatus = 1
	ContentStatus_TOP ContentStatus = 2
)

var ContentStatus_name = map[int32]string{
	0: "NEW",
	1: "REL",
	2: "TOP",
}

var ContentStatus_value = map[string]int32{
	"NEW": 0,
	"REL": 1,
	"TOP": 2,
}

func (x ContentStatus) String() string {
	return proto.EnumName(ContentStatus_name, int32(x))
}

func (ContentStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56d9f74966f40d04, []int{0}
}

// Content holds the most updated information about the interactions with
// the content.
type Content struct {
	// Are considered interactions with the content the following actions:
	// upvote, comment, comment upvote, subcomment and subcomment upvote.
	Interactions uint32 `protobuf:"varint,1,opt,name=interactions,proto3" json:"interactions,omitempty"`
	// LastUpdated indicates when it was the last time this content
	// received an interaction.
	LastUpdated *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	// AvgUpdateTime indicates the average difference in time between
	// interactions.
	AvgUpdateTime float64 `protobuf:"fixed64,3,opt,name=avg_update_time,json=avgUpdateTime,proto3" json:"avg_update_time,omitempty"`
	// DataKey holds the key referencing the data.
	DataKey              string   `protobuf:"bytes,4,opt,name=data_key,json=dataKey,proto3" json:"data_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d9f74966f40d04, []int{0}
}

func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetInteractions() uint32 {
	if m != nil {
		return m.Interactions
	}
	return 0
}

func (m *Content) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

func (m *Content) GetAvgUpdateTime() float64 {
	if m != nil {
		return m.AvgUpdateTime
	}
	return 0
}

func (m *Content) GetDataKey() string {
	if m != nil {
		return m.DataKey
	}
	return ""
}

// GeneralContent holds the most updated information about the interactions
// with the content, along with the section id it belongs to.
type GeneralContent struct {
	SectionId            string   `protobuf:"bytes,1,opt,name=section_id,json=sectionId,proto3" json:"section_id,omitempty"`
	Content              *Content `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneralContent) Reset()         { *m = GeneralContent{} }
func (m *GeneralContent) String() string { return proto.CompactTextString(m) }
func (*GeneralContent) ProtoMessage()    {}
func (*GeneralContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_56d9f74966f40d04, []int{1}
}

func (m *GeneralContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralContent.Unmarshal(m, b)
}
func (m *GeneralContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralContent.Marshal(b, m, deterministic)
}
func (m *GeneralContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralContent.Merge(m, src)
}
func (m *GeneralContent) XXX_Size() int {
	return xxx_messageInfo_GeneralContent.Size(m)
}
func (m *GeneralContent) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralContent.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralContent proto.InternalMessageInfo

func (m *GeneralContent) GetSectionId() string {
	if m != nil {
		return m.SectionId
	}
	return ""
}

func (m *GeneralContent) GetContent() *Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterEnum("metadata.ContentStatus", ContentStatus_name, ContentStatus_value)
	proto.RegisterType((*Content)(nil), "metadata.Content")
	proto.RegisterType((*GeneralContent)(nil), "metadata.GeneralContent")
}

func init() {
	proto.RegisterFile("metadata.proto", fileDescriptor_56d9f74966f40d04)
}

var fileDescriptor_56d9f74966f40d04 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xdd, 0x4a, 0xf3, 0x40,
	0x10, 0x86, 0xbf, 0x6d, 0x3f, 0x4c, 0x3b, 0xfd, 0xb1, 0xee, 0x51, 0x2c, 0x88, 0xa1, 0x07, 0x12,
	0xaa, 0x26, 0xa0, 0xc7, 0x9e, 0x28, 0x45, 0x44, 0x51, 0x59, 0x2b, 0x82, 0x08, 0x61, 0x9b, 0x8c,
	0xdb, 0x60, 0x92, 0x2d, 0xc9, 0xa4, 0xd0, 0xab, 0xf2, 0x16, 0x65, 0xf3, 0x23, 0x78, 0xb6, 0xf3,
	0xf2, 0xcc, 0xbe, 0x0f, 0x03, 0xe3, 0x14, 0x49, 0x46, 0x92, 0xa4, 0xb7, 0xc9, 0x35, 0x69, 0xde,
	0x6b, 0xe7, 0xe9, 0xb1, 0xd2, 0x5a, 0x25, 0xe8, 0x57, 0xf9, 0xaa, 0xfc, 0xf4, 0x29, 0x4e, 0xb1,
	0x20, 0x99, 0x6e, 0x6a, 0x74, 0xf6, 0xcd, 0xc0, 0xba, 0xd1, 0x19, 0x61, 0x46, 0x7c, 0x06, 0xc3,
	0x38, 0x23, 0xcc, 0x65, 0x48, 0xb1, 0xce, 0x0a, 0x9b, 0x39, 0xcc, 0x1d, 0x89, 0x3f, 0x19, 0xbf,
	0x82, 0x61, 0x22, 0x0b, 0x0a, 0xca, 0x4d, 0x24, 0x09, 0x23, 0xbb, 0xe3, 0x30, 0x77, 0x70, 0x31,
	0xf5, 0xea, 0x1e, 0xaf, 0xed, 0xf1, 0x96, 0x6d, 0x8f, 0x18, 0x18, 0xfe, 0xb5, 0xc6, 0xf9, 0x09,
	0xec, 0xcb, 0xad, 0x6a, 0xb6, 0x03, 0x23, 0x63, 0x77, 0x1d, 0xe6, 0x32, 0x31, 0x92, 0x5b, 0x55,
	0x43, 0x66, 0x93, 0x1f, 0x42, 0xcf, 0xf8, 0x07, 0x5f, 0xb8, 0xb3, 0xff, 0x3b, 0xcc, 0xed, 0x0b,
	0xcb, 0xcc, 0xf7, 0xb8, 0x9b, 0x7d, 0xc0, 0xf8, 0x16, 0x33, 0xcc, 0x65, 0xd2, 0x7a, 0x1f, 0x01,
	0x14, 0x58, 0xf9, 0x05, 0x71, 0x54, 0x59, 0xf7, 0x45, 0xbf, 0x49, 0xee, 0x22, 0x7e, 0x0a, 0x56,
	0x58, 0x93, 0x8d, 0xed, 0x81, 0xf7, 0x7b, 0xaf, 0xe6, 0x0b, 0xd1, 0x12, 0xf3, 0x39, 0x8c, 0x9a,
	0xec, 0x85, 0x24, 0x95, 0x05, 0xb7, 0xa0, 0xfb, 0xb8, 0x78, 0x9b, 0xfc, 0x33, 0x0f, 0xb1, 0x78,
	0x98, 0x30, 0xf3, 0x58, 0x3e, 0x3d, 0x4f, 0x3a, 0xd7, 0x67, 0xef, 0x73, 0x15, 0xd3, 0xba, 0x5c,
	0x79, 0xa1, 0x4e, 0xfd, 0xa4, 0x8c, 0x0b, 0x55, 0x6e, 0xd1, 0x0f, 0xd7, 0x98, 0xeb, 0xea, 0x10,
	0xe7, 0x4a, 0xfb, 0x6d, 0xd5, 0x6a, 0xaf, 0x8a, 0x2e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9e,
	0xfe, 0x1e, 0xe8, 0xad, 0x01, 0x00, 0x00,
}
