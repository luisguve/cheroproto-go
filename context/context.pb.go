// Code generated by protoc-gen-go. DO NOT EDIT.
// source: context.proto

// package context defines different context types that hold information
// about the place at which an event occurs or where content should be
// stored in and retrieved from.

package context

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

type Context struct {
	// Types that are valid to be assigned to Ctx:
	//	*Context_SectionCtx
	//	*Context_ThreadCtx
	//	*Context_CommentCtx
	//	*Context_SubcommentCtx
	Ctx                  isContext_Ctx `protobuf_oneof:"ctx"`
	SectionId            string        `protobuf:"bytes,5,opt,name=section_id,json=sectionId,proto3" json:"section_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Context) Reset()         { *m = Context{} }
func (m *Context) String() string { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()    {}
func (*Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_b64063be2fc89884, []int{0}
}

func (m *Context) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Context.Unmarshal(m, b)
}
func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Context.Marshal(b, m, deterministic)
}
func (m *Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Context.Merge(m, src)
}
func (m *Context) XXX_Size() int {
	return xxx_messageInfo_Context.Size(m)
}
func (m *Context) XXX_DiscardUnknown() {
	xxx_messageInfo_Context.DiscardUnknown(m)
}

var xxx_messageInfo_Context proto.InternalMessageInfo

type isContext_Ctx interface {
	isContext_Ctx()
}

type Context_SectionCtx struct {
	SectionCtx *Section `protobuf:"bytes,1,opt,name=section_ctx,json=sectionCtx,proto3,oneof"`
}

type Context_ThreadCtx struct {
	ThreadCtx *Thread `protobuf:"bytes,2,opt,name=thread_ctx,json=threadCtx,proto3,oneof"`
}

type Context_CommentCtx struct {
	CommentCtx *Comment `protobuf:"bytes,3,opt,name=comment_ctx,json=commentCtx,proto3,oneof"`
}

type Context_SubcommentCtx struct {
	SubcommentCtx *Subcomment `protobuf:"bytes,4,opt,name=subcomment_ctx,json=subcommentCtx,proto3,oneof"`
}

func (*Context_SectionCtx) isContext_Ctx() {}

func (*Context_ThreadCtx) isContext_Ctx() {}

func (*Context_CommentCtx) isContext_Ctx() {}

func (*Context_SubcommentCtx) isContext_Ctx() {}

func (m *Context) GetCtx() isContext_Ctx {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Context) GetSectionCtx() *Section {
	if x, ok := m.GetCtx().(*Context_SectionCtx); ok {
		return x.SectionCtx
	}
	return nil
}

func (m *Context) GetThreadCtx() *Thread {
	if x, ok := m.GetCtx().(*Context_ThreadCtx); ok {
		return x.ThreadCtx
	}
	return nil
}

func (m *Context) GetCommentCtx() *Comment {
	if x, ok := m.GetCtx().(*Context_CommentCtx); ok {
		return x.CommentCtx
	}
	return nil
}

func (m *Context) GetSubcommentCtx() *Subcomment {
	if x, ok := m.GetCtx().(*Context_SubcommentCtx); ok {
		return x.SubcommentCtx
	}
	return nil
}

func (m *Context) GetSectionId() string {
	if m != nil {
		return m.SectionId
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Context) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Context_SectionCtx)(nil),
		(*Context_ThreadCtx)(nil),
		(*Context_CommentCtx)(nil),
		(*Context_SubcommentCtx)(nil),
	}
}

// A Section is the highest level and just has the name of the section.
type Section struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Section) Reset()         { *m = Section{} }
func (m *Section) String() string { return proto.CompactTextString(m) }
func (*Section) ProtoMessage()    {}
func (*Section) Descriptor() ([]byte, []int) {
	return fileDescriptor_b64063be2fc89884, []int{1}
}

func (m *Section) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Section.Unmarshal(m, b)
}
func (m *Section) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Section.Marshal(b, m, deterministic)
}
func (m *Section) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Section.Merge(m, src)
}
func (m *Section) XXX_Size() int {
	return xxx_messageInfo_Section.Size(m)
}
func (m *Section) XXX_DiscardUnknown() {
	xxx_messageInfo_Section.DiscardUnknown(m)
}

var xxx_messageInfo_Section proto.InternalMessageInfo

func (m *Section) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// A Thread is inside of a section and has a thread identifier unique to
// that section.
type Thread struct {
	SectionCtx           *Section `protobuf:"bytes,1,opt,name=section_ctx,json=sectionCtx,proto3" json:"section_ctx,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Thread) Reset()         { *m = Thread{} }
func (m *Thread) String() string { return proto.CompactTextString(m) }
func (*Thread) ProtoMessage()    {}
func (*Thread) Descriptor() ([]byte, []int) {
	return fileDescriptor_b64063be2fc89884, []int{2}
}

func (m *Thread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Thread.Unmarshal(m, b)
}
func (m *Thread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Thread.Marshal(b, m, deterministic)
}
func (m *Thread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Thread.Merge(m, src)
}
func (m *Thread) XXX_Size() int {
	return xxx_messageInfo_Thread.Size(m)
}
func (m *Thread) XXX_DiscardUnknown() {
	xxx_messageInfo_Thread.DiscardUnknown(m)
}

var xxx_messageInfo_Thread proto.InternalMessageInfo

func (m *Thread) GetSectionCtx() *Section {
	if m != nil {
		return m.SectionCtx
	}
	return nil
}

func (m *Thread) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// A Comment is inside of a thread and has a comment identifier unique to
// that thread.
type Comment struct {
	ThreadCtx            *Thread  `protobuf:"bytes,1,opt,name=thread_ctx,json=threadCtx,proto3" json:"thread_ctx,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_b64063be2fc89884, []int{3}
}

func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetThreadCtx() *Thread {
	if m != nil {
		return m.ThreadCtx
	}
	return nil
}

func (m *Comment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// A Subcomment is inside of a comment and has a subcomment identifier
// unique to that comment.
type Subcomment struct {
	CommentCtx           *Comment `protobuf:"bytes,1,opt,name=comment_ctx,json=commentCtx,proto3" json:"comment_ctx,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subcomment) Reset()         { *m = Subcomment{} }
func (m *Subcomment) String() string { return proto.CompactTextString(m) }
func (*Subcomment) ProtoMessage()    {}
func (*Subcomment) Descriptor() ([]byte, []int) {
	return fileDescriptor_b64063be2fc89884, []int{4}
}

func (m *Subcomment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subcomment.Unmarshal(m, b)
}
func (m *Subcomment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subcomment.Marshal(b, m, deterministic)
}
func (m *Subcomment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subcomment.Merge(m, src)
}
func (m *Subcomment) XXX_Size() int {
	return xxx_messageInfo_Subcomment.Size(m)
}
func (m *Subcomment) XXX_DiscardUnknown() {
	xxx_messageInfo_Subcomment.DiscardUnknown(m)
}

var xxx_messageInfo_Subcomment proto.InternalMessageInfo

func (m *Subcomment) GetCommentCtx() *Comment {
	if m != nil {
		return m.CommentCtx
	}
	return nil
}

func (m *Subcomment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Context)(nil), "context.Context")
	proto.RegisterType((*Section)(nil), "context.Section")
	proto.RegisterType((*Thread)(nil), "context.Thread")
	proto.RegisterType((*Comment)(nil), "context.Comment")
	proto.RegisterType((*Subcomment)(nil), "context.Subcomment")
}

