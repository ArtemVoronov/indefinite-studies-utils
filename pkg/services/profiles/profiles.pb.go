// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/services/profiles/profiles.proto

package profiles

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

type ValidateCredentialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ValidateCredentialsRequest) Reset() {
	*x = ValidateCredentialsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_profiles_profiles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateCredentialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCredentialsRequest) ProtoMessage() {}

func (x *ValidateCredentialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_profiles_profiles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCredentialsRequest.ProtoReflect.Descriptor instead.
func (*ValidateCredentialsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_services_profiles_profiles_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateCredentialsRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *ValidateCredentialsRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ValidateCredentialsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid string `protobuf:"bytes,1,opt,name=userUuid,proto3" json:"userUuid,omitempty"`
	IsValid  bool   `protobuf:"varint,2,opt,name=isValid,proto3" json:"isValid,omitempty"`
	Role     string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *ValidateCredentialsReply) Reset() {
	*x = ValidateCredentialsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_profiles_profiles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateCredentialsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCredentialsReply) ProtoMessage() {}

func (x *ValidateCredentialsReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_profiles_profiles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCredentialsReply.ProtoReflect.Descriptor instead.
func (*ValidateCredentialsReply) Descriptor() ([]byte, []int) {
	return file_pkg_services_profiles_profiles_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateCredentialsReply) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *ValidateCredentialsReply) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

func (x *ValidateCredentialsReply) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_profiles_profiles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_profiles_profiles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_pkg_services_profiles_profiles_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid           string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Login          string               `protobuf:"bytes,3,opt,name=login,proto3" json:"login,omitempty"`
	Email          string               `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Role           string               `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	State          string               `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	CreateDate     *timestamp.Timestamp `protobuf:"bytes,7,opt,name=createDate,proto3" json:"createDate,omitempty"`
	LastUpdateDate *timestamp.Timestamp `protobuf:"bytes,8,opt,name=lastUpdateDate,proto3" json:"lastUpdateDate,omitempty"`
}

func (x *GetUserReply) Reset() {
	*x = GetUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_profiles_profiles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReply) ProtoMessage() {}

func (x *GetUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_profiles_profiles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReply.ProtoReflect.Descriptor instead.
func (*GetUserReply) Descriptor() ([]byte, []int) {
	return file_pkg_services_profiles_profiles_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserReply) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetUserReply) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *GetUserReply) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserReply) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *GetUserReply) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *GetUserReply) GetCreateDate() *timestamp.Timestamp {
	if x != nil {
		return x.CreateDate
	}
	return nil
}

func (x *GetUserReply) GetLastUpdateDate() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdateDate
	}
	return nil
}

type GetUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Shard  int32 `protobuf:"varint,3,opt,name=shard,proto3" json:"shard,omitempty"`
}

func (x *GetUsersRequest) Reset() {
	*x = GetUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_profiles_profiles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRequest) ProtoMessage() {}

func (x *GetUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_profiles_profiles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return file_pkg_services_profiles_profiles_proto_rawDescGZIP(), []int{4}
}

func (x *GetUsersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetUsersRequest) GetShard() int32 {
	if x != nil {
		return x.Shard
	}
	return 0
}

type GetUsersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count       int32           `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Offset      int32           `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit       int32           `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	ShardsCount int32           `protobuf:"varint,4,opt,name=shardsCount,proto3" json:"shardsCount,omitempty"`
	Users       []*GetUserReply `protobuf:"bytes,5,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUsersReply) Reset() {
	*x = GetUsersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_profiles_profiles_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersReply) ProtoMessage() {}

func (x *GetUsersReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_profiles_profiles_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersReply.ProtoReflect.Descriptor instead.
func (*GetUsersReply) Descriptor() ([]byte, []int) {
	return file_pkg_services_profiles_profiles_proto_rawDescGZIP(), []int{5}
}

func (x *GetUsersReply) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetUsersReply) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetUsersReply) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetUsersReply) GetShardsCount() int32 {
	if x != nil {
		return x.ShardsCount
	}
	return 0
}

func (x *GetUsersReply) GetUsers() []*GetUserReply {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_pkg_services_profiles_profiles_proto protoreflect.FileDescriptor

var file_pkg_services_profiles_profiles_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4e, 0x0a, 0x1a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x64, 0x0a, 0x18, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x88, 0x02,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x55, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x22,
	0xa3, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0xbf, 0x02, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x13, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x74, 0x65, 0x6d, 0x56, 0x6f, 0x72, 0x6f, 0x6e,
	0x6f, 0x76, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x2d, 0x73, 0x74,
	0x75, 0x64, 0x69, 0x65, 0x73, 0x2d, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_services_profiles_profiles_proto_rawDescOnce sync.Once
	file_pkg_services_profiles_profiles_proto_rawDescData = file_pkg_services_profiles_profiles_proto_rawDesc
)

func file_pkg_services_profiles_profiles_proto_rawDescGZIP() []byte {
	file_pkg_services_profiles_profiles_proto_rawDescOnce.Do(func() {
		file_pkg_services_profiles_profiles_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_services_profiles_profiles_proto_rawDescData)
	})
	return file_pkg_services_profiles_profiles_proto_rawDescData
}

var file_pkg_services_profiles_profiles_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_services_profiles_profiles_proto_goTypes = []interface{}{
	(*ValidateCredentialsRequest)(nil), // 0: profiles.ValidateCredentialsRequest
	(*ValidateCredentialsReply)(nil),   // 1: profiles.ValidateCredentialsReply
	(*GetUserRequest)(nil),             // 2: profiles.GetUserRequest
	(*GetUserReply)(nil),               // 3: profiles.GetUserReply
	(*GetUsersRequest)(nil),            // 4: profiles.GetUsersRequest
	(*GetUsersReply)(nil),              // 5: profiles.GetUsersReply
	(*timestamp.Timestamp)(nil),        // 6: google.protobuf.Timestamp
}
var file_pkg_services_profiles_profiles_proto_depIdxs = []int32{
	6, // 0: profiles.GetUserReply.createDate:type_name -> google.protobuf.Timestamp
	6, // 1: profiles.GetUserReply.lastUpdateDate:type_name -> google.protobuf.Timestamp
	3, // 2: profiles.GetUsersReply.users:type_name -> profiles.GetUserReply
	0, // 3: profiles.ProfilesService.ValidateCredentials:input_type -> profiles.ValidateCredentialsRequest
	2, // 4: profiles.ProfilesService.GetUser:input_type -> profiles.GetUserRequest
	4, // 5: profiles.ProfilesService.GetUsers:input_type -> profiles.GetUsersRequest
	2, // 6: profiles.ProfilesService.GetUsersStream:input_type -> profiles.GetUserRequest
	1, // 7: profiles.ProfilesService.ValidateCredentials:output_type -> profiles.ValidateCredentialsReply
	3, // 8: profiles.ProfilesService.GetUser:output_type -> profiles.GetUserReply
	5, // 9: profiles.ProfilesService.GetUsers:output_type -> profiles.GetUsersReply
	3, // 10: profiles.ProfilesService.GetUsersStream:output_type -> profiles.GetUserReply
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_services_profiles_profiles_proto_init() }
func file_pkg_services_profiles_profiles_proto_init() {
	if File_pkg_services_profiles_profiles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_services_profiles_profiles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateCredentialsRequest); i {
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
		file_pkg_services_profiles_profiles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateCredentialsReply); i {
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
		file_pkg_services_profiles_profiles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_pkg_services_profiles_profiles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReply); i {
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
		file_pkg_services_profiles_profiles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRequest); i {
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
		file_pkg_services_profiles_profiles_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersReply); i {
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
			RawDescriptor: file_pkg_services_profiles_profiles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_services_profiles_profiles_proto_goTypes,
		DependencyIndexes: file_pkg_services_profiles_profiles_proto_depIdxs,
		MessageInfos:      file_pkg_services_profiles_profiles_proto_msgTypes,
	}.Build()
	File_pkg_services_profiles_profiles_proto = out.File
	file_pkg_services_profiles_profiles_proto_rawDesc = nil
	file_pkg_services_profiles_profiles_proto_goTypes = nil
	file_pkg_services_profiles_profiles_proto_depIdxs = nil
}
