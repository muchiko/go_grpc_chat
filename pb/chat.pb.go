// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package pb

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

type Request struct {
	Cmd                  string   `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	RoomId               string   `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *Request) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *Request) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Payload struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "pb.Request")
	proto.RegisterType((*Payload)(nil), "pb.Payload")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xf2, 0xe3, 0x62, 0x0f, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0x4e, 0xce, 0x4d, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0xc4, 0xb9, 0xd8, 0x8b, 0xf2, 0xf3, 0x73, 0xe3, 0x33, 0x53,
	0x24, 0x98, 0xc0, 0xa2, 0x6c, 0x20, 0xae, 0x67, 0x8a, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71,
	0x71, 0x62, 0x7a, 0xaa, 0x04, 0x33, 0x58, 0x02, 0xc6, 0x55, 0x52, 0xe6, 0x62, 0x0f, 0x48, 0xac,
	0xcc, 0xc9, 0x4f, 0x44, 0x51, 0xc4, 0x88, 0xa2, 0xc8, 0xc8, 0x8a, 0x8b, 0xdb, 0x39, 0x23, 0xb1,
	0x24, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x48, 0x9b, 0x8b, 0x33, 0xa4, 0x28, 0x31, 0xaf,
	0xb8, 0x20, 0xbf, 0xa8, 0x44, 0x88, 0x5b, 0xaf, 0x20, 0x49, 0x0f, 0xea, 0x24, 0x29, 0x30, 0x07,
	0x6a, 0x9e, 0x12, 0x83, 0x06, 0xa3, 0x01, 0x63, 0x12, 0x1b, 0xd8, 0xed, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xef, 0x13, 0x00, 0xad, 0xc9, 0x00, 0x00, 0x00,
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
	Transport(ctx context.Context, opts ...grpc.CallOption) (ChatService_TransportClient, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Transport(ctx context.Context, opts ...grpc.CallOption) (ChatService_TransportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/pb.ChatService/Transport", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceTransportClient{stream}
	return x, nil
}

type ChatService_TransportClient interface {
	Send(*Request) error
	Recv() (*Payload, error)
	grpc.ClientStream
}

type chatServiceTransportClient struct {
	grpc.ClientStream
}

func (x *chatServiceTransportClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceTransportClient) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	Transport(ChatService_TransportServer) error
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) Transport(srv ChatService_TransportServer) error {
	return status.Errorf(codes.Unimplemented, "method Transport not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_Transport_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Transport(&chatServiceTransportServer{stream})
}

type ChatService_TransportServer interface {
	Send(*Payload) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type chatServiceTransportServer struct {
	grpc.ServerStream
}

func (x *chatServiceTransportServer) Send(m *Payload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceTransportServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Transport",
			Handler:       _ChatService_Transport_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}