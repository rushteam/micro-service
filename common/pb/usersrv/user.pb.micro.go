// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: usersrv/user.proto

package usersrv

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	//用户注册
	Signup(ctx context.Context, in *SignupReq, opts ...client.CallOption) (*UserInfo, error)
	//用户登录(账号+密码)
	Signin(ctx context.Context, in *SigninReq, opts ...client.CallOption) (*AuthRsp, error)
	//用户登录(手机号+验证码)
	SigninByPhoneCaptcha(ctx context.Context, in *SigninByPhoneCaptchaReq, opts ...client.CallOption) (*AuthRsp, error)
	//第三方授权登陆 OAuth
	Authorize(ctx context.Context, in *AuthorizeReq, opts ...client.CallOption) (*AuthRsp, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Signup(ctx context.Context, in *SignupReq, opts ...client.CallOption) (*UserInfo, error) {
	req := c.c.NewRequest(c.name, "UserService.Signup", in)
	out := new(UserInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Signin(ctx context.Context, in *SigninReq, opts ...client.CallOption) (*AuthRsp, error) {
	req := c.c.NewRequest(c.name, "UserService.Signin", in)
	out := new(AuthRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SigninByPhoneCaptcha(ctx context.Context, in *SigninByPhoneCaptchaReq, opts ...client.CallOption) (*AuthRsp, error) {
	req := c.c.NewRequest(c.name, "UserService.SigninByPhoneCaptcha", in)
	out := new(AuthRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Authorize(ctx context.Context, in *AuthorizeReq, opts ...client.CallOption) (*AuthRsp, error) {
	req := c.c.NewRequest(c.name, "UserService.Authorize", in)
	out := new(AuthRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	//用户注册
	Signup(context.Context, *SignupReq, *UserInfo) error
	//用户登录(账号+密码)
	Signin(context.Context, *SigninReq, *AuthRsp) error
	//用户登录(手机号+验证码)
	SigninByPhoneCaptcha(context.Context, *SigninByPhoneCaptchaReq, *AuthRsp) error
	//第三方授权登陆 OAuth
	Authorize(context.Context, *AuthorizeReq, *AuthRsp) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Signup(ctx context.Context, in *SignupReq, out *UserInfo) error
		Signin(ctx context.Context, in *SigninReq, out *AuthRsp) error
		SigninByPhoneCaptcha(ctx context.Context, in *SigninByPhoneCaptchaReq, out *AuthRsp) error
		Authorize(ctx context.Context, in *AuthorizeReq, out *AuthRsp) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Signup(ctx context.Context, in *SignupReq, out *UserInfo) error {
	return h.UserServiceHandler.Signup(ctx, in, out)
}

func (h *userServiceHandler) Signin(ctx context.Context, in *SigninReq, out *AuthRsp) error {
	return h.UserServiceHandler.Signin(ctx, in, out)
}

func (h *userServiceHandler) SigninByPhoneCaptcha(ctx context.Context, in *SigninByPhoneCaptchaReq, out *AuthRsp) error {
	return h.UserServiceHandler.SigninByPhoneCaptcha(ctx, in, out)
}

func (h *userServiceHandler) Authorize(ctx context.Context, in *AuthorizeReq, out *AuthRsp) error {
	return h.UserServiceHandler.Authorize(ctx, in, out)
}
