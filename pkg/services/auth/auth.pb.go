// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/services/auth/auth.proto

package auth

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

type VerifyTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VerifyTokenRequest) Reset() {
	*x = VerifyTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenRequest) ProtoMessage() {}

func (x *VerifyTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenRequest.ProtoReflect.Descriptor instead.
func (*VerifyTokenRequest) Descriptor() ([]byte, []int) {
	return file_pkg_services_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerifyTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid   bool `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
	IsExpired bool `protobuf:"varint,2,opt,name=isExpired,proto3" json:"isExpired,omitempty"`
}

func (x *VerifyTokenReply) Reset() {
	*x = VerifyTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenReply) ProtoMessage() {}

func (x *VerifyTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenReply.ProtoReflect.Descriptor instead.
func (*VerifyTokenReply) Descriptor() ([]byte, []int) {
	return file_pkg_services_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyTokenReply) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

func (x *VerifyTokenReply) GetIsExpired() bool {
	if x != nil {
		return x.IsExpired
	}
	return false
}

type GetTokenClaimsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetTokenClaimsRequest) Reset() {
	*x = GetTokenClaimsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenClaimsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenClaimsRequest) ProtoMessage() {}

func (x *GetTokenClaimsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenClaimsRequest.ProtoReflect.Descriptor instead.
func (*GetTokenClaimsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_services_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *GetTokenClaimsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetTokenClaimsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetTokenClaimsReply) Reset() {
	*x = GetTokenClaimsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenClaimsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenClaimsReply) ProtoMessage() {}

func (x *GetTokenClaimsReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenClaimsReply.ProtoReflect.Descriptor instead.
func (*GetTokenClaimsReply) Descriptor() ([]byte, []int) {
	return file_pkg_services_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *GetTokenClaimsReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetTokenClaimsReply) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_pkg_services_auth_auth_proto protoreflect.FileDescriptor

var file_pkg_services_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x22, 0x2a, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x4a, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x22, 0x2d, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x39, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x32, 0x9c, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x1b, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6c, 0x61, 0x69, 0x6d,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x74, 0x65, 0x6d, 0x56, 0x6f, 0x72, 0x6f, 0x6e, 0x6f, 0x76,
	0x2f, 0x69, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x2d, 0x73, 0x74, 0x75, 0x64,
	0x69, 0x65, 0x73, 0x2d, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_services_auth_auth_proto_rawDescOnce sync.Once
	file_pkg_services_auth_auth_proto_rawDescData = file_pkg_services_auth_auth_proto_rawDesc
)

func file_pkg_services_auth_auth_proto_rawDescGZIP() []byte {
	file_pkg_services_auth_auth_proto_rawDescOnce.Do(func() {
		file_pkg_services_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_services_auth_auth_proto_rawDescData)
	})
	return file_pkg_services_auth_auth_proto_rawDescData
}

var file_pkg_services_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_services_auth_auth_proto_goTypes = []interface{}{
	(*VerifyTokenRequest)(nil),    // 0: auth.VerifyTokenRequest
	(*VerifyTokenReply)(nil),      // 1: auth.VerifyTokenReply
	(*GetTokenClaimsRequest)(nil), // 2: auth.GetTokenClaimsRequest
	(*GetTokenClaimsReply)(nil),   // 3: auth.GetTokenClaimsReply
}
var file_pkg_services_auth_auth_proto_depIdxs = []int32{
	0, // 0: auth.AuthService.VerifyToken:input_type -> auth.VerifyTokenRequest
	2, // 1: auth.AuthService.GetTokenClaims:input_type -> auth.GetTokenClaimsRequest
	1, // 2: auth.AuthService.VerifyToken:output_type -> auth.VerifyTokenReply
	3, // 3: auth.AuthService.GetTokenClaims:output_type -> auth.GetTokenClaimsReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_services_auth_auth_proto_init() }
func file_pkg_services_auth_auth_proto_init() {
	if File_pkg_services_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_services_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenRequest); i {
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
		file_pkg_services_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenReply); i {
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
		file_pkg_services_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenClaimsRequest); i {
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
		file_pkg_services_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenClaimsReply); i {
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
			RawDescriptor: file_pkg_services_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_services_auth_auth_proto_goTypes,
		DependencyIndexes: file_pkg_services_auth_auth_proto_depIdxs,
		MessageInfos:      file_pkg_services_auth_auth_proto_msgTypes,
	}.Build()
	File_pkg_services_auth_auth_proto = out.File
	file_pkg_services_auth_auth_proto_rawDesc = nil
	file_pkg_services_auth_auth_proto_goTypes = nil
	file_pkg_services_auth_auth_proto_depIdxs = nil
}
