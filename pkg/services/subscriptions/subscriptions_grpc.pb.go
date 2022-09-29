// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/services/subscriptions/subscriptions.proto

package subscriptions

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

// SubscriptionsServiceClient is the client API for SubscriptionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionsServiceClient interface {
	PutEvent(ctx context.Context, in *PutEventRequest, opts ...grpc.CallOption) (*PutEventReply, error)
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventReply, error)
}

type subscriptionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionsServiceClient(cc grpc.ClientConnInterface) SubscriptionsServiceClient {
	return &subscriptionsServiceClient{cc}
}

func (c *subscriptionsServiceClient) PutEvent(ctx context.Context, in *PutEventRequest, opts ...grpc.CallOption) (*PutEventReply, error) {
	out := new(PutEventReply)
	err := c.cc.Invoke(ctx, "/subscriptions.SubscriptionsService/PutEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionsServiceClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventReply, error) {
	out := new(GetEventReply)
	err := c.cc.Invoke(ctx, "/subscriptions.SubscriptionsService/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionsServiceServer is the server API for SubscriptionsService service.
// All implementations must embed UnimplementedSubscriptionsServiceServer
// for forward compatibility
type SubscriptionsServiceServer interface {
	PutEvent(context.Context, *PutEventRequest) (*PutEventReply, error)
	GetEvent(context.Context, *GetEventRequest) (*GetEventReply, error)
	mustEmbedUnimplementedSubscriptionsServiceServer()
}

// UnimplementedSubscriptionsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriptionsServiceServer struct {
}

func (UnimplementedSubscriptionsServiceServer) PutEvent(context.Context, *PutEventRequest) (*PutEventReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutEvent not implemented")
}
func (UnimplementedSubscriptionsServiceServer) GetEvent(context.Context, *GetEventRequest) (*GetEventReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedSubscriptionsServiceServer) mustEmbedUnimplementedSubscriptionsServiceServer() {}

// UnsafeSubscriptionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionsServiceServer will
// result in compilation errors.
type UnsafeSubscriptionsServiceServer interface {
	mustEmbedUnimplementedSubscriptionsServiceServer()
}

func RegisterSubscriptionsServiceServer(s grpc.ServiceRegistrar, srv SubscriptionsServiceServer) {
	s.RegisterService(&SubscriptionsService_ServiceDesc, srv)
}

func _SubscriptionsService_PutEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionsServiceServer).PutEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscriptions.SubscriptionsService/PutEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionsServiceServer).PutEvent(ctx, req.(*PutEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionsService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionsServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscriptions.SubscriptionsService/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionsServiceServer).GetEvent(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionsService_ServiceDesc is the grpc.ServiceDesc for SubscriptionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subscriptions.SubscriptionsService",
	HandlerType: (*SubscriptionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutEvent",
			Handler:    _SubscriptionsService_PutEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _SubscriptionsService_GetEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/services/subscriptions/subscriptions.proto",
}
