// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keys.proto

package rusk

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

type SecretKey struct {
	A                    []byte   `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    []byte   `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecretKey) Reset()         { *m = SecretKey{} }
func (m *SecretKey) String() string { return proto.CompactTextString(m) }
func (*SecretKey) ProtoMessage()    {}
func (*SecretKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9084e97af2346a26, []int{0}
}

func (m *SecretKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretKey.Unmarshal(m, b)
}
func (m *SecretKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretKey.Marshal(b, m, deterministic)
}
func (m *SecretKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretKey.Merge(m, src)
}
func (m *SecretKey) XXX_Size() int {
	return xxx_messageInfo_SecretKey.Size(m)
}
func (m *SecretKey) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretKey.DiscardUnknown(m)
}

var xxx_messageInfo_SecretKey proto.InternalMessageInfo

func (m *SecretKey) GetA() []byte {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *SecretKey) GetB() []byte {
	if m != nil {
		return m.B
	}
	return nil
}

type ViewKey struct {
	A                    []byte   `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	BG                   []byte   `protobuf:"bytes,2,opt,name=b_g,json=bG,proto3" json:"b_g,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewKey) Reset()         { *m = ViewKey{} }
func (m *ViewKey) String() string { return proto.CompactTextString(m) }
func (*ViewKey) ProtoMessage()    {}
func (*ViewKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9084e97af2346a26, []int{1}
}

func (m *ViewKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewKey.Unmarshal(m, b)
}
func (m *ViewKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewKey.Marshal(b, m, deterministic)
}
func (m *ViewKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewKey.Merge(m, src)
}
func (m *ViewKey) XXX_Size() int {
	return xxx_messageInfo_ViewKey.Size(m)
}
func (m *ViewKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewKey.DiscardUnknown(m)
}

var xxx_messageInfo_ViewKey proto.InternalMessageInfo

func (m *ViewKey) GetA() []byte {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *ViewKey) GetBG() []byte {
	if m != nil {
		return m.BG
	}
	return nil
}

type PublicKey struct {
	AG                   []byte   `protobuf:"bytes,1,opt,name=a_g,json=aG,proto3" json:"a_g,omitempty"`
	BG                   []byte   `protobuf:"bytes,2,opt,name=b_g,json=bG,proto3" json:"b_g,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}
func (*PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9084e97af2346a26, []int{2}
}

func (m *PublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKey.Unmarshal(m, b)
}
func (m *PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKey.Marshal(b, m, deterministic)
}
func (m *PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKey.Merge(m, src)
}
func (m *PublicKey) XXX_Size() int {
	return xxx_messageInfo_PublicKey.Size(m)
}
func (m *PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKey proto.InternalMessageInfo

func (m *PublicKey) GetAG() []byte {
	if m != nil {
		return m.AG
	}
	return nil
}

func (m *PublicKey) GetBG() []byte {
	if m != nil {
		return m.BG
	}
	return nil
}

type StealthAddress struct {
	RG                   []byte   `protobuf:"bytes,1,opt,name=r_g,json=rG,proto3" json:"r_g,omitempty"`
	PkR                  []byte   `protobuf:"bytes,2,opt,name=pk_r,json=pkR,proto3" json:"pk_r,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StealthAddress) Reset()         { *m = StealthAddress{} }
func (m *StealthAddress) String() string { return proto.CompactTextString(m) }
func (*StealthAddress) ProtoMessage()    {}
func (*StealthAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_9084e97af2346a26, []int{3}
}

func (m *StealthAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StealthAddress.Unmarshal(m, b)
}
func (m *StealthAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StealthAddress.Marshal(b, m, deterministic)
}
func (m *StealthAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StealthAddress.Merge(m, src)
}
func (m *StealthAddress) XXX_Size() int {
	return xxx_messageInfo_StealthAddress.Size(m)
}
func (m *StealthAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_StealthAddress.DiscardUnknown(m)
}

var xxx_messageInfo_StealthAddress proto.InternalMessageInfo

func (m *StealthAddress) GetRG() []byte {
	if m != nil {
		return m.RG
	}
	return nil
}

func (m *StealthAddress) GetPkR() []byte {
	if m != nil {
		return m.PkR
	}
	return nil
}

type GenerateKeysRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateKeysRequest) Reset()         { *m = GenerateKeysRequest{} }
func (m *GenerateKeysRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateKeysRequest) ProtoMessage()    {}
func (*GenerateKeysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9084e97af2346a26, []int{4}
}

func (m *GenerateKeysRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeysRequest.Unmarshal(m, b)
}
func (m *GenerateKeysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeysRequest.Marshal(b, m, deterministic)
}
func (m *GenerateKeysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeysRequest.Merge(m, src)
}
func (m *GenerateKeysRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateKeysRequest.Size(m)
}
func (m *GenerateKeysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeysRequest proto.InternalMessageInfo

type GenerateKeysResponse struct {
	Sk                   *SecretKey `protobuf:"bytes,1,opt,name=sk,proto3" json:"sk,omitempty"`
	Vk                   *ViewKey   `protobuf:"bytes,2,opt,name=vk,proto3" json:"vk,omitempty"`
	Pk                   *PublicKey `protobuf:"bytes,3,opt,name=pk,proto3" json:"pk,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GenerateKeysResponse) Reset()         { *m = GenerateKeysResponse{} }
func (m *GenerateKeysResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateKeysResponse) ProtoMessage()    {}
func (*GenerateKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9084e97af2346a26, []int{5}
}

func (m *GenerateKeysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeysResponse.Unmarshal(m, b)
}
func (m *GenerateKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeysResponse.Marshal(b, m, deterministic)
}
func (m *GenerateKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeysResponse.Merge(m, src)
}
func (m *GenerateKeysResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateKeysResponse.Size(m)
}
func (m *GenerateKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeysResponse proto.InternalMessageInfo

func (m *GenerateKeysResponse) GetSk() *SecretKey {
	if m != nil {
		return m.Sk
	}
	return nil
}

func (m *GenerateKeysResponse) GetVk() *ViewKey {
	if m != nil {
		return m.Vk
	}
	return nil
}

func (m *GenerateKeysResponse) GetPk() *PublicKey {
	if m != nil {
		return m.Pk
	}
	return nil
}

func init() {
	proto.RegisterType((*SecretKey)(nil), "rusk.SecretKey")
	proto.RegisterType((*ViewKey)(nil), "rusk.ViewKey")
	proto.RegisterType((*PublicKey)(nil), "rusk.PublicKey")
	proto.RegisterType((*StealthAddress)(nil), "rusk.StealthAddress")
	proto.RegisterType((*GenerateKeysRequest)(nil), "rusk.GenerateKeysRequest")
	proto.RegisterType((*GenerateKeysResponse)(nil), "rusk.GenerateKeysResponse")
}

func init() { proto.RegisterFile("keys.proto", fileDescriptor_9084e97af2346a26) }

var fileDescriptor_9084e97af2346a26 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x4f, 0x02, 0x31,
	0x10, 0xc5, 0xd9, 0x2e, 0xd1, 0x30, 0xa0, 0xc4, 0x8a, 0x06, 0x49, 0x8c, 0xa6, 0x17, 0x89, 0x09,
	0x4b, 0x82, 0xde, 0x3c, 0xe1, 0x65, 0x0f, 0x5c, 0xcc, 0x92, 0x78, 0xf0, 0x42, 0xda, 0x65, 0x02,
	0xa4, 0xc0, 0xd6, 0xb6, 0x0b, 0xe1, 0x6b, 0xf8, 0x89, 0x4d, 0x77, 0x57, 0xe4, 0xdf, 0xad, 0xd3,
	0xf9, 0xbd, 0xce, 0x7b, 0xcd, 0x00, 0x48, 0xdc, 0x98, 0x40, 0xe9, 0xc4, 0x26, 0xb4, 0xac, 0x53,
	0x23, 0xd9, 0x13, 0x54, 0x86, 0x18, 0x6b, 0xb4, 0x03, 0xdc, 0xd0, 0x1a, 0x78, 0xbc, 0xe9, 0x3d,
	0x7a, 0xed, 0x5a, 0xe4, 0x71, 0x57, 0x89, 0x26, 0xc9, 0x2b, 0xc1, 0xda, 0x70, 0xfe, 0x39, 0xc3,
	0xf5, 0x31, 0x56, 0x07, 0x5f, 0x8c, 0x26, 0x05, 0x48, 0x44, 0xc8, 0x3a, 0x50, 0xf9, 0x48, 0xc5,
	0x7c, 0x16, 0x3b, 0xb6, 0x0e, 0x3e, 0x1f, 0x4d, 0x0a, 0x9a, 0xf0, 0xf0, 0x18, 0x7f, 0x85, 0xcb,
	0xa1, 0x45, 0x3e, 0xb7, 0xd3, 0xfe, 0x78, 0xac, 0xd1, 0x18, 0x87, 0xe8, 0x7f, 0x8d, 0x0e, 0xe9,
	0x15, 0x94, 0x95, 0x1c, 0xe9, 0x42, 0xe4, 0x2b, 0x19, 0xb1, 0x1b, 0xb8, 0x0e, 0x71, 0x89, 0x9a,
	0x5b, 0x1c, 0xe0, 0xc6, 0x44, 0xf8, 0x9d, 0xa2, 0xb1, 0x6c, 0x0d, 0x8d, 0xfd, 0x6b, 0xa3, 0x92,
	0xa5, 0x41, 0xfa, 0x00, 0xc4, 0xc8, 0xec, 0xc5, 0x6a, 0xaf, 0x1e, 0xb8, 0xe4, 0xc1, 0x36, 0x76,
	0x44, 0x8c, 0xa4, 0xf7, 0x40, 0x56, 0x32, 0x1b, 0x50, 0xed, 0x5d, 0xe4, 0x40, 0x11, 0x37, 0x22,
	0x2b, 0xe9, 0xf4, 0x4a, 0x36, 0xfd, 0x5d, 0xfd, 0x36, 0x63, 0x44, 0x94, 0xec, 0xfd, 0x78, 0x50,
	0x76, 0x13, 0x69, 0x08, 0xb5, 0x5d, 0x07, 0xf4, 0x2e, 0xa7, 0x4f, 0x98, 0x6d, 0xb5, 0x4e, 0xb5,
	0x72, 0xc3, 0xac, 0x44, 0xfb, 0x70, 0xfb, 0xd7, 0x39, 0xfc, 0x9f, 0x03, 0x03, 0xad, 0x46, 0x91,
	0x68, 0x0f, 0x63, 0xa5, 0xf7, 0xe7, 0xaf, 0xf6, 0x64, 0x66, 0xa7, 0xa9, 0x08, 0xe2, 0x64, 0xd1,
	0x1d, 0xa7, 0x46, 0x76, 0x96, 0x68, 0xd7, 0x89, 0x96, 0x5d, 0x27, 0xe8, 0x98, 0x78, 0x8a, 0x0b,
	0xfe, 0xe6, 0xce, 0xe2, 0x2c, 0xdb, 0x8a, 0x97, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0x2a,
	0xab, 0x05, 0x23, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeysClient is the client API for Keys service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeysClient interface {
	GenerateKeys(ctx context.Context, in *GenerateKeysRequest, opts ...grpc.CallOption) (*GenerateKeysResponse, error)
	GenerateStealthAddress(ctx context.Context, in *PublicKey, opts ...grpc.CallOption) (*StealthAddress, error)
}

type keysClient struct {
	cc *grpc.ClientConn
}

func NewKeysClient(cc *grpc.ClientConn) KeysClient {
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

func (*UnimplementedKeysServer) GenerateKeys(ctx context.Context, req *GenerateKeysRequest) (*GenerateKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKeys not implemented")
}
func (*UnimplementedKeysServer) GenerateStealthAddress(ctx context.Context, req *PublicKey) (*StealthAddress, error) {
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
