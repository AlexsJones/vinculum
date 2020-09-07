// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: command.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CommandSyn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeepAlive int64 `protobuf:"varint,1,opt,name=keepAlive,proto3" json:"keepAlive,omitempty"`
}

func (x *CommandSyn) Reset() {
	*x = CommandSyn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandSyn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandSyn) ProtoMessage() {}

func (x *CommandSyn) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandSyn.ProtoReflect.Descriptor instead.
func (*CommandSyn) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{0}
}

func (x *CommandSyn) GetKeepAlive() int64 {
	if x != nil {
		return x.KeepAlive
	}
	return 0
}

type CommandAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommandAck) Reset() {
	*x = CommandAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandAck) ProtoMessage() {}

func (x *CommandAck) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandAck.ProtoReflect.Descriptor instead.
func (*CommandAck) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{1}
}

var File_command_proto protoreflect.FileDescriptor

var file_command_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x53, 0x79, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69,
	0x76, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x63, 0x6b,
	0x32, 0x3f, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x34, 0x0a, 0x06, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x79, 0x6e, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_command_proto_rawDescOnce sync.Once
	file_command_proto_rawDescData = file_command_proto_rawDesc
)

func file_command_proto_rawDescGZIP() []byte {
	file_command_proto_rawDescOnce.Do(func() {
		file_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_command_proto_rawDescData)
	})
	return file_command_proto_rawDescData
}

var file_command_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_command_proto_goTypes = []interface{}{
	(*CommandSyn)(nil), // 0: proto.CommandSyn
	(*CommandAck)(nil), // 1: proto.CommandAck
}
var file_command_proto_depIdxs = []int32{
	0, // 0: proto.Command.Stream:input_type -> proto.CommandSyn
	1, // 1: proto.Command.Stream:output_type -> proto.CommandAck
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_command_proto_init() }
func file_command_proto_init() {
	if File_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandSyn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandAck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_command_proto_goTypes,
		DependencyIndexes: file_command_proto_depIdxs,
		MessageInfos:      file_command_proto_msgTypes,
	}.Build()
	File_command_proto = out.File
	file_command_proto_rawDesc = nil
	file_command_proto_goTypes = nil
	file_command_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommandClient is the client API for Command service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommandClient interface {
	// Leader to client stream
	Stream(ctx context.Context, opts ...grpc.CallOption) (Command_StreamClient, error)
}

type commandClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandClient(cc grpc.ClientConnInterface) CommandClient {
	return &commandClient{cc}
}

func (c *commandClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Command_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Command_serviceDesc.Streams[0], "/proto.Command/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &commandStreamClient{stream}
	return x, nil
}

type Command_StreamClient interface {
	Send(*CommandSyn) error
	Recv() (*CommandAck, error)
	grpc.ClientStream
}

type commandStreamClient struct {
	grpc.ClientStream
}

func (x *commandStreamClient) Send(m *CommandSyn) error {
	return x.ClientStream.SendMsg(m)
}

func (x *commandStreamClient) Recv() (*CommandAck, error) {
	m := new(CommandAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommandServer is the server API for Command service.
type CommandServer interface {
	// Leader to client stream
	Stream(Command_StreamServer) error
}

// UnimplementedCommandServer can be embedded to have forward compatible implementations.
type UnimplementedCommandServer struct {
}

func (*UnimplementedCommandServer) Stream(Command_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}

func RegisterCommandServer(s *grpc.Server, srv CommandServer) {
	s.RegisterService(&_Command_serviceDesc, srv)
}

func _Command_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CommandServer).Stream(&commandStreamServer{stream})
}

type Command_StreamServer interface {
	Send(*CommandAck) error
	Recv() (*CommandSyn, error)
	grpc.ServerStream
}

type commandStreamServer struct {
	grpc.ServerStream
}

func (x *commandStreamServer) Send(m *CommandAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *commandStreamServer) Recv() (*CommandSyn, error) {
	m := new(CommandSyn)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Command_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Command",
	HandlerType: (*CommandServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Command_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "command.proto",
}
