// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: gateway/v1/comment.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CommentService_CreateComment_FullMethodName     = "/gateway.v1.CommentService/CreateComment"
	CommentService_RemoveComment_FullMethodName     = "/gateway.v1.CommentService/RemoveComment"
	CommentService_ListComment4Video_FullMethodName = "/gateway.v1.CommentService/ListComment4Video"
	CommentService_ListChildComment_FullMethodName  = "/gateway.v1.CommentService/ListChildComment"
)

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 评论
type CommentServiceClient interface {
	// 创建评论
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	// 删除评论
	RemoveComment(ctx context.Context, in *RemoveCommentRequest, opts ...grpc.CallOption) (*RemoveCommentResponse, error)
	// 列出视频的评论
	ListComment4Video(ctx context.Context, in *ListComment4VideoRequest, opts ...grpc.CallOption) (*ListComment4VideoResponse, error)
	ListChildComment(ctx context.Context, in *ListChildCommentRequest, opts ...grpc.CallOption) (*ListChildCommentResponse, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_CreateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) RemoveComment(ctx context.Context, in *RemoveCommentRequest, opts ...grpc.CallOption) (*RemoveCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_RemoveComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) ListComment4Video(ctx context.Context, in *ListComment4VideoRequest, opts ...grpc.CallOption) (*ListComment4VideoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListComment4VideoResponse)
	err := c.cc.Invoke(ctx, CommentService_ListComment4Video_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) ListChildComment(ctx context.Context, in *ListChildCommentRequest, opts ...grpc.CallOption) (*ListChildCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListChildCommentResponse)
	err := c.cc.Invoke(ctx, CommentService_ListChildComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
// All implementations must embed UnimplementedCommentServiceServer
// for forward compatibility.
//
// 评论
type CommentServiceServer interface {
	// 创建评论
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	// 删除评论
	RemoveComment(context.Context, *RemoveCommentRequest) (*RemoveCommentResponse, error)
	// 列出视频的评论
	ListComment4Video(context.Context, *ListComment4VideoRequest) (*ListComment4VideoResponse, error)
	ListChildComment(context.Context, *ListChildCommentRequest) (*ListChildCommentResponse, error)
	mustEmbedUnimplementedCommentServiceServer()
}

// UnimplementedCommentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommentServiceServer struct{}

func (UnimplementedCommentServiceServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentServiceServer) RemoveComment(context.Context, *RemoveCommentRequest) (*RemoveCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveComment not implemented")
}
func (UnimplementedCommentServiceServer) ListComment4Video(context.Context, *ListComment4VideoRequest) (*ListComment4VideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComment4Video not implemented")
}
func (UnimplementedCommentServiceServer) ListChildComment(context.Context, *ListChildCommentRequest) (*ListChildCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChildComment not implemented")
}
func (UnimplementedCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {}
func (UnimplementedCommentServiceServer) testEmbeddedByValue()                        {}

// UnsafeCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServiceServer will
// result in compilation errors.
type UnsafeCommentServiceServer interface {
	mustEmbedUnimplementedCommentServiceServer()
}

func RegisterCommentServiceServer(s grpc.ServiceRegistrar, srv CommentServiceServer) {
	// If the following call pancis, it indicates UnimplementedCommentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CommentService_ServiceDesc, srv)
}

func _CommentService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_RemoveComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).RemoveComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_RemoveComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).RemoveComment(ctx, req.(*RemoveCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_ListComment4Video_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListComment4VideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).ListComment4Video(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_ListComment4Video_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).ListComment4Video(ctx, req.(*ListComment4VideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_ListChildComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChildCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).ListChildComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_ListChildComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).ListChildComment(ctx, req.(*ListChildCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentService_ServiceDesc is the grpc.ServiceDesc for CommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.v1.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComment",
			Handler:    _CommentService_CreateComment_Handler,
		},
		{
			MethodName: "RemoveComment",
			Handler:    _CommentService_RemoveComment_Handler,
		},
		{
			MethodName: "ListComment4Video",
			Handler:    _CommentService_ListComment4Video_Handler,
		},
		{
			MethodName: "ListChildComment",
			Handler:    _CommentService_ListChildComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway/v1/comment.proto",
}