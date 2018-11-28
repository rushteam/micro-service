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
	return fileDescriptor_be3d914a4924265a, []int{0}
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

type LoginReq struct {
	Platform             string   `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	Login                string   `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3d914a4924265a, []int{1}
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

func (m *LoginReq) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginRsp struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	AccessToken          string   `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRsp) Reset()         { *m = LoginRsp{} }
func (m *LoginRsp) String() string { return proto.CompactTextString(m) }
func (*LoginRsp) ProtoMessage()    {}
func (*LoginRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3d914a4924265a, []int{2}
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

func (m *LoginRsp) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
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
	return fileDescriptor_be3d914a4924265a, []int{3}
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
	return fileDescriptor_be3d914a4924265a, []int{4}
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

type UserReq struct {
	Jwt                  string   `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserReq) Reset()         { *m = UserReq{} }
func (m *UserReq) String() string { return proto.CompactTextString(m) }
func (*UserReq) ProtoMessage()    {}
func (*UserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3d914a4924265a, []int{5}
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

func (m *UserReq) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

type UpdateReq struct {
	Jwt                  string    `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	Userinfo             *Userinfo `protobuf:"bytes,2,opt,name=userinfo,proto3" json:"userinfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateReq) Reset()         { *m = UpdateReq{} }
func (m *UpdateReq) String() string { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()    {}
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3d914a4924265a, []int{6}
}

func (m *UpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReq.Unmarshal(m, b)
}
func (m *UpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReq.Marshal(b, m, deterministic)
}
func (m *UpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReq.Merge(m, src)
}
func (m *UpdateReq) XXX_Size() int {
	return xxx_messageInfo_UpdateReq.Size(m)
}
func (m *UpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReq proto.InternalMessageInfo

func (m *UpdateReq) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

func (m *UpdateReq) GetUserinfo() *Userinfo {
	if m != nil {
		return m.Userinfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Userinfo)(nil), "user_srv.Userinfo")
	proto.RegisterType((*LoginReq)(nil), "user_srv.LoginReq")
	proto.RegisterType((*LoginRsp)(nil), "user_srv.LoginRsp")
	proto.RegisterType((*CreateReq)(nil), "user_srv.CreateReq")
	proto.RegisterType((*UserRsp)(nil), "user_srv.UserRsp")
	proto.RegisterType((*UserReq)(nil), "user_srv.UserReq")
	proto.RegisterType((*UpdateReq)(nil), "user_srv.UpdateReq")
}

func init() { proto.RegisterFile("user_srv/user.proto", fileDescriptor_be3d914a4924265a) }

var fileDescriptor_be3d914a4924265a = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x3d, 0xaf, 0xd3, 0x30,
	0x14, 0x25, 0xcd, 0x6b, 0x5e, 0x72, 0xcb, 0x00, 0x7e, 0x08, 0x45, 0x05, 0xa4, 0xe2, 0xe9, 0x4d,
	0x41, 0x2f, 0x4c, 0x4c, 0xa8, 0x62, 0x2d, 0x4b, 0xa0, 0x12, 0x5b, 0x64, 0x12, 0xb7, 0x32, 0x4d,
	0x93, 0xd4, 0x76, 0xdb, 0xbf, 0xca, 0xc8, 0x4f, 0x41, 0xd7, 0x1f, 0x29, 0x85, 0x20, 0xc4, 0xe6,
	0x73, 0xce, 0xfd, 0x3a, 0xf7, 0xca, 0x70, 0x77, 0x54, 0x5c, 0x96, 0x4a, 0x9e, 0xde, 0xe0, 0x23,
	0xeb, 0x65, 0xa7, 0x3b, 0x12, 0x7b, 0x92, 0xfe, 0x08, 0x20, 0x5e, 0x2b, 0x2e, 0x45, 0xbb, 0xe9,
	0xc8, 0x13, 0x08, 0x8f, 0xa2, 0x4e, 0x83, 0x45, 0x70, 0x1f, 0x16, 0xf8, 0x24, 0x73, 0x88, 0x5b,
	0x51, 0xed, 0x5a, 0xb6, 0xe7, 0xe9, 0x64, 0x11, 0xdc, 0x27, 0xc5, 0x80, 0xc9, 0x4b, 0x48, 0x36,
	0x42, 0x2a, 0x6d, 0xc4, 0xd0, 0x88, 0x17, 0x02, 0x33, 0x1b, 0xe6, 0xc4, 0x1b, 0x9b, 0xe9, 0x31,
	0x79, 0x0e, 0xd1, 0x96, 0xb7, 0x35, 0x97, 0xe9, 0xd4, 0x28, 0x0e, 0x21, 0xcf, 0x4e, 0x4c, 0x33,
	0x99, 0x46, 0x96, 0xb7, 0x88, 0xbc, 0x02, 0x38, 0xf6, 0x35, 0xd3, 0xbc, 0x2e, 0x99, 0x4e, 0x6f,
	0x6d, 0x2b, 0xc7, 0x2c, 0x35, 0xca, 0x95, 0xe4, 0x5e, 0x8e, 0xad, 0xec, 0x98, 0xa5, 0xa6, 0x5f,
	0x20, 0x5e, 0x75, 0x5b, 0xd1, 0x16, 0xfc, 0x80, 0x53, 0xf5, 0x0d, 0xd3, 0x9b, 0x4e, 0xee, 0x8d,
	0xcd, 0xa4, 0x18, 0x30, 0x79, 0x06, 0xd3, 0x06, 0xe3, 0x9c, 0x51, 0x0b, 0x4c, 0x06, 0x53, 0xea,
	0xdc, 0xc9, 0xda, 0x99, 0x1c, 0x30, 0x7d, 0xef, 0x2b, 0xab, 0x7e, 0x64, 0x77, 0xaf, 0xe1, 0x31,
	0xab, 0x2a, 0xae, 0x54, 0xa9, 0xbb, 0x1d, 0xf7, 0x65, 0x67, 0x96, 0xfb, 0x8c, 0x14, 0x6d, 0x21,
	0xf9, 0x60, 0xe6, 0xc4, 0xd9, 0x32, 0x30, 0x67, 0xc1, 0x4b, 0x98, 0x32, 0xb3, 0x9c, 0x64, 0xfe,
	0x4e, 0x99, 0xbf, 0x51, 0x31, 0xc4, 0x90, 0x07, 0x00, 0x33, 0x62, 0xd9, 0x08, 0xa5, 0xd3, 0xc9,
	0x22, 0xbc, 0xce, 0xf0, 0x9e, 0x8b, 0xc4, 0x44, 0xad, 0x84, 0xd2, 0xf4, 0x1d, 0xdc, 0x62, 0x21,
	0x9c, 0xf7, 0x3f, 0xbb, 0xd1, 0x17, 0x2e, 0x95, 0x1f, 0xd0, 0xea, 0xb7, 0xb3, 0x76, 0xfb, 0xc3,
	0x27, 0xfd, 0x08, 0xc9, 0xda, 0x9c, 0x63, 0x54, 0xbe, 0xea, 0x35, 0xf9, 0x77, 0xaf, 0xfc, 0x7b,
	0x00, 0x33, 0xa4, 0x3f, 0x71, 0x79, 0x12, 0x15, 0x27, 0x0f, 0x30, 0x35, 0x6e, 0xc8, 0x88, 0xbd,
	0xf9, 0x1f, 0x9c, 0xea, 0xe9, 0x23, 0x92, 0x43, 0x64, 0x37, 0x4b, 0xee, 0x2e, 0xfa, 0xb0, 0xeb,
	0xf9, 0xd3, 0xeb, 0xfe, 0x36, 0x27, 0x83, 0x1b, 0x04, 0xe4, 0x77, 0xf1, 0x6f, 0xf1, 0x39, 0x44,
	0xd6, 0xf5, 0xaf, 0x3d, 0x86, 0x3d, 0x8c, 0xe6, 0x7c, 0x8d, 0xcc, 0x07, 0x7c, 0xfb, 0x33, 0x00,
	0x00, 0xff, 0xff, 0x57, 0x10, 0x05, 0x00, 0x97, 0x03, 0x00, 0x00,
}
