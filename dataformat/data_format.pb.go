// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data_format.proto

// package dataformat defines the format in which the data will be stored
// in the database.

package dataformat

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	context "github.com/luisguve/cheroproto-go/context"
	metadata "github.com/luisguve/cheroproto-go/metadata"
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

type Notif_NotifType int32

const (
	Notif_UNKNOWN           Notif_NotifType = 0
	Notif_UPVOTE            Notif_NotifType = 1
	Notif_UPVOTE_COMMENT    Notif_NotifType = 2
	Notif_UPVOTE_SUBCOMMENT Notif_NotifType = 3
	Notif_COMMENT           Notif_NotifType = 4
	Notif_SUBCOMMENT        Notif_NotifType = 5
	Notif_MENTION           Notif_NotifType = 6
	Notif_UPDATE_THREAD     Notif_NotifType = 7
	Notif_UPDATE_COMMENT    Notif_NotifType = 8
	Notif_UPDATE_SUBCOMMENT Notif_NotifType = 9
)

var Notif_NotifType_name = map[int32]string{
	0: "UNKNOWN",
	1: "UPVOTE",
	2: "UPVOTE_COMMENT",
	3: "UPVOTE_SUBCOMMENT",
	4: "COMMENT",
	5: "SUBCOMMENT",
	6: "MENTION",
	7: "UPDATE_THREAD",
	8: "UPDATE_COMMENT",
	9: "UPDATE_SUBCOMMENT",
}

var Notif_NotifType_value = map[string]int32{
	"UNKNOWN":           0,
	"UPVOTE":            1,
	"UPVOTE_COMMENT":    2,
	"UPVOTE_SUBCOMMENT": 3,
	"COMMENT":           4,
	"SUBCOMMENT":        5,
	"MENTION":           6,
	"UPDATE_THREAD":     7,
	"UPDATE_COMMENT":    8,
	"UPDATE_SUBCOMMENT": 9,
}

func (x Notif_NotifType) String() string {
	return proto.EnumName(Notif_NotifType_name, int32(x))
}

func (Notif_NotifType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f3ef411ac4eb1018, []int{0, 0}
}

