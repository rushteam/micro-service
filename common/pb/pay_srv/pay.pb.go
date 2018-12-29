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
	Extra                string   `protobuf:"bytes,10,opt,name=extra,proto3" json:"extra,omitempty"`
	NotifyUrl            string   `protobuf:"bytes,11,opt,name=notify_url,json=notifyUrl,proto3" json:"notify_url,omitempty"`
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

func (m *CreateReq) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func (m *CreateReq) GetNotifyUrl() string {
	if m != nil {
		return m.NotifyUrl
	}
	return ""
}

type PayField struct {
	AppId      string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	OutTradeNo string `protobuf:"bytes,2,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
	TradeNo    string `protobuf:"bytes,3,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no,omitempty"`
	TotalFee   int64  `protobuf:"varint,4,opt,name=total_fee,json=totalFee,proto3" json:"total_fee,omitempty"`
	//    string trade_type = 5;   //交易方式 (JSAPI\NATIVE)
	CodeUrl              string   `protobuf:"bytes,5,opt,name=code_url,json=codeUrl,proto3" json:"code_url,omitempty"`
	FieldStr             string   `protobuf:"bytes,6,opt,name=field_str,json=fieldStr,proto3" json:"field_str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayField) Reset()         { *m = PayField{} }
func (m *PayField) String() string { return proto.CompactTextString(m) }
func (*PayField) ProtoMessage()    {}
func (*PayField) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddea8d583b895584, []int{1}
}

func (m *PayField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayField.Unmarshal(m, b)
}
func (m *PayField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayField.Marshal(b, m, deterministic)
}
func (m *PayField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayField.Merge(m, src)
}
func (m *PayField) XXX_Size() int {
	return xxx_messageInfo_PayField.Size(m)
}
func (m *PayField) XXX_DiscardUnknown() {
	xxx_messageInfo_PayField.DiscardUnknown(m)
}

var xxx_messageInfo_PayField proto.InternalMessageInfo

func (m *PayField) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *PayField) GetOutTradeNo() string {
	if m != nil {
		return m.OutTradeNo
	}
	return ""
}

func (m *PayField) GetTradeNo() string {
	if m != nil {
		return m.TradeNo
	}
	return ""
}

func (m *PayField) GetTotalFee() int64 {
	if m != nil {
		return m.TotalFee
	}
	return 0
}

func (m *PayField) GetCodeUrl() string {
	if m != nil {
		return m.CodeUrl
	}
	return ""
}

func (m *PayField) GetFieldStr() string {
	if m != nil {
		return m.FieldStr
	}
	return ""
}

type PayRsp struct {
	Provider             string    `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Channel              string    `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	PayField             *PayField `protobuf:"bytes,3,opt,name=pay_field,json=payField,proto3" json:"pay_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PayRsp) Reset()         { *m = PayRsp{} }
func (m *PayRsp) String() string { return proto.CompactTextString(m) }
func (*PayRsp) ProtoMessage()    {}
func (*PayRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddea8d583b895584, []int{2}
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

func (m *PayRsp) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *PayRsp) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *PayRsp) GetPayField() *PayField {
	if m != nil {
		return m.PayField
	}
	return nil
}

type NotifyReq struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Channel              string   `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Raw                  string   `protobuf:"bytes,3,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyReq) Reset()         { *m = NotifyReq{} }
func (m *NotifyReq) String() string { return proto.CompactTextString(m) }
func (*NotifyReq) ProtoMessage()    {}
func (*NotifyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddea8d583b895584, []int{3}
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

func (m *NotifyReq) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *NotifyReq) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *NotifyReq) GetRaw() string {
	if m != nil {
		return m.Raw
	}
	return ""
}

type NotifyRsp struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	OutTradeNo           string   `protobuf:"bytes,2,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyRsp) Reset()         { *m = NotifyRsp{} }
func (m *NotifyRsp) String() string { return proto.CompactTextString(m) }
func (*NotifyRsp) ProtoMessage()    {}
func (*NotifyRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddea8d583b895584, []int{4}
}

func (m *NotifyRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyRsp.Unmarshal(m, b)
}
func (m *NotifyRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyRsp.Marshal(b, m, deterministic)
}
func (m *NotifyRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyRsp.Merge(m, src)
}
func (m *NotifyRsp) XXX_Size() int {
	return xxx_messageInfo_NotifyRsp.Size(m)
}
func (m *NotifyRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyRsp.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyRsp proto.InternalMessageInfo

func (m *NotifyRsp) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *NotifyRsp) GetOutTradeNo() string {
	if m != nil {
		return m.OutTradeNo
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
	return fileDescriptor_ddea8d583b895584, []int{5}
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

//支付通知事件
type NotifyEvent struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyEvent) Reset()         { *m = NotifyEvent{} }
func (m *NotifyEvent) String() string { return proto.CompactTextString(m) }
func (*NotifyEvent) ProtoMessage()    {}
func (*NotifyEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddea8d583b895584, []int{6}
}

func (m *NotifyEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyEvent.Unmarshal(m, b)
}
func (m *NotifyEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyEvent.Marshal(b, m, deterministic)
}
func (m *NotifyEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyEvent.Merge(m, src)
}
func (m *NotifyEvent) XXX_Size() int {
	return xxx_messageInfo_NotifyEvent.Size(m)
}
func (m *NotifyEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyEvent.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyEvent proto.InternalMessageInfo

func (m *NotifyEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NotifyEvent) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *NotifyEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NotifyEvent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateReq)(nil), "pay_srv.CreateReq")
	proto.RegisterType((*PayField)(nil), "pay_srv.PayField")
	proto.RegisterType((*PayRsp)(nil), "pay_srv.PayRsp")
	proto.RegisterType((*NotifyReq)(nil), "pay_srv.NotifyReq")
	proto.RegisterType((*NotifyRsp)(nil), "pay_srv.NotifyRsp")
	proto.RegisterType((*QueryReq)(nil), "pay_srv.QueryReq")
	proto.RegisterType((*NotifyEvent)(nil), "pay_srv.NotifyEvent")
}

