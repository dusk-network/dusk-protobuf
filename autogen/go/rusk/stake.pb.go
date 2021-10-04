// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stake.proto

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

type Stake struct {
	Amount               uint64   `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	StartHeight          uint64   `protobuf:"fixed64,2,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	EndHeight            uint64   `protobuf:"fixed64,3,opt,name=end_height,json=endHeight,proto3" json:"end_height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stake) Reset()         { *m = Stake{} }
func (m *Stake) String() string { return proto.CompactTextString(m) }
func (*Stake) ProtoMessage()    {}
func (*Stake) Descriptor() ([]byte, []int) {
	return fileDescriptor_988c95942a06b53d, []int{0}
}

func (m *Stake) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stake.Unmarshal(m, b)
}
func (m *Stake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stake.Marshal(b, m, deterministic)
}
func (m *Stake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stake.Merge(m, src)
}
func (m *Stake) XXX_Size() int {
	return xxx_messageInfo_Stake.Size(m)
}
func (m *Stake) XXX_DiscardUnknown() {
	xxx_messageInfo_Stake.DiscardUnknown(m)
}

var xxx_messageInfo_Stake proto.InternalMessageInfo

func (m *Stake) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Stake) GetStartHeight() uint64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *Stake) GetEndHeight() uint64 {
	if m != nil {
		return m.EndHeight
	}
	return 0
}

type StakeTransactionRequest struct {
	Value                uint64   `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	PublicKeyBls         []byte   `protobuf:"bytes,2,opt,name=public_key_bls,json=publicKeyBls,proto3" json:"public_key_bls,omitempty"`
	GasLimit             uint64   `protobuf:"fixed64,3,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasPrice             uint64   `protobuf:"fixed64,4,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StakeTransactionRequest) Reset()         { *m = StakeTransactionRequest{} }
func (m *StakeTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*StakeTransactionRequest) ProtoMessage()    {}
func (*StakeTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_988c95942a06b53d, []int{1}
}

func (m *StakeTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakeTransactionRequest.Unmarshal(m, b)
}
func (m *StakeTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakeTransactionRequest.Marshal(b, m, deterministic)
}
func (m *StakeTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeTransactionRequest.Merge(m, src)
}
func (m *StakeTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_StakeTransactionRequest.Size(m)
}
func (m *StakeTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StakeTransactionRequest proto.InternalMessageInfo

func (m *StakeTransactionRequest) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *StakeTransactionRequest) GetPublicKeyBls() []byte {
	if m != nil {
		return m.PublicKeyBls
	}
	return nil
}

func (m *StakeTransactionRequest) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *StakeTransactionRequest) GetGasPrice() uint64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

