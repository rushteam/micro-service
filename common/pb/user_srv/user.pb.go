// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_srv/user.proto

package user_srv

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

type LoginReq struct {
	Platform             string   `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	Openid               string   `protobuf:"bytes,2,opt,name=openid,proto3" json:"openid,omitempty"`
	AccessToken          string   `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3d914a4924265a, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *LoginReq) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *LoginReq) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type LoginRsp struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRsp) Reset()         { *m = LoginRsp{} }
func (m *LoginRsp) String() string { return proto.CompactTextString(m) }
func (*LoginRsp) ProtoMessage()    {}
func (*LoginRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3d914a4924265a, []int{1}
}

func (m *LoginRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRsp.Unmarshal(m, b)
}
func (m *LoginRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRsp.Marshal(b, m, deterministic)
}
func (m *LoginRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRsp.Merge(m, src)
}
func (m *LoginRsp) XXX_Size() int {
	return xxx_messageInfo_LoginRsp.Size(m)
}
func (m *LoginRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRsp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRsp proto.InternalMessageInfo

func (m *LoginRsp) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type UserReq struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserReq) Reset()         { *m = UserReq{} }
func (m *UserReq) String() string { return proto.CompactTextString(m) }
func (*UserReq) ProtoMessage()    {}
func (*UserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3d914a4924265a, []int{2}
}

