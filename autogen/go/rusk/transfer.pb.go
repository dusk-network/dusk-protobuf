// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transfer.proto

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

type TransferTransactionRequest struct {
	Value                uint64   `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	Recipient            []byte   `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	GasLimit             uint64   `protobuf:"fixed64,3,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasPrice             uint64   `protobuf:"fixed64,4,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferTransactionRequest) Reset()         { *m = TransferTransactionRequest{} }
func (m *TransferTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*TransferTransactionRequest) ProtoMessage()    {}
func (*TransferTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96c3e6bcafb460d3, []int{0}
}

func (m *TransferTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferTransactionRequest.Unmarshal(m, b)
}
func (m *TransferTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferTransactionRequest.Marshal(b, m, deterministic)
}
func (m *TransferTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferTransactionRequest.Merge(m, src)
}
func (m *TransferTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_TransferTransactionRequest.Size(m)
}
func (m *TransferTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferTransactionRequest proto.InternalMessageInfo

func (m *TransferTransactionRequest) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TransferTransactionRequest) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *TransferTransactionRequest) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *TransferTransactionRequest) GetGasPrice() uint64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func init() {
	proto.RegisterType((*TransferTransactionRequest)(nil), "rusk.TransferTransactionRequest")
}

func init() { proto.RegisterFile("transfer.proto", fileDescriptor_96c3e6bcafb460d3) }

var fileDescriptor_96c3e6bcafb460d3 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x29, 0x4a, 0xcc,
	0x2b, 0x4e, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2a, 0x2d, 0xce,
	0x96, 0x12, 0x04, 0x8b, 0x26, 0x26, 0x97, 0x64, 0xe6, 0xe7, 0x41, 0x24, 0x94, 0xba, 0x18, 0xb9,
	0xa4, 0x42, 0xa0, 0x6a, 0x43, 0x10, 0xb2, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x22,
	0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x6c, 0x41, 0x10, 0x8e,
	0x90, 0x0c, 0x17, 0x67, 0x51, 0x6a, 0x72, 0x66, 0x41, 0x66, 0x6a, 0x5e, 0x89, 0x04, 0x93, 0x02,
	0xa3, 0x06, 0x4f, 0x10, 0x42, 0x40, 0x48, 0x9a, 0x8b, 0x33, 0x3d, 0xb1, 0x38, 0x3e, 0x27, 0x33,
	0x37, 0xb3, 0x44, 0x82, 0x19, 0xac, 0x8f, 0x23, 0x3d, 0xb1, 0xd8, 0x07, 0xc4, 0x87, 0x49, 0x16,
	0x14, 0x65, 0x26, 0xa7, 0x4a, 0xb0, 0xc0, 0x25, 0x03, 0x40, 0x7c, 0xa3, 0x00, 0x2e, 0x0e, 0x98,
	0x5b, 0x84, 0x5c, 0xb8, 0xb8, 0xfd, 0x52, 0xcb, 0xe1, 0x5c, 0x05, 0x3d, 0x90, 0x0f, 0xf4, 0x70,
	0x3b, 0x55, 0x4a, 0x10, 0x49, 0x05, 0x44, 0x46, 0x89, 0xc1, 0x49, 0x2b, 0x4a, 0x23, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0xa5, 0xb4, 0x38, 0x5b, 0x37, 0x2f, 0xb5,
	0xa4, 0x3c, 0xbf, 0x28, 0x5b, 0x1f, 0xa4, 0x5a, 0xb7, 0x38, 0x39, 0x23, 0x35, 0x37, 0xd1, 0x1a,
	0xc4, 0x4e, 0x62, 0x03, 0x87, 0x88, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x39, 0x49, 0xec,
	0x3c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransferClient is the client API for Transfer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransferClient interface {
	// Generate a new Transfer transaction.
	NewTransfer(ctx context.Context, in *TransferTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
}

type transferClient struct {
	cc *grpc.ClientConn
}

func NewTransferClient(cc *grpc.ClientConn) TransferClient {
	return &transferClient{cc}
}

func (c *transferClient) NewTransfer(ctx context.Context, in *TransferTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/rusk.Transfer/NewTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransferServer is the server API for Transfer service.
type TransferServer interface {
	// Generate a new Transfer transaction.
	NewTransfer(context.Context, *TransferTransactionRequest) (*Transaction, error)
}

// UnimplementedTransferServer can be embedded to have forward compatible implementations.
type UnimplementedTransferServer struct {
}

func (*UnimplementedTransferServer) NewTransfer(ctx context.Context, req *TransferTransactionRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewTransfer not implemented")
}

func RegisterTransferServer(s *grpc.Server, srv TransferServer) {
	s.RegisterService(&_Transfer_serviceDesc, srv)
}

func _Transfer_NewTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServer).NewTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusk.Transfer/NewTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServer).NewTransfer(ctx, req.(*TransferTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transfer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rusk.Transfer",
	HandlerType: (*TransferServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewTransfer",
			Handler:    _Transfer_NewTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfer.proto",
}
