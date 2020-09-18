// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: ld.proto

package ld

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

type BreakdownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID  string `protobuf:"bytes,1,opt,name=projectID,proto3" json:"projectID,omitempty"`
	CommitHash string `protobuf:"bytes,2,opt,name=commitHash,proto3" json:"commitHash,omitempty"`
}

func (x *BreakdownRequest) Reset() {
	*x = BreakdownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BreakdownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakdownRequest) ProtoMessage() {}

func (x *BreakdownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreakdownRequest.ProtoReflect.Descriptor instead.
func (*BreakdownRequest) Descriptor() ([]byte, []int) {
	return file_ld_proto_rawDescGZIP(), []int{0}
}

func (x *BreakdownRequest) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *BreakdownRequest) GetCommitHash() string {
	if x != nil {
		return x.CommitHash
	}
	return ""
}

type BreakdownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Breakdown []*Language `protobuf:"bytes,1,rep,name=breakdown,proto3" json:"breakdown,omitempty"`
}

func (x *BreakdownResponse) Reset() {
	*x = BreakdownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BreakdownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakdownResponse) ProtoMessage() {}

func (x *BreakdownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreakdownResponse.ProtoReflect.Descriptor instead.
func (*BreakdownResponse) Descriptor() ([]byte, []int) {
	return file_ld_proto_rawDescGZIP(), []int{1}
}

func (x *BreakdownResponse) GetBreakdown() []*Language {
	if x != nil {
		return x.Breakdown
	}
	return nil
}

type Language struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Files []string `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *Language) Reset() {
	*x = Language{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Language) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Language) ProtoMessage() {}

func (x *Language) ProtoReflect() protoreflect.Message {
	mi := &file_ld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Language.ProtoReflect.Descriptor instead.
func (*Language) Descriptor() ([]byte, []int) {
	return file_ld_proto_rawDescGZIP(), []int{2}
}

func (x *Language) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Language) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_ld_proto protoreflect.FileDescriptor

var file_ld_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x10, 0x42, 0x72,
	0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x48, 0x61, 0x73, 0x68, 0x22, 0x3c, 0x0a, 0x11,
	0x42, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x09, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52,
	0x09, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x22, 0x34, 0x0a, 0x08, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x32, 0x43, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x32, 0x0a, 0x09, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x11,
	0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x6c, 0x64, 0x3b, 0x6c, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ld_proto_rawDescOnce sync.Once
	file_ld_proto_rawDescData = file_ld_proto_rawDesc
)

func file_ld_proto_rawDescGZIP() []byte {
	file_ld_proto_rawDescOnce.Do(func() {
		file_ld_proto_rawDescData = protoimpl.X.CompressGZIP(file_ld_proto_rawDescData)
	})
	return file_ld_proto_rawDescData
}

var file_ld_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ld_proto_goTypes = []interface{}{
	(*BreakdownRequest)(nil),  // 0: BreakdownRequest
	(*BreakdownResponse)(nil), // 1: BreakdownResponse
	(*Language)(nil),          // 2: Language
}
var file_ld_proto_depIdxs = []int32{
	2, // 0: BreakdownResponse.breakdown:type_name -> Language
	0, // 1: UsedLanguages.Breakdown:input_type -> BreakdownRequest
	1, // 2: UsedLanguages.Breakdown:output_type -> BreakdownResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ld_proto_init() }
func file_ld_proto_init() {
	if File_ld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BreakdownRequest); i {
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
		file_ld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BreakdownResponse); i {
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
		file_ld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Language); i {
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
			RawDescriptor: file_ld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ld_proto_goTypes,
		DependencyIndexes: file_ld_proto_depIdxs,
		MessageInfos:      file_ld_proto_msgTypes,
	}.Build()
	File_ld_proto = out.File
	file_ld_proto_rawDesc = nil
	file_ld_proto_goTypes = nil
	file_ld_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UsedLanguagesClient is the client API for UsedLanguages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsedLanguagesClient interface {
	Breakdown(ctx context.Context, in *BreakdownRequest, opts ...grpc.CallOption) (*BreakdownResponse, error)
}

type usedLanguagesClient struct {
	cc grpc.ClientConnInterface
}

func NewUsedLanguagesClient(cc grpc.ClientConnInterface) UsedLanguagesClient {
	return &usedLanguagesClient{cc}
}

func (c *usedLanguagesClient) Breakdown(ctx context.Context, in *BreakdownRequest, opts ...grpc.CallOption) (*BreakdownResponse, error) {
	out := new(BreakdownResponse)
	err := c.cc.Invoke(ctx, "/UsedLanguages/Breakdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsedLanguagesServer is the server API for UsedLanguages service.
type UsedLanguagesServer interface {
	Breakdown(context.Context, *BreakdownRequest) (*BreakdownResponse, error)
}

// UnimplementedUsedLanguagesServer can be embedded to have forward compatible implementations.
type UnimplementedUsedLanguagesServer struct {
}

func (*UnimplementedUsedLanguagesServer) Breakdown(context.Context, *BreakdownRequest) (*BreakdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Breakdown not implemented")
}

func RegisterUsedLanguagesServer(s *grpc.Server, srv UsedLanguagesServer) {
	s.RegisterService(&_UsedLanguages_serviceDesc, srv)
}

func _UsedLanguages_Breakdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BreakdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsedLanguagesServer).Breakdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UsedLanguages/Breakdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsedLanguagesServer).Breakdown(ctx, req.(*BreakdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UsedLanguages_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UsedLanguages",
	HandlerType: (*UsedLanguagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Breakdown",
			Handler:    _UsedLanguages_Breakdown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ld.proto",
}
