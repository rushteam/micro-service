// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pay_srv/pay.proto

package pay_srv

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
	OutPayNo             string   `protobuf:"bytes,3,opt,name=out_pay_no,json=outPayNo,proto3" json:"out_pay_no,omitempty"`
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
	return fileDescriptor_pay_6e474df051a0d67a, []int{0}
}
func (m *CreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReq.Unmarshal(m, b)
}
func (m *CreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReq.Marshal(b, m, deterministic)
}
func (dst *CreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReq.Merge(dst, src)
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

func (m *CreateReq) GetOutPayNo() string {
	if m != nil {
		return m.OutPayNo
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
	return fileDescriptor_pay_6e474df051a0d67a, []int{1}
}
func (m *PayField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayField.Unmarshal(m, b)
}
func (m *PayField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayField.Marshal(b, m, deterministic)
}
func (dst *PayField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayField.Merge(dst, src)
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
	return fileDescriptor_pay_6e474df051a0d67a, []int{2}
}
func (m *PayRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayRsp.Unmarshal(m, b)
}
func (m *PayRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayRsp.Marshal(b, m, deterministic)
}
func (dst *PayRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayRsp.Merge(dst, src)
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
	return fileDescriptor_pay_6e474df051a0d67a, []int{3}
}
func (m *NotifyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyReq.Unmarshal(m, b)
}
func (m *NotifyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyReq.Marshal(b, m, deterministic)
}
func (dst *NotifyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyReq.Merge(dst, src)
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
	OutPayNo             string   `protobuf:"bytes,2,opt,name=out_pay_no,json=outPayNo,proto3" json:"out_pay_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyRsp) Reset()         { *m = NotifyRsp{} }
func (m *NotifyRsp) String() string { return proto.CompactTextString(m) }
func (*NotifyRsp) ProtoMessage()    {}
func (*NotifyRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_pay_6e474df051a0d67a, []int{4}
}
func (m *NotifyRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyRsp.Unmarshal(m, b)
}
func (m *NotifyRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyRsp.Marshal(b, m, deterministic)
}
func (dst *NotifyRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyRsp.Merge(dst, src)
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

func (m *NotifyRsp) GetOutPayNo() string {
	if m != nil {
		return m.OutPayNo
	}
	return ""
}

type QueryReq struct {
	UniTradeNo           string   `protobuf:"bytes,3,opt,name=uni_trade_no,json=uniTradeNo,proto3" json:"uni_trade_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryReq) Reset()         { *m = QueryReq{} }
func (m *QueryReq) String() string { return proto.CompactTextString(m) }
func (*QueryReq) ProtoMessage()    {}
func (*QueryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_pay_6e474df051a0d67a, []int{5}
}
func (m *QueryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryReq.Unmarshal(m, b)
}
func (m *QueryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryReq.Marshal(b, m, deterministic)
}
func (dst *QueryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReq.Merge(dst, src)
}
func (m *QueryReq) XXX_Size() int {
	return xxx_messageInfo_QueryReq.Size(m)
}
func (m *QueryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReq.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReq proto.InternalMessageInfo

func (m *QueryReq) GetUniTradeNo() string {
	if m != nil {
		return m.UniTradeNo
	}
	return ""
}

// 支付通知事件
type NotifyEvent struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	PayNo                string   `protobuf:"bytes,4,opt,name=pay_no,json=payNo,proto3" json:"pay_no,omitempty"`
	Url                  string   `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Body                 string   `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyEvent) Reset()         { *m = NotifyEvent{} }
func (m *NotifyEvent) String() string { return proto.CompactTextString(m) }
func (*NotifyEvent) ProtoMessage()    {}
func (*NotifyEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_pay_6e474df051a0d67a, []int{6}
}
func (m *NotifyEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyEvent.Unmarshal(m, b)
}
func (m *NotifyEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyEvent.Marshal(b, m, deterministic)
}
func (dst *NotifyEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyEvent.Merge(dst, src)
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

func (m *NotifyEvent) GetPayNo() string {
	if m != nil {
		return m.PayNo
	}
	return ""
}

func (m *NotifyEvent) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *NotifyEvent) GetBody() string {
	if m != nil {
		return m.Body
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

func init() { proto.RegisterFile("pay_srv/pay.proto", fileDescriptor_pay_6e474df051a0d67a) }

var fileDescriptor_pay_6e474df051a0d67a = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0xfe, 0xed, 0xd4, 0x8e, 0x3d, 0xa9, 0x7e, 0xda, 0x15, 0x87, 0x6d, 0x29, 0x52, 0xf0, 0x55,
	0x2f, 0x20, 0x95, 0xca, 0x13, 0x54, 0x88, 0x4a, 0xbd, 0xa9, 0x82, 0x9b, 0x5e, 0x5b, 0x5b, 0x7b,
	0x22, 0x0c, 0xce, 0xee, 0xb2, 0x5e, 0x07, 0xfc, 0x0e, 0x3c, 0x04, 0xef, 0xc0, 0x2b, 0xf0, 0x60,
	0x68, 0x0f, 0x71, 0xda, 0x00, 0xe2, 0x6e, 0xe6, 0x1b, 0xcd, 0xe9, 0xfb, 0x66, 0x17, 0x0e, 0x25,
	0xeb, 0x8b, 0x56, 0xad, 0xcf, 0x24, 0xeb, 0x67, 0x52, 0x09, 0x2d, 0xc8, 0xd8, 0x43, 0xd9, 0xcf,
	0x10, 0xd2, 0xb7, 0x0a, 0x99, 0xc6, 0x1c, 0x3f, 0x93, 0xe7, 0x90, 0x96, 0x4d, 0x8d, 0x5c, 0x17,
	0x75, 0x45, 0x83, 0x69, 0x70, 0x9a, 0xe6, 0x89, 0x03, 0xae, 0x2a, 0xf2, 0x12, 0xf6, 0x59, 0x59,
	0x62, 0xdb, 0x16, 0x5a, 0x7c, 0x42, 0x4e, 0x43, 0x1b, 0x9f, 0x38, 0x6c, 0x61, 0x20, 0x72, 0x02,
	0x20, 0x3a, 0x5d, 0x98, 0xe2, 0x5c, 0xd0, 0x91, 0x2b, 0x20, 0x3a, 0x3d, 0x67, 0xfd, 0xb5, 0x20,
	0x14, 0xc6, 0xe5, 0x07, 0xc6, 0x39, 0x36, 0x74, 0xcf, 0x86, 0x36, 0xae, 0xe9, 0xab, 0x85, 0x66,
	0x4d, 0xb1, 0x44, 0xa4, 0xd1, 0x34, 0x38, 0x1d, 0xe5, 0x89, 0x05, 0x2e, 0x11, 0xc9, 0x0b, 0x00,
	0xad, 0x58, 0x85, 0x85, 0xee, 0x25, 0xd2, 0xd8, 0x66, 0xa6, 0x16, 0x59, 0xf4, 0x12, 0x4d, 0xd5,
	0xb6, 0xbb, 0xfb, 0x88, 0xa5, 0xa6, 0x63, 0x57, 0xd5, 0xbb, 0xe4, 0x19, 0x8c, 0x97, 0x4a, 0xac,
	0x8a, 0x5a, 0xd2, 0xc4, 0x46, 0x62, 0xe3, 0x5e, 0x49, 0xbb, 0xa6, 0xdd, 0xb9, 0x60, 0x9a, 0xa6,
	0xae, 0x9d, 0x03, 0x2e, 0x34, 0x79, 0x0c, 0x11, 0x7e, 0xd5, 0x8a, 0x51, 0xb0, 0x39, 0xce, 0x31,
	0x43, 0x70, 0xa1, 0xeb, 0x65, 0x5f, 0x74, 0xaa, 0xa1, 0x13, 0x37, 0x84, 0x43, 0x6e, 0x55, 0x93,
	0xfd, 0x08, 0x20, 0x99, 0xb3, 0xfe, 0xb2, 0xc6, 0xa6, 0x22, 0x4f, 0x20, 0x66, 0x52, 0x6e, 0x29,
	0x8c, 0x98, 0x94, 0x57, 0x15, 0x99, 0xc2, 0xbe, 0x21, 0xc7, 0xed, 0xc2, 0x85, 0xe7, 0xcf, 0x10,
	0xb6, 0x30, 0xd0, 0xb5, 0x20, 0x47, 0x90, 0x0c, 0x51, 0x47, 0xde, 0x58, 0xfb, 0xd0, 0x03, 0x86,
	0xf6, 0x76, 0x18, 0x3a, 0x82, 0xa4, 0x14, 0x15, 0xda, 0xd1, 0x22, 0xcf, 0xac, 0xa8, 0xf0, 0x56,
	0x59, 0x66, 0x97, 0x66, 0xa8, 0xa2, 0xd5, 0xca, 0x73, 0x97, 0x58, 0xe0, 0x46, 0xab, 0x8c, 0x43,
	0x3c, 0x67, 0x7d, 0xde, 0x4a, 0x72, 0x0c, 0x89, 0x54, 0x62, 0x5d, 0x57, 0xa8, 0x36, 0xba, 0x6f,
	0xfc, 0xfb, 0xb2, 0x85, 0x0f, 0x65, 0x9b, 0x41, 0x6a, 0xa4, 0xb6, 0xf5, 0xec, 0xc0, 0x93, 0xf3,
	0xc3, 0x99, 0xbf, 0xac, 0xd9, 0x86, 0x8e, 0x3c, 0x91, 0xde, 0xca, 0x16, 0x90, 0x5e, 0x5b, 0xca,
	0xfe, 0x79, 0x6b, 0x7f, 0xef, 0x79, 0x00, 0x23, 0xc5, 0xbe, 0x78, 0x7a, 0x8c, 0x99, 0x5d, 0x0c,
	0x55, 0x5b, 0x49, 0x9e, 0x42, 0xac, 0xb0, 0xed, 0x1a, 0xed, 0x4b, 0x7a, 0x6f, 0xe7, 0x32, 0xc3,
	0x87, 0x97, 0x99, 0xbd, 0x82, 0xe4, 0x7d, 0x87, 0xca, 0xce, 0x35, 0x85, 0xfd, 0x8e, 0xd7, 0xc5,
	0x8e, 0x10, 0xd0, 0xf1, 0xda, 0xcb, 0x94, 0x7d, 0x0b, 0x60, 0xe2, 0x3a, 0xbe, 0x5b, 0x23, 0xd7,
	0xe4, 0x7f, 0x08, 0x87, 0x15, 0xc2, 0xba, 0x22, 0x27, 0x90, 0xea, 0x7a, 0x85, 0xad, 0x66, 0x2b,
	0x69, 0x5b, 0x8d, 0xf2, 0x2d, 0x40, 0x08, 0xec, 0x71, 0xb6, 0x42, 0x5f, 0xd7, 0xda, 0xe6, 0x62,
	0xfc, 0x64, 0xee, 0x61, 0x44, 0xd2, 0x3e, 0x98, 0x03, 0x18, 0x6d, 0x25, 0x35, 0xa6, 0x49, 0xbe,
	0x13, 0x55, 0xef, 0x95, 0xb4, 0xf6, 0xf9, 0xf7, 0x00, 0x60, 0xce, 0xfa, 0x1b, 0x54, 0xeb, 0xba,
	0x44, 0x72, 0x06, 0xb1, 0x7b, 0xd0, 0x84, 0x0c, 0x5a, 0x0c, 0x2f, 0xfc, 0xf8, 0xd1, 0x7d, 0x7d,
	0xf2, 0x56, 0x66, 0xff, 0x91, 0x73, 0x88, 0xdd, 0x36, 0xf7, 0x12, 0x06, 0x99, 0x8e, 0x7f, 0xc3,
	0x6c, 0xce, 0x6b, 0x88, 0x2c, 0x61, 0x64, 0xab, 0xf7, 0x86, 0xc0, 0x3f, 0xb4, 0xb8, 0x8b, 0xed,
	0xaf, 0xf3, 0xe6, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x91, 0x02, 0xdd, 0x8a, 0x04, 0x00,
	0x00,
}
