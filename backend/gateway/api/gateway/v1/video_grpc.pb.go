// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: gateway/v1/video.proto

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
	ShortVideoCoreVideoService_PreSign4UploadVideo_FullMethodName     = "/gateway.v1.ShortVideoCoreVideoService/PreSign4UploadVideo"
	ShortVideoCoreVideoService_PreSign4UploadCover_FullMethodName     = "/gateway.v1.ShortVideoCoreVideoService/PreSign4UploadCover"
	ShortVideoCoreVideoService_ReportFinishUpload_FullMethodName      = "/gateway.v1.ShortVideoCoreVideoService/ReportFinishUpload"
	ShortVideoCoreVideoService_ReportVideoFinishUpload_FullMethodName = "/gateway.v1.ShortVideoCoreVideoService/ReportVideoFinishUpload"
	ShortVideoCoreVideoService_FeedShortVideo_FullMethodName          = "/gateway.v1.ShortVideoCoreVideoService/FeedShortVideo"
	ShortVideoCoreVideoService_GetVideoById_FullMethodName            = "/gateway.v1.ShortVideoCoreVideoService/GetVideoById"
	ShortVideoCoreVideoService_ListPublishedVideo_FullMethodName      = "/gateway.v1.ShortVideoCoreVideoService/ListPublishedVideo"
)

// ShortVideoCoreVideoServiceClient is the client API for ShortVideoCoreVideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortVideoCoreVideoServiceClient interface {
	// 预注册上传视频
	PreSign4UploadVideo(ctx context.Context, in *PreSign4UploadVideoRequest, opts ...grpc.CallOption) (*PreSign4UploadVideoResponse, error)
	// 预注册上传封面
	PreSign4UploadCover(ctx context.Context, in *PreSign4UploadRequest, opts ...grpc.CallOption) (*PreSign4UploadResponse, error)
	// 通用确认上传完成
	ReportFinishUpload(ctx context.Context, in *ReportFinishUploadRequest, opts ...grpc.CallOption) (*ReportFinishUploadResponse, error)
	// 确认视频上传完成
	ReportVideoFinishUpload(ctx context.Context, in *ReportVideoFinishUploadRequest, opts ...grpc.CallOption) (*ReportVideoFinishUploadResponse, error)
	// 刷视频
	FeedShortVideo(ctx context.Context, in *FeedShortVideoRequest, opts ...grpc.CallOption) (*FeedShortVideoResponse, error)
	// 获取视频信息
	GetVideoById(ctx context.Context, in *GetVideoByIdRequest, opts ...grpc.CallOption) (*GetVideoByIdResponse, error)
	// 获取当前用户的发布视频列表
	ListPublishedVideo(ctx context.Context, in *ListPublishedVideoRequest, opts ...grpc.CallOption) (*ListPublishedVideoResponse, error)
}

type shortVideoCoreVideoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShortVideoCoreVideoServiceClient(cc grpc.ClientConnInterface) ShortVideoCoreVideoServiceClient {
	return &shortVideoCoreVideoServiceClient{cc}
}

