// Code generated by protoc-gen-go. DO NOT EDIT.
// source: command-service.proto

package grpc_commands

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

type CommandResponse struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	StandardOutput       string   `protobuf:"bytes,2,opt,name=standardOutput,proto3" json:"standardOutput,omitempty"`
	ErrorOutput          string   `protobuf:"bytes,3,opt,name=errorOutput,proto3" json:"errorOutput,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandResponse) Reset()         { *m = CommandResponse{} }
func (m *CommandResponse) String() string { return proto.CompactTextString(m) }
func (*CommandResponse) ProtoMessage()    {}
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e9667d4afd994f, []int{0}
}

func (m *CommandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandResponse.Unmarshal(m, b)
}
func (m *CommandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandResponse.Marshal(b, m, deterministic)
}
func (m *CommandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandResponse.Merge(m, src)
}
func (m *CommandResponse) XXX_Size() int {
	return xxx_messageInfo_CommandResponse.Size(m)
}
func (m *CommandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommandResponse proto.InternalMessageInfo

func (m *CommandResponse) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *CommandResponse) GetStandardOutput() string {
	if m != nil {
		return m.StandardOutput
	}
	return ""
}

func (m *CommandResponse) GetErrorOutput() string {
	if m != nil {
		return m.ErrorOutput
	}
	return ""
}

type CommandRequest struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Params               string   `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandRequest) Reset()         { *m = CommandRequest{} }
func (m *CommandRequest) String() string { return proto.CompactTextString(m) }
func (*CommandRequest) ProtoMessage()    {}
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3e9667d4afd994f, []int{1}
}

func (m *CommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandRequest.Unmarshal(m, b)
}
func (m *CommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandRequest.Marshal(b, m, deterministic)
}
func (m *CommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandRequest.Merge(m, src)
}
func (m *CommandRequest) XXX_Size() int {
	return xxx_messageInfo_CommandRequest.Size(m)
}
func (m *CommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommandRequest proto.InternalMessageInfo

func (m *CommandRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *CommandRequest) GetParams() string {
	if m != nil {
		return m.Params
	}
	return ""
}

func init() {
	proto.RegisterType((*CommandResponse)(nil), "CommandResponse")
	proto.RegisterType((*CommandRequest)(nil), "CommandRequest")
}

func init() { proto.RegisterFile("command-service.proto", fileDescriptor_f3e9667d4afd994f) }

var fileDescriptor_f3e9667d4afd994f = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0xd1, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x57, 0x2a, 0xe5, 0xe2, 0x77, 0x86, 0x48, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x0a, 0x49, 0x70, 0xb1, 0x43, 0xd5, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42,
	0x6a, 0x5c, 0x7c, 0xc5, 0x25, 0x89, 0x79, 0x29, 0x89, 0x45, 0x29, 0xfe, 0xa5, 0x25, 0x05, 0xa5,
	0x25, 0x12, 0x4c, 0x60, 0x05, 0x68, 0xa2, 0x42, 0x0a, 0x5c, 0xdc, 0xa9, 0x45, 0x45, 0xf9, 0x45,
	0x50, 0x45, 0xcc, 0x60, 0x45, 0xc8, 0x42, 0x4a, 0x4e, 0x5c, 0x7c, 0x70, 0x6b, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0xf0, 0xd8, 0x2a, 0xc6, 0xc5, 0x56, 0x90, 0x58, 0x94, 0x98, 0x5b, 0x0c, 0xb5, 0x0d,
	0xca, 0x33, 0xf2, 0x80, 0x9b, 0x11, 0x0c, 0xf1, 0x92, 0x90, 0x19, 0x17, 0xb7, 0x73, 0x62, 0x4e,
	0x0e, 0x54, 0x54, 0x88, 0x5f, 0x0f, 0xd5, 0x0e, 0x29, 0x01, 0x3d, 0x34, 0xbf, 0x2a, 0x31, 0x68,
	0x30, 0x1a, 0x30, 0x26, 0xb1, 0x81, 0xc3, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x54, 0x4e,
	0x32, 0x06, 0x24, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommandServiceClient is the client API for CommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommandServiceClient interface {
	CallCommand(ctx context.Context, opts ...grpc.CallOption) (CommandService_CallCommandClient, error)
}

type commandServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommandServiceClient(cc *grpc.ClientConn) CommandServiceClient {
	return &commandServiceClient{cc}
}

func (c *commandServiceClient) CallCommand(ctx context.Context, opts ...grpc.CallOption) (CommandService_CallCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CommandService_serviceDesc.Streams[0], "/CommandService/CallCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &commandServiceCallCommandClient{stream}
	return x, nil
}

type CommandService_CallCommandClient interface {
	Send(*CommandRequest) error
	Recv() (*CommandResponse, error)
	grpc.ClientStream
}

type commandServiceCallCommandClient struct {
	grpc.ClientStream
}

func (x *commandServiceCallCommandClient) Send(m *CommandRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *commandServiceCallCommandClient) Recv() (*CommandResponse, error) {
	m := new(CommandResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommandServiceServer is the server API for CommandService service.
type CommandServiceServer interface {
	CallCommand(CommandService_CallCommandServer) error
}

// UnimplementedCommandServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommandServiceServer struct {
}

func (*UnimplementedCommandServiceServer) CallCommand(srv CommandService_CallCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method CallCommand not implemented")
}

func RegisterCommandServiceServer(s *grpc.Server, srv CommandServiceServer) {
	s.RegisterService(&_CommandService_serviceDesc, srv)
}

func _CommandService_CallCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CommandServiceServer).CallCommand(&commandServiceCallCommandServer{stream})
}

type CommandService_CallCommandServer interface {
	Send(*CommandResponse) error
	Recv() (*CommandRequest, error)
	grpc.ServerStream
}

type commandServiceCallCommandServer struct {
	grpc.ServerStream
}

func (x *commandServiceCallCommandServer) Send(m *CommandResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *commandServiceCallCommandServer) Recv() (*CommandRequest, error) {
	m := new(CommandRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CommandService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CommandService",
	HandlerType: (*CommandServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CallCommand",
			Handler:       _CommandService_CallCommand_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "command-service.proto",
}
