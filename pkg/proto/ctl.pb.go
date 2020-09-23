// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: ctl.proto

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

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command     string      `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	CommandType CommandType `protobuf:"varint,2,opt,name=commandType,proto3,enum=proto.CommandType" json:"commandType,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_ctl_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Input) GetCommandType() CommandType {
	if x != nil {
		return x.CommandType
	}
	return CommandType_Shell
}

type Update struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandType CommandType `protobuf:"varint,1,opt,name=commandType,proto3,enum=proto.CommandType" json:"commandType,omitempty"`
	Error       string      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Response    string      `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Update) Reset() {
	*x = Update{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ctl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Update) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Update) ProtoMessage() {}

func (x *Update) ProtoReflect() protoreflect.Message {
	mi := &file_ctl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Update.ProtoReflect.Descriptor instead.
func (*Update) Descriptor() ([]byte, []int) {
	return file_ctl_proto_rawDescGZIP(), []int{1}
}

func (x *Update) GetCommandType() CommandType {
	if x != nil {
		return x.CommandType
	}
	return CommandType_Shell
}

func (x *Update) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Update) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_ctl_proto protoreflect.FileDescriptor

var file_ctl_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x74, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x57, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x34, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x70, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x2e, 0x0a, 0x03, 0x43, 0x54,
	0x4c, 0x12, 0x27, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ctl_proto_rawDescOnce sync.Once
	file_ctl_proto_rawDescData = file_ctl_proto_rawDesc
)

func file_ctl_proto_rawDescGZIP() []byte {
	file_ctl_proto_rawDescOnce.Do(func() {
		file_ctl_proto_rawDescData = protoimpl.X.CompressGZIP(file_ctl_proto_rawDescData)
	})
	return file_ctl_proto_rawDescData
}

var file_ctl_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ctl_proto_goTypes = []interface{}{
	(*Input)(nil),    // 0: proto.Input
	(*Update)(nil),   // 1: proto.Update
	(CommandType)(0), // 2: proto.CommandType
}
var file_ctl_proto_depIdxs = []int32{
	2, // 0: proto.Input.commandType:type_name -> proto.CommandType
	2, // 1: proto.Update.commandType:type_name -> proto.CommandType
	0, // 2: proto.CTL.Send:input_type -> proto.Input
	1, // 3: proto.CTL.Send:output_type -> proto.Update
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ctl_proto_init() }
func file_ctl_proto_init() {
	if File_ctl_proto != nil {
		return
	}
	file_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ctl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_ctl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Update); i {
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
			RawDescriptor: file_ctl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ctl_proto_goTypes,
		DependencyIndexes: file_ctl_proto_depIdxs,
		MessageInfos:      file_ctl_proto_msgTypes,
	}.Build()
	File_ctl_proto = out.File
	file_ctl_proto_rawDesc = nil
	file_ctl_proto_goTypes = nil
	file_ctl_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CTLClient is the client API for CTL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CTLClient interface {
	Send(ctx context.Context, in *Input, opts ...grpc.CallOption) (CTL_SendClient, error)
}

type cTLClient struct {
	cc grpc.ClientConnInterface
}

func NewCTLClient(cc grpc.ClientConnInterface) CTLClient {
	return &cTLClient{cc}
}

func (c *cTLClient) Send(ctx context.Context, in *Input, opts ...grpc.CallOption) (CTL_SendClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CTL_serviceDesc.Streams[0], "/proto.CTL/Send", opts...)
	if err != nil {
		return nil, err
	}
	x := &cTLSendClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CTL_SendClient interface {
	Recv() (*Update, error)
	grpc.ClientStream
}

type cTLSendClient struct {
	grpc.ClientStream
}

func (x *cTLSendClient) Recv() (*Update, error) {
	m := new(Update)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CTLServer is the server API for CTL service.
type CTLServer interface {
	Send(*Input, CTL_SendServer) error
}

// UnimplementedCTLServer can be embedded to have forward compatible implementations.
type UnimplementedCTLServer struct {
}

func (*UnimplementedCTLServer) Send(*Input, CTL_SendServer) error {
	return status.Errorf(codes.Unimplemented, "method Send not implemented")
}

func RegisterCTLServer(s *grpc.Server, srv CTLServer) {
	s.RegisterService(&_CTL_serviceDesc, srv)
}

func _CTL_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Input)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CTLServer).Send(m, &cTLSendServer{stream})
}

type CTL_SendServer interface {
	Send(*Update) error
	grpc.ServerStream
}

type cTLSendServer struct {
	grpc.ServerStream
}

func (x *cTLSendServer) Send(m *Update) error {
	return x.ServerStream.SendMsg(m)
}

var _CTL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CTL",
	HandlerType: (*CTLServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Send",
			Handler:       _CTL_Send_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ctl.proto",
}
