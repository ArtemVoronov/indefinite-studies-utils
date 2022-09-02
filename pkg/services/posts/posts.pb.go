// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/services/posts/posts.proto

package posts

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GetPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_posts_posts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_posts_posts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_pkg_services_posts_posts_proto_rawDescGZIP(), []int{0}
}

func (x *GetPostRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPostReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId       int32                `protobuf:"varint,2,opt,name=authorId,proto3" json:"authorId,omitempty"`
	Text           string               `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	PreviewText    string               `protobuf:"bytes,4,opt,name=previewText,proto3" json:"previewText,omitempty"`
	Topic          string               `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
	State          string               `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	CreateDate     *timestamp.Timestamp `protobuf:"bytes,7,opt,name=createDate,proto3" json:"createDate,omitempty"`
	LastUpdateDate *timestamp.Timestamp `protobuf:"bytes,8,opt,name=lastUpdateDate,proto3" json:"lastUpdateDate,omitempty"`
}

func (x *GetPostReply) Reset() {
	*x = GetPostReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_posts_posts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostReply) ProtoMessage() {}

func (x *GetPostReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_posts_posts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostReply.ProtoReflect.Descriptor instead.
func (*GetPostReply) Descriptor() ([]byte, []int) {
	return file_pkg_services_posts_posts_proto_rawDescGZIP(), []int{1}
}

func (x *GetPostReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetPostReply) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *GetPostReply) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GetPostReply) GetPreviewText() string {
	if x != nil {
		return x.PreviewText
	}
	return ""
}

func (x *GetPostReply) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *GetPostReply) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *GetPostReply) GetCreateDate() *timestamp.Timestamp {
	if x != nil {
		return x.CreateDate
	}
	return nil
}

func (x *GetPostReply) GetLastUpdateDate() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdateDate
	}
	return nil
}

type GetPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetPostsRequest) Reset() {
	*x = GetPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_posts_posts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsRequest) ProtoMessage() {}

func (x *GetPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_posts_posts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsRequest.ProtoReflect.Descriptor instead.
func (*GetPostsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_services_posts_posts_proto_rawDescGZIP(), []int{2}
}

func (x *GetPostsRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetPostsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*GetPostReply `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetPostsReply) Reset() {
	*x = GetPostsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_posts_posts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsReply) ProtoMessage() {}

func (x *GetPostsReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_posts_posts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsReply.ProtoReflect.Descriptor instead.
func (*GetPostsReply) Descriptor() ([]byte, []int) {
	return file_pkg_services_posts_posts_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostsReply) GetPosts() []*GetPostReply {
	if x != nil {
		return x.Posts
	}
	return nil
}

type GetCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCommentRequest) Reset() {
	*x = GetCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_posts_posts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentRequest) ProtoMessage() {}

func (x *GetCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_posts_posts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentRequest.ProtoReflect.Descriptor instead.
func (*GetCommentRequest) Descriptor() ([]byte, []int) {
	return file_pkg_services_posts_posts_proto_rawDescGZIP(), []int{4}
}

func (x *GetCommentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId        int32                `protobuf:"varint,2,opt,name=authorId,proto3" json:"authorId,omitempty"`
	PostId          int32                `protobuf:"varint,3,opt,name=postId,proto3" json:"postId,omitempty"`
	LinkedCommentId int32                `protobuf:"varint,5,opt,name=linkedCommentId,proto3" json:"linkedCommentId,omitempty"`
	Text            string               `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty"`
	State           string               `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	CreateDate      *timestamp.Timestamp `protobuf:"bytes,8,opt,name=createDate,proto3" json:"createDate,omitempty"`
	LastUpdateDate  *timestamp.Timestamp `protobuf:"bytes,9,opt,name=lastUpdateDate,proto3" json:"lastUpdateDate,omitempty"`
}

func (x *GetCommentReply) Reset() {
	*x = GetCommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_posts_posts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentReply) ProtoMessage() {}

func (x *GetCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_posts_posts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentReply.ProtoReflect.Descriptor instead.
func (*GetCommentReply) Descriptor() ([]byte, []int) {
	return file_pkg_services_posts_posts_proto_rawDescGZIP(), []int{5}
}

func (x *GetCommentReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCommentReply) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *GetCommentReply) GetPostId() int32 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *GetCommentReply) GetLinkedCommentId() int32 {
	if x != nil {
		return x.LinkedCommentId
	}
	return 0
}

func (x *GetCommentReply) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GetCommentReply) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *GetCommentReply) GetCreateDate() *timestamp.Timestamp {
	if x != nil {
		return x.CreateDate
	}
	return nil
}

func (x *GetCommentReply) GetLastUpdateDate() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdateDate
	}
	return nil
}

type GetCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetCommentsRequest) Reset() {
	*x = GetCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_posts_posts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsRequest) ProtoMessage() {}

func (x *GetCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_posts_posts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsRequest.ProtoReflect.Descriptor instead.
func (*GetCommentsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_services_posts_posts_proto_rawDescGZIP(), []int{6}
}

func (x *GetCommentsRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetCommentsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*GetCommentReply `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *GetCommentsReply) Reset() {
	*x = GetCommentsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_posts_posts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsReply) ProtoMessage() {}

func (x *GetCommentsReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_posts_posts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsReply.ProtoReflect.Descriptor instead.
func (*GetCommentsReply) Descriptor() ([]byte, []int) {
	return file_pkg_services_posts_posts_proto_rawDescGZIP(), []int{7}
}

func (x *GetCommentsReply) GetComments() []*GetCommentReply {
	if x != nil {
		return x.Comments
	}
	return nil
}

var File_pkg_services_posts_posts_proto protoreflect.FileDescriptor

var file_pkg_services_posts_posts_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9c, 0x02, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x54, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x54, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x42, 0x02, 0x10, 0x01, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x22, 0x3a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x23,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xa9, 0x02, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x6c,
	0x69, 0x6e, 0x6b, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x6c,
	0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22,
	0x2a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x05, 0x42, 0x02, 0x10, 0x01, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x46, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x32, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x32, 0x9b, 0x03, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x15, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x40, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x19,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x72, 0x74, 0x65, 0x6d, 0x56, 0x6f, 0x72, 0x6f, 0x6e, 0x6f, 0x76, 0x2f, 0x69, 0x6e, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x2d, 0x73, 0x74, 0x75, 0x64, 0x69, 0x65, 0x73, 0x2d,
	0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_services_posts_posts_proto_rawDescOnce sync.Once
	file_pkg_services_posts_posts_proto_rawDescData = file_pkg_services_posts_posts_proto_rawDesc
)

func file_pkg_services_posts_posts_proto_rawDescGZIP() []byte {
	file_pkg_services_posts_posts_proto_rawDescOnce.Do(func() {
		file_pkg_services_posts_posts_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_services_posts_posts_proto_rawDescData)
	})
	return file_pkg_services_posts_posts_proto_rawDescData
}