type FindStakeRequest struct {
	Pk                   []byte   `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindStakeRequest) Reset()         { *m = FindStakeRequest{} }
func (m *FindStakeRequest) String() string { return proto.CompactTextString(m) }
func (*FindStakeRequest) ProtoMessage()    {}
func (*FindStakeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_988c95942a06b53d, []int{2}
}

func (m *FindStakeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindStakeRequest.Unmarshal(m, b)
}
func (m *FindStakeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindStakeRequest.Marshal(b, m, deterministic)
}
func (m *FindStakeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindStakeRequest.Merge(m, src)
}
func (m *FindStakeRequest) XXX_Size() int {
	return xxx_messageInfo_FindStakeRequest.Size(m)
}
func (m *FindStakeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindStakeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindStakeRequest proto.InternalMessageInfo

func (m *FindStakeRequest) GetPk() []byte {
	if m != nil {
		return m.Pk
	}
	return nil
}

type FindStakeResponse struct {
	Stakes               []*Stake `protobuf:"bytes,1,rep,name=stakes,proto3" json:"stakes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindStakeResponse) Reset()         { *m = FindStakeResponse{} }
func (m *FindStakeResponse) String() string { return proto.CompactTextString(m) }
func (*FindStakeResponse) ProtoMessage()    {}
func (*FindStakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_988c95942a06b53d, []int{3}
}

func (m *FindStakeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindStakeResponse.Unmarshal(m, b)
}
func (m *FindStakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindStakeResponse.Marshal(b, m, deterministic)
}
func (m *FindStakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindStakeResponse.Merge(m, src)
}
func (m *FindStakeResponse) XXX_Size() int {
	return xxx_messageInfo_FindStakeResponse.Size(m)
}
func (m *FindStakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindStakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindStakeResponse proto.InternalMessageInfo

func (m *FindStakeResponse) GetStakes() []*Stake {
	if m != nil {
		return m.Stakes
	}
	return nil
}

func init() {
	proto.RegisterType((*Stake)(nil), "rusk.Stake")
	proto.RegisterType((*StakeTransactionRequest)(nil), "rusk.StakeTransactionRequest")
	proto.RegisterType((*FindStakeRequest)(nil), "rusk.FindStakeRequest")
	proto.RegisterType((*FindStakeResponse)(nil), "rusk.FindStakeResponse")
}

func init() { proto.RegisterFile("stake.proto", fileDescriptor_988c95942a06b53d) }

var fileDescriptor_988c95942a06b53d = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4f, 0x0b, 0xda, 0x40,
	0x14, 0xc4, 0x8d, 0x7f, 0x82, 0xbe, 0x04, 0xa9, 0x4b, 0xd1, 0x60, 0x11, 0x6c, 0xda, 0x83, 0x14,
	0x8c, 0x60, 0x2f, 0x85, 0x82, 0x07, 0x0f, 0xa5, 0xd0, 0x52, 0x4a, 0xec, 0xa9, 0x97, 0xb0, 0x49,
	0x1e, 0xc9, 0x92, 0x64, 0x93, 0x66, 0x37, 0x8a, 0x5f, 0xa2, 0xfd, 0xca, 0x25, 0xbb, 0xc6, 0x4a,
	0xbd, 0xe5, 0xcd, 0x6f, 0x96, 0x19, 0x86, 0x80, 0x25, 0x24, 0xcd, 0xd0, 0xab, 0xea, 0x52, 0x96,
	0x64, 0x58, 0x37, 0x22, 0x5b, 0xce, 0x64, 0x4d, 0xb9, 0xa0, 0x91, 0x64, 0x25, 0xd7, 0xc0, 0xa5,
	0x30, 0x3a, 0xb5, 0x3e, 0x32, 0x07, 0x93, 0x16, 0x65, 0xc3, 0xa5, 0x63, 0xac, 0x8d, 0x8d, 0xe9,
	0xdf, 0x2e, 0xf2, 0x1a, 0x6c, 0x21, 0x69, 0x2d, 0x83, 0x14, 0x59, 0x92, 0x4a, 0xa7, 0xaf, 0xa8,
	0xa5, 0xb4, 0xcf, 0x4a, 0x22, 0x2b, 0x00, 0xe4, 0x71, 0x67, 0x18, 0x28, 0xc3, 0x04, 0x79, 0xac,
	0xb1, 0xfb, 0xc7, 0x80, 0x85, 0xca, 0xf8, 0xf1, 0x2f, 0xdd, 0xc7, 0x5f, 0x0d, 0x0a, 0x49, 0x5e,
	0xc2, 0xe8, 0x4c, 0xf3, 0x06, 0x6f, 0xa1, 0xfa, 0x20, 0x6f, 0x61, 0x5a, 0x35, 0x61, 0xce, 0xa2,
	0x20, 0xc3, 0x6b, 0x10, 0xe6, 0x42, 0xa5, 0xda, 0xbe, 0xad, 0xd5, 0x2f, 0x78, 0x3d, 0xe6, 0x82,
	0xbc, 0x82, 0x49, 0x42, 0x45, 0x90, 0xb3, 0x82, 0x75, 0xa9, 0xe3, 0x84, 0x8a, 0xaf, 0xed, 0xdd,
	0xc1, 0xaa, 0x66, 0x11, 0x3a, 0xc3, 0x3b, 0xfc, 0xde, 0xde, 0xae, 0x0b, 0x2f, 0x3e, 0x31, 0x1e,
	0xab, 0x52, 0x5d, 0x93, 0x29, 0xf4, 0xab, 0x4c, 0xd5, 0xb0, 0xfd, 0x7e, 0x95, 0xb9, 0x1f, 0x60,
	0xf6, 0xe0, 0x11, 0x55, 0xc9, 0x05, 0x92, 0x37, 0x60, 0xaa, 0x55, 0x85, 0x63, 0xac, 0x07, 0x1b,
	0x6b, 0x6f, 0x79, 0xed, 0xae, 0x9e, 0x36, 0xdd, 0xd0, 0xfe, 0xb7, 0x01, 0xb6, 0x52, 0x4e, 0x58,
	0x9f, 0x59, 0x84, 0xe4, 0x00, 0xe3, 0x6f, 0x78, 0xd1, 0x33, 0xaf, 0x1e, 0x5e, 0x3c, 0xef, 0xb1,
	0x9c, 0x69, 0xfc, 0x40, 0xdc, 0x1e, 0x39, 0xc0, 0xe4, 0x5e, 0x85, 0xcc, 0xb5, 0xe3, 0xff, 0xfe,
	0xcb, 0xc5, 0x93, 0xae, 0x3b, 0xbb, 0xbd, 0xe3, 0xbb, 0x9f, 0x9b, 0x84, 0xc9, 0xb4, 0x09, 0xbd,
	0xa8, 0x2c, 0x76, 0x71, 0x23, 0xb2, 0x2d, 0x47, 0x79, 0x29, 0xeb, 0x6c, 0xd7, 0xbe, 0xd9, 0x8a,
	0x28, 0xc5, 0x82, 0x7e, 0x6c, 0xbf, 0x43, 0x53, 0xfd, 0x16, 0xef, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x79, 0xf0, 0xd9, 0xae, 0x3e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StakeServiceClient is the client API for StakeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StakeServiceClient interface {
	// Generate a new Stake transaction.
	NewStake(ctx context.Context, in *StakeTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
	// Find all stakes related to a provisioner public key.
	FindStake(ctx context.Context, in *FindStakeRequest, opts ...grpc.CallOption) (*FindStakeResponse, error)
}

type stakeServiceClient struct {
	cc *grpc.ClientConn
}

func NewStakeServiceClient(cc *grpc.ClientConn) StakeServiceClient {
	return &stakeServiceClient{cc}
}

func (c *stakeServiceClient) NewStake(ctx context.Context, in *StakeTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/rusk.StakeService/NewStake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakeServiceClient) FindStake(ctx context.Context, in *FindStakeRequest, opts ...grpc.CallOption) (*FindStakeResponse, error) {
	out := new(FindStakeResponse)
	err := c.cc.Invoke(ctx, "/rusk.StakeService/FindStake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StakeServiceServer is the server API for StakeService service.
type StakeServiceServer interface {
	// Generate a new Stake transaction.
	NewStake(context.Context, *StakeTransactionRequest) (*Transaction, error)
	// Find all stakes related to a provisioner public key.
	FindStake(context.Context, *FindStakeRequest) (*FindStakeResponse, error)
}

// UnimplementedStakeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStakeServiceServer struct {
}

func (*UnimplementedStakeServiceServer) NewStake(ctx context.Context, req *StakeTransactionRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewStake not implemented")
}
func (*UnimplementedStakeServiceServer) FindStake(ctx context.Context, req *FindStakeRequest) (*FindStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindStake not implemented")
}

func RegisterStakeServiceServer(s *grpc.Server, srv StakeServiceServer) {
	s.RegisterService(&_StakeService_serviceDesc, srv)
}

func _StakeService_NewStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StakeTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakeServiceServer).NewStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusk.StakeService/NewStake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakeServiceServer).NewStake(ctx, req.(*StakeTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StakeService_FindStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindStakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakeServiceServer).FindStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusk.StakeService/FindStake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakeServiceServer).FindStake(ctx, req.(*FindStakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StakeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rusk.StakeService",
	HandlerType: (*StakeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewStake",
			Handler:    _StakeService_NewStake_Handler,
		},
		{
			MethodName: "FindStake",
			Handler:    _StakeService_FindStake_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stake.proto",
}