func init() {
	proto.RegisterFile("context.proto", fileDescriptor_b64063be2fc89884)
}

var fileDescriptor_b64063be2fc89884 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0x87, 0x1b, 0xf7, 0x6d, 0xab, 0x5c, 0xd5, 0xbe, 0xc8, 0x2c, 0x65, 0x40, 0xaa, 0x32, 0x15,
	0x21, 0x52, 0xa0, 0x2b, 0x53, 0xb3, 0x24, 0x62, 0x40, 0x32, 0x4c, 0x2c, 0x15, 0xb1, 0xad, 0xc4,
	0x12, 0x89, 0x51, 0xe2, 0xa0, 0x7c, 0x05, 0xbe, 0x35, 0xb2, 0xe3, 0x36, 0x7f, 0xd4, 0x81, 0x2d,
	0xf1, 0xdd, 0x73, 0xfe, 0xdd, 0x23, 0xc3, 0x82, 0xca, 0x5c, 0xf1, 0x5a, 0xf9, 0x5f, 0x85, 0x54,
	0x12, 0xcf, 0xec, 0xaf, 0xf7, 0x83, 0x60, 0x16, 0x34, 0xdf, 0x78, 0x07, 0xf3, 0x92, 0x53, 0x25,
	0x64, 0x7e, 0xa0, 0xaa, 0x5e, 0x39, 0x6b, 0x67, 0x33, 0x7f, 0xbc, 0xf0, 0x8f, 0xe4, 0x6b, 0x53,
	0x0b, 0x47, 0x04, 0x6c, 0x5b, 0xa0, 0x6a, 0x7c, 0x0f, 0xa0, 0xd2, 0x82, 0x7f, 0x30, 0xc3, 0x20,
	0xc3, 0xfc, 0x3f, 0x31, 0x6f, 0xa6, 0x14, 0x8e, 0x88, 0xdb, 0x34, 0x69, 0x62, 0x07, 0x73, 0x2a,
	0xb3, 0x8c, 0xe7, 0xca, 0x20, 0xe3, 0xc1, 0x35, 0x41, 0x53, 0xd3, 0xd7, 0xd8, 0x36, 0x0d, 0x3d,
	0xc1, 0xb2, 0xac, 0xe2, 0x2e, 0xf7, 0xcf, 0x70, 0x97, 0x6d, 0xbc, 0x53, 0x39, 0x1c, 0x91, 0x45,
	0xdb, 0xac, 0xe9, 0x6b, 0x38, 0x46, 0x3e, 0x08, 0xb6, 0x9a, 0xac, 0x9d, 0x8d, 0x4b, 0x5c, 0x7b,
	0x12, 0xb1, 0xfd, 0x04, 0xc6, 0x54, 0xd5, 0xde, 0x15, 0xcc, 0xec, 0x8e, 0x78, 0x09, 0x28, 0x62,
	0xc6, 0x80, 0x4b, 0x50, 0xc4, 0xbc, 0x67, 0x98, 0x36, 0xab, 0xe0, 0x87, 0x3f, 0x49, 0xea, 0x29,
	0x5a, 0x02, 0x12, 0xcc, 0xa8, 0x71, 0x09, 0x12, 0xcc, 0x8b, 0xb4, 0x72, 0x93, 0x0d, 0xfb, 0x3d,
	0x7b, 0xce, 0x59, 0x7b, 0x5d, 0x77, 0xc3, 0x51, 0x2f, 0x00, 0xed, 0xde, 0x3a, 0x5b, 0xd7, 0x90,
	0x73, 0xde, 0x6c, 0xcf, 0xeb, 0x60, 0xe0, 0xfe, 0xf6, 0xfd, 0x26, 0x11, 0x2a, 0xad, 0x62, 0x9f,
	0xca, 0x6c, 0xfb, 0x59, 0x89, 0x32, 0xa9, 0xbe, 0xf9, 0x96, 0xa6, 0xbc, 0x90, 0xe6, 0xe9, 0xdc,
	0x25, 0x72, 0x6b, 0x07, 0xc6, 0x53, 0x73, 0xb2, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x60, 0x60,
	0x9a, 0x84, 0x5d, 0x02, 0x00, 0x00,
}
