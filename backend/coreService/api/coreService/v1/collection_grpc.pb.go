// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: coreService/v1/collection.proto

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
	CollectionService_CreateCollection_FullMethodName          = "/coreService.v1.CollectionService/CreateCollection"
	CollectionService_GetCollectionById_FullMethodName         = "/coreService.v1.CollectionService/GetCollectionById"
	CollectionService_RemoveCollection_FullMethodName          = "/coreService.v1.CollectionService/RemoveCollection"
	CollectionService_ListCollection_FullMethodName            = "/coreService.v1.CollectionService/ListCollection"
	CollectionService_UpdateCollection_FullMethodName          = "/coreService.v1.CollectionService/UpdateCollection"
	CollectionService_AddVideo2Collection_FullMethodName       = "/coreService.v1.CollectionService/AddVideo2Collection"
	CollectionService_RemoveVideoFromCollection_FullMethodName = "/coreService.v1.CollectionService/RemoveVideoFromCollection"
	CollectionService_ListCollectionVideo_FullMethodName       = "/coreService.v1.CollectionService/ListCollectionVideo"
	CollectionService_IsCollected_FullMethodName               = "/coreService.v1.CollectionService/IsCollected"
	CollectionService_CountCollect4Video_FullMethodName        = "/coreService.v1.CollectionService/CountCollect4Video"
)

// CollectionServiceClient is the client API for CollectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectionServiceClient interface {
	CreateCollection(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error)
	GetCollectionById(ctx context.Context, in *GetCollectionByIdRequest, opts ...grpc.CallOption) (*GetCollectionByIdResponse, error)
	RemoveCollection(ctx context.Context, in *RemoveCollectionRequest, opts ...grpc.CallOption) (*RemoveCollectionResponse, error)
	ListCollection(ctx context.Context, in *ListCollectionRequest, opts ...grpc.CallOption) (*ListCollectionResponse, error)
	UpdateCollection(ctx context.Context, in *UpdateCollectionRequest, opts ...grpc.CallOption) (*UpdateCollectionResponse, error)
	AddVideo2Collection(ctx context.Context, in *AddVideo2CollectionRequest, opts ...grpc.CallOption) (*AddVideo2CollectionResponse, error)
	RemoveVideoFromCollection(ctx context.Context, in *RemoveVideoFromCollectionRequest, opts ...grpc.CallOption) (*RemoveVideoFromCollectionResponse, error)
	ListCollectionVideo(ctx context.Context, in *ListCollectionVideoRequest, opts ...grpc.CallOption) (*ListCollectionVideoResponse, error)
	IsCollected(ctx context.Context, in *IsCollectedRequest, opts ...grpc.CallOption) (*IsCollectedResponse, error)
	CountCollect4Video(ctx context.Context, in *CountCollect4VideoRequest, opts ...grpc.CallOption) (*CountCollect4VideoResponse, error)
}

type collectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectionServiceClient(cc grpc.ClientConnInterface) CollectionServiceClient {
	return &collectionServiceClient{cc}
}

