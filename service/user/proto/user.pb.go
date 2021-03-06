// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package yak_user

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegisterUserRequest struct {
	Record               string   `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterUserRequest) Reset()         { *m = RegisterUserRequest{} }
func (m *RegisterUserRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterUserRequest) ProtoMessage()    {}
func (*RegisterUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *RegisterUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterUserRequest.Unmarshal(m, b)
}
func (m *RegisterUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterUserRequest.Marshal(b, m, deterministic)
}
func (m *RegisterUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUserRequest.Merge(m, src)
}
func (m *RegisterUserRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterUserRequest.Size(m)
}
func (m *RegisterUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUserRequest proto.InternalMessageInfo

func (m *RegisterUserRequest) GetRecord() string {
	if m != nil {
		return m.Record
	}
	return ""
}

func (m *RegisterUserRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RegisterUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterUserResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterUserResponse) Reset()         { *m = RegisterUserResponse{} }
func (m *RegisterUserResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterUserResponse) ProtoMessage()    {}
func (*RegisterUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *RegisterUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterUserResponse.Unmarshal(m, b)
}
func (m *RegisterUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterUserResponse.Marshal(b, m, deterministic)
}
func (m *RegisterUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUserResponse.Merge(m, src)
}
func (m *RegisterUserResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterUserResponse.Size(m)
}
func (m *RegisterUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUserResponse proto.InternalMessageInfo

type LoginUserRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginUserRequest) Reset()         { *m = LoginUserRequest{} }
func (m *LoginUserRequest) String() string { return proto.CompactTextString(m) }
func (*LoginUserRequest) ProtoMessage()    {}
func (*LoginUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *LoginUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginUserRequest.Unmarshal(m, b)
}
func (m *LoginUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginUserRequest.Marshal(b, m, deterministic)
}
func (m *LoginUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginUserRequest.Merge(m, src)
}
func (m *LoginUserRequest) XXX_Size() int {
	return xxx_messageInfo_LoginUserRequest.Size(m)
}
func (m *LoginUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginUserRequest proto.InternalMessageInfo

func (m *LoginUserRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginUserResponse struct {
	Record               string   `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginUserResponse) Reset()         { *m = LoginUserResponse{} }
