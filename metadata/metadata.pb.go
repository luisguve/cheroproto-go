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
	DataKey string `protobuf:"bytes,4,opt,name=data_key,json=dataKey,proto3" json:"data_key,omitempty"`
	// Diff holds the accumulated time difference between interactions.
	Diff                 int64    `protobuf:"varint,5,opt,name=diff,proto3" json:"diff,omitempty"`
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

func (m *Content) GetDiff() int64 {
	if m != nil {
		return m.Diff
	}
	return 0
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
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xcd, 0x36, 0xed, 0x96, 0xfd, 0x71, 0xe6, 0xa9, 0x0e, 0xc4, 0xb2, 0x07, 0x29, 0x53,
	0x5b, 0xd0, 0x67, 0x5f, 0x94, 0x21, 0xa2, 0xa8, 0xc4, 0x89, 0x20, 0x42, 0xc9, 0xda, 0xbb, 0x2c,
	0xd8, 0x36, 0xa3, 0xbd, 0x1d, 0xec, 0xf3, 0xf9, 0xc5, 0x24, 0xfd, 0x23, 0xf8, 0x76, 0x73, 0xf8,
	0xdd, 0x9c, 0x73, 0x2e, 0x1d, 0x25, 0x80, 0x22, 0x12, 0x28, 0xbc, 0x4d, 0xa6, 0x51, 0xb3, 0x6e,
	0xf3, 0x9e, 0x9c, 0x4a, 0xad, 0x65, 0x0c, 0x7e, 0xa9, 0x2f, 0x8b, 0x95, 0x8f, 0x2a, 0x81, 0x1c,
	0x45, 0xb2, 0xa9, 0xd0, 0xe9, 0x0f, 0xa1, 0xd6, 0x9d, 0x4e, 0x11, 0x52, 0x64, 0x53, 0x3a, 0x50,
	0x29, 0x42, 0x26, 0x42, 0x54, 0x3a, 0xcd, 0x6d, 0xe2, 0x10, 0x77, 0xc8, 0xff, 0x69, 0xec, 0x86,
	0x0e, 0x62, 0x91, 0x63, 0x50, 0x6c, 0x22, 0x81, 0x10, 0xd9, 0x2d, 0x87, 0xb8, 0xfd, 0xab, 0x89,
	0x57, 0xf9, 0x78, 0x8d, 0x8f, 0xb7, 0x68, 0x7c, 0x78, 0xdf, 0xf0, 0xef, 0x15, 0xce, 0xce, 0xe8,
	0xa1, 0xd8, 0xca, 0x7a, 0x3b, 0x30, 0x61, 0xec, 0xb6, 0x43, 0x5c, 0xc2, 0x87, 0x62, 0x2b, 0x2b,
	0xc8, 0x6c, 0xb2, 0x63, 0xda, 0x35, 0xf9, 0x83, 0x6f, 0xd8, 0xd9, 0x1d, 0x87, 0xb8, 0x3d, 0x6e,
	0x99, 0xf7, 0x23, 0xec, 0x18, 0xa3, 0x9d, 0x48, 0xad, 0x56, 0xf6, 0xbe, 0x43, 0xdc, 0x36, 0x2f,
	0xe7, 0xe9, 0x17, 0x1d, 0xdd, 0x43, 0x0a, 0x99, 0x88, 0x9b, 0x2e, 0x27, 0x94, 0xe6, 0x50, 0x66,
	0x0e, 0x54, 0x54, 0x36, 0xe9, 0xf1, 0x5e, 0xad, 0x3c, 0x44, 0xec, 0x9c, 0x5a, 0x61, 0x45, 0xd6,
	0x0d, 0x8e, 0xbc, 0xbf, 0x1b, 0xd6, 0x5f, 0xf0, 0x86, 0x98, 0xcd, 0xe8, 0xb0, 0xd6, 0xde, 0x50,
	0x60, 0x91, 0x33, 0x8b, 0xb6, 0x9f, 0xe7, 0x1f, 0xe3, 0x3d, 0x33, 0xf0, 0xf9, 0xd3, 0x98, 0x98,
	0x61, 0xf1, 0xf2, 0x3a, 0x6e, 0xdd, 0x5e, 0x7c, 0xce, 0xa4, 0xc2, 0x75, 0xb1, 0xf4, 0x42, 0x9d,
	0xf8, 0x71, 0xa1, 0x72, 0x59, 0x6c, 0xc1, 0x0f, 0xd7, 0x90, 0xe9, 0xf2, 0x38, 0x97, 0x52, 0xfb,
	0x8d, 0xd5, 0xf2, 0xa0, 0x94, 0xae, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x11, 0x11, 0xbd, 0x28,
	0xc1, 0x01, 0x00, 0x00,
}