func (c *collectionServiceClient) CreateCollection(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCollectionResponse)
	err := c.cc.Invoke(ctx, CollectionService_CreateCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) GetCollectionById(ctx context.Context, in *GetCollectionByIdRequest, opts ...grpc.CallOption) (*GetCollectionByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCollectionByIdResponse)
	err := c.cc.Invoke(ctx, CollectionService_GetCollectionById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) RemoveCollection(ctx context.Context, in *RemoveCollectionRequest, opts ...grpc.CallOption) (*RemoveCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveCollectionResponse)
	err := c.cc.Invoke(ctx, CollectionService_RemoveCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) ListCollection(ctx context.Context, in *ListCollectionRequest, opts ...grpc.CallOption) (*ListCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCollectionResponse)
	err := c.cc.Invoke(ctx, CollectionService_ListCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) UpdateCollection(ctx context.Context, in *UpdateCollectionRequest, opts ...grpc.CallOption) (*UpdateCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCollectionResponse)
	err := c.cc.Invoke(ctx, CollectionService_UpdateCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) AddVideo2Collection(ctx context.Context, in *AddVideo2CollectionRequest, opts ...grpc.CallOption) (*AddVideo2CollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddVideo2CollectionResponse)
	err := c.cc.Invoke(ctx, CollectionService_AddVideo2Collection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) RemoveVideoFromCollection(ctx context.Context, in *RemoveVideoFromCollectionRequest, opts ...grpc.CallOption) (*RemoveVideoFromCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveVideoFromCollectionResponse)
	err := c.cc.Invoke(ctx, CollectionService_RemoveVideoFromCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) ListCollectionVideo(ctx context.Context, in *ListCollectionVideoRequest, opts ...grpc.CallOption) (*ListCollectionVideoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCollectionVideoResponse)
	err := c.cc.Invoke(ctx, CollectionService_ListCollectionVideo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) IsCollected(ctx context.Context, in *IsCollectedRequest, opts ...grpc.CallOption) (*IsCollectedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsCollectedResponse)
	err := c.cc.Invoke(ctx, CollectionService_IsCollected_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) CountCollect4Video(ctx context.Context, in *CountCollect4VideoRequest, opts ...grpc.CallOption) (*CountCollect4VideoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountCollect4VideoResponse)
	err := c.cc.Invoke(ctx, CollectionService_CountCollect4Video_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectionServiceServer is the server API for CollectionService service.
// All implementations must embed UnimplementedCollectionServiceServer
// for forward compatibility.
type CollectionServiceServer interface {
	CreateCollection(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error)
	GetCollectionById(context.Context, *GetCollectionByIdRequest) (*GetCollectionByIdResponse, error)
	RemoveCollection(context.Context, *RemoveCollectionRequest) (*RemoveCollectionResponse, error)
	ListCollection(context.Context, *ListCollectionRequest) (*ListCollectionResponse, error)
	UpdateCollection(context.Context, *UpdateCollectionRequest) (*UpdateCollectionResponse, error)
	AddVideo2Collection(context.Context, *AddVideo2CollectionRequest) (*AddVideo2CollectionResponse, error)
	RemoveVideoFromCollection(context.Context, *RemoveVideoFromCollectionRequest) (*RemoveVideoFromCollectionResponse, error)
	ListCollectionVideo(context.Context, *ListCollectionVideoRequest) (*ListCollectionVideoResponse, error)
	IsCollected(context.Context, *IsCollectedRequest) (*IsCollectedResponse, error)
	CountCollect4Video(context.Context, *CountCollect4VideoRequest) (*CountCollect4VideoResponse, error)
	mustEmbedUnimplementedCollectionServiceServer()
}

// UnimplementedCollectionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCollectionServiceServer struct{}

func (UnimplementedCollectionServiceServer) CreateCollection(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (UnimplementedCollectionServiceServer) GetCollectionById(context.Context, *GetCollectionByIdRequest) (*GetCollectionByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollectionById not implemented")
}
func (UnimplementedCollectionServiceServer) RemoveCollection(context.Context, *RemoveCollectionRequest) (*RemoveCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCollection not implemented")
}
func (UnimplementedCollectionServiceServer) ListCollection(context.Context, *ListCollectionRequest) (*ListCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollection not implemented")
}
func (UnimplementedCollectionServiceServer) UpdateCollection(context.Context, *UpdateCollectionRequest) (*UpdateCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCollection not implemented")
}
func (UnimplementedCollectionServiceServer) AddVideo2Collection(context.Context, *AddVideo2CollectionRequest) (*AddVideo2CollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVideo2Collection not implemented")
}
func (UnimplementedCollectionServiceServer) RemoveVideoFromCollection(context.Context, *RemoveVideoFromCollectionRequest) (*RemoveVideoFromCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveVideoFromCollection not implemented")
}
func (UnimplementedCollectionServiceServer) ListCollectionVideo(context.Context, *ListCollectionVideoRequest) (*ListCollectionVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollectionVideo not implemented")
}
func (UnimplementedCollectionServiceServer) IsCollected(context.Context, *IsCollectedRequest) (*IsCollectedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsCollected not implemented")
}
func (UnimplementedCollectionServiceServer) CountCollect4Video(context.Context, *CountCollect4VideoRequest) (*CountCollect4VideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountCollect4Video not implemented")
}
func (UnimplementedCollectionServiceServer) mustEmbedUnimplementedCollectionServiceServer() {}
func (UnimplementedCollectionServiceServer) testEmbeddedByValue()                           {}

// UnsafeCollectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectionServiceServer will
// result in compilation errors.
type UnsafeCollectionServiceServer interface {
	mustEmbedUnimplementedCollectionServiceServer()
}

func RegisterCollectionServiceServer(s grpc.ServiceRegistrar, srv CollectionServiceServer) {
	// If the following call pancis, it indicates UnimplementedCollectionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CollectionService_ServiceDesc, srv)
}

func _CollectionService_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_CreateCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).CreateCollection(ctx, req.(*CreateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_GetCollectionById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectionByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).GetCollectionById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_GetCollectionById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).GetCollectionById(ctx, req.(*GetCollectionByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_RemoveCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).RemoveCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_RemoveCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).RemoveCollection(ctx, req.(*RemoveCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_ListCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).ListCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_ListCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).ListCollection(ctx, req.(*ListCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_UpdateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).UpdateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_UpdateCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).UpdateCollection(ctx, req.(*UpdateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_AddVideo2Collection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVideo2CollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).AddVideo2Collection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_AddVideo2Collection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).AddVideo2Collection(ctx, req.(*AddVideo2CollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_RemoveVideoFromCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveVideoFromCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).RemoveVideoFromCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_RemoveVideoFromCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).RemoveVideoFromCollection(ctx, req.(*RemoveVideoFromCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_ListCollectionVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCollectionVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).ListCollectionVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_ListCollectionVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).ListCollectionVideo(ctx, req.(*ListCollectionVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_IsCollected_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsCollectedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).IsCollected(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_IsCollected_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).IsCollected(ctx, req.(*IsCollectedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_CountCollect4Video_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountCollect4VideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).CountCollect4Video(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_CountCollect4Video_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).CountCollect4Video(ctx, req.(*CountCollect4VideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CollectionService_ServiceDesc is the grpc.ServiceDesc for CollectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coreService.v1.CollectionService",
	HandlerType: (*CollectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCollection",
			Handler:    _CollectionService_CreateCollection_Handler,
		},
		{
			MethodName: "GetCollectionById",
			Handler:    _CollectionService_GetCollectionById_Handler,
		},
		{
			MethodName: "RemoveCollection",
			Handler:    _CollectionService_RemoveCollection_Handler,
		},
		{
			MethodName: "ListCollection",
			Handler:    _CollectionService_ListCollection_Handler,
		},
		{
			MethodName: "UpdateCollection",
			Handler:    _CollectionService_UpdateCollection_Handler,
		},
		{
			MethodName: "AddVideo2Collection",
			Handler:    _CollectionService_AddVideo2Collection_Handler,
		},
		{
			MethodName: "RemoveVideoFromCollection",
			Handler:    _CollectionService_RemoveVideoFromCollection_Handler,
		},
		{
			MethodName: "ListCollectionVideo",
			Handler:    _CollectionService_ListCollectionVideo_Handler,
		},
		{
			MethodName: "IsCollected",
			Handler:    _CollectionService_IsCollected_Handler,
		},
		{
			MethodName: "CountCollect4Video",
			Handler:    _CollectionService_CountCollect4Video_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coreService/v1/collection.proto",
}
