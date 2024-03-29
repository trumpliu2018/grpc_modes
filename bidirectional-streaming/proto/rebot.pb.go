// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rebot.proto

package grpc_bidirectional_stream

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type ChatMessage struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_rebot_3911565cae177e50, []int{0}
}
func (m *ChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessage.Unmarshal(m, b)
}
func (m *ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessage.Marshal(b, m, deterministic)
}
func (dst *ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessage.Merge(dst, src)
}
func (m *ChatMessage) XXX_Size() int {
	return xxx_messageInfo_ChatMessage.Size(m)
}
func (m *ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessage proto.InternalMessageInfo

func (m *ChatMessage) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ChatMessage)(nil), "grpc.bidirectional.stream.ChatMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	GetChatMsg(ctx context.Context, opts ...grpc.CallOption) (ChatService_GetChatMsgClient, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) GetChatMsg(ctx context.Context, opts ...grpc.CallOption) (ChatService_GetChatMsgClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/grpc.bidirectional.stream.ChatService/GetChatMsg", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceGetChatMsgClient{stream}
	return x, nil
}

type ChatService_GetChatMsgClient interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatServiceGetChatMsgClient struct {
	grpc.ClientStream
}

func (x *chatServiceGetChatMsgClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceGetChatMsgClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	GetChatMsg(ChatService_GetChatMsgServer) error
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_GetChatMsg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).GetChatMsg(&chatServiceGetChatMsgServer{stream})
}

type ChatService_GetChatMsgServer interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type chatServiceGetChatMsgServer struct {
	grpc.ServerStream
}

func (x *chatServiceGetChatMsgServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceGetChatMsgServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.bidirectional.stream.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetChatMsg",
			Handler:       _ChatService_GetChatMsg_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rebot.proto",
}

func init() { proto.RegisterFile("rebot.proto", fileDescriptor_rebot_3911565cae177e50) }

var fileDescriptor_rebot_3911565cae177e50 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x4a, 0x4d, 0xca,
	0x2f, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4c, 0x2f, 0x2a, 0x48, 0xd6, 0x4b, 0xca,
	0x4c, 0xc9, 0x2c, 0x4a, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0x4b, 0xcc, 0xd1, 0x2b, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0x55, 0x92, 0xe7, 0xe2, 0x76, 0xce, 0x48, 0x2c, 0xf1, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
	0x4f, 0x15, 0x12, 0xe0, 0x62, 0xf6, 0x2d, 0x4e, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02,
	0x31, 0x8d, 0xf2, 0x21, 0x0a, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x12, 0xb8, 0xb8,
	0xdc, 0x53, 0x4b, 0xc0, 0x5a, 0x8a, 0xd3, 0x85, 0xd4, 0xf4, 0x70, 0x9a, 0xac, 0x87, 0x64, 0xac,
	0x14, 0x91, 0xea, 0x34, 0x18, 0x0d, 0x18, 0x93, 0xd8, 0xc0, 0x6e, 0x36, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x21, 0x2b, 0x8c, 0x05, 0xc2, 0x00, 0x00, 0x00,
}
