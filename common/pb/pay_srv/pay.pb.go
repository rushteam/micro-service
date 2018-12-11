// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pay_srv/pay.proto

package pay_srv

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type CreateReq struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	AccessToken          string   `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	OutTradeNo           string   `protobuf:"bytes,3,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
	Channel              string   `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"`
	TotalFee             int64    `protobuf:"varint,5,opt,name=total_fee,json=totalFee,proto3" json:"total_fee,omitempty"`
	TradeType            string   `protobuf:"bytes,6,opt,name=trade_type,json=tradeType,proto3" json:"trade_type,omitempty"`
	Subject              string   `protobuf:"bytes,7,opt,name=subject,proto3" json:"subject,omitempty"`
	FromIp               string   `protobuf:"bytes,8,opt,name=from_ip,json=fromIp,proto3" json:"from_ip,omitempty"`
	CreateAt             int64    `protobuf:"varint,9,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	OpenId               string   `protobuf:"bytes,10,opt,name=open_id,json=openId,proto3" json:"open_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateReq) Reset()         { *m = CreateReq{} }
func (m *CreateReq) String() string { return proto.CompactTextString(m) }
func (*CreateReq) ProtoMessage()    {}
func (*CreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddea8d583b895584, []int{0}
}

func (m *CreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReq.Unmarshal(m, b)
}
func (m *CreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReq.Marshal(b, m, deterministic)
}
func (m *CreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReq.Merge(m, src)
}
func (m *CreateReq) XXX_Size() int {
	return xxx_messageInfo_CreateReq.Size(m)
}
func (m *CreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReq proto.InternalMessageInfo

func (m *CreateReq) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *CreateReq) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *CreateReq) GetOutTradeNo() string {
	if m != nil {
		return m.OutTradeNo
	}
	return ""
}

func (m *CreateReq) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *CreateReq) GetTotalFee() int64 {
	if m != nil {
		return m.TotalFee
	}
	return 0
}

func (m *CreateReq) GetTradeType() string {
	if m != nil {
		return m.TradeType
	}
	return ""
}

func (m *CreateReq) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *CreateReq) GetFromIp() string {
	if m != nil {
		return m.FromIp
	}
	return ""
}

func (m *CreateReq) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (m *CreateReq) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

type PayRsp struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	OutTradeNo           string   `protobuf:"bytes,3,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
	Channel              string   `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"`
	TotalFee             int64    `protobuf:"varint,5,opt,name=total_fee,json=totalFee,proto3" json:"total_fee,omitempty"`
	TradeType            int64    `protobuf:"varint,6,opt,name=trade_type,json=tradeType,proto3" json:"trade_type,omitempty"`
	Subject              string   `protobuf:"bytes,7,opt,name=subject,proto3" json:"subject,omitempty"`
	FromIp               string   `protobuf:"bytes,8,opt,name=from_ip,json=fromIp,proto3" json:"from_ip,omitempty"`
	CreateAt             int64    `protobuf:"varint,9,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	PayState             int64    `protobuf:"varint,10,opt,name=pay_state,json=payState,proto3" json:"pay_state,omitempty"`
	PayAt                int64    `protobuf:"varint,11,opt,name=pay_at,json=payAt,proto3" json:"pay_at,omitempty"`
	TradeNo              string   `protobuf:"bytes,12,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no,omitempty"`
	PayField             string   `protobuf:"bytes,13,opt,name=pay_field,json=payField,proto3" json:"pay_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayRsp) Reset()         { *m = PayRsp{} }
func (m *PayRsp) String() string { return proto.CompactTextString(m) }
func (*PayRsp) ProtoMessage()    {}
func (*PayRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddea8d583b895584, []int{1}
}

func (m *PayRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayRsp.Unmarshal(m, b)
}
func (m *PayRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayRsp.Marshal(b, m, deterministic)
}
func (m *PayRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayRsp.Merge(m, src)
}
func (m *PayRsp) XXX_Size() int {
	return xxx_messageInfo_PayRsp.Size(m)
}
func (m *PayRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_PayRsp.DiscardUnknown(m)
}

var xxx_messageInfo_PayRsp proto.InternalMessageInfo

func (m *PayRsp) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *PayRsp) GetOutTradeNo() string {
	if m != nil {
		return m.OutTradeNo
	}
	return ""
}

func (m *PayRsp) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *PayRsp) GetTotalFee() int64 {
	if m != nil {
		return m.TotalFee
	}
	return 0
}

func (m *PayRsp) GetTradeType() int64 {
	if m != nil {
		return m.TradeType
	}
	return 0
}

