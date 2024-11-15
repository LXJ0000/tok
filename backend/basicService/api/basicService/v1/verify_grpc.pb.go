// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: basicService/v1/verify.proto

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
	AuthService_CreateVerificationCode_FullMethodName   = "/basicService.v1.AuthService/CreateVerificationCode"
	AuthService_ValidateVerificationCode_FullMethodName = "/basicService.v1.AuthService/ValidateVerificationCode"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	CreateVerificationCode(ctx context.Context, in *CreateVerificationCodeRequest, opts ...grpc.CallOption) (*CreateVerificationCodeResponse, error)
	ValidateVerificationCode(ctx context.Context, in *ValidateVerificationCodeRequest, opts ...grpc.CallOption) (*ValidateVerificationCodeResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) CreateVerificationCode(ctx context.Context, in *CreateVerificationCodeRequest, opts ...grpc.CallOption) (*CreateVerificationCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateVerificationCodeResponse)
	err := c.cc.Invoke(ctx, AuthService_CreateVerificationCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ValidateVerificationCode(ctx context.Context, in *ValidateVerificationCodeRequest, opts ...grpc.CallOption) (*ValidateVerificationCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateVerificationCodeResponse)
	err := c.cc.Invoke(ctx, AuthService_ValidateVerificationCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
type AuthServiceServer interface {
	CreateVerificationCode(context.Context, *CreateVerificationCodeRequest) (*CreateVerificationCodeResponse, error)
	ValidateVerificationCode(context.Context, *ValidateVerificationCodeRequest) (*ValidateVerificationCodeResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) CreateVerificationCode(context.Context, *CreateVerificationCodeRequest) (*CreateVerificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVerificationCode not implemented")
}
func (UnimplementedAuthServiceServer) ValidateVerificationCode(context.Context, *ValidateVerificationCodeRequest) (*ValidateVerificationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateVerificationCode not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_CreateVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVerificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CreateVerificationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateVerificationCode(ctx, req.(*CreateVerificationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ValidateVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateVerificationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ValidateVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ValidateVerificationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ValidateVerificationCode(ctx, req.(*ValidateVerificationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "basicService.v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVerificationCode",
			Handler:    _AuthService_CreateVerificationCode_Handler,
		},
		{
			MethodName: "ValidateVerificationCode",
			Handler:    _AuthService_ValidateVerificationCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basicService/v1/verify.proto",
}