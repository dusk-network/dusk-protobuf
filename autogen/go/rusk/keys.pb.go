// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: keys.proto

package rusk

import (
	context "context"
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

type SecretKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *SecretKey) Reset() {
	*x = SecretKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretKey) ProtoMessage() {}

func (x *SecretKey) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretKey.ProtoReflect.Descriptor instead.
func (*SecretKey) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{0}
}

func (x *SecretKey) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type ViewKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ViewKey) Reset() {
	*x = ViewKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewKey) ProtoMessage() {}

func (x *ViewKey) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewKey.ProtoReflect.Descriptor instead.
func (*ViewKey) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{1}
}

func (x *ViewKey) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{2}
}

func (x *PublicKey) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type StealthAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *StealthAddress) Reset() {
	*x = StealthAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StealthAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StealthAddress) ProtoMessage() {}

func (x *StealthAddress) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StealthAddress.ProtoReflect.Descriptor instead.
func (*StealthAddress) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{3}
}

func (x *StealthAddress) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type GenerateKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GenerateKeysRequest) Reset() {
	*x = GenerateKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateKeysRequest) ProtoMessage() {}

func (x *GenerateKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateKeysRequest.ProtoReflect.Descriptor instead.
func (*GenerateKeysRequest) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{4}
}

type GenerateKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sk *SecretKey `protobuf:"bytes,1,opt,name=sk,proto3" json:"sk,omitempty"`
	Vk *ViewKey   `protobuf:"bytes,2,opt,name=vk,proto3" json:"vk,omitempty"`
	Pk *PublicKey `protobuf:"bytes,3,opt,name=pk,proto3" json:"pk,omitempty"`
}

func (x *GenerateKeysResponse) Reset() {
	*x = GenerateKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateKeysResponse) ProtoMessage() {}

func (x *GenerateKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateKeysResponse.ProtoReflect.Descriptor instead.
func (*GenerateKeysResponse) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{5}
}

func (x *GenerateKeysResponse) GetSk() *SecretKey {
	if x != nil {
		return x.Sk
	}
	return nil
}

func (x *GenerateKeysResponse) GetVk() *ViewKey {
	if x != nil {
		return x.Vk
	}
	return nil
}

func (x *GenerateKeysResponse) GetPk() *PublicKey {
	if x != nil {
		return x.Pk
	}
	return nil
}

var File_keys_proto protoreflect.FileDescriptor

var file_keys_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x75,
	0x73, 0x6b, 0x22, 0x25, 0x0a, 0x09, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x23, 0x0a, 0x07, 0x56, 0x69, 0x65,
	0x77, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x25,
	0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x77, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x02, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72,
	0x75, 0x73, 0x6b, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x02, 0x73,
	0x6b, 0x12, 0x1d, 0x0a, 0x02, 0x76, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x72, 0x75, 0x73, 0x6b, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x4b, 0x65, 0x79, 0x52, 0x02, 0x76, 0x6b,
	0x12, 0x1f, 0x0a, 0x02, 0x70, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72,
	0x75, 0x73, 0x6b, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x02, 0x70,
	0x6b, 0x32, 0x92, 0x01, 0x0a, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x19, 0x2e, 0x72, 0x75, 0x73,
	0x6b, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0f, 0x2e,
	0x72, 0x75, 0x73, 0x6b, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x1a, 0x14,
	0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x53, 0x74, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keys_proto_rawDescOnce sync.Once
	file_keys_proto_rawDescData = file_keys_proto_rawDesc
)

func file_keys_proto_rawDescGZIP() []byte {
	file_keys_proto_rawDescOnce.Do(func() {
		file_keys_proto_rawDescData = protoimpl.X.CompressGZIP(file_keys_proto_rawDescData)
	})
	return file_keys_proto_rawDescData
}