func (m *UserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReq.Unmarshal(m, b)
}
func (m *UserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReq.Marshal(b, m, deterministic)
}
func (m *UserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReq.Merge(m, src)
}
func (m *UserReq) XXX_Size() int {
	return xxx_messageInfo_UserReq.Size(m)
}
func (m *UserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserReq proto.InternalMessageInfo

func (m *UserReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type UserRsp struct {
	Userinfo             *Userinfo `protobuf:"bytes,1,opt,name=userinfo,proto3" json:"userinfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserRsp) Reset()         { *m = UserRsp{} }
func (m *UserRsp) String() string { return proto.CompactTextString(m) }
func (*UserRsp) ProtoMessage()    {}
func (*UserRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3d914a4924265a, []int{3}
}

func (m *UserRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRsp.Unmarshal(m, b)
}
func (m *UserRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRsp.Marshal(b, m, deterministic)
}
func (m *UserRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRsp.Merge(m, src)
}
func (m *UserRsp) XXX_Size() int {
	return xxx_messageInfo_UserRsp.Size(m)
}
func (m *UserRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRsp.DiscardUnknown(m)
}

var xxx_messageInfo_UserRsp proto.InternalMessageInfo

func (m *UserRsp) GetUserinfo() *Userinfo {
	if m != nil {
		return m.Userinfo
	}
	return nil
}

type Userinfo struct {
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

func (m *Userinfo) Reset()         { *m = Userinfo{} }
func (m *Userinfo) String() string { return proto.CompactTextString(m) }
func (*Userinfo) ProtoMessage()    {}
func (*Userinfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3d914a4924265a, []int{4}
}

func (m *Userinfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Userinfo.Unmarshal(m, b)
}
func (m *Userinfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Userinfo.Marshal(b, m, deterministic)
}
func (m *Userinfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Userinfo.Merge(m, src)
}
func (m *Userinfo) XXX_Size() int {
	return xxx_messageInfo_Userinfo.Size(m)
}
func (m *Userinfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Userinfo.DiscardUnknown(m)
}

var xxx_messageInfo_Userinfo proto.InternalMessageInfo

func (m *Userinfo) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Userinfo) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Userinfo) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *Userinfo) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *Userinfo) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *Userinfo) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Userinfo) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Userinfo) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type CreateReq struct {
	Userinfo             *Userinfo   `protobuf:"bytes,1,opt,name=userinfo,proto3" json:"userinfo,omitempty"`
	LoginList            []*LoginReq `protobuf:"bytes,2,rep,name=login_list,json=loginList,proto3" json:"login_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateReq) Reset()         { *m = CreateReq{} }
func (m *CreateReq) String() string { return proto.CompactTextString(m) }
func (*CreateReq) ProtoMessage()    {}
func (*CreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3d914a4924265a, []int{5}
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

func (m *CreateReq) GetUserinfo() *Userinfo {
	if m != nil {
		return m.Userinfo
	}
	return nil
}

func (m *CreateReq) GetLoginList() []*LoginReq {
	if m != nil {
		return m.LoginList
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "user_srv.LoginReq")
	proto.RegisterType((*LoginRsp)(nil), "user_srv.LoginRsp")
	proto.RegisterType((*UserReq)(nil), "user_srv.UserReq")
	proto.RegisterType((*UserRsp)(nil), "user_srv.UserRsp")
	proto.RegisterType((*Userinfo)(nil), "user_srv.Userinfo")
	proto.RegisterType((*CreateReq)(nil), "user_srv.CreateReq")
}

func init() { proto.RegisterFile("user_srv/user.proto", fileDescriptor_be3d914a4924265a) }

var fileDescriptor_be3d914a4924265a = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x52, 0xfa, 0x30,
	0x10, 0xc6, 0xff, 0x50, 0x28, 0xed, 0xf2, 0x3f, 0x68, 0x98, 0x71, 0x3a, 0x88, 0x33, 0xd8, 0x93,
	0xa7, 0x2a, 0x78, 0xf2, 0x88, 0x5e, 0x39, 0x55, 0x39, 0x77, 0x62, 0x1b, 0x98, 0x0c, 0x25, 0xad,
	0x49, 0xe0, 0xe1, 0x7c, 0x1a, 0x1f, 0xc5, 0xd9, 0xa4, 0x2d, 0x83, 0xe8, 0x0c, 0xdc, 0xfa, 0x7d,
	0xbf, 0xfd, 0xb2, 0xd9, 0x0d, 0xc0, 0x60, 0xab, 0x98, 0x4c, 0x94, 0xdc, 0xdd, 0xe3, 0x47, 0x54,
	0xca, 0x42, 0x17, 0xc4, 0xab, 0xcd, 0x90, 0x82, 0x37, 0x2f, 0x56, 0x5c, 0xc4, 0xec, 0x83, 0x0c,
	0xc1, 0x2b, 0x73, 0xaa, 0x97, 0x85, 0xdc, 0x04, 0xad, 0x71, 0xeb, 0xce, 0x8f, 0x1b, 0x4d, 0xae,
	0xc0, 0x2d, 0x4a, 0x26, 0x78, 0x16, 0xb4, 0x0d, 0xa9, 0x14, 0xb9, 0x85, 0xff, 0x34, 0x4d, 0x99,
	0x52, 0x89, 0x2e, 0xd6, 0x4c, 0x04, 0x8e, 0xa1, 0x7d, 0xeb, 0xbd, 0xa1, 0x15, 0x8e, 0xea, 0x16,
	0xaa, 0x24, 0x17, 0xe0, 0x6c, 0x79, 0x66, 0x4e, 0x77, 0x62, 0xfc, 0x0c, 0xaf, 0xa1, 0xb7, 0x50,
	0x4c, 0x62, 0xff, 0x63, 0xf8, 0x54, 0x41, 0x55, 0x92, 0x08, 0xcc, 0xa5, 0xb9, 0x58, 0x16, 0xa6,
	0xa2, 0x3f, 0x25, 0x51, 0x3d, 0x45, 0xb4, 0xa8, 0x48, 0xdc, 0xd4, 0x84, 0x5f, 0x2d, 0xf0, 0x6a,
	0xfb, 0xf8, 0x64, 0x9c, 0x55, 0xf0, 0x74, 0x2d, 0xe8, 0x86, 0x55, 0x13, 0x35, 0x9a, 0x8c, 0xc0,
	0x5f, 0x72, 0xa9, 0xb4, 0x81, 0x76, 0xa0, 0xbd, 0x81, 0xc9, 0x9c, 0x56, 0xb0, 0x63, 0x93, 0xb5,
	0xc6, 0x2d, 0xad, 0x98, 0xc8, 0x98, 0x0c, 0xba, 0x76, 0x4b, 0x56, 0xa1, 0x4f, 0x77, 0x54, 0x53,
	0x19, 0xb8, 0xd6, 0xb7, 0x8a, 0xdc, 0x00, 0x6c, 0xcb, 0x8c, 0x6a, 0x96, 0x25, 0x54, 0x07, 0x3d,
	0xdb, 0xaa, 0x72, 0x66, 0x1a, 0x71, 0x2a, 0x59, 0x8d, 0x3d, 0x8b, 0x2b, 0x67, 0xa6, 0x43, 0x01,
	0xfe, 0x8b, 0x11, 0xb8, 0xbc, 0x33, 0xf7, 0x43, 0x26, 0x00, 0x39, 0xbe, 0x4a, 0x92, 0x73, 0xa5,
	0x83, 0xf6, 0xd8, 0x39, 0x4c, 0xd4, 0x3f, 0x8a, 0xd8, 0x37, 0x55, 0x73, 0xae, 0xf4, 0xf4, 0xb3,
	0x0d, 0x7d, 0x3c, 0xe9, 0x95, 0xc9, 0x1d, 0x4f, 0x19, 0x99, 0x40, 0xd7, 0x94, 0x91, 0x5f, 0x72,
	0xc3, 0x23, 0x4f, 0x95, 0xe1, 0x3f, 0x32, 0x05, 0xd7, 0x5e, 0x99, 0x0c, 0xf6, 0xbc, 0x19, 0x62,
	0x78, 0x79, 0x78, 0x65, 0x9b, 0x79, 0x80, 0xce, 0x33, 0x17, 0xd9, 0x59, 0x09, 0x77, 0x21, 0x4c,
	0xe6, 0x27, 0xfe, 0x2b, 0x11, 0x41, 0x07, 0xc5, 0xc9, 0xf5, 0xd8, 0xc1, 0x3c, 0xd3, 0xa9, 0x89,
	0x77, 0xd7, 0xfc, 0xf3, 0x1e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x82, 0xdb, 0x99, 0x90,
	0x03, 0x00, 0x00,
}
