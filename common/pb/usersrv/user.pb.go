// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usersrv/user.proto

package usersrv

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

// message LoginReq {
//     string platform = 1; //qq/wx/weibo/username_pwd/phone_pwd/email_pwd/phone_sms/
//     string loginname = 2; //openid/phone/email/username
//     string password = 3;//access_token/pasword/captcha
// }
type LoginByPasswordReq struct {
	Loginname            string   `protobuf:"bytes,1,opt,name=loginname,proto3" json:"loginname,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginByPasswordReq) Reset()         { *m = LoginByPasswordReq{} }
func (m *LoginByPasswordReq) String() string { return proto.CompactTextString(m) }
func (*LoginByPasswordReq) ProtoMessage()    {}
func (*LoginByPasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_e8a4ea4c84caebe2, []int{0}
}
func (m *LoginByPasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginByPasswordReq.Unmarshal(m, b)
}
func (m *LoginByPasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginByPasswordReq.Marshal(b, m, deterministic)
}
func (dst *LoginByPasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginByPasswordReq.Merge(dst, src)
}
func (m *LoginByPasswordReq) XXX_Size() int {
	return xxx_messageInfo_LoginByPasswordReq.Size(m)
}
func (m *LoginByPasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginByPasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginByPasswordReq proto.InternalMessageInfo

func (m *LoginByPasswordReq) GetLoginname() string {
	if m != nil {
		return m.Loginname
	}
	return ""
}

func (m *LoginByPasswordReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginByCaptchaReq struct {
	Loginname            string   `protobuf:"bytes,1,opt,name=loginname,proto3" json:"loginname,omitempty"`
	Captcha              string   `protobuf:"bytes,2,opt,name=captcha,proto3" json:"captcha,omitempty"`
	CaptchaId            int64    `protobuf:"varint,3,opt,name=captcha_id,json=captchaId,proto3" json:"captcha_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginByCaptchaReq) Reset()         { *m = LoginByCaptchaReq{} }
func (m *LoginByCaptchaReq) String() string { return proto.CompactTextString(m) }
func (*LoginByCaptchaReq) ProtoMessage()    {}
func (*LoginByCaptchaReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_e8a4ea4c84caebe2, []int{1}
}
func (m *LoginByCaptchaReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginByCaptchaReq.Unmarshal(m, b)
}
func (m *LoginByCaptchaReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginByCaptchaReq.Marshal(b, m, deterministic)
}
func (dst *LoginByCaptchaReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginByCaptchaReq.Merge(dst, src)
}
func (m *LoginByCaptchaReq) XXX_Size() int {
	return xxx_messageInfo_LoginByCaptchaReq.Size(m)
}
func (m *LoginByCaptchaReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginByCaptchaReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginByCaptchaReq proto.InternalMessageInfo

func (m *LoginByCaptchaReq) GetLoginname() string {
	if m != nil {
		return m.Loginname
	}
	return ""
}

func (m *LoginByCaptchaReq) GetCaptcha() string {
	if m != nil {
		return m.Captcha
	}
	return ""
}

func (m *LoginByCaptchaReq) GetCaptchaId() int64 {
	if m != nil {
		return m.CaptchaId
	}
	return 0
}

type LoginByOAuthReq struct {
	Platform             string            `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	Openid               string            `protobuf:"bytes,2,opt,name=openid,proto3" json:"openid,omitempty"`
	AccessToken          string            `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Fileds               map[string]string `protobuf:"bytes,4,rep,name=fileds,proto3" json:"fileds,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LoginByOAuthReq) Reset()         { *m = LoginByOAuthReq{} }
func (m *LoginByOAuthReq) String() string { return proto.CompactTextString(m) }
func (*LoginByOAuthReq) ProtoMessage()    {}
func (*LoginByOAuthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_e8a4ea4c84caebe2, []int{2}
}
func (m *LoginByOAuthReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginByOAuthReq.Unmarshal(m, b)
}
func (m *LoginByOAuthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginByOAuthReq.Marshal(b, m, deterministic)
}
func (dst *LoginByOAuthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginByOAuthReq.Merge(dst, src)
}
func (m *LoginByOAuthReq) XXX_Size() int {
	return xxx_messageInfo_LoginByOAuthReq.Size(m)
}
func (m *LoginByOAuthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginByOAuthReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginByOAuthReq proto.InternalMessageInfo

func (m *LoginByOAuthReq) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *LoginByOAuthReq) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *LoginByOAuthReq) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *LoginByOAuthReq) GetFileds() map[string]string {
	if m != nil {
		return m.Fileds
	}
	return nil
}

type AuthRsp struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Uid                  int64    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRsp) Reset()         { *m = AuthRsp{} }
func (m *AuthRsp) String() string { return proto.CompactTextString(m) }
func (*AuthRsp) ProtoMessage()    {}
func (*AuthRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_e8a4ea4c84caebe2, []int{3}
}
func (m *AuthRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRsp.Unmarshal(m, b)
}
func (m *AuthRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRsp.Marshal(b, m, deterministic)
}
func (dst *AuthRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRsp.Merge(dst, src)
}
func (m *AuthRsp) XXX_Size() int {
	return xxx_messageInfo_AuthRsp.Size(m)
}
func (m *AuthRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRsp.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRsp proto.InternalMessageInfo

func (m *AuthRsp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthRsp) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type UserReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserReq) Reset()         { *m = UserReq{} }
func (m *UserReq) String() string { return proto.CompactTextString(m) }
func (*UserReq) ProtoMessage()    {}
func (*UserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_e8a4ea4c84caebe2, []int{4}
}
func (m *UserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReq.Unmarshal(m, b)
}
func (m *UserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReq.Marshal(b, m, deterministic)
}
func (dst *UserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReq.Merge(dst, src)
}
func (m *UserReq) XXX_Size() int {
	return xxx_messageInfo_UserReq.Size(m)
}
func (m *UserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserReq proto.InternalMessageInfo

func (m *UserReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// user
type UerRsp struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Firstname            string   `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Gender               string   `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Avatar               string   `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt            string   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UerRsp) Reset()         { *m = UerRsp{} }
func (m *UerRsp) String() string { return proto.CompactTextString(m) }
func (*UerRsp) ProtoMessage()    {}
func (*UerRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_e8a4ea4c84caebe2, []int{5}
}
func (m *UerRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UerRsp.Unmarshal(m, b)
}
func (m *UerRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UerRsp.Marshal(b, m, deterministic)
}
func (dst *UerRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UerRsp.Merge(dst, src)
}
func (m *UerRsp) XXX_Size() int {
	return xxx_messageInfo_UerRsp.Size(m)
}
func (m *UerRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_UerRsp.DiscardUnknown(m)
}

var xxx_messageInfo_UerRsp proto.InternalMessageInfo

func (m *UerRsp) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UerRsp) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UerRsp) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *UerRsp) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *UerRsp) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UerRsp) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UerRsp) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *UerRsp) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type UserRsp struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Firstname            string   `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Gender               string   `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Avatar               string   `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt            string   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRsp) Reset()         { *m = UserRsp{} }
func (m *UserRsp) String() string { return proto.CompactTextString(m) }
func (*UserRsp) ProtoMessage()    {}
func (*UserRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_e8a4ea4c84caebe2, []int{6}
}
func (m *UserRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRsp.Unmarshal(m, b)
}
func (m *UserRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRsp.Marshal(b, m, deterministic)
}
func (dst *UserRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRsp.Merge(dst, src)
}
func (m *UserRsp) XXX_Size() int {
	return xxx_messageInfo_UserRsp.Size(m)
}
func (m *UserRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRsp.DiscardUnknown(m)
}

var xxx_messageInfo_UserRsp proto.InternalMessageInfo

func (m *UserRsp) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UserRsp) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserRsp) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *UserRsp) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *UserRsp) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UserRsp) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UserRsp) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *UserRsp) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

// //register user
// message RegisterReq {
//     LoginReq login = 1;
//     string nickname = 2;
//     string firstname = 3;
//     string lastname = 4;
//     string gender = 5;
//     string avatar = 6;
//     // repeated LoginReq login_list = 1;
// }
type BindReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindReq) Reset()         { *m = BindReq{} }
func (m *BindReq) String() string { return proto.CompactTextString(m) }
func (*BindReq) ProtoMessage()    {}
func (*BindReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_e8a4ea4c84caebe2, []int{7}
}
func (m *BindReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindReq.Unmarshal(m, b)
}
func (m *BindReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindReq.Marshal(b, m, deterministic)
}
func (dst *BindReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindReq.Merge(dst, src)
}
func (m *BindReq) XXX_Size() int {
	return xxx_messageInfo_BindReq.Size(m)
}
func (m *BindReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BindReq.DiscardUnknown(m)
}

var xxx_messageInfo_BindReq proto.InternalMessageInfo

func (m *BindReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UnbindReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnbindReq) Reset()         { *m = UnbindReq{} }
func (m *UnbindReq) String() string { return proto.CompactTextString(m) }
func (*UnbindReq) ProtoMessage()    {}
func (*UnbindReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_e8a4ea4c84caebe2, []int{8}
}
func (m *UnbindReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnbindReq.Unmarshal(m, b)
}
func (m *UnbindReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnbindReq.Marshal(b, m, deterministic)
}
func (dst *UnbindReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnbindReq.Merge(dst, src)
}
func (m *UnbindReq) XXX_Size() int {
	return xxx_messageInfo_UnbindReq.Size(m)
}
func (m *UnbindReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UnbindReq.DiscardUnknown(m)
}

var xxx_messageInfo_UnbindReq proto.InternalMessageInfo

func (m *UnbindReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UnbindReq) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type UpdateReq struct {
	Token                string            `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Fileds               map[string]string `protobuf:"bytes,2,rep,name=fileds,proto3" json:"fileds,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateReq) Reset()         { *m = UpdateReq{} }
func (m *UpdateReq) String() string { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()    {}
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_e8a4ea4c84caebe2, []int{9}
}
func (m *UpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReq.Unmarshal(m, b)
}
func (m *UpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReq.Marshal(b, m, deterministic)
}
func (dst *UpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReq.Merge(dst, src)
}
func (m *UpdateReq) XXX_Size() int {
	return xxx_messageInfo_UpdateReq.Size(m)
}
func (m *UpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReq proto.InternalMessageInfo

func (m *UpdateReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UpdateReq) GetFileds() map[string]string {
	if m != nil {
		return m.Fileds
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginByPasswordReq)(nil), "usersrv.LoginByPasswordReq")
	proto.RegisterType((*LoginByCaptchaReq)(nil), "usersrv.LoginByCaptchaReq")
	proto.RegisterType((*LoginByOAuthReq)(nil), "usersrv.LoginByOAuthReq")
	proto.RegisterMapType((map[string]string)(nil), "usersrv.LoginByOAuthReq.FiledsEntry")
	proto.RegisterType((*AuthRsp)(nil), "usersrv.AuthRsp")
	proto.RegisterType((*UserReq)(nil), "usersrv.UserReq")
	proto.RegisterType((*UerRsp)(nil), "usersrv.UerRsp")
	proto.RegisterType((*UserRsp)(nil), "usersrv.UserRsp")
	proto.RegisterType((*BindReq)(nil), "usersrv.BindReq")
	proto.RegisterType((*UnbindReq)(nil), "usersrv.UnbindReq")
	proto.RegisterType((*UpdateReq)(nil), "usersrv.UpdateReq")
	proto.RegisterMapType((map[string]string)(nil), "usersrv.UpdateReq.FiledsEntry")
}

func init() { proto.RegisterFile("usersrv/user.proto", fileDescriptor_user_e8a4ea4c84caebe2) }

var fileDescriptor_user_e8a4ea4c84caebe2 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xe3, 0xd4, 0xae, 0x27, 0x15, 0x94, 0x11, 0x42, 0x56, 0xf8, 0x68, 0xb0, 0x38, 0xe4,
	0x50, 0x05, 0x28, 0x02, 0x01, 0xea, 0x81, 0x04, 0x81, 0x84, 0x84, 0x00, 0x19, 0x72, 0x8e, 0xb6,
	0xf6, 0xa6, 0x35, 0x49, 0x6d, 0xb3, 0xbb, 0x09, 0xca, 0xcf, 0x40, 0xe2, 0xf7, 0x81, 0xf8, 0x27,
	0x68, 0x77, 0xc7, 0x4e, 0x88, 0x92, 0x72, 0xe1, 0xc4, 0xa9, 0xfb, 0xde, 0xcc, 0x9b, 0x9d, 0xcc,
	0xf3, 0x6c, 0x01, 0x67, 0x92, 0x0b, 0x29, 0xe6, 0xf7, 0xf5, 0xdf, 0x5e, 0x29, 0x0a, 0x55, 0xa0,
	0x4f, 0x5c, 0xf4, 0x0e, 0xf0, 0x6d, 0x71, 0x96, 0xe5, 0x83, 0xc5, 0x07, 0x26, 0xe5, 0xd7, 0x42,
	0xa4, 0x31, 0xff, 0x82, 0xb7, 0x20, 0x98, 0x6a, 0x36, 0x67, 0x17, 0x3c, 0x74, 0x3a, 0x4e, 0x37,
	0x88, 0x97, 0x04, 0xb6, 0x61, 0xaf, 0xa4, 0xe4, 0xb0, 0x61, 0x82, 0x35, 0x8e, 0x3e, 0xc3, 0x35,
	0xaa, 0xf7, 0x92, 0x95, 0x2a, 0x39, 0x67, 0x7f, 0x2f, 0x17, 0x82, 0x9f, 0xd8, 0x5c, 0xaa, 0x56,
	0x41, 0xbc, 0x0d, 0x40, 0xc7, 0x51, 0x96, 0x86, 0x6e, 0xc7, 0xe9, 0xba, 0x71, 0x40, 0xcc, 0x9b,
	0x34, 0xfa, 0xe5, 0xc0, 0x55, 0xba, 0xec, 0x7d, 0x7f, 0xa6, 0xce, 0xf5, 0x55, 0xba, 0xb7, 0x29,
	0x53, 0xe3, 0x42, 0x5c, 0xd0, 0x4d, 0x35, 0xc6, 0x1b, 0xe0, 0x15, 0x25, 0xcf, 0xb3, 0xaa, 0x6b,
	0x42, 0x78, 0x17, 0xf6, 0x59, 0x92, 0x70, 0x29, 0x47, 0xaa, 0x98, 0xf0, 0xdc, 0x5c, 0x14, 0xc4,
	0x2d, 0xcb, 0x7d, 0xd2, 0x14, 0x9e, 0x80, 0x37, 0xce, 0xa6, 0x3c, 0x95, 0x61, 0xb3, 0xe3, 0x76,
	0x5b, 0xc7, 0xf7, 0x7a, 0x34, 0xc0, 0xde, 0x5a, 0x03, 0xbd, 0xd7, 0x26, 0xed, 0x55, 0xae, 0xc4,
	0x22, 0x26, 0x4d, 0xfb, 0x19, 0xb4, 0x56, 0x68, 0x3c, 0x00, 0x77, 0xc2, 0x17, 0xd4, 0x9e, 0x3e,
	0xe2, 0x75, 0xd8, 0x9d, 0xb3, 0xe9, 0x8c, 0x53, 0x63, 0x16, 0x3c, 0x6f, 0x3c, 0x75, 0xa2, 0x87,
	0xe0, 0x9b, 0xca, 0xb2, 0xd4, 0x49, 0xb6, 0x3f, 0x2b, 0xb4, 0x40, 0x17, 0x9b, 0xd1, 0x2f, 0x72,
	0x63, 0x7d, 0x8c, 0x0e, 0xc1, 0x1f, 0x4a, 0x2e, 0xf4, 0x34, 0x36, 0x4a, 0xa2, 0x1f, 0x0e, 0x78,
	0x43, 0x2e, 0x74, 0x4d, 0x52, 0x3b, 0xb5, 0x5a, 0x0f, 0x30, 0xcf, 0x92, 0x89, 0xb1, 0x8a, 0xcc,
	0xad, 0xb0, 0xf6, 0x71, 0x9c, 0x09, 0xa9, 0x4c, 0xd0, 0x4e, 0x69, 0x49, 0x68, 0xe5, 0x94, 0x51,
	0xb0, 0x69, 0x95, 0x15, 0xd6, 0xa3, 0x3f, 0xe3, 0x79, 0xca, 0x45, 0xb8, 0x6b, 0x47, 0x6f, 0x91,
	0xe6, 0xd9, 0x9c, 0x29, 0x26, 0x42, 0xcf, 0xf2, 0x16, 0x69, 0xe7, 0x67, 0x65, 0xca, 0x14, 0x4f,
	0x47, 0x4c, 0x85, 0xbe, 0xbd, 0x8a, 0x98, 0xbe, 0x32, 0x1f, 0x86, 0xe0, 0x55, 0x78, 0xcf, 0x86,
	0x89, 0xe9, 0xab, 0xe8, 0xa7, 0x43, 0x23, 0xf8, 0x6f, 0x7f, 0xe1, 0x21, 0xf8, 0x83, 0x2c, 0x4f,
	0xb7, 0x7b, 0xfc, 0x18, 0x82, 0x61, 0x7e, 0x7a, 0x59, 0x0a, 0x22, 0x34, 0xd5, 0xa2, 0xac, 0x66,
	0x60, 0xce, 0xd1, 0x77, 0x07, 0x82, 0xa1, 0x69, 0x62, 0xbb, 0xee, 0x49, 0xbd, 0x0b, 0x0d, 0xb3,
	0x0b, 0x77, 0xea, 0x5d, 0xa8, 0x95, 0xff, 0x78, 0x0b, 0x8e, 0xbf, 0xb9, 0xd0, 0xd2, 0x86, 0x7e,
	0xe4, 0x62, 0x9e, 0x25, 0x1c, 0x07, 0xf5, 0xe2, 0x57, 0xaf, 0x16, 0xde, 0x5c, 0xdf, 0xc8, 0x95,
	0xf7, 0xac, 0x7d, 0x50, 0x07, 0x69, 0x99, 0xa2, 0x1d, 0x7c, 0x01, 0x57, 0xfe, 0x7c, 0xa9, 0xb0,
	0xbd, 0x5e, 0x62, 0xf9, 0x84, 0x6d, 0xac, 0x70, 0x02, 0xfb, 0xab, 0xdb, 0x8f, 0xe1, 0xb6, 0x47,
	0x61, 0xa3, 0xfa, 0x08, 0x9a, 0xda, 0x42, 0x5c, 0xc6, 0xc8, 0xd1, 0x95, 0x6c, 0xfa, 0x88, 0xa3,
	0x1d, 0x7c, 0x00, 0x9e, 0xf5, 0x13, 0x71, 0x19, 0xad, 0x0c, 0xde, 0xa8, 0x38, 0x82, 0xa6, 0x06,
	0xb8, 0x16, 0xbb, 0xa4, 0xbe, 0x71, 0x6f, 0xb5, 0x7e, 0x65, 0xe7, 0x26, 0xc5, 0xa9, 0x67, 0xfe,
	0x93, 0x3c, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xab, 0x8a, 0xa6, 0x5f, 0x06, 0x00, 0x00,
}
