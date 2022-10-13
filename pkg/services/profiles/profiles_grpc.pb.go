// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/services/profiles/profiles.proto

package profiles

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProfilesServiceClient is the client API for ProfilesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfilesServiceClient interface {
	ValidateCredentials(ctx context.Context, in *ValidateCredentialsRequest, opts ...grpc.CallOption) (*ValidateCredentialsReply, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersReply, error)
	GetUsersStream(ctx context.Context, opts ...grpc.CallOption) (ProfilesService_GetUsersStreamClient, error)
}

type profilesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilesServiceClient(cc grpc.ClientConnInterface) ProfilesServiceClient {
	return &profilesServiceClient{cc}
}

func (c *profilesServiceClient) ValidateCredentials(ctx context.Context, in *ValidateCredentialsRequest, opts ...grpc.CallOption) (*ValidateCredentialsReply, error) {
	out := new(ValidateCredentialsReply)
	err := c.cc.Invoke(ctx, "/profiles.ProfilesService/ValidateCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error) {
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, "/profiles.ProfilesService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersReply, error) {
	out := new(GetUsersReply)
	err := c.cc.Invoke(ctx, "/profiles.ProfilesService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesServiceClient) GetUsersStream(ctx context.Context, opts ...grpc.CallOption) (ProfilesService_GetUsersStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProfilesService_ServiceDesc.Streams[0], "/profiles.ProfilesService/GetUsersStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &profilesServiceGetUsersStreamClient{stream}
	return x, nil
}

type ProfilesService_GetUsersStreamClient interface {
	Send(*GetUserRequest) error
	Recv() (*GetUserReply, error)
	grpc.ClientStream
}

type profilesServiceGetUsersStreamClient struct {
	grpc.ClientStream
}

func (x *profilesServiceGetUsersStreamClient) Send(m *GetUserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *profilesServiceGetUsersStreamClient) Recv() (*GetUserReply, error) {
	m := new(GetUserReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProfilesServiceServer is the server API for ProfilesService service.
// All implementations must embed UnimplementedProfilesServiceServer
// for forward compatibility
type ProfilesServiceServer interface {
	ValidateCredentials(context.Context, *ValidateCredentialsRequest) (*ValidateCredentialsReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersReply, error)
	GetUsersStream(ProfilesService_GetUsersStreamServer) error
	mustEmbedUnimplementedProfilesServiceServer()
}

// UnimplementedProfilesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfilesServiceServer struct {
}

func (UnimplementedProfilesServiceServer) ValidateCredentials(context.Context, *ValidateCredentialsRequest) (*ValidateCredentialsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCredentials not implemented")
}
func (UnimplementedProfilesServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedProfilesServiceServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedProfilesServiceServer) GetUsersStream(ProfilesService_GetUsersStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsersStream not implemented")
}
func (UnimplementedProfilesServiceServer) mustEmbedUnimplementedProfilesServiceServer() {}

// UnsafeProfilesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfilesServiceServer will
// result in compilation errors.
type UnsafeProfilesServiceServer interface {
	mustEmbedUnimplementedProfilesServiceServer()
}

func RegisterProfilesServiceServer(s grpc.ServiceRegistrar, srv ProfilesServiceServer) {
	s.RegisterService(&ProfilesService_ServiceDesc, srv)
}

func _ProfilesService_ValidateCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServiceServer).ValidateCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.ProfilesService/ValidateCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServiceServer).ValidateCredentials(ctx, req.(*ValidateCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilesService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.ProfilesService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilesService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.ProfilesService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilesService_GetUsersStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProfilesServiceServer).GetUsersStream(&profilesServiceGetUsersStreamServer{stream})
}

type ProfilesService_GetUsersStreamServer interface {
	Send(*GetUserReply) error
	Recv() (*GetUserRequest, error)
	grpc.ServerStream
}

type profilesServiceGetUsersStreamServer struct {
	grpc.ServerStream
}

func (x *profilesServiceGetUsersStreamServer) Send(m *GetUserReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *profilesServiceGetUsersStreamServer) Recv() (*GetUserRequest, error) {
	m := new(GetUserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProfilesService_ServiceDesc is the grpc.ServiceDesc for ProfilesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfilesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profiles.ProfilesService",
	HandlerType: (*ProfilesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateCredentials",
			Handler:    _ProfilesService_ValidateCredentials_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _ProfilesService_GetUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _ProfilesService_GetUsers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUsersStream",
			Handler:       _ProfilesService_GetUsersStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/services/profiles/profiles.proto",
}