var file_pkg_services_posts_posts_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_services_posts_posts_proto_goTypes = []interface{}{
	(*GetPostRequest)(nil),      // 0: posts.GetPostRequest
	(*GetPostReply)(nil),        // 1: posts.GetPostReply
	(*GetPostsRequest)(nil),     // 2: posts.GetPostsRequest
	(*GetPostsReply)(nil),       // 3: posts.GetPostsReply
	(*GetCommentRequest)(nil),   // 4: posts.GetCommentRequest
	(*GetCommentReply)(nil),     // 5: posts.GetCommentReply
	(*GetCommentsRequest)(nil),  // 6: posts.GetCommentsRequest
	(*GetCommentsReply)(nil),    // 7: posts.GetCommentsReply
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_pkg_services_posts_posts_proto_depIdxs = []int32{
	8,  // 0: posts.GetPostReply.createDate:type_name -> google.protobuf.Timestamp
	8,  // 1: posts.GetPostReply.lastUpdateDate:type_name -> google.protobuf.Timestamp
	1,  // 2: posts.GetPostsReply.posts:type_name -> posts.GetPostReply
	8,  // 3: posts.GetCommentReply.createDate:type_name -> google.protobuf.Timestamp
	8,  // 4: posts.GetCommentReply.lastUpdateDate:type_name -> google.protobuf.Timestamp
	5,  // 5: posts.GetCommentsReply.comments:type_name -> posts.GetCommentReply
	0,  // 6: posts.PostsService.GetPost:input_type -> posts.GetPostRequest
	2,  // 7: posts.PostsService.GetPosts:input_type -> posts.GetPostsRequest
	0,  // 8: posts.PostsService.GetPostsStream:input_type -> posts.GetPostRequest
	4,  // 9: posts.PostsService.GetComment:input_type -> posts.GetCommentRequest
	6,  // 10: posts.PostsService.GetComments:input_type -> posts.GetCommentsRequest
	4,  // 11: posts.PostsService.GetCommentsStream:input_type -> posts.GetCommentRequest
	1,  // 12: posts.PostsService.GetPost:output_type -> posts.GetPostReply
	3,  // 13: posts.PostsService.GetPosts:output_type -> posts.GetPostsReply
	1,  // 14: posts.PostsService.GetPostsStream:output_type -> posts.GetPostReply
	5,  // 15: posts.PostsService.GetComment:output_type -> posts.GetCommentReply
	7,  // 16: posts.PostsService.GetComments:output_type -> posts.GetCommentsReply
	5,  // 17: posts.PostsService.GetCommentsStream:output_type -> posts.GetCommentReply
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_services_posts_posts_proto_init() }
func file_pkg_services_posts_posts_proto_init() {
	if File_pkg_services_posts_posts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_services_posts_posts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostRequest); i {
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
		file_pkg_services_posts_posts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostReply); i {
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
		file_pkg_services_posts_posts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsRequest); i {
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
		file_pkg_services_posts_posts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsReply); i {
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
		file_pkg_services_posts_posts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentRequest); i {
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
		file_pkg_services_posts_posts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentReply); i {
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
		file_pkg_services_posts_posts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsRequest); i {
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
		file_pkg_services_posts_posts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsReply); i {
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
			RawDescriptor: file_pkg_services_posts_posts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_services_posts_posts_proto_goTypes,
		DependencyIndexes: file_pkg_services_posts_posts_proto_depIdxs,
		MessageInfos:      file_pkg_services_posts_posts_proto_msgTypes,
	}.Build()
	File_pkg_services_posts_posts_proto = out.File
	file_pkg_services_posts_posts_proto_rawDesc = nil
	file_pkg_services_posts_posts_proto_goTypes = nil
	file_pkg_services_posts_posts_proto_depIdxs = nil
}
