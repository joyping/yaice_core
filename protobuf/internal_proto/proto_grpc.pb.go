// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto_grpc.proto

package internal_proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//客户端请求服务器->服务消息
type C_ServiceMsgRequest struct {
	MsgHandlerNumber     int32     `protobuf:"varint,1,opt,name=msgHandlerNumber,proto3" json:"msgHandlerNumber,omitempty"`
	Body                 *C2S_Body `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *C_ServiceMsgRequest) Reset()         { *m = C_ServiceMsgRequest{} }
func (m *C_ServiceMsgRequest) String() string { return proto.CompactTextString(m) }
func (*C_ServiceMsgRequest) ProtoMessage()    {}
func (*C_ServiceMsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38909006c036dd04, []int{0}
}

func (m *C_ServiceMsgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C_ServiceMsgRequest.Unmarshal(m, b)
}
func (m *C_ServiceMsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C_ServiceMsgRequest.Marshal(b, m, deterministic)
}
func (m *C_ServiceMsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C_ServiceMsgRequest.Merge(m, src)
}
func (m *C_ServiceMsgRequest) XXX_Size() int {
	return xxx_messageInfo_C_ServiceMsgRequest.Size(m)
}
func (m *C_ServiceMsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_C_ServiceMsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_C_ServiceMsgRequest proto.InternalMessageInfo

func (m *C_ServiceMsgRequest) GetMsgHandlerNumber() int32 {
	if m != nil {
		return m.MsgHandlerNumber
	}
	return 0
}

func (m *C_ServiceMsgRequest) GetBody() *C2S_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

//客户端收到服务器->
type S_ServiceMsgReply struct {
	MsgHandlerNumber     int32     `protobuf:"varint,1,opt,name=msgHandlerNumber,proto3" json:"msgHandlerNumber,omitempty"`
	Body                 *S2C_Body `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *S_ServiceMsgReply) Reset()         { *m = S_ServiceMsgReply{} }
func (m *S_ServiceMsgReply) String() string { return proto.CompactTextString(m) }
func (*S_ServiceMsgReply) ProtoMessage()    {}
func (*S_ServiceMsgReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_38909006c036dd04, []int{1}
}

func (m *S_ServiceMsgReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S_ServiceMsgReply.Unmarshal(m, b)
}
func (m *S_ServiceMsgReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S_ServiceMsgReply.Marshal(b, m, deterministic)
}
func (m *S_ServiceMsgReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S_ServiceMsgReply.Merge(m, src)
}
func (m *S_ServiceMsgReply) XXX_Size() int {
	return xxx_messageInfo_S_ServiceMsgReply.Size(m)
}
func (m *S_ServiceMsgReply) XXX_DiscardUnknown() {
	xxx_messageInfo_S_ServiceMsgReply.DiscardUnknown(m)
}

var xxx_messageInfo_S_ServiceMsgReply proto.InternalMessageInfo

func (m *S_ServiceMsgReply) GetMsgHandlerNumber() int32 {
	if m != nil {
		return m.MsgHandlerNumber
	}
	return 0
}

func (m *S_ServiceMsgReply) GetBody() *S2C_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*C_ServiceMsgRequest)(nil), "internal_proto.C_ServiceMsgRequest")
	proto.RegisterType((*S_ServiceMsgReply)(nil), "internal_proto.S_ServiceMsgReply")
}

func init() { proto.RegisterFile("proto_grpc.proto", fileDescriptor_38909006c036dd04) }

var fileDescriptor_38909006c036dd04 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xcf, 0x4a, 0x85, 0x40,
	0x14, 0xc6, 0x9b, 0xa8, 0x16, 0x27, 0x92, 0x7b, 0x27, 0x08, 0x91, 0xbb, 0xb8, 0xd9, 0xe6, 0x12,
	0x21, 0x61, 0x6f, 0x90, 0x9b, 0x36, 0x45, 0x38, 0x0f, 0x20, 0x3a, 0x1e, 0x44, 0xd0, 0x19, 0x3b,
	0x33, 0x06, 0xf3, 0x8c, 0xbd, 0x54, 0xe4, 0x1f, 0xe8, 0xa6, 0xcb, 0x76, 0x9e, 0xef, 0x73, 0xce,
	0xef, 0xfb, 0x0e, 0x6c, 0x3a, 0xd2, 0x56, 0x67, 0x15, 0x75, 0x32, 0x1a, 0x3e, 0xb9, 0x57, 0x2b,
	0x8b, 0xa4, 0xf2, 0x26, 0x1b, 0xe6, 0xc0, 0x23, 0x34, 0x9d, 0x56, 0x06, 0x47, 0x3f, 0xb8, 0x22,
	0xfc, 0xe8, 0xd1, 0xd8, 0x69, 0xe4, 0xe3, 0x02, 0x63, 0xa9, 0x97, 0x93, 0x16, 0x6a, 0xb8, 0x4e,
	0x32, 0x81, 0xf4, 0x59, 0x4b, 0x7c, 0x35, 0x55, 0x3a, 0x3e, 0xe0, 0xf7, 0xb0, 0x69, 0x4d, 0xf5,
	0x92, 0xab, 0xb2, 0x41, 0x7a, 0xeb, 0xdb, 0x02, 0xc9, 0x67, 0x7b, 0x76, 0x38, 0x4f, 0x17, 0x3a,
	0x7f, 0x80, 0xb3, 0x42, 0x97, 0xce, 0x3f, 0xdd, 0xb3, 0xc3, 0x65, 0xec, 0x47, 0xc7, 0xa1, 0xa2,
	0x24, 0x16, 0xd9, 0xb3, 0x2e, 0x5d, 0x3a, 0xfc, 0x15, 0xb6, 0xb0, 0x15, 0x47, 0xc0, 0xae, 0x71,
	0xff, 0x89, 0x13, 0x71, 0xf2, 0x0b, 0x17, 0x7f, 0x31, 0xf0, 0x26, 0x5a, 0xa2, 0x95, 0x42, 0x69,
	0x79, 0x09, 0x37, 0x29, 0x56, 0xb5, 0xb1, 0x48, 0x93, 0x33, 0xb7, 0xbe, 0x5b, 0x64, 0x5f, 0x9e,
	0x26, 0xb8, 0x5d, 0x10, 0xff, 0xd6, 0x09, 0x4f, 0x1e, 0x19, 0x17, 0xb0, 0x15, 0x4e, 0xc9, 0xf7,
	0x26, 0x77, 0x48, 0x33, 0x60, 0xb7, 0x76, 0x9c, 0x39, 0x4c, 0xb0, 0x5b, 0xeb, 0x32, 0xbb, 0x3f,
	0x4b, 0x8b, 0x8b, 0x41, 0x7f, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x6b, 0xb8, 0xee, 0x0b,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceConnectClient is the client API for ServiceConnect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceConnectClient interface {
	RegisterServiceRequest(ctx context.Context, in *C_ServiceMsgRequest, opts ...grpc.CallOption) (ServiceConnect_RegisterServiceRequestClient, error)
	SyncPlayerRequest(ctx context.Context, in *C2S_Register, opts ...grpc.CallOption) (ServiceConnect_SyncPlayerRequestClient, error)
}