func (m *PayRsp) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *PayRsp) GetFromIp() string {
	if m != nil {
		return m.FromIp
	}
	return ""
}

func (m *PayRsp) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (m *PayRsp) GetPayState() int64 {
	if m != nil {
		return m.PayState
	}
	return 0
}

func (m *PayRsp) GetPayAt() int64 {
	if m != nil {
		return m.PayAt
	}
	return 0
}

func (m *PayRsp) GetTradeNo() string {
	if m != nil {
		return m.TradeNo
	}
	return ""
}

func (m *PayRsp) GetPayField() string {
	if m != nil {
		return m.PayField
	}
	return ""
}

type NotifyReq struct {
	PayType              int64    `protobuf:"varint,1,opt,name=pay_type,json=payType,proto3" json:"pay_type,omitempty"`
	TotalFee             int64    `protobuf:"varint,2,opt,name=total_fee,json=totalFee,proto3" json:"total_fee,omitempty"`
	OutTradeNo           string   `protobuf:"bytes,3,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
	Raw                  string   `protobuf:"bytes,4,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyReq) Reset()         { *m = NotifyReq{} }
func (m *NotifyReq) String() string { return proto.CompactTextString(m) }
func (*NotifyReq) ProtoMessage()    {}
func (*NotifyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddea8d583b895584, []int{2}
}

func (m *NotifyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyReq.Unmarshal(m, b)
}
func (m *NotifyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyReq.Marshal(b, m, deterministic)
}
func (m *NotifyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyReq.Merge(m, src)
}
func (m *NotifyReq) XXX_Size() int {
	return xxx_messageInfo_NotifyReq.Size(m)
}
func (m *NotifyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyReq.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyReq proto.InternalMessageInfo

func (m *NotifyReq) GetPayType() int64 {
	if m != nil {
		return m.PayType
	}
	return 0
}

func (m *NotifyReq) GetTotalFee() int64 {
	if m != nil {
		return m.TotalFee
	}
	return 0
}

func (m *NotifyReq) GetOutTradeNo() string {
	if m != nil {
		return m.OutTradeNo
	}
	return ""
}

func (m *NotifyReq) GetRaw() string {
	if m != nil {
		return m.Raw
	}
	return ""
}

type QueryReq struct {
	OutTradeNo           string   `protobuf:"bytes,3,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryReq) Reset()         { *m = QueryReq{} }
func (m *QueryReq) String() string { return proto.CompactTextString(m) }
func (*QueryReq) ProtoMessage()    {}
func (*QueryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddea8d583b895584, []int{3}
}

func (m *QueryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryReq.Unmarshal(m, b)
}
func (m *QueryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryReq.Marshal(b, m, deterministic)
}
func (m *QueryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReq.Merge(m, src)
}
func (m *QueryReq) XXX_Size() int {
	return xxx_messageInfo_QueryReq.Size(m)
}
func (m *QueryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReq.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReq proto.InternalMessageInfo

func (m *QueryReq) GetOutTradeNo() string {
	if m != nil {
		return m.OutTradeNo
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateReq)(nil), "pay_srv.CreateReq")
	proto.RegisterType((*PayRsp)(nil), "pay_srv.PayRsp")
	proto.RegisterType((*NotifyReq)(nil), "pay_srv.NotifyReq")
	proto.RegisterType((*QueryReq)(nil), "pay_srv.QueryReq")
}

func init() { proto.RegisterFile("pay_srv/pay.proto", fileDescriptor_ddea8d583b895584) }

var fileDescriptor_ddea8d583b895584 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x31, 0x71, 0xec, 0x69, 0x10, 0x74, 0x25, 0xc4, 0x96, 0x0a, 0x29, 0xf8, 0xd4, 0x03,
	0xa4, 0x12, 0x7c, 0x41, 0x85, 0x54, 0x29, 0x97, 0xaa, 0xb8, 0xb9, 0x5b, 0x53, 0x7b, 0x22, 0x0c,
	0xc1, 0xbb, 0xd8, 0x93, 0x56, 0xfb, 0x31, 0xfc, 0x01, 0x3f, 0xc4, 0xdf, 0xa0, 0x9d, 0x75, 0x2d,
	0x40, 0x15, 0x39, 0xd1, 0x9b, 0xe7, 0x3d, 0xbf, 0xd9, 0x79, 0x4f, 0x33, 0x70, 0x68, 0xd1, 0x95,
	0x7d, 0x77, 0x73, 0x6a, 0xd1, 0x2d, 0x6d, 0x67, 0xd8, 0xa8, 0xd9, 0x00, 0xe5, 0x3f, 0x26, 0x90,
	0x7d, 0xe8, 0x08, 0x99, 0x0a, 0xfa, 0xa6, 0x8e, 0x21, 0xab, 0xb6, 0x0d, 0xb5, 0x5c, 0x36, 0xb5,
	0x8e, 0x16, 0xd1, 0x49, 0x56, 0xa4, 0x01, 0x58, 0xd5, 0xea, 0x35, 0xcc, 0xb1, 0xaa, 0xa8, 0xef,
	0x4b, 0x36, 0x5f, 0xa8, 0xd5, 0x13, 0xe1, 0x0f, 0x02, 0xb6, 0xf6, 0x90, 0x5a, 0xc0, 0xdc, 0xec,
	0xb8, 0xe4, 0x0e, 0x6b, 0x2a, 0x5b, 0xa3, 0x63, 0xf9, 0x05, 0xcc, 0x8e, 0xd7, 0x1e, 0xba, 0x30,
	0x4a, 0xc3, 0xac, 0xfa, 0x84, 0x6d, 0x4b, 0x5b, 0xfd, 0x58, 0xc8, 0xbb, 0xd2, 0xbf, 0xcd, 0x86,
	0x71, 0x5b, 0x6e, 0x88, 0xf4, 0x74, 0x11, 0x9d, 0xc4, 0x45, 0x2a, 0xc0, 0x39, 0x91, 0x7a, 0x05,
	0x10, 0x9a, 0xb2, 0xb3, 0xa4, 0x13, 0x51, 0x66, 0x82, 0xac, 0x9d, 0x25, 0xdf, 0xb5, 0xdf, 0x5d,
	0x7f, 0xa6, 0x8a, 0xf5, 0x2c, 0x74, 0x1d, 0x4a, 0xf5, 0x02, 0x66, 0x9b, 0xce, 0x7c, 0x2d, 0x1b,
	0xab, 0x53, 0x61, 0x12, 0x5f, 0xae, 0xac, 0x58, 0x15, 0xdf, 0x25, 0xb2, 0xce, 0xc2, 0x73, 0x01,
	0x38, 0x13, 0x95, 0xb1, 0xd4, 0xfa, 0x14, 0x20, 0xa8, 0x7c, 0xb9, 0xaa, 0xf3, 0x9f, 0x13, 0x48,
	0x2e, 0xd1, 0x15, 0xbd, 0xfd, 0x77, 0x56, 0x0f, 0x18, 0x44, 0xfc, 0x3f, 0x83, 0x38, 0x86, 0x4c,
	0x36, 0x85, 0x91, 0x49, 0xa2, 0x88, 0x8b, 0xd4, 0xa2, 0xbb, 0xf2, 0xb5, 0x7a, 0x0e, 0x89, 0x27,
	0x91, 0xf5, 0x81, 0x30, 0x53, 0x8b, 0xee, 0x8c, 0xd5, 0x11, 0xa4, 0xa3, 0xef, 0x79, 0x18, 0x82,
	0x07, 0xd3, 0x43, 0xbb, 0x4d, 0x43, 0xdb, 0x5a, 0x3f, 0x09, 0x99, 0x59, 0x74, 0xe7, 0xbe, 0xce,
	0x6f, 0x21, 0xbb, 0x30, 0xdc, 0x6c, 0x9c, 0xdf, 0xc4, 0x23, 0xf0, 0x44, 0x70, 0x19, 0x49, 0x77,
	0xbf, 0xb2, 0xe2, 0xf1, 0x8f, 0x7c, 0x26, 0x7f, 0xe5, 0xb3, 0x3f, 0xf8, 0x67, 0x10, 0x77, 0x78,
	0x3b, 0x84, 0xee, 0x3f, 0xf3, 0x37, 0x90, 0x7e, 0xdc, 0x51, 0x27, 0xef, 0xee, 0xd5, 0xbf, 0xfb,
	0x1e, 0x01, 0x5c, 0xa2, 0xbb, 0xa2, 0xee, 0xa6, 0xa9, 0x48, 0x9d, 0x42, 0x12, 0xee, 0x47, 0xa9,
	0xe5, 0x70, 0x54, 0xcb, 0xf1, 0xa0, 0x5e, 0x3e, 0x1d, 0xb1, 0xb0, 0x35, 0xf9, 0x23, 0x2f, 0x08,
	0x36, 0x7f, 0x13, 0x8c, 0xbe, 0xef, 0x13, 0xbc, 0x85, 0xa9, 0x8c, 0xa7, 0x0e, 0x47, 0xee, 0x6e,
	0xdc, 0x7b, 0x7e, 0xbf, 0x4e, 0xe4, 0xc2, 0xdf, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x37, 0xbb,
	0x2c, 0x36, 0xf6, 0x03, 0x00, 0x00,
}