var file_keys_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_keys_proto_goTypes = []interface{}{
	(*SecretKey)(nil),            // 0: rusk.SecretKey
	(*ViewKey)(nil),              // 1: rusk.ViewKey
	(*PublicKey)(nil),            // 2: rusk.PublicKey
	(*StealthAddress)(nil),       // 3: rusk.StealthAddress
	(*GenerateKeysRequest)(nil),  // 4: rusk.GenerateKeysRequest
	(*GenerateKeysResponse)(nil), // 5: rusk.GenerateKeysResponse
}
var file_keys_proto_depIdxs = []int32{
	0, // 0: rusk.GenerateKeysResponse.sk:type_name -> rusk.SecretKey
	1, // 1: rusk.GenerateKeysResponse.vk:type_name -> rusk.ViewKey
	2, // 2: rusk.GenerateKeysResponse.pk:type_name -> rusk.PublicKey
	4, // 3: rusk.Keys.GenerateKeys:input_type -> rusk.GenerateKeysRequest
	2, // 4: rusk.Keys.GenerateStealthAddress:input_type -> rusk.PublicKey
	5, // 5: rusk.Keys.GenerateKeys:output_type -> rusk.GenerateKeysResponse
	3, // 6: rusk.Keys.GenerateStealthAddress:output_type -> rusk.StealthAddress
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_keys_proto_init() }
func file_keys_proto_init() {
	if File_keys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretKey); i {
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
		file_keys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewKey); i {
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
		file_keys_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicKey); i {
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
		file_keys_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StealthAddress); i {
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
		file_keys_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateKeysRequest); i {
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
		file_keys_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateKeysResponse); i {
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
			RawDescriptor: file_keys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keys_proto_goTypes,
		DependencyIndexes: file_keys_proto_depIdxs,
		MessageInfos:      file_keys_proto_msgTypes,
	}.Build()
	File_keys_proto = out.File
	file_keys_proto_rawDesc = nil
	file_keys_proto_goTypes = nil
	file_keys_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeysClient is the client API for Keys service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeysClient interface {
	GenerateKeys(ctx context.Context, in *GenerateKeysRequest, opts ...grpc.CallOption) (*GenerateKeysResponse, error)
	GenerateStealthAddress(ctx context.Context, in *PublicKey, opts ...grpc.CallOption) (*StealthAddress, error)
}

type keysClient struct {
	cc grpc.ClientConnInterface
}

func NewKeysClient(cc grpc.ClientConnInterface) KeysClient {
	return &keysClient{cc}
}

func (c *keysClient) GenerateKeys(ctx context.Context, in *GenerateKeysRequest, opts ...grpc.CallOption) (*GenerateKeysResponse, error) {
	out := new(GenerateKeysResponse)
	err := c.cc.Invoke(ctx, "/rusk.Keys/GenerateKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysClient) GenerateStealthAddress(ctx context.Context, in *PublicKey, opts ...grpc.CallOption) (*StealthAddress, error) {
	out := new(StealthAddress)
	err := c.cc.Invoke(ctx, "/rusk.Keys/GenerateStealthAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeysServer is the server API for Keys service.
type KeysServer interface {
	GenerateKeys(context.Context, *GenerateKeysRequest) (*GenerateKeysResponse, error)
	GenerateStealthAddress(context.Context, *PublicKey) (*StealthAddress, error)
}

// UnimplementedKeysServer can be embedded to have forward compatible implementations.
type UnimplementedKeysServer struct {
}

func (*UnimplementedKeysServer) GenerateKeys(context.Context, *GenerateKeysRequest) (*GenerateKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKeys not implemented")
}
func (*UnimplementedKeysServer) GenerateStealthAddress(context.Context, *PublicKey) (*StealthAddress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateStealthAddress not implemented")
}

func RegisterKeysServer(s *grpc.Server, srv KeysServer) {
	s.RegisterService(&_Keys_serviceDesc, srv)
}

func _Keys_GenerateKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServer).GenerateKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusk.Keys/GenerateKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServer).GenerateKeys(ctx, req.(*GenerateKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keys_GenerateStealthAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServer).GenerateStealthAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusk.Keys/GenerateStealthAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServer).GenerateStealthAddress(ctx, req.(*PublicKey))
	}
	return interceptor(ctx, in, info, handler)
}

var _Keys_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rusk.Keys",
	HandlerType: (*KeysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKeys",
			Handler:    _Keys_GenerateKeys_Handler,
		},
		{
			MethodName: "GenerateStealthAddress",
			Handler:    _Keys_GenerateStealthAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keys.proto",
}