func (c *shortVideoCoreVideoServiceClient) PreSign4UploadVideo(ctx context.Context, in *PreSign4UploadVideoRequest, opts ...grpc.CallOption) (*PreSign4UploadVideoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PreSign4UploadVideoResponse)
	err := c.cc.Invoke(ctx, ShortVideoCoreVideoService_PreSign4UploadVideo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortVideoCoreVideoServiceClient) PreSign4UploadCover(ctx context.Context, in *PreSign4UploadRequest, opts ...grpc.CallOption) (*PreSign4UploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PreSign4UploadResponse)
	err := c.cc.Invoke(ctx, ShortVideoCoreVideoService_PreSign4UploadCover_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortVideoCoreVideoServiceClient) ReportFinishUpload(ctx context.Context, in *ReportFinishUploadRequest, opts ...grpc.CallOption) (*ReportFinishUploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReportFinishUploadResponse)
	err := c.cc.Invoke(ctx, ShortVideoCoreVideoService_ReportFinishUpload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortVideoCoreVideoServiceClient) ReportVideoFinishUpload(ctx context.Context, in *ReportVideoFinishUploadRequest, opts ...grpc.CallOption) (*ReportVideoFinishUploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReportVideoFinishUploadResponse)
	err := c.cc.Invoke(ctx, ShortVideoCoreVideoService_ReportVideoFinishUpload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortVideoCoreVideoServiceClient) FeedShortVideo(ctx context.Context, in *FeedShortVideoRequest, opts ...grpc.CallOption) (*FeedShortVideoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FeedShortVideoResponse)
	err := c.cc.Invoke(ctx, ShortVideoCoreVideoService_FeedShortVideo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortVideoCoreVideoServiceClient) GetVideoById(ctx context.Context, in *GetVideoByIdRequest, opts ...grpc.CallOption) (*GetVideoByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVideoByIdResponse)
	err := c.cc.Invoke(ctx, ShortVideoCoreVideoService_GetVideoById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortVideoCoreVideoServiceClient) ListPublishedVideo(ctx context.Context, in *ListPublishedVideoRequest, opts ...grpc.CallOption) (*ListPublishedVideoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPublishedVideoResponse)
	err := c.cc.Invoke(ctx, ShortVideoCoreVideoService_ListPublishedVideo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortVideoCoreVideoServiceServer is the server API for ShortVideoCoreVideoService service.
// All implementations must embed UnimplementedShortVideoCoreVideoServiceServer
// for forward compatibility.
type ShortVideoCoreVideoServiceServer interface {
	// 预注册上传视频
	PreSign4UploadVideo(context.Context, *PreSign4UploadVideoRequest) (*PreSign4UploadVideoResponse, error)
	// 预注册上传封面
	PreSign4UploadCover(context.Context, *PreSign4UploadRequest) (*PreSign4UploadResponse, error)
	// 通用确认上传完成
	ReportFinishUpload(context.Context, *ReportFinishUploadRequest) (*ReportFinishUploadResponse, error)
	// 确认视频上传完成
	ReportVideoFinishUpload(context.Context, *ReportVideoFinishUploadRequest) (*ReportVideoFinishUploadResponse, error)
	// 刷视频
	FeedShortVideo(context.Context, *FeedShortVideoRequest) (*FeedShortVideoResponse, error)
	// 获取视频信息
	GetVideoById(context.Context, *GetVideoByIdRequest) (*GetVideoByIdResponse, error)
	// 获取当前用户的发布视频列表
	ListPublishedVideo(context.Context, *ListPublishedVideoRequest) (*ListPublishedVideoResponse, error)
	mustEmbedUnimplementedShortVideoCoreVideoServiceServer()
}

// UnimplementedShortVideoCoreVideoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShortVideoCoreVideoServiceServer struct{}

func (UnimplementedShortVideoCoreVideoServiceServer) PreSign4UploadVideo(context.Context, *PreSign4UploadVideoRequest) (*PreSign4UploadVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreSign4UploadVideo not implemented")
}
func (UnimplementedShortVideoCoreVideoServiceServer) PreSign4UploadCover(context.Context, *PreSign4UploadRequest) (*PreSign4UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreSign4UploadCover not implemented")
}
func (UnimplementedShortVideoCoreVideoServiceServer) ReportFinishUpload(context.Context, *ReportFinishUploadRequest) (*ReportFinishUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportFinishUpload not implemented")
}
func (UnimplementedShortVideoCoreVideoServiceServer) ReportVideoFinishUpload(context.Context, *ReportVideoFinishUploadRequest) (*ReportVideoFinishUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportVideoFinishUpload not implemented")
}
func (UnimplementedShortVideoCoreVideoServiceServer) FeedShortVideo(context.Context, *FeedShortVideoRequest) (*FeedShortVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedShortVideo not implemented")
}
func (UnimplementedShortVideoCoreVideoServiceServer) GetVideoById(context.Context, *GetVideoByIdRequest) (*GetVideoByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoById not implemented")
}
func (UnimplementedShortVideoCoreVideoServiceServer) ListPublishedVideo(context.Context, *ListPublishedVideoRequest) (*ListPublishedVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPublishedVideo not implemented")
}
func (UnimplementedShortVideoCoreVideoServiceServer) mustEmbedUnimplementedShortVideoCoreVideoServiceServer() {
}
func (UnimplementedShortVideoCoreVideoServiceServer) testEmbeddedByValue() {}

// UnsafeShortVideoCoreVideoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortVideoCoreVideoServiceServer will
// result in compilation errors.
type UnsafeShortVideoCoreVideoServiceServer interface {
	mustEmbedUnimplementedShortVideoCoreVideoServiceServer()
}

func RegisterShortVideoCoreVideoServiceServer(s grpc.ServiceRegistrar, srv ShortVideoCoreVideoServiceServer) {
	// If the following call pancis, it indicates UnimplementedShortVideoCoreVideoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ShortVideoCoreVideoService_ServiceDesc, srv)
}

func _ShortVideoCoreVideoService_PreSign4UploadVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreSign4UploadVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortVideoCoreVideoServiceServer).PreSign4UploadVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortVideoCoreVideoService_PreSign4UploadVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortVideoCoreVideoServiceServer).PreSign4UploadVideo(ctx, req.(*PreSign4UploadVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortVideoCoreVideoService_PreSign4UploadCover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreSign4UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortVideoCoreVideoServiceServer).PreSign4UploadCover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortVideoCoreVideoService_PreSign4UploadCover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortVideoCoreVideoServiceServer).PreSign4UploadCover(ctx, req.(*PreSign4UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortVideoCoreVideoService_ReportFinishUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportFinishUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortVideoCoreVideoServiceServer).ReportFinishUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortVideoCoreVideoService_ReportFinishUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortVideoCoreVideoServiceServer).ReportFinishUpload(ctx, req.(*ReportFinishUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortVideoCoreVideoService_ReportVideoFinishUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportVideoFinishUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortVideoCoreVideoServiceServer).ReportVideoFinishUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortVideoCoreVideoService_ReportVideoFinishUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortVideoCoreVideoServiceServer).ReportVideoFinishUpload(ctx, req.(*ReportVideoFinishUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortVideoCoreVideoService_FeedShortVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedShortVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortVideoCoreVideoServiceServer).FeedShortVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortVideoCoreVideoService_FeedShortVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortVideoCoreVideoServiceServer).FeedShortVideo(ctx, req.(*FeedShortVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortVideoCoreVideoService_GetVideoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortVideoCoreVideoServiceServer).GetVideoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortVideoCoreVideoService_GetVideoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortVideoCoreVideoServiceServer).GetVideoById(ctx, req.(*GetVideoByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortVideoCoreVideoService_ListPublishedVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPublishedVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortVideoCoreVideoServiceServer).ListPublishedVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortVideoCoreVideoService_ListPublishedVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortVideoCoreVideoServiceServer).ListPublishedVideo(ctx, req.(*ListPublishedVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortVideoCoreVideoService_ServiceDesc is the grpc.ServiceDesc for ShortVideoCoreVideoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortVideoCoreVideoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.v1.ShortVideoCoreVideoService",
	HandlerType: (*ShortVideoCoreVideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PreSign4UploadVideo",
			Handler:    _ShortVideoCoreVideoService_PreSign4UploadVideo_Handler,
		},
		{
			MethodName: "PreSign4UploadCover",
			Handler:    _ShortVideoCoreVideoService_PreSign4UploadCover_Handler,
		},
		{
			MethodName: "ReportFinishUpload",
			Handler:    _ShortVideoCoreVideoService_ReportFinishUpload_Handler,
		},
		{
			MethodName: "ReportVideoFinishUpload",
			Handler:    _ShortVideoCoreVideoService_ReportVideoFinishUpload_Handler,
		},
		{
			MethodName: "FeedShortVideo",
			Handler:    _ShortVideoCoreVideoService_FeedShortVideo_Handler,
		},
		{
			MethodName: "GetVideoById",
			Handler:    _ShortVideoCoreVideoService_GetVideoById_Handler,
		},
		{
			MethodName: "ListPublishedVideo",
			Handler:    _ShortVideoCoreVideoService_ListPublishedVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway/v1/video.proto",
}
