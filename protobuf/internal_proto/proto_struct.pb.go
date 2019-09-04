// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto_struct.proto

package internal_proto

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

//用户信息
type S_UserInfo struct {
	//用户id
	UserGuid string `protobuf:"bytes,1,opt,name=userGuid,proto3" json:"userGuid,omitempty"`
	//用户等级
	UserLv int32 `protobuf:"varint,2,opt,name=userLv,proto3" json:"userLv,omitempty"`
	//用户昵称
	UserName string `protobuf:"bytes,3,opt,name=userName,proto3" json:"userName,omitempty"`
	//Vip等级
	VipLV                int32    `protobuf:"varint,4,opt,name=vipLV,proto3" json:"vipLV,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S_UserInfo) Reset()         { *m = S_UserInfo{} }
func (m *S_UserInfo) String() string { return proto.CompactTextString(m) }
func (*S_UserInfo) ProtoMessage()    {}
func (*S_UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc4bef41498765a3, []int{0}
}

func (m *S_UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S_UserInfo.Unmarshal(m, b)
}
func (m *S_UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S_UserInfo.Marshal(b, m, deterministic)
}
func (m *S_UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S_UserInfo.Merge(m, src)
}
func (m *S_UserInfo) XXX_Size() int {
	return xxx_messageInfo_S_UserInfo.Size(m)
}
func (m *S_UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_S_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_S_UserInfo proto.InternalMessageInfo

func (m *S_UserInfo) GetUserGuid() string {
	if m != nil {
		return m.UserGuid
	}
	return ""
}

func (m *S_UserInfo) GetUserLv() int32 {
	if m != nil {
		return m.UserLv
	}
	return 0
}

func (m *S_UserInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *S_UserInfo) GetVipLV() int32 {
	if m != nil {
		return m.VipLV
	}
	return 0
}

func init() {
	proto.RegisterType((*S_UserInfo)(nil), "internal_proto.S_UserInfo")
}

func init() { proto.RegisterFile("proto_struct.proto", fileDescriptor_fc4bef41498765a3) }

var fileDescriptor_fc4bef41498765a3 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x8f, 0x2f, 0x2e, 0x29, 0x2a, 0x4d, 0x2e, 0xd1, 0x03, 0x73, 0x84, 0xf8, 0x32, 0xf3, 0x4a,
	0x52, 0x8b, 0xf2, 0x12, 0x73, 0xe2, 0xc1, 0x7c, 0xa5, 0x22, 0x2e, 0xae, 0xe0, 0xf8, 0xd0, 0xe2,
	0xd4, 0x22, 0xcf, 0xbc, 0xb4, 0x7c, 0x21, 0x29, 0x2e, 0x8e, 0xd2, 0xe2, 0xd4, 0x22, 0xf7, 0xd2,
	0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38, 0x5f, 0x48, 0x8c, 0x8b, 0x0d, 0xc4,
	0xf6, 0x29, 0x93, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0xf2, 0x60, 0x7a, 0xfc, 0x12, 0x73,
	0x53, 0x25, 0x98, 0x11, 0x7a, 0x40, 0x7c, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xcc, 0x02, 0x9f, 0x30,
	0x09, 0x16, 0xb0, 0x16, 0x08, 0x27, 0x89, 0x0d, 0x6c, 0xb5, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x64, 0xd2, 0x2f, 0xfe, 0xa0, 0x00, 0x00, 0x00,
}
