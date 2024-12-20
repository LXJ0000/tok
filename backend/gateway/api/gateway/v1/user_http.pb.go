// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v3.21.12
// source: gateway/v1/user.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserServiceBindUserVoucher = "/gateway.v1.UserService/BindUserVoucher"
const OperationUserServiceGetUserInfo = "/gateway.v1.UserService/GetUserInfo"
const OperationUserServiceGetVerificationCode = "/gateway.v1.UserService/GetVerificationCode"
const OperationUserServiceLogin = "/gateway.v1.UserService/Login"
const OperationUserServiceRegister = "/gateway.v1.UserService/Register"
const OperationUserServiceUnbindUserVoucher = "/gateway.v1.UserService/UnbindUserVoucher"
const OperationUserServiceUpdateUserInfo = "/gateway.v1.UserService/UpdateUserInfo"

type UserServiceHTTPServer interface {
	BindUserVoucher(context.Context, *BindUserVoucherRequest) (*BindUserVoucherResponse, error)
	// GetUserInfo 获取用户信息
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	// GetVerificationCode 获取验证码
	GetVerificationCode(context.Context, *GetVerificationCodeRequest) (*GetVerificationCodeResponse, error)
	// Login 登录
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Register 注册
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	UnbindUserVoucher(context.Context, *UnbindUserVoucherRequest) (*UnbindUserVoucherResponse, error)
	// UpdateUserInfo 更新用户信息
	UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error)
}

func RegisterUserServiceHTTPServer(s *http.Server, srv UserServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/user/code", _UserService_GetVerificationCode0_HTTP_Handler(srv))
	r.POST("/user/register", _UserService_Register0_HTTP_Handler(srv))
	r.POST("/user/login", _UserService_Login0_HTTP_Handler(srv))
	r.GET("/user/info", _UserService_GetUserInfo0_HTTP_Handler(srv))
	r.PUT("/user/info", _UserService_UpdateUserInfo0_HTTP_Handler(srv))
	r.POST("/user/voucher", _UserService_BindUserVoucher0_HTTP_Handler(srv))
	r.DELETE("/user/voucher", _UserService_UnbindUserVoucher0_HTTP_Handler(srv))
}

func _UserService_GetVerificationCode0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetVerificationCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetVerificationCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetVerificationCode(ctx, req.(*GetVerificationCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetVerificationCodeResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_Register0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_Login0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetUserInfo0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfo(ctx, req.(*GetUserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_UpdateUserInfo0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserInfoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceUpdateUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserInfo(ctx, req.(*UpdateUserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_BindUserVoucher0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BindUserVoucherRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceBindUserVoucher)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BindUserVoucher(ctx, req.(*BindUserVoucherRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BindUserVoucherResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_UnbindUserVoucher0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnbindUserVoucherRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceUnbindUserVoucher)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UnbindUserVoucher(ctx, req.(*UnbindUserVoucherRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UnbindUserVoucherResponse)
		return ctx.Result(200, reply)
	}
}

type UserServiceHTTPClient interface {
	BindUserVoucher(ctx context.Context, req *BindUserVoucherRequest, opts ...http.CallOption) (rsp *BindUserVoucherResponse, err error)
	GetUserInfo(ctx context.Context, req *GetUserInfoRequest, opts ...http.CallOption) (rsp *GetUserInfoResponse, err error)
	GetVerificationCode(ctx context.Context, req *GetVerificationCodeRequest, opts ...http.CallOption) (rsp *GetVerificationCodeResponse, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginResponse, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterResponse, err error)
	UnbindUserVoucher(ctx context.Context, req *UnbindUserVoucherRequest, opts ...http.CallOption) (rsp *UnbindUserVoucherResponse, err error)
	UpdateUserInfo(ctx context.Context, req *UpdateUserInfoRequest, opts ...http.CallOption) (rsp *UpdateUserInfoResponse, err error)
}

type UserServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewUserServiceHTTPClient(client *http.Client) UserServiceHTTPClient {
	return &UserServiceHTTPClientImpl{client}
}

func (c *UserServiceHTTPClientImpl) BindUserVoucher(ctx context.Context, in *BindUserVoucherRequest, opts ...http.CallOption) (*BindUserVoucherResponse, error) {
	var out BindUserVoucherResponse
	pattern := "/user/voucher"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceBindUserVoucher))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...http.CallOption) (*GetUserInfoResponse, error) {
	var out GetUserInfoResponse
	pattern := "/user/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) GetVerificationCode(ctx context.Context, in *GetVerificationCodeRequest, opts ...http.CallOption) (*GetVerificationCodeResponse, error) {
	var out GetVerificationCodeResponse
	pattern := "/user/code"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceGetVerificationCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginResponse, error) {
	var out LoginResponse
	pattern := "/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterResponse, error) {
	var out RegisterResponse
	pattern := "/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) UnbindUserVoucher(ctx context.Context, in *UnbindUserVoucherRequest, opts ...http.CallOption) (*UnbindUserVoucherResponse, error) {
	var out UnbindUserVoucherResponse
	pattern := "/user/voucher"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceUnbindUserVoucher))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...http.CallOption) (*UpdateUserInfoResponse, error) {
	var out UpdateUserInfoResponse
	pattern := "/user/info"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceUpdateUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
