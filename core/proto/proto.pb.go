// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto.proto

package _proto

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

//注册
type C2GGameRegister struct {
	GroupId              string   `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	TypeName             string   `protobuf:"bytes,2,opt,name=typeName,proto3" json:"typeName,omitempty"`
	Pid                  string   `protobuf:"bytes,3,opt,name=pid,proto3" json:"pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2GGameRegister) Reset()         { *m = C2GGameRegister{} }
func (m *C2GGameRegister) String() string { return proto.CompactTextString(m) }
func (*C2GGameRegister) ProtoMessage()    {}
func (*C2GGameRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{0}
}

func (m *C2GGameRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2GGameRegister.Unmarshal(m, b)
}
func (m *C2GGameRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2GGameRegister.Marshal(b, m, deterministic)
}
func (m *C2GGameRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2GGameRegister.Merge(m, src)
}
func (m *C2GGameRegister) XXX_Size() int {
	return xxx_messageInfo_C2GGameRegister.Size(m)
}
func (m *C2GGameRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_C2GGameRegister.DiscardUnknown(m)
}

var xxx_messageInfo_C2GGameRegister proto.InternalMessageInfo

func (m *C2GGameRegister) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *C2GGameRegister) GetTypeName() string {
	if m != nil {
		return m.TypeName
	}
	return ""
}

func (m *C2GGameRegister) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

//接收注册的消息
type S2CGameRegister struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CGameRegister) Reset()         { *m = S2CGameRegister{} }
func (m *S2CGameRegister) String() string { return proto.CompactTextString(m) }
func (*S2CGameRegister) ProtoMessage()    {}
func (*S2CGameRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{1}
}

func (m *S2CGameRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CGameRegister.Unmarshal(m, b)
}
func (m *S2CGameRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CGameRegister.Marshal(b, m, deterministic)
}
func (m *S2CGameRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CGameRegister.Merge(m, src)
}
func (m *S2CGameRegister) XXX_Size() int {
	return xxx_messageInfo_S2CGameRegister.Size(m)
}
func (m *S2CGameRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CGameRegister.DiscardUnknown(m)
}

var xxx_messageInfo_S2CGameRegister proto.InternalMessageInfo

func init() {
	proto.RegisterType((*C2GGameRegister)(nil), "_proto.c2g_game_register")
	proto.RegisterType((*S2CGameRegister)(nil), "_proto.s2c_game_register")
}

func init() { proto.RegisterFile("proto.proto", fileDescriptor_2fcc84b9998d60d8) }

var fileDescriptor_2fcc84b9998d60d8 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x03, 0x93, 0x42, 0x6c, 0xf1, 0x60, 0x5a, 0x29, 0x9a, 0x4b, 0x30, 0xd9, 0x28, 0x3d,
	0x3e, 0x3d, 0x31, 0x37, 0x35, 0xbe, 0x28, 0x35, 0x3d, 0xb3, 0xb8, 0x24, 0xb5, 0x48, 0x48, 0x82,
	0x8b, 0x3d, 0xbd, 0x28, 0xbf, 0xb4, 0xc0, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08,
	0xc6, 0x15, 0x92, 0xe2, 0xe2, 0x28, 0xa9, 0x2c, 0x48, 0xf5, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02,
	0x4b, 0xc1, 0xf9, 0x42, 0x02, 0x5c, 0xcc, 0x05, 0x99, 0x29, 0x12, 0xcc, 0x60, 0x61, 0x10, 0x53,
	0x49, 0x98, 0x4b, 0xb0, 0xd8, 0x28, 0x19, 0xd5, 0xf0, 0x24, 0x36, 0xb0, 0xc5, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd1, 0x79, 0x1c, 0x20, 0x8f, 0x00, 0x00, 0x00,
}