// Notif is the format in which users get their notifications.
type Notif struct {
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// example message: "User1234 and 2 more have commented out your thread"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// example subject: "On your thread: Improving your skills at writing..."
	Subject string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	// example id: {section}#{content_id}#{subject}
	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	// url to content related to the notification
	Permalink string `protobuf:"bytes,5,opt,name=permalink,proto3" json:"permalink,omitempty"`
	// internal details on the notification
	Details              *Notif_NotifDetails `protobuf:"bytes,6,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Notif) Reset()         { *m = Notif{} }
func (m *Notif) String() string { return proto.CompactTextString(m) }
func (*Notif) ProtoMessage()    {}
func (*Notif) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3ef411ac4eb1018, []int{0}
}

func (m *Notif) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notif.Unmarshal(m, b)
}
func (m *Notif) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notif.Marshal(b, m, deterministic)
}
func (m *Notif) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notif.Merge(m, src)
}
func (m *Notif) XXX_Size() int {
	return xxx_messageInfo_Notif.Size(m)
}
func (m *Notif) XXX_DiscardUnknown() {
	xxx_messageInfo_Notif.DiscardUnknown(m)
}

var xxx_messageInfo_Notif proto.InternalMessageInfo

func (m *Notif) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Notif) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Notif) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Notif) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Notif) GetPermalink() string {
	if m != nil {
		return m.Permalink
	}
	return ""
}

func (m *Notif) GetDetails() *Notif_NotifDetails {
	if m != nil {
		return m.Details
	}
	return nil
}

type Notif_NotifDetails struct {
	Type                 Notif_NotifType `protobuf:"varint,1,opt,name=type,proto3,enum=dataformat.Notif_NotifType" json:"type,omitempty"`
	LastUserIdInvolved   string          `protobuf:"bytes,2,opt,name=last_user_id_involved,json=lastUserIdInvolved,proto3" json:"last_user_id_involved,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Notif_NotifDetails) Reset()         { *m = Notif_NotifDetails{} }
func (m *Notif_NotifDetails) String() string { return proto.CompactTextString(m) }
func (*Notif_NotifDetails) ProtoMessage()    {}
func (*Notif_NotifDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3ef411ac4eb1018, []int{0, 0}
}

func (m *Notif_NotifDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notif_NotifDetails.Unmarshal(m, b)
}
func (m *Notif_NotifDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notif_NotifDetails.Marshal(b, m, deterministic)
}
func (m *Notif_NotifDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notif_NotifDetails.Merge(m, src)
}
func (m *Notif_NotifDetails) XXX_Size() int {
	return xxx_messageInfo_Notif_NotifDetails.Size(m)
}
func (m *Notif_NotifDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_Notif_NotifDetails.DiscardUnknown(m)
}

var xxx_messageInfo_Notif_NotifDetails proto.InternalMessageInfo

func (m *Notif_NotifDetails) GetType() Notif_NotifType {
	if m != nil {
		return m.Type
	}
	return Notif_UNKNOWN
}

func (m *Notif_NotifDetails) GetLastUserIdInvolved() string {
	if m != nil {
		return m.LastUserIdInvolved
	}
	return ""
}

// Activity holds the threads created, comments and subcomments of a given
// user
type Activity struct {
	ThreadsCreated       []*context.Thread     `protobuf:"bytes,1,rep,name=threads_created,json=threadsCreated,proto3" json:"threads_created,omitempty"`
	Comments             []*context.Comment    `protobuf:"bytes,2,rep,name=comments,proto3" json:"comments,omitempty"`
	Subcomments          []*context.Subcomment `protobuf:"bytes,3,rep,name=subcomments,proto3" json:"subcomments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Activity) Reset()         { *m = Activity{} }
func (m *Activity) String() string { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()    {}
func (*Activity) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3ef411ac4eb1018, []int{1}
}

func (m *Activity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Activity.Unmarshal(m, b)
}
func (m *Activity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Activity.Marshal(b, m, deterministic)
}
func (m *Activity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Activity.Merge(m, src)
}
func (m *Activity) XXX_Size() int {
	return xxx_messageInfo_Activity.Size(m)
}
func (m *Activity) XXX_DiscardUnknown() {
	xxx_messageInfo_Activity.DiscardUnknown(m)
}

var xxx_messageInfo_Activity proto.InternalMessageInfo

func (m *Activity) GetThreadsCreated() []*context.Thread {
	if m != nil {
		return m.ThreadsCreated
	}
	return nil
}

func (m *Activity) GetComments() []*context.Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

func (m *Activity) GetSubcomments() []*context.Subcomment {
	if m != nil {
		return m.Subcomments
	}
	return nil
}

// Public data of users
type BasicUserData struct {
	Alias                string   `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	PicUrl               string   `protobuf:"bytes,3,opt,name=pic_url,json=picUrl,proto3" json:"pic_url,omitempty"`
	About                string   `protobuf:"bytes,4,opt,name=about,proto3" json:"about,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BasicUserData) Reset()         { *m = BasicUserData{} }
func (m *BasicUserData) String() string { return proto.CompactTextString(m) }
func (*BasicUserData) ProtoMessage()    {}
func (*BasicUserData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3ef411ac4eb1018, []int{2}
}

func (m *BasicUserData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BasicUserData.Unmarshal(m, b)
}
func (m *BasicUserData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BasicUserData.Marshal(b, m, deterministic)
}
func (m *BasicUserData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasicUserData.Merge(m, src)
}
func (m *BasicUserData) XXX_Size() int {
	return xxx_messageInfo_BasicUserData.Size(m)
}
func (m *BasicUserData) XXX_DiscardUnknown() {
	xxx_messageInfo_BasicUserData.DiscardUnknown(m)
}

var xxx_messageInfo_BasicUserData proto.InternalMessageInfo

func (m *BasicUserData) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *BasicUserData) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *BasicUserData) GetPicUrl() string {
	if m != nil {
		return m.PicUrl
	}
	return ""
}

func (m *BasicUserData) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

type PrivateData struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             []byte   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrivateData) Reset()         { *m = PrivateData{} }
func (m *PrivateData) String() string { return proto.CompactTextString(m) }
func (*PrivateData) ProtoMessage()    {}
func (*PrivateData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3ef411ac4eb1018, []int{3}
}

func (m *PrivateData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateData.Unmarshal(m, b)
}
func (m *PrivateData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateData.Marshal(b, m, deterministic)
}
func (m *PrivateData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateData.Merge(m, src)
}
func (m *PrivateData) XXX_Size() int {
	return xxx_messageInfo_PrivateData.Size(m)
}
func (m *PrivateData) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateData.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateData proto.InternalMessageInfo

func (m *PrivateData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *PrivateData) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

type User struct {
	BasicUserData *BasicUserData `protobuf:"bytes,1,opt,name=basic_user_data,json=basicUserData,proto3" json:"basic_user_data,omitempty"`
	PrivateData   *PrivateData   `protobuf:"bytes,2,opt,name=private_data,json=privateData,proto3" json:"private_data,omitempty"`
	// ids of users that are following this user.
	FollowersIds []string `protobuf:"bytes,3,rep,name=followers_ids,json=followersIds,proto3" json:"followers_ids,omitempty"`
	// ids of users that this user is following.
	FollowingIds []string `protobuf:"bytes,4,rep,name=following_ids,json=followingIds,proto3" json:"following_ids,omitempty"`
	// unread notifications of this user.
	UnreadNotifs []*Notif `protobuf:"bytes,5,rep,name=unread_notifs,json=unreadNotifs,proto3" json:"unread_notifs,omitempty"`
	// read notifications of this user.
	ReadNotifs []*Notif `protobuf:"bytes,6,rep,name=read_notifs,json=readNotifs,proto3" json:"read_notifs,omitempty"`
	// ids of content this user has created recently.
	RecentActivity *Activity `protobuf:"bytes,7,opt,name=recent_activity,json=recentActivity,proto3" json:"recent_activity,omitempty"`
	// ids of unactive content this user has created.
	OldActivity *Activity `protobuf:"bytes,8,opt,name=old_activity,json=oldActivity,proto3" json:"old_activity,omitempty"`
	// ids of threads that this user has saved to read later.
	SavedThreads []*context.Thread `protobuf:"bytes,9,rep,name=saved_threads,json=savedThreads,proto3" json:"saved_threads,omitempty"`
	// the last time this user created a thread.
	LastTimeCreated      *timestamp.Timestamp `protobuf:"bytes,10,opt,name=last_time_created,json=lastTimeCreated,proto3" json:"last_time_created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3ef411ac4eb1018, []int{4}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetBasicUserData() *BasicUserData {
	if m != nil {
		return m.BasicUserData
	}
	return nil
}

func (m *User) GetPrivateData() *PrivateData {
	if m != nil {
		return m.PrivateData
	}
	return nil
}

func (m *User) GetFollowersIds() []string {
	if m != nil {
		return m.FollowersIds
	}
	return nil
}

func (m *User) GetFollowingIds() []string {
	if m != nil {
		return m.FollowingIds
	}
	return nil
}

func (m *User) GetUnreadNotifs() []*Notif {
	if m != nil {
		return m.UnreadNotifs
	}
	return nil
}

func (m *User) GetReadNotifs() []*Notif {
	if m != nil {
		return m.ReadNotifs
	}
	return nil
}

func (m *User) GetRecentActivity() *Activity {
	if m != nil {
		return m.RecentActivity
	}
	return nil
}

func (m *User) GetOldActivity() *Activity {
	if m != nil {
		return m.OldActivity
	}
	return nil
}

func (m *User) GetSavedThreads() []*context.Thread {
	if m != nil {
		return m.SavedThreads
	}
	return nil
}

func (m *User) GetLastTimeCreated() *timestamp.Timestamp {
	if m != nil {
		return m.LastTimeCreated
	}
	return nil
}

type Content struct {
	Title       string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content     string               `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	FtFile      string               `protobuf:"bytes,3,opt,name=ft_file,json=ftFile,proto3" json:"ft_file,omitempty"`
	PublishDate *timestamp.Timestamp `protobuf:"bytes,4,opt,name=publish_date,json=publishDate,proto3" json:"publish_date,omitempty"`
	AuthorId    string               `protobuf:"bytes,5,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	// id contains the THREAD ID
	Id string `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	// section_name is the name of the section
	SectionName string `protobuf:"bytes,7,opt,name=section_name,json=sectionName,proto3" json:"section_name,omitempty"`
	// section_id holds the id of the section
	SectionId  string   `protobuf:"bytes,8,opt,name=section_id,json=sectionId,proto3" json:"section_id,omitempty"`
	Topvotes   uint32   `protobuf:"varint,9,opt,name=topvotes,proto3" json:"topvotes,omitempty"`
	Replies    uint32   `protobuf:"varint,10,opt,name=replies,proto3" json:"replies,omitempty"`
	VoterIds   []string `protobuf:"bytes,11,rep,name=voter_ids,json=voterIds,proto3" json:"voter_ids,omitempty"`
	ReplierIds []string `protobuf:"bytes,12,rep,name=replier_ids,json=replierIds,proto3" json:"replier_ids,omitempty"`
	// list of users who saved this content
	UsersWhoSaved []string `protobuf:"bytes,13,rep,name=users_who_saved,json=usersWhoSaved,proto3" json:"users_who_saved,omitempty"`
	// Metadata holds information about the interactions with this content.
	Metadata *metadata.Content `protobuf:"bytes,14,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// permalink is the full url to the content, including fragments
	Permalink            string   `protobuf:"bytes,15,opt,name=permalink,proto3" json:"permalink,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3ef411ac4eb1018, []int{5}
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

func (m *Content) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Content) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Content) GetFtFile() string {
	if m != nil {
		return m.FtFile
	}
	return ""
}

func (m *Content) GetPublishDate() *timestamp.Timestamp {
	if m != nil {
		return m.PublishDate
	}
	return nil
}

func (m *Content) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Content) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Content) GetSectionName() string {
	if m != nil {
		return m.SectionName
	}
	return ""
}

func (m *Content) GetSectionId() string {
	if m != nil {
		return m.SectionId
	}
	return ""
}

func (m *Content) GetTopvotes() uint32 {
	if m != nil {
		return m.Topvotes
	}
	return 0
}

func (m *Content) GetReplies() uint32 {
	if m != nil {
		return m.Replies
	}
	return 0
}

func (m *Content) GetVoterIds() []string {
	if m != nil {
		return m.VoterIds
	}
	return nil
}

func (m *Content) GetReplierIds() []string {
	if m != nil {
		return m.ReplierIds
	}
	return nil
}

func (m *Content) GetUsersWhoSaved() []string {
	if m != nil {
		return m.UsersWhoSaved
	}
	return nil
}

func (m *Content) GetMetadata() *metadata.Content {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Content) GetPermalink() string {
	if m != nil {
		return m.Permalink
	}
	return ""
}

func init() {
	proto.RegisterEnum("dataformat.Notif_NotifType", Notif_NotifType_name, Notif_NotifType_value)
	proto.RegisterType((*Notif)(nil), "dataformat.Notif")
	proto.RegisterType((*Notif_NotifDetails)(nil), "dataformat.Notif.NotifDetails")
	proto.RegisterType((*Activity)(nil), "dataformat.Activity")
	proto.RegisterType((*BasicUserData)(nil), "dataformat.BasicUserData")
	proto.RegisterType((*PrivateData)(nil), "dataformat.PrivateData")
	proto.RegisterType((*User)(nil), "dataformat.User")
	proto.RegisterType((*Content)(nil), "dataformat.Content")
}

func init() {
	proto.RegisterFile("data_format.proto", fileDescriptor_f3ef411ac4eb1018)
}

var fileDescriptor_f3ef411ac4eb1018 = []byte{
	// 1007 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdb, 0x8e, 0xe3, 0x44,
	0x10, 0x25, 0x99, 0x4c, 0x12, 0x97, 0xe3, 0x5c, 0x9a, 0x5d, 0xad, 0xc9, 0x02, 0x3b, 0x04, 0x09,
	0xcd, 0xc3, 0xae, 0x23, 0x86, 0xdb, 0x08, 0x69, 0x85, 0xe6, 0xb6, 0x22, 0x42, 0x9b, 0x19, 0x79,
	0x12, 0x56, 0xe2, 0xc5, 0xea, 0xd8, 0x9d, 0xa4, 0xc1, 0x76, 0x5b, 0xee, 0x76, 0x86, 0xf9, 0x06,
	0xbe, 0x84, 0x47, 0x9e, 0xf8, 0x26, 0xfe, 0x02, 0xf5, 0xc5, 0x8e, 0x07, 0xd8, 0xe5, 0x25, 0x4a,
	0x9d, 0x3a, 0xa7, 0xbb, 0xba, 0xfb, 0x54, 0x19, 0x46, 0x11, 0x16, 0x38, 0x58, 0xb3, 0x3c, 0xc1,
	0xc2, 0xcb, 0x72, 0x26, 0x18, 0x02, 0x09, 0x69, 0x64, 0xfc, 0x6c, 0xc3, 0xd8, 0x26, 0x26, 0x53,
	0x95, 0x59, 0x15, 0xeb, 0xa9, 0xa0, 0x09, 0xe1, 0x02, 0x27, 0x99, 0x26, 0x8f, 0x9d, 0x90, 0xa5,
	0x82, 0xfc, 0x6a, 0xb4, 0xe3, 0x7e, 0x42, 0x04, 0x96, 0x7a, 0x1d, 0x4f, 0x7e, 0x6b, 0xc1, 0xe1,
	0x9c, 0x09, 0xba, 0x46, 0xa7, 0x60, 0x55, 0x5a, 0xb7, 0x71, 0xd4, 0x38, 0xb6, 0x4f, 0xc6, 0x9e,
	0x5e, 0xdd, 0x2b, 0x57, 0xf7, 0x16, 0x25, 0xc3, 0xdf, 0x93, 0x91, 0x0b, 0x9d, 0x84, 0x70, 0x8e,
	0x37, 0xc4, 0x6d, 0x1e, 0x35, 0x8e, 0x2d, 0xbf, 0x0c, 0x65, 0x86, 0x17, 0xab, 0x9f, 0x49, 0x28,
	0xdc, 0x03, 0x9d, 0x31, 0x21, 0xea, 0x43, 0x93, 0x46, 0x6e, 0x4b, 0x81, 0x4d, 0x1a, 0xa1, 0x0f,
	0xc1, 0xca, 0x48, 0x9e, 0xe0, 0x98, 0xa6, 0xbf, 0xb8, 0x87, 0x0a, 0xde, 0x03, 0xe8, 0x14, 0x3a,
	0x11, 0x11, 0x98, 0xc6, 0xdc, 0x6d, 0xab, 0xca, 0x3e, 0xf6, 0xf6, 0x77, 0xe0, 0xa9, 0xfa, 0xf5,
	0xef, 0xa5, 0x66, 0xf9, 0x25, 0x7d, 0x9c, 0x43, 0xaf, 0x9e, 0x40, 0x53, 0x68, 0x89, 0xfb, 0x8c,
	0xa8, 0x03, 0xf6, 0x4f, 0x9e, 0xbe, 0x65, 0x99, 0xc5, 0x7d, 0x46, 0x7c, 0x45, 0x44, 0x9f, 0xc3,
	0xe3, 0x18, 0x73, 0x11, 0x14, 0x9c, 0xe4, 0x01, 0x8d, 0x02, 0x9a, 0xee, 0x58, 0xbc, 0x23, 0x91,
	0x39, 0x2a, 0x92, 0xc9, 0x25, 0x27, 0xf9, 0x2c, 0x9a, 0x99, 0xcc, 0xe4, 0xcf, 0x06, 0x58, 0xd5,
	0x32, 0xc8, 0x86, 0xce, 0x72, 0xfe, 0xc3, 0xfc, 0xfa, 0xcd, 0x7c, 0xf8, 0x1e, 0x02, 0x68, 0x2f,
	0x6f, 0x7e, 0xbc, 0x5e, 0x5c, 0x0d, 0x1b, 0x08, 0x41, 0x5f, 0xff, 0x0f, 0x2e, 0xae, 0x5f, 0xbf,
	0xbe, 0x9a, 0x2f, 0x86, 0x4d, 0xf4, 0x18, 0x46, 0x06, 0xbb, 0x5d, 0x9e, 0x97, 0xf0, 0x81, 0x5c,
	0xa3, 0x0c, 0x5a, 0xa8, 0x0f, 0x50, 0x4b, 0x1e, 0xca, 0xa4, 0xfc, 0x37, 0xbb, 0x9e, 0x0f, 0xdb,
	0x68, 0x04, 0xce, 0xf2, 0xe6, 0xf2, 0x6c, 0x71, 0x15, 0x2c, 0xbe, 0xf7, 0xaf, 0xce, 0x2e, 0x87,
	0x1d, 0xbd, 0x8f, 0x82, 0x4a, 0x4d, 0x57, 0xef, 0xa3, 0xb0, 0xda, 0x52, 0xd6, 0xe4, 0xf7, 0x06,
	0x74, 0xcf, 0x42, 0x41, 0x77, 0x54, 0xdc, 0xa3, 0x53, 0x18, 0x88, 0x6d, 0x4e, 0x70, 0xc4, 0x83,
	0x30, 0x27, 0x58, 0x90, 0xc8, 0x6d, 0x1c, 0x1d, 0x1c, 0xdb, 0x27, 0x03, 0xaf, 0xf4, 0xd4, 0x42,
	0xe5, 0xfd, 0xbe, 0xe1, 0x5d, 0x68, 0x1a, 0x7a, 0x0e, 0xdd, 0x90, 0x25, 0x09, 0x49, 0x05, 0x77,
	0x9b, 0x4a, 0x32, 0xac, 0x24, 0x17, 0x3a, 0xe1, 0x57, 0x0c, 0xf4, 0x15, 0xd8, 0xbc, 0x58, 0x55,
	0x82, 0x03, 0x25, 0x78, 0xbf, 0x12, 0xdc, 0x56, 0x39, 0xbf, 0xce, 0x9b, 0x64, 0xe0, 0x9c, 0x63,
	0x4e, 0x43, 0x79, 0xf9, 0x97, 0x58, 0x60, 0xf4, 0x08, 0x0e, 0x71, 0x4c, 0x31, 0x57, 0x6f, 0x6b,
	0xf9, 0x3a, 0x40, 0x63, 0xe8, 0xca, 0xa7, 0x4b, 0x71, 0x52, 0xba, 0xb3, 0x8a, 0xd1, 0x13, 0xe8,
	0x64, 0x34, 0x0c, 0x8a, 0x3c, 0x36, 0xf6, 0x6c, 0x67, 0x34, 0x5c, 0xe6, 0xb1, 0x5a, 0x6a, 0xc5,
	0x0a, 0x61, 0x0c, 0xaa, 0x83, 0xc9, 0x77, 0x60, 0xdf, 0xe4, 0x74, 0x87, 0x05, 0x29, 0xf7, 0x23,
	0x09, 0xa6, 0x71, 0xb9, 0x9f, 0x0a, 0xe4, 0x7e, 0x19, 0xe6, 0xfc, 0x8e, 0xe5, 0xda, 0x22, 0x3d,
	0xbf, 0x8a, 0x27, 0x7f, 0xb4, 0xa0, 0x25, 0xcb, 0x45, 0x67, 0x30, 0x58, 0xc9, 0xda, 0xb5, 0xab,
	0xa4, 0x07, 0x4d, 0xc7, 0x7d, 0x50, 0x37, 0xe4, 0x83, 0xe3, 0xf9, 0xce, 0xea, 0xc1, 0x69, 0xbf,
	0x85, 0x5e, 0xa6, 0x8b, 0xd1, 0xfa, 0xa6, 0xd2, 0x3f, 0xa9, 0xeb, 0x6b, 0xc5, 0xfa, 0x76, 0x56,
	0xab, 0xfc, 0x53, 0x70, 0xd6, 0x2c, 0x8e, 0xd9, 0x1d, 0xc9, 0x79, 0x40, 0x23, 0x7d, 0xe7, 0x96,
	0xdf, 0xab, 0xc0, 0x59, 0xc4, 0xf7, 0x24, 0x9a, 0x6e, 0x14, 0xa9, 0x55, 0x27, 0xd1, 0x74, 0x23,
	0x49, 0x5f, 0x83, 0x53, 0xa4, 0xf2, 0xed, 0x83, 0x54, 0x1a, 0x9e, 0xbb, 0x87, 0xea, 0xf5, 0x46,
	0xff, 0xea, 0x2b, 0xbf, 0xa7, 0x79, 0x2a, 0xe0, 0xe8, 0x04, 0xec, 0xba, 0xaa, 0xfd, 0x36, 0x15,
	0xd4, 0x34, 0x2f, 0x61, 0x90, 0x93, 0x90, 0xa4, 0x22, 0xc0, 0xc6, 0xa2, 0x6e, 0x47, 0x1d, 0xfa,
	0x51, 0x5d, 0x57, 0xda, 0xd7, 0xef, 0x6b, 0x72, 0x65, 0xe7, 0x6f, 0xa0, 0xc7, 0xe2, 0x68, 0xaf,
	0xed, 0xbe, 0x43, 0x6b, 0xb3, 0x38, 0xaa, 0x84, 0x5f, 0x82, 0xc3, 0xf1, 0x8e, 0x44, 0x81, 0x71,
	0xb9, 0x6b, 0xfd, 0x77, 0x17, 0xf4, 0x14, 0x4b, 0x07, 0x1c, 0xbd, 0x82, 0x91, 0x9a, 0x1b, 0x72,
	0x4c, 0x56, 0xfd, 0x03, 0xff, 0x3b, 0x56, 0x07, 0x52, 0x24, 0x43, 0xd3, 0x4b, 0x93, 0xbf, 0x0e,
	0xa0, 0x73, 0x21, 0x37, 0x4a, 0x85, 0x74, 0x9c, 0xa0, 0x22, 0x26, 0xa5, 0xe3, 0x54, 0x20, 0x87,
	0x6c, 0xa8, 0x09, 0xe5, 0xf8, 0x35, 0xa1, 0xf4, 0xf7, 0x5a, 0x04, 0x6b, 0x1a, 0x93, 0xd2, 0xdf,
	0x6b, 0xf1, 0x8a, 0xc6, 0x04, 0xbd, 0x84, 0x5e, 0x56, 0xac, 0x62, 0xca, 0xb7, 0xd2, 0x3c, 0x44,
	0xd9, 0xfc, 0xdd, 0x75, 0xd9, 0x86, 0x7f, 0x89, 0x05, 0x41, 0x4f, 0xc1, 0xc2, 0x85, 0xd8, 0x32,
	0x39, 0x10, 0xcd, 0xb0, 0xee, 0x6a, 0x60, 0x16, 0x99, 0xc9, 0xde, 0xae, 0x26, 0xfb, 0x27, 0xd0,
	0xe3, 0x24, 0x14, 0x94, 0xa5, 0x81, 0x6a, 0xc2, 0x8e, 0xca, 0xd8, 0x06, 0x9b, 0xcb, 0x3e, 0xfc,
	0x08, 0xa0, 0xa4, 0xd0, 0x48, 0x3d, 0x8c, 0xe5, 0x5b, 0x06, 0x99, 0x45, 0xb2, 0xa5, 0x04, 0xcb,
	0x76, 0x4c, 0x10, 0x79, 0xf7, 0x8d, 0x63, 0xc7, 0xaf, 0x62, 0x79, 0xf8, 0x9c, 0x64, 0x31, 0x25,
	0x5c, 0x5d, 0xae, 0xe3, 0x97, 0xa1, 0x2c, 0x52, 0x52, 0x72, 0xe5, 0x5d, 0x5b, 0x79, 0xb7, 0xab,
	0x00, 0xe9, 0xdb, 0x67, 0xd2, 0x7f, 0x92, 0xa7, 0xd3, 0x3d, 0x95, 0x06, 0x03, 0x49, 0xc2, 0x67,
	0x30, 0x90, 0xbd, 0xc9, 0x83, 0xbb, 0x2d, 0x0b, 0xd4, 0xc3, 0xba, 0x8e, 0x22, 0x39, 0x0a, 0x7e,
	0xb3, 0x65, 0xb7, 0x12, 0x44, 0x2f, 0xa0, 0x5b, 0x7e, 0x51, 0xdd, 0xbe, 0xba, 0xc5, 0x91, 0x57,
	0x7d, 0x62, 0xcd, 0xbb, 0xf9, 0x15, 0xe5, 0xe1, 0x67, 0x6e, 0xf0, 0x8f, 0xcf, 0xdc, 0xb9, 0xf7,
	0xd3, 0xf3, 0x0d, 0x15, 0xdb, 0x62, 0xe5, 0x85, 0x2c, 0x99, 0xc6, 0x05, 0xe5, 0x9b, 0x62, 0x47,
	0xa6, 0xe1, 0x96, 0xe4, 0x4c, 0xbd, 0xc9, 0x8b, 0x0d, 0x9b, 0xee, 0xfd, 0xba, 0x6a, 0x2b, 0xf0,
	0x8b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xd9, 0x36, 0x99, 0x24, 0x08, 0x00, 0x00,
}