type serviceConnectClient struct {
	cc *grpc.ClientConn
}

func NewServiceConnectClient(cc *grpc.ClientConn) ServiceConnectClient {
	return &serviceConnectClient{cc}
}

func (c *serviceConnectClient) RegisterServiceRequest(ctx context.Context, in *C_ServiceMsgRequest, opts ...grpc.CallOption) (ServiceConnect_RegisterServiceRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServiceConnect_serviceDesc.Streams[0], "/internal_proto.ServiceConnect/RegisterServiceRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceConnectRegisterServiceRequestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceConnect_RegisterServiceRequestClient interface {
	Recv() (*S_ServiceMsgReply, error)
	grpc.ClientStream
}

type serviceConnectRegisterServiceRequestClient struct {
	grpc.ClientStream
}

func (x *serviceConnectRegisterServiceRequestClient) Recv() (*S_ServiceMsgReply, error) {
	m := new(S_ServiceMsgReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceConnectClient) SyncPlayerRequest(ctx context.Context, in *C2S_Register, opts ...grpc.CallOption) (ServiceConnect_SyncPlayerRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServiceConnect_serviceDesc.Streams[1], "/internal_proto.ServiceConnect/SyncPlayerRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceConnectSyncPlayerRequestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceConnect_SyncPlayerRequestClient interface {
	Recv() (*S2C_Register, error)
	grpc.ClientStream
}

type serviceConnectSyncPlayerRequestClient struct {
	grpc.ClientStream
}

func (x *serviceConnectSyncPlayerRequestClient) Recv() (*S2C_Register, error) {
	m := new(S2C_Register)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceConnectServer is the server API for ServiceConnect service.
type ServiceConnectServer interface {
	RegisterServiceRequest(*C_ServiceMsgRequest, ServiceConnect_RegisterServiceRequestServer) error
	SyncPlayerRequest(*C2S_Register, ServiceConnect_SyncPlayerRequestServer) error
}

// UnimplementedServiceConnectServer can be embedded to have forward compatible implementations.
type UnimplementedServiceConnectServer struct {
}

func (*UnimplementedServiceConnectServer) RegisterServiceRequest(req *C_ServiceMsgRequest, srv ServiceConnect_RegisterServiceRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterServiceRequest not implemented")
}
func (*UnimplementedServiceConnectServer) SyncPlayerRequest(req *C2S_Register, srv ServiceConnect_SyncPlayerRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncPlayerRequest not implemented")
}

func RegisterServiceConnectServer(s *grpc.Server, srv ServiceConnectServer) {
	s.RegisterService(&_ServiceConnect_serviceDesc, srv)
}

func _ServiceConnect_RegisterServiceRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(C_ServiceMsgRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceConnectServer).RegisterServiceRequest(m, &serviceConnectRegisterServiceRequestServer{stream})
}

type ServiceConnect_RegisterServiceRequestServer interface {
	Send(*S_ServiceMsgReply) error
	grpc.ServerStream
}

type serviceConnectRegisterServiceRequestServer struct {
	grpc.ServerStream
}

func (x *serviceConnectRegisterServiceRequestServer) Send(m *S_ServiceMsgReply) error {
	return x.ServerStream.SendMsg(m)
}

func _ServiceConnect_SyncPlayerRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(C2S_Register)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceConnectServer).SyncPlayerRequest(m, &serviceConnectSyncPlayerRequestServer{stream})
}

type ServiceConnect_SyncPlayerRequestServer interface {
	Send(*S2C_Register) error
	grpc.ServerStream
}

type serviceConnectSyncPlayerRequestServer struct {
	grpc.ServerStream
}

func (x *serviceConnectSyncPlayerRequestServer) Send(m *S2C_Register) error {
	return x.ServerStream.SendMsg(m)
}

var _ServiceConnect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "internal_proto.ServiceConnect",
	HandlerType: (*ServiceConnectServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterServiceRequest",
			Handler:       _ServiceConnect_RegisterServiceRequest_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SyncPlayerRequest",
			Handler:       _ServiceConnect_SyncPlayerRequest_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto_grpc.proto",
}
