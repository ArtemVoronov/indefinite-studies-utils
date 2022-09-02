// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/services/posts/posts.proto

package posts

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

// PostsServiceClient is the client API for PostsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostsServiceClient interface {
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostReply, error)
	GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsReply, error)
	GetPostsStream(ctx context.Context, opts ...grpc.CallOption) (PostsService_GetPostsStreamClient, error)
	GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentReply, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsReply, error)
	GetCommentsStream(ctx context.Context, opts ...grpc.CallOption) (PostsService_GetCommentsStreamClient, error)
}

type postsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostsServiceClient(cc grpc.ClientConnInterface) PostsServiceClient {
	return &postsServiceClient{cc}
}

func (c *postsServiceClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostReply, error) {
	out := new(GetPostReply)
	err := c.cc.Invoke(ctx, "/posts.PostsService/GetPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsReply, error) {
	out := new(GetPostsReply)
	err := c.cc.Invoke(ctx, "/posts.PostsService/GetPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) GetPostsStream(ctx context.Context, opts ...grpc.CallOption) (PostsService_GetPostsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &PostsService_ServiceDesc.Streams[0], "/posts.PostsService/GetPostsStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &postsServiceGetPostsStreamClient{stream}
	return x, nil
}

type PostsService_GetPostsStreamClient interface {
	Send(*GetPostRequest) error
	Recv() (*GetPostReply, error)
	grpc.ClientStream
}

type postsServiceGetPostsStreamClient struct {
	grpc.ClientStream
}

func (x *postsServiceGetPostsStreamClient) Send(m *GetPostRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *postsServiceGetPostsStreamClient) Recv() (*GetPostReply, error) {
	m := new(GetPostReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *postsServiceClient) GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentReply, error) {
	out := new(GetCommentReply)
	err := c.cc.Invoke(ctx, "/posts.PostsService/GetComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsReply, error) {
	out := new(GetCommentsReply)
	err := c.cc.Invoke(ctx, "/posts.PostsService/GetComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsServiceClient) GetCommentsStream(ctx context.Context, opts ...grpc.CallOption) (PostsService_GetCommentsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &PostsService_ServiceDesc.Streams[1], "/posts.PostsService/GetCommentsStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &postsServiceGetCommentsStreamClient{stream}
	return x, nil
}

type PostsService_GetCommentsStreamClient interface {
	Send(*GetCommentRequest) error
	Recv() (*GetCommentReply, error)
	grpc.ClientStream
}

type postsServiceGetCommentsStreamClient struct {
	grpc.ClientStream
}

func (x *postsServiceGetCommentsStreamClient) Send(m *GetCommentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *postsServiceGetCommentsStreamClient) Recv() (*GetCommentReply, error) {
	m := new(GetCommentReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PostsServiceServer is the server API for PostsService service.
// All implementations must embed UnimplementedPostsServiceServer
// for forward compatibility
type PostsServiceServer interface {
	GetPost(context.Context, *GetPostRequest) (*GetPostReply, error)
	GetPosts(context.Context, *GetPostsRequest) (*GetPostsReply, error)
	GetPostsStream(PostsService_GetPostsStreamServer) error
	GetComment(context.Context, *GetCommentRequest) (*GetCommentReply, error)
	GetComments(context.Context, *GetCommentsRequest) (*GetCommentsReply, error)
	GetCommentsStream(PostsService_GetCommentsStreamServer) error
	mustEmbedUnimplementedPostsServiceServer()
}

// UnimplementedPostsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostsServiceServer struct {
}

func (UnimplementedPostsServiceServer) GetPost(context.Context, *GetPostRequest) (*GetPostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPostsServiceServer) GetPosts(context.Context, *GetPostsRequest) (*GetPostsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPosts not implemented")
}
func (UnimplementedPostsServiceServer) GetPostsStream(PostsService_GetPostsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPostsStream not implemented")
}
func (UnimplementedPostsServiceServer) GetComment(context.Context, *GetCommentRequest) (*GetCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedPostsServiceServer) GetComments(context.Context, *GetCommentsRequest) (*GetCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedPostsServiceServer) GetCommentsStream(PostsService_GetCommentsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCommentsStream not implemented")
}
func (UnimplementedPostsServiceServer) mustEmbedUnimplementedPostsServiceServer() {}

// UnsafePostsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostsServiceServer will
// result in compilation errors.
type UnsafePostsServiceServer interface {
	mustEmbedUnimplementedPostsServiceServer()
}

func RegisterPostsServiceServer(s grpc.ServiceRegistrar, srv PostsServiceServer) {
	s.RegisterService(&PostsService_ServiceDesc, srv)
}

func _PostsService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts.PostsService/GetPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_GetPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).GetPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts.PostsService/GetPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).GetPosts(ctx, req.(*GetPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_GetPostsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PostsServiceServer).GetPostsStream(&postsServiceGetPostsStreamServer{stream})
}

type PostsService_GetPostsStreamServer interface {
	Send(*GetPostReply) error
	Recv() (*GetPostRequest, error)
	grpc.ServerStream
}

type postsServiceGetPostsStreamServer struct {
	grpc.ServerStream
}

func (x *postsServiceGetPostsStreamServer) Send(m *GetPostReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *postsServiceGetPostsStreamServer) Recv() (*GetPostRequest, error) {
	m := new(GetPostRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PostsService_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts.PostsService/GetComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).GetComment(ctx, req.(*GetCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServiceServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posts.PostsService/GetComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServiceServer).GetComments(ctx, req.(*GetCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostsService_GetCommentsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PostsServiceServer).GetCommentsStream(&postsServiceGetCommentsStreamServer{stream})
}

type PostsService_GetCommentsStreamServer interface {
	Send(*GetCommentReply) error
	Recv() (*GetCommentRequest, error)
	grpc.ServerStream
}

type postsServiceGetCommentsStreamServer struct {
	grpc.ServerStream
}

func (x *postsServiceGetCommentsStreamServer) Send(m *GetCommentReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *postsServiceGetCommentsStreamServer) Recv() (*GetCommentRequest, error) {
	m := new(GetCommentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PostsService_ServiceDesc is the grpc.ServiceDesc for PostsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "posts.PostsService",
	HandlerType: (*PostsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPost",
			Handler:    _PostsService_GetPost_Handler,
		},
		{
			MethodName: "GetPosts",
			Handler:    _PostsService_GetPosts_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _PostsService_GetComment_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _PostsService_GetComments_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPostsStream",
			Handler:       _PostsService_GetPostsStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetCommentsStream",
			Handler:       _PostsService_GetCommentsStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/services/posts/posts.proto",
}
