// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: sync.proto

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

type CheckName int32

const (
	CheckName_HealthCheck CheckName = 0
)

// Enum value maps for CheckName.
var (
	CheckName_name = map[int32]string{
		0: "HealthCheck",
	}
	CheckName_value = map[string]int32{
		"HealthCheck": 0,
	}
)

func (x CheckName) Enum() *CheckName {
	p := new(CheckName)
	*p = x
	return p
}

func (x CheckName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckName) Descriptor() protoreflect.EnumDescriptor {
	return file_sync_proto_enumTypes[0].Descriptor()
}

func (CheckName) Type() protoreflect.EnumType {
	return &file_sync_proto_enumTypes[0]
}

func (x CheckName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckName.Descriptor instead.
func (CheckName) EnumDescriptor() ([]byte, []int) {
	return file_sync_proto_rawDescGZIP(), []int{0}
}

type SyncSyn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandName CheckName `protobuf:"varint,1,opt,name=commandName,proto3,enum=proto.CheckName" json:"commandName,omitempty"`
	CommandArgs string    `protobuf:"bytes,2,opt,name=commandArgs,proto3" json:"commandArgs,omitempty"`
}

func (x *SyncSyn) Reset() {
	*x = SyncSyn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncSyn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncSyn) ProtoMessage() {}

func (x *SyncSyn) ProtoReflect() protoreflect.Message {
	mi := &file_sync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncSyn.ProtoReflect.Descriptor instead.
func (*SyncSyn) Descriptor() ([]byte, []int) {
	return file_sync_proto_rawDescGZIP(), []int{0}
}

func (x *SyncSyn) GetCommandName() CheckName {
	if x != nil {
		return x.CommandName
	}
	return CheckName_HealthCheck
}

func (x *SyncSyn) GetCommandArgs() string {
	if x != nil {
		return x.CommandArgs
	}
	return ""
}

type SyncAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandName CheckName `protobuf:"varint,1,opt,name=commandName,proto3,enum=proto.CheckName" json:"commandName,omitempty"`
	Response    string    `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	Error       string    `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SyncAck) Reset() {
	*x = SyncAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncAck) ProtoMessage() {}

func (x *SyncAck) ProtoReflect() protoreflect.Message {
	mi := &file_sync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncAck.ProtoReflect.Descriptor instead.
func (*SyncAck) Descriptor() ([]byte, []int) {
	return file_sync_proto_rawDescGZIP(), []int{1}
}

func (x *SyncAck) GetCommandName() CheckName {
	if x != nil {
		return x.CommandName
	}
	return CheckName_HealthCheck
}

func (x *SyncAck) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *SyncAck) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_sync_proto protoreflect.FileDescriptor

var file_sync_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x07, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x79, 0x6e, 0x12, 0x32,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x72, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x41, 0x72, 0x67, 0x73, 0x22, 0x6f, 0x0a, 0x07, 0x53, 0x79, 0x6e, 0x63, 0x41, 0x63, 0x6b, 0x12,
	0x32, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x10, 0x00, 0x32, 0x30, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x28, 0x0a, 0x04, 0x53,
	0x65, 0x6e, 0x64, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x53, 0x79, 0x6e, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x41, 0x63, 0x6b, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sync_proto_rawDescOnce sync.Once
	file_sync_proto_rawDescData = file_sync_proto_rawDesc
)

func file_sync_proto_rawDescGZIP() []byte {
	file_sync_proto_rawDescOnce.Do(func() {
		file_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_proto_rawDescData)
	})
	return file_sync_proto_rawDescData
}

var file_sync_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sync_proto_goTypes = []interface{}{
	(CheckName)(0),  // 0: proto.CheckName
	(*SyncSyn)(nil), // 1: proto.SyncSyn
	(*SyncAck)(nil), // 2: proto.SyncAck
}
var file_sync_proto_depIdxs = []int32{
	0, // 0: proto.SyncSyn.commandName:type_name -> proto.CheckName
	0, // 1: proto.SyncAck.commandName:type_name -> proto.CheckName
	1, // 2: proto.Sync.Send:input_type -> proto.SyncSyn
	2, // 3: proto.Sync.Send:output_type -> proto.SyncAck
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sync_proto_init() }
func file_sync_proto_init() {
	if File_sync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncSyn); i {
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
		file_sync_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncAck); i {
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
			RawDescriptor: file_sync_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sync_proto_goTypes,
		DependencyIndexes: file_sync_proto_depIdxs,
		EnumInfos:         file_sync_proto_enumTypes,
		MessageInfos:      file_sync_proto_msgTypes,
	}.Build()
	File_sync_proto = out.File
	file_sync_proto_rawDesc = nil
	file_sync_proto_goTypes = nil
	file_sync_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SyncClient is the client API for Sync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncClient interface {
	Send(ctx context.Context, in *SyncSyn, opts ...grpc.CallOption) (*SyncAck, error)
}

type syncClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncClient(cc grpc.ClientConnInterface) SyncClient {
	return &syncClient{cc}
}

func (c *syncClient) Send(ctx context.Context, in *SyncSyn, opts ...grpc.CallOption) (*SyncAck, error) {
	out := new(SyncAck)
	err := c.cc.Invoke(ctx, "/proto.Sync/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncServer is the server API for Sync service.
type SyncServer interface {
	Send(context.Context, *SyncSyn) (*SyncAck, error)
}

// UnimplementedSyncServer can be embedded to have forward compatible implementations.
type UnimplementedSyncServer struct {
}

func (*UnimplementedSyncServer) Send(context.Context, *SyncSyn) (*SyncAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}

func RegisterSyncServer(s *grpc.Server, srv SyncServer) {
	s.RegisterService(&_Sync_serviceDesc, srv)
}

func _Sync_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncSyn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Sync/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServer).Send(ctx, req.(*SyncSyn))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Sync",
	HandlerType: (*SyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Sync_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sync.proto",
}