func init() { proto.RegisterFile("pay_srv/pay.proto", fileDescriptor_ddea8d583b895584) }

var fileDescriptor_ddea8d583b895584 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0x7f, 0x4d, 0xdb, 0x34, 0x39, 0x9d, 0x7e, 0x30, 0x8b, 0x3f, 0x5e, 0x01, 0xa9, 0xe4,
	0x6a, 0x17, 0xd0, 0x49, 0xe5, 0x09, 0x10, 0xda, 0xa4, 0xde, 0x4c, 0x25, 0xeb, 0xae, 0x23, 0x2f,
	0x39, 0x85, 0x40, 0x1a, 0x1b, 0xdb, 0x29, 0xe4, 0x6d, 0x78, 0x07, 0x5e, 0x82, 0xc7, 0x42, 0xb1,
	0x9d, 0xb4, 0x5b, 0x41, 0xbd, 0xf3, 0xf9, 0x3a, 0xe7, 0xdf, 0xe7, 0x1c, 0x07, 0x4e, 0x05, 0xab,
	0x13, 0x25, 0xb7, 0x17, 0x82, 0xd5, 0x33, 0x21, 0xb9, 0xe6, 0x64, 0xe4, 0xa4, 0xe8, 0xb7, 0x07,
	0xe1, 0x07, 0x89, 0x4c, 0x63, 0x8c, 0xdf, 0xc8, 0x0b, 0x08, 0xd3, 0x22, 0xc7, 0x52, 0x27, 0x79,
	0x46, 0x7b, 0xd3, 0xde, 0x79, 0x18, 0x07, 0x56, 0x58, 0x64, 0xe4, 0x35, 0x9c, 0xb0, 0x34, 0x45,
	0xa5, 0x12, 0xcd, 0xbf, 0x62, 0x49, 0x3d, 0x73, 0x3f, 0xb6, 0xda, 0xaa, 0x91, 0xc8, 0x14, 0x4e,
	0x78, 0xa5, 0x13, 0x2d, 0x59, 0x86, 0x49, 0xc9, 0x69, 0xdf, 0x7c, 0x02, 0xbc, 0xd2, 0xab, 0x46,
	0xba, 0xe6, 0x84, 0xc2, 0x28, 0xfd, 0xcc, 0xca, 0x12, 0x0b, 0x3a, 0x30, 0x97, 0xad, 0xd9, 0xe4,
	0xd6, 0x5c, 0xb3, 0x22, 0x59, 0x23, 0xd2, 0xe1, 0xb4, 0x77, 0xde, 0x8f, 0x03, 0x23, 0x5c, 0x21,
	0x92, 0x57, 0x00, 0x36, 0xa8, 0xae, 0x05, 0x52, 0xdf, 0x78, 0x86, 0x46, 0x59, 0xd5, 0x02, 0x9b,
	0xa8, 0xaa, 0xba, 0xfb, 0x82, 0xa9, 0xa6, 0x23, 0x1b, 0xd5, 0x99, 0xe4, 0x39, 0x8c, 0xd6, 0x92,
	0x6f, 0x92, 0x5c, 0xd0, 0xc0, 0xdc, 0xf8, 0x8d, 0xb9, 0x10, 0xa6, 0x55, 0xd3, 0x77, 0xc2, 0x34,
	0x0d, 0x6d, 0x3a, 0x2b, 0xbc, 0xd7, 0xe4, 0x09, 0x0c, 0xf1, 0x87, 0x96, 0x8c, 0x82, 0xf1, 0xb1,
	0x46, 0x53, 0x44, 0xc9, 0x75, 0xbe, 0xae, 0x93, 0x4a, 0x16, 0x74, 0x6c, 0x8b, 0xb0, 0xca, 0xad,
	0x2c, 0xa2, 0x5f, 0x3d, 0x08, 0x96, 0xac, 0xbe, 0xca, 0xb1, 0xc8, 0xc8, 0x53, 0xf0, 0x99, 0x10,
	0x3b, 0x8c, 0x43, 0x26, 0xc4, 0x22, 0x3b, 0x00, 0xe4, 0x1d, 0x00, 0x3a, 0x83, 0xe0, 0x01, 0xbe,
	0x91, 0x76, 0x57, 0xf7, 0x08, 0x0d, 0x1e, 0x10, 0x3a, 0x83, 0x20, 0xe5, 0x19, 0x9a, 0xd2, 0x86,
	0x8e, 0x2c, 0xcf, 0xf0, 0x56, 0x1a, 0xb2, 0xeb, 0xa6, 0xa8, 0x44, 0x69, 0xe9, 0xd8, 0x05, 0x46,
	0xb8, 0xd1, 0x32, 0x2a, 0xc1, 0x5f, 0xb2, 0x3a, 0x56, 0x82, 0x4c, 0x20, 0x10, 0x92, 0x6f, 0xf3,
	0x0c, 0x65, 0x3b, 0xfb, 0xd6, 0xde, 0x1f, 0x9b, 0x77, 0x7f, 0x6c, 0x33, 0x08, 0x9b, 0x5d, 0x32,
	0xf1, 0x4c, 0xc1, 0xe3, 0xf9, 0xe9, 0xcc, 0x6d, 0xd7, 0xac, 0xc5, 0x11, 0x07, 0xc2, 0x9d, 0xa2,
	0x15, 0x84, 0xd7, 0x06, 0xd9, 0xd1, 0x7d, 0xfb, 0x77, 0xce, 0xc7, 0xd0, 0x97, 0xec, 0xbb, 0xc3,
	0xd3, 0x1c, 0xa3, 0xcb, 0x2e, 0xaa, 0x12, 0xe4, 0x19, 0xf8, 0x12, 0x55, 0x55, 0x68, 0x17, 0xd2,
	0x59, 0xc7, 0xe1, 0x47, 0x6f, 0x20, 0xf8, 0x58, 0xa1, 0x34, 0xb5, 0x1d, 0xdd, 0xe5, 0x28, 0x87,
	0xb1, 0x4d, 0x7a, 0xb9, 0xc5, 0x52, 0x93, 0xff, 0xc1, 0xeb, 0xba, 0xf0, 0xf2, 0x8c, 0xbc, 0x84,
	0x50, 0xe7, 0x1b, 0x54, 0x9a, 0x6d, 0x84, 0xc9, 0xd5, 0x8f, 0x77, 0x02, 0x21, 0x30, 0x28, 0xd9,
	0x06, 0x5d, 0x58, 0x73, 0x6e, 0x3a, 0xde, 0xa0, 0x52, 0xec, 0x13, 0xb6, 0x8f, 0xc3, 0x99, 0xf3,
	0x9f, 0x3d, 0x80, 0x25, 0xab, 0x6f, 0x50, 0x6e, 0xf3, 0x14, 0xc9, 0x05, 0xf8, 0xf6, 0xd1, 0x12,
	0xd2, 0xb1, 0xee, 0x5e, 0xf1, 0xe4, 0xd1, 0x3e, 0xff, 0x58, 0x89, 0xe8, 0x3f, 0x32, 0x07, 0xdf,
	0x96, 0xba, 0xe7, 0xd0, 0x8d, 0x61, 0x72, 0xa0, 0x19, 0x9f, 0xb7, 0x30, 0x34, 0x30, 0xc8, 0x6e,
	0x9e, 0x2d, 0x9c, 0xbf, 0xa4, 0xb8, 0xf3, 0xcd, 0x9f, 0xe5, 0xdd, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa0, 0x8e, 0x3e, 0x60, 0x6e, 0x04, 0x00, 0x00,
}
