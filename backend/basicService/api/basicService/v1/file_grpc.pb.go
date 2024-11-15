// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: basicService/v1/file.proto

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
	FileService_PreSignGet_FullMethodName                 = "/basicService.v1.FileService/PreSignGet"
	FileService_PreSignPut_FullMethodName                 = "/basicService.v1.FileService/PreSignPut"
	FileService_ReportUploaded_FullMethodName             = "/basicService.v1.FileService/ReportUploaded"
	FileService_PreSignSlicingPut_FullMethodName          = "/basicService.v1.FileService/PreSignSlicingPut"
	FileService_GetProgressRate4SlicingPut_FullMethodName = "/basicService.v1.FileService/GetProgressRate4SlicingPut"
	FileService_MergeFileParts_FullMethodName             = "/basicService.v1.FileService/MergeFileParts"
	FileService_RemoveFile_FullMethodName                 = "/basicService.v1.FileService/RemoveFile"
	FileService_GetFileInfoById_FullMethodName            = "/basicService.v1.FileService/GetFileInfoById"
)

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	// pre sign a file url for user get it
	PreSignGet(ctx context.Context, in *PreSignGetRequest, opts ...grpc.CallOption) (*PreSignGetResponse, error)
	// pre sign a file url for user put it
	PreSignPut(ctx context.Context, in *PreSignPutRequest, opts ...grpc.CallOption) (*PreSignPutResponse, error)
	// report a file has been uploaded
	ReportUploaded(ctx context.Context, in *ReportUploadedRequest, opts ...grpc.CallOption) (*ReportUploadedResponse, error)
	// pre sign a file url for user put it with slicing
	PreSignSlicingPut(ctx context.Context, in *PreSignSlicingPutRequest, opts ...grpc.CallOption) (*PreSignSlicingPutResponse, error)
	// get upload progress rate for slicing put
	GetProgressRate4SlicingPut(ctx context.Context, in *GetProgressRate4SlicingPutRequest, opts ...grpc.CallOption) (*GetProgressRate4SlicingPutResponse, error)
	// merge a slicing uploading file
	MergeFileParts(ctx context.Context, in *MergeFilePartsRequest, opts ...grpc.CallOption) (*MergeFilePartsResponse, error)
	// remove a file
	RemoveFile(ctx context.Context, in *RemoveFileRequest, opts ...grpc.CallOption) (*RemoveFileResponse, error)
	GetFileInfoById(ctx context.Context, in *GetFileInfoByIdRequest, opts ...grpc.CallOption) (*GetFileInfoByIdResponse, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) PreSignGet(ctx context.Context, in *PreSignGetRequest, opts ...grpc.CallOption) (*PreSignGetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PreSignGetResponse)
	err := c.cc.Invoke(ctx, FileService_PreSignGet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) PreSignPut(ctx context.Context, in *PreSignPutRequest, opts ...grpc.CallOption) (*PreSignPutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PreSignPutResponse)
	err := c.cc.Invoke(ctx, FileService_PreSignPut_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) ReportUploaded(ctx context.Context, in *ReportUploadedRequest, opts ...grpc.CallOption) (*ReportUploadedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReportUploadedResponse)
	err := c.cc.Invoke(ctx, FileService_ReportUploaded_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) PreSignSlicingPut(ctx context.Context, in *PreSignSlicingPutRequest, opts ...grpc.CallOption) (*PreSignSlicingPutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PreSignSlicingPutResponse)
	err := c.cc.Invoke(ctx, FileService_PreSignSlicingPut_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetProgressRate4SlicingPut(ctx context.Context, in *GetProgressRate4SlicingPutRequest, opts ...grpc.CallOption) (*GetProgressRate4SlicingPutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProgressRate4SlicingPutResponse)
	err := c.cc.Invoke(ctx, FileService_GetProgressRate4SlicingPut_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) MergeFileParts(ctx context.Context, in *MergeFilePartsRequest, opts ...grpc.CallOption) (*MergeFilePartsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MergeFilePartsResponse)
	err := c.cc.Invoke(ctx, FileService_MergeFileParts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) RemoveFile(ctx context.Context, in *RemoveFileRequest, opts ...grpc.CallOption) (*RemoveFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveFileResponse)
	err := c.cc.Invoke(ctx, FileService_RemoveFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFileInfoById(ctx context.Context, in *GetFileInfoByIdRequest, opts ...grpc.CallOption) (*GetFileInfoByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFileInfoByIdResponse)
	err := c.cc.Invoke(ctx, FileService_GetFileInfoById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility.
type FileServiceServer interface {
	// pre sign a file url for user get it
	PreSignGet(context.Context, *PreSignGetRequest) (*PreSignGetResponse, error)
	// pre sign a file url for user put it
	PreSignPut(context.Context, *PreSignPutRequest) (*PreSignPutResponse, error)
	// report a file has been uploaded
	ReportUploaded(context.Context, *ReportUploadedRequest) (*ReportUploadedResponse, error)
	// pre sign a file url for user put it with slicing
	PreSignSlicingPut(context.Context, *PreSignSlicingPutRequest) (*PreSignSlicingPutResponse, error)
	// get upload progress rate for slicing put
	GetProgressRate4SlicingPut(context.Context, *GetProgressRate4SlicingPutRequest) (*GetProgressRate4SlicingPutResponse, error)
	// merge a slicing uploading file
	MergeFileParts(context.Context, *MergeFilePartsRequest) (*MergeFilePartsResponse, error)
	// remove a file
	RemoveFile(context.Context, *RemoveFileRequest) (*RemoveFileResponse, error)
	GetFileInfoById(context.Context, *GetFileInfoByIdRequest) (*GetFileInfoByIdResponse, error)
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFileServiceServer struct{}

func (UnimplementedFileServiceServer) PreSignGet(context.Context, *PreSignGetRequest) (*PreSignGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreSignGet not implemented")
}
func (UnimplementedFileServiceServer) PreSignPut(context.Context, *PreSignPutRequest) (*PreSignPutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreSignPut not implemented")
}
func (UnimplementedFileServiceServer) ReportUploaded(context.Context, *ReportUploadedRequest) (*ReportUploadedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportUploaded not implemented")
}
func (UnimplementedFileServiceServer) PreSignSlicingPut(context.Context, *PreSignSlicingPutRequest) (*PreSignSlicingPutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreSignSlicingPut not implemented")
}
func (UnimplementedFileServiceServer) GetProgressRate4SlicingPut(context.Context, *GetProgressRate4SlicingPutRequest) (*GetProgressRate4SlicingPutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProgressRate4SlicingPut not implemented")
}
func (UnimplementedFileServiceServer) MergeFileParts(context.Context, *MergeFilePartsRequest) (*MergeFilePartsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MergeFileParts not implemented")
}
func (UnimplementedFileServiceServer) RemoveFile(context.Context, *RemoveFileRequest) (*RemoveFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFile not implemented")
}
func (UnimplementedFileServiceServer) GetFileInfoById(context.Context, *GetFileInfoByIdRequest) (*GetFileInfoByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileInfoById not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}
func (UnimplementedFileServiceServer) testEmbeddedByValue()                     {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	// If the following call pancis, it indicates UnimplementedFileServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_PreSignGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreSignGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).PreSignGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_PreSignGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).PreSignGet(ctx, req.(*PreSignGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_PreSignPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreSignPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).PreSignPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_PreSignPut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).PreSignPut(ctx, req.(*PreSignPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_ReportUploaded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportUploadedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).ReportUploaded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_ReportUploaded_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).ReportUploaded(ctx, req.(*ReportUploadedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_PreSignSlicingPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreSignSlicingPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).PreSignSlicingPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_PreSignSlicingPut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).PreSignSlicingPut(ctx, req.(*PreSignSlicingPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetProgressRate4SlicingPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProgressRate4SlicingPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetProgressRate4SlicingPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_GetProgressRate4SlicingPut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetProgressRate4SlicingPut(ctx, req.(*GetProgressRate4SlicingPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_MergeFileParts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeFilePartsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).MergeFileParts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_MergeFileParts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).MergeFileParts(ctx, req.(*MergeFilePartsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_RemoveFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).RemoveFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_RemoveFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).RemoveFile(ctx, req.(*RemoveFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFileInfoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileInfoByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFileInfoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_GetFileInfoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFileInfoById(ctx, req.(*GetFileInfoByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "basicService.v1.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PreSignGet",
			Handler:    _FileService_PreSignGet_Handler,
		},
		{
			MethodName: "PreSignPut",
			Handler:    _FileService_PreSignPut_Handler,
		},
		{
			MethodName: "ReportUploaded",
			Handler:    _FileService_ReportUploaded_Handler,
		},
		{
			MethodName: "PreSignSlicingPut",
			Handler:    _FileService_PreSignSlicingPut_Handler,
		},
		{
			MethodName: "GetProgressRate4SlicingPut",
			Handler:    _FileService_GetProgressRate4SlicingPut_Handler,
		},
		{
			MethodName: "MergeFileParts",
			Handler:    _FileService_MergeFileParts_Handler,
		},
		{
			MethodName: "RemoveFile",
			Handler:    _FileService_RemoveFile_Handler,
		},
		{
			MethodName: "GetFileInfoById",
			Handler:    _FileService_GetFileInfoById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basicService/v1/file.proto",
}