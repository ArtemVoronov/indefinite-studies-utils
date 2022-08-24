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

// ProfilesServiceServer is the server API for ProfilesService service.
// All implementations must embed UnimplementedProfilesServiceServer
// for forward compatibility
type ProfilesServiceServer interface {
	ValidateCredentials(context.Context, *ValidateCredentialsRequest) (*ValidateCredentialsReply, error)
	mustEmbedUnimplementedProfilesServiceServer()
}

// UnimplementedProfilesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfilesServiceServer struct {
}

func (UnimplementedProfilesServiceServer) ValidateCredentials(context.Context, *ValidateCredentialsRequest) (*ValidateCredentialsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCredentials not implemented")
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/services/profiles/profiles.proto",
}
