// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pay_srv/pay.proto

/*
Package pay_srv is a generated protocol buffer package.

It is generated from these files:
	pay_srv/pay.proto

It has these top-level messages:
	CreateReq
	PayField
	PayRsp
	NotifyReq
	NotifyRsp
	QueryReq
*/
package pay_srv

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PayService service

type PayService interface {
	// 创建支付单
	Create(ctx context.Context, in *CreateReq, opts ...client.CallOption) (*PayRsp, error)
	// 完成支付
	Notify(ctx context.Context, in *NotifyReq, opts ...client.CallOption) (*NotifyRsp, error)
	// 支付详情
	Query(ctx context.Context, in *QueryReq, opts ...client.CallOption) (*PayRsp, error)
}

type payService struct {
	c    client.Client
	name string
}

func NewPayService(name string, c client.Client) PayService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "pay_srv"
	}
	return &payService{
		c:    c,
		name: name,
	}
}

func (c *payService) Create(ctx context.Context, in *CreateReq, opts ...client.CallOption) (*PayRsp, error) {
	req := c.c.NewRequest(c.name, "PayService.Create", in)
	out := new(PayRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payService) Notify(ctx context.Context, in *NotifyReq, opts ...client.CallOption) (*NotifyRsp, error) {
	req := c.c.NewRequest(c.name, "PayService.Notify", in)
	out := new(NotifyRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payService) Query(ctx context.Context, in *QueryReq, opts ...client.CallOption) (*PayRsp, error) {
	req := c.c.NewRequest(c.name, "PayService.Query", in)
	out := new(PayRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PayService service

type PayServiceHandler interface {
	// 创建支付单
	Create(context.Context, *CreateReq, *PayRsp) error
	// 完成支付
	Notify(context.Context, *NotifyReq, *NotifyRsp) error
	// 支付详情
	Query(context.Context, *QueryReq, *PayRsp) error
}

func RegisterPayServiceHandler(s server.Server, hdlr PayServiceHandler, opts ...server.HandlerOption) error {
	type payService interface {
		Create(ctx context.Context, in *CreateReq, out *PayRsp) error
		Notify(ctx context.Context, in *NotifyReq, out *NotifyRsp) error
		Query(ctx context.Context, in *QueryReq, out *PayRsp) error
	}
	type PayService struct {
		payService
	}
	h := &payServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PayService{h}, opts...))
}

type payServiceHandler struct {
	PayServiceHandler
}

func (h *payServiceHandler) Create(ctx context.Context, in *CreateReq, out *PayRsp) error {
	return h.PayServiceHandler.Create(ctx, in, out)
}

func (h *payServiceHandler) Notify(ctx context.Context, in *NotifyReq, out *NotifyRsp) error {
	return h.PayServiceHandler.Notify(ctx, in, out)
}

func (h *payServiceHandler) Query(ctx context.Context, in *QueryReq, out *PayRsp) error {
	return h.PayServiceHandler.Query(ctx, in, out)
}