func (m *LoginUserResponse) String() string { return proto.CompactTextString(m) }
func (*LoginUserResponse) ProtoMessage()    {}
func (*LoginUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *LoginUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginUserResponse.Unmarshal(m, b)
}
func (m *LoginUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginUserResponse.Marshal(b, m, deterministic)
}
func (m *LoginUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginUserResponse.Merge(m, src)
}
func (m *LoginUserResponse) XXX_Size() int {
	return xxx_messageInfo_LoginUserResponse.Size(m)
}
func (m *LoginUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginUserResponse proto.InternalMessageInfo

func (m *LoginUserResponse) GetRecord() string {
	if m != nil {
		return m.Record
	}
	return ""
}

func (m *LoginUserResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginUserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginUserResponse) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type ResetUserRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	NewPassword          string   `protobuf:"bytes,3,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	Record               string   `protobuf:"bytes,4,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetUserRequest) Reset()         { *m = ResetUserRequest{} }
func (m *ResetUserRequest) String() string { return proto.CompactTextString(m) }
func (*ResetUserRequest) ProtoMessage()    {}
func (*ResetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *ResetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetUserRequest.Unmarshal(m, b)
}
func (m *ResetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetUserRequest.Marshal(b, m, deterministic)
}
func (m *ResetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetUserRequest.Merge(m, src)
}
func (m *ResetUserRequest) XXX_Size() int {
	return xxx_messageInfo_ResetUserRequest.Size(m)
}
func (m *ResetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResetUserRequest proto.InternalMessageInfo

func (m *ResetUserRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *ResetUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ResetUserRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

func (m *ResetUserRequest) GetRecord() string {
	if m != nil {
		return m.Record
	}
	return ""
}

type ResetUserResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetUserResponse) Reset()         { *m = ResetUserResponse{} }
func (m *ResetUserResponse) String() string { return proto.CompactTextString(m) }
func (*ResetUserResponse) ProtoMessage()    {}
func (*ResetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *ResetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetUserResponse.Unmarshal(m, b)
}
func (m *ResetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetUserResponse.Marshal(b, m, deterministic)
}
func (m *ResetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetUserResponse.Merge(m, src)
}
func (m *ResetUserResponse) XXX_Size() int {
	return xxx_messageInfo_ResetUserResponse.Size(m)
}
func (m *ResetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResetUserResponse proto.InternalMessageInfo

type UpdateUserProfileRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	UserEmail            string   `protobuf:"bytes,2,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	UserPhone            string   `protobuf:"bytes,3,opt,name=userPhone,proto3" json:"userPhone,omitempty"`
	Record               string   `protobuf:"bytes,4,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserProfileRequest) Reset()         { *m = UpdateUserProfileRequest{} }
func (m *UpdateUserProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserProfileRequest) ProtoMessage()    {}
func (*UpdateUserProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UpdateUserProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserProfileRequest.Unmarshal(m, b)
}
func (m *UpdateUserProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserProfileRequest.Merge(m, src)
}
func (m *UpdateUserProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserProfileRequest.Size(m)
}
func (m *UpdateUserProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserProfileRequest proto.InternalMessageInfo

func (m *UpdateUserProfileRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UpdateUserProfileRequest) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *UpdateUserProfileRequest) GetUserPhone() string {
	if m != nil {
		return m.UserPhone
	}
	return ""
}

func (m *UpdateUserProfileRequest) GetRecord() string {
	if m != nil {
		return m.Record
	}
	return ""
}

type UpdateUserProfileResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserProfileResponse) Reset()         { *m = UpdateUserProfileResponse{} }
func (m *UpdateUserProfileResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserProfileResponse) ProtoMessage()    {}
func (*UpdateUserProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UpdateUserProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserProfileResponse.Unmarshal(m, b)
}
func (m *UpdateUserProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserProfileResponse.Marshal(b, m, deterministic)
}
func (m *UpdateUserProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserProfileResponse.Merge(m, src)
}
func (m *UpdateUserProfileResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserProfileResponse.Size(m)
}
func (m *UpdateUserProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserProfileResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterUserRequest)(nil), "yak.user.RegisterUserRequest")
	proto.RegisterType((*RegisterUserResponse)(nil), "yak.user.RegisterUserResponse")
	proto.RegisterType((*LoginUserRequest)(nil), "yak.user.LoginUserRequest")
	proto.RegisterType((*LoginUserResponse)(nil), "yak.user.LoginUserResponse")
	proto.RegisterType((*ResetUserRequest)(nil), "yak.user.ResetUserRequest")
	proto.RegisterType((*ResetUserResponse)(nil), "yak.user.ResetUserResponse")
	proto.RegisterType((*UpdateUserProfileRequest)(nil), "yak.user.UpdateUserProfileRequest")
	proto.RegisterType((*UpdateUserProfileResponse)(nil), "yak.user.UpdateUserProfileResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0xa5, 0x80, 0x04, 0x06, 0x0f, 0xb0, 0x10, 0x52, 0x8b, 0x1a, 0xb2, 0x5e, 0x3c, 0xf5, 0xa0,
	0xdf, 0xa0, 0x07, 0x63, 0x94, 0xd4, 0x70, 0x35, 0xa9, 0x30, 0x62, 0x23, 0x74, 0xeb, 0x6e, 0x91,
	0x78, 0xf7, 0xe4, 0xc7, 0xf9, 0x4d, 0x66, 0x77, 0x5b, 0xba, 0x94, 0x96, 0x18, 0xe3, 0xad, 0xf3,
	0xde, 0xee, 0xcc, 0x9b, 0xb7, 0xaf, 0x00, 0x2b, 0x81, 0xdc, 0x8d, 0x38, 0x8b, 0x19, 0x69, 0x7e,
	0xf8, 0xaf, 0xae, 0xac, 0x29, 0x42, 0xcf, 0xc3, 0x79, 0x20, 0x62, 0xe4, 0x13, 0x81, 0xdc, 0xc3,
	0xb7, 0x15, 0x8a, 0x98, 0x0c, 0xa0, 0xc1, 0x71, 0xca, 0xf8, 0xcc, 0xb6, 0x46, 0xd6, 0x79, 0xcb,
	0x4b, 0x2a, 0xe2, 0x40, 0x53, 0x5e, 0xbb, 0xf3, 0x97, 0x68, 0x57, 0x15, 0xb3, 0xa9, 0x25, 0x17,
	0xf9, 0x42, 0xac, 0xe5, 0xad, 0x9a, 0xe6, 0xd2, 0x9a, 0x0e, 0xa0, 0xbf, 0x3d, 0x46, 0x44, 0x2c,
	0x14, 0x48, 0x6f, 0xa0, 0x73, 0xcb, 0xe6, 0x41, 0x68, 0xce, 0x36, 0x67, 0x58, 0x7b, 0x66, 0x54,
	0x73, 0x33, 0x04, 0x74, 0x8d, 0x5e, 0x7a, 0xc0, 0x9f, 0x16, 0xe9, 0xc3, 0x01, 0x2e, 0xfd, 0x60,
	0x91, 0x6c, 0xa1, 0x0b, 0x89, 0x46, 0x2f, 0x2c, 0x44, 0xbb, 0xae, 0x51, 0x55, 0xd0, 0x4f, 0x0b,
	0x3a, 0x1e, 0x0a, 0x8c, 0xff, 0x61, 0x03, 0x32, 0x82, 0x76, 0x88, 0xeb, 0xf1, 0xb6, 0x89, 0x26,
	0x64, 0xac, 0x53, 0x37, 0xd7, 0xa1, 0x3d, 0xe8, 0x1a, 0x2a, 0x12, 0x73, 0xbf, 0x2c, 0xb0, 0x27,
	0xd1, 0xcc, 0x8f, 0x51, 0xc2, 0x63, 0xce, 0x9e, 0x83, 0x05, 0xfe, 0x46, 0xe3, 0x31, 0xb4, 0xe4,
	0xf7, 0x95, 0x32, 0x41, 0x8b, 0xcc, 0x80, 0x94, 0x1d, 0x2b, 0x33, 0x6a, 0x19, 0xab, 0x80, 0x52,
	0x85, 0x43, 0x38, 0x2a, 0xd0, 0xa2, 0x95, 0x5e, 0x7c, 0x57, 0xa1, 0x2d, 0xf1, 0x07, 0xe4, 0xef,
	0xc1, 0x14, 0xc9, 0x3d, 0x1c, 0x9a, 0x71, 0x21, 0x27, 0x6e, 0x1a, 0x58, 0xb7, 0x20, 0xad, 0xce,
	0x69, 0x19, 0x9d, 0x18, 0x51, 0x21, 0xd7, 0xd0, 0xda, 0x64, 0x83, 0x38, 0xd9, 0xf1, 0x7c, 0xf8,
	0x9c, 0x61, 0x21, 0x67, 0xf6, 0xd9, 0xf8, 0x6c, 0xf6, 0xc9, 0x47, 0xc0, 0xec, 0xb3, 0xfb, 0x30,
	0x15, 0xf2, 0x08, 0xdd, 0x1d, 0x37, 0x08, 0xcd, 0xee, 0x94, 0x3d, 0x9b, 0x73, 0xb6, 0xf7, 0x4c,
	0xda, 0xff, 0xa9, 0xa1, 0xfe, 0xf3, 0xcb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x07, 0x6e,
	0x3e, 0xf5, 0x03, 0x00, 0x00,
}
