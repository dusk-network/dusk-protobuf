// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blindbid.proto

package rusk

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GenerateScoreRequest struct {
	K                    []byte   `protobuf:"bytes,1,opt,name=k,proto3" json:"k,omitempty"`
	Seed                 []byte   `protobuf:"bytes,2,opt,name=seed,proto3" json:"seed,omitempty"`
	Secret               []byte   `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	Round                uint32   `protobuf:"varint,4,opt,name=round,proto3" json:"round,omitempty"`
	Step                 uint32   `protobuf:"varint,5,opt,name=step,proto3" json:"step,omitempty"`
	IndexStoredBid       uint64   `protobuf:"varint,6,opt,name=index_stored_bid,json=indexStoredBid,proto3" json:"index_stored_bid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateScoreRequest) Reset()         { *m = GenerateScoreRequest{} }
func (m *GenerateScoreRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateScoreRequest) ProtoMessage()    {}
func (*GenerateScoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_490c3399c7b8c3f4, []int{0}
}
func (m *GenerateScoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateScoreRequest.Unmarshal(m, b)
}
func (m *GenerateScoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateScoreRequest.Marshal(b, m, deterministic)
}
func (m *GenerateScoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateScoreRequest.Merge(m, src)
}
func (m *GenerateScoreRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateScoreRequest.Size(m)
}
func (m *GenerateScoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateScoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateScoreRequest proto.InternalMessageInfo

func (m *GenerateScoreRequest) GetK() []byte {
	if m != nil {
		return m.K
	}
	return nil
}

func (m *GenerateScoreRequest) GetSeed() []byte {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *GenerateScoreRequest) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *GenerateScoreRequest) GetRound() uint32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *GenerateScoreRequest) GetStep() uint32 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *GenerateScoreRequest) GetIndexStoredBid() uint64 {
	if m != nil {
		return m.IndexStoredBid
	}
	return 0
}

type GenerateScoreResponse struct {
	BlindbidProof        []byte   `protobuf:"bytes,1,opt,name=blindbid_proof,json=blindbidProof,proto3" json:"blindbid_proof,omitempty"`
	Score                []byte   `protobuf:"bytes,2,opt,name=score,proto3" json:"score,omitempty"`
	ProverIdentity       []byte   `protobuf:"bytes,4,opt,name=prover_identity,json=proverIdentity,proto3" json:"prover_identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateScoreResponse) Reset()         { *m = GenerateScoreResponse{} }
func (m *GenerateScoreResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateScoreResponse) ProtoMessage()    {}
func (*GenerateScoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_490c3399c7b8c3f4, []int{1}
}
func (m *GenerateScoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateScoreResponse.Unmarshal(m, b)
}
func (m *GenerateScoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateScoreResponse.Marshal(b, m, deterministic)
}
func (m *GenerateScoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateScoreResponse.Merge(m, src)
}
func (m *GenerateScoreResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateScoreResponse.Size(m)
}
func (m *GenerateScoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateScoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateScoreResponse proto.InternalMessageInfo

func (m *GenerateScoreResponse) GetBlindbidProof() []byte {
	if m != nil {
		return m.BlindbidProof
	}
	return nil
}

func (m *GenerateScoreResponse) GetScore() []byte {
	if m != nil {
		return m.Score
	}
	return nil
}

func (m *GenerateScoreResponse) GetProverIdentity() []byte {
	if m != nil {
		return m.ProverIdentity
	}
	return nil
}

type VerifyScoreRequest struct {
	Proof                []byte   `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
	Score                []byte   `protobuf:"bytes,2,opt,name=score,proto3" json:"score,omitempty"`
	Seed                 []byte   `protobuf:"bytes,3,opt,name=seed,proto3" json:"seed,omitempty"`
	ProverId             []byte   `protobuf:"bytes,4,opt,name=prover_id,json=proverId,proto3" json:"prover_id,omitempty"`
	Round                uint64   `protobuf:"fixed64,5,opt,name=round,proto3" json:"round,omitempty"`
	Step                 uint32   `protobuf:"varint,6,opt,name=step,proto3" json:"step,omitempty"`
	IndexStoredBid       uint64   `protobuf:"varint,7,opt,name=index_stored_bid,json=indexStoredBid,proto3" json:"index_stored_bid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyScoreRequest) Reset()         { *m = VerifyScoreRequest{} }
func (m *VerifyScoreRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyScoreRequest) ProtoMessage()    {}
func (*VerifyScoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_490c3399c7b8c3f4, []int{2}
}
func (m *VerifyScoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyScoreRequest.Unmarshal(m, b)
}
func (m *VerifyScoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyScoreRequest.Marshal(b, m, deterministic)
}
func (m *VerifyScoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyScoreRequest.Merge(m, src)
}
func (m *VerifyScoreRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyScoreRequest.Size(m)
}
func (m *VerifyScoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyScoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyScoreRequest proto.InternalMessageInfo

func (m *VerifyScoreRequest) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *VerifyScoreRequest) GetScore() []byte {
	if m != nil {
		return m.Score
	}
	return nil
}

func (m *VerifyScoreRequest) GetSeed() []byte {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *VerifyScoreRequest) GetProverId() []byte {
	if m != nil {
		return m.ProverId
	}
	return nil
}

func (m *VerifyScoreRequest) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *VerifyScoreRequest) GetStep() uint32 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *VerifyScoreRequest) GetIndexStoredBid() uint64 {
	if m != nil {
		return m.IndexStoredBid
	}
	return 0
}

type VerifyScoreResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyScoreResponse) Reset()         { *m = VerifyScoreResponse{} }
func (m *VerifyScoreResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyScoreResponse) ProtoMessage()    {}
func (*VerifyScoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_490c3399c7b8c3f4, []int{3}
}
func (m *VerifyScoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyScoreResponse.Unmarshal(m, b)
}
func (m *VerifyScoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyScoreResponse.Marshal(b, m, deterministic)
}
func (m *VerifyScoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyScoreResponse.Merge(m, src)
}
func (m *VerifyScoreResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyScoreResponse.Size(m)
}
func (m *VerifyScoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyScoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyScoreResponse proto.InternalMessageInfo

func (m *VerifyScoreResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*GenerateScoreRequest)(nil), "rusk.GenerateScoreRequest")
	proto.RegisterType((*GenerateScoreResponse)(nil), "rusk.GenerateScoreResponse")
	proto.RegisterType((*VerifyScoreRequest)(nil), "rusk.VerifyScoreRequest")
	proto.RegisterType((*VerifyScoreResponse)(nil), "rusk.VerifyScoreResponse")
}

func init() { proto.RegisterFile("blindbid.proto", fileDescriptor_490c3399c7b8c3f4) }

var fileDescriptor_490c3399c7b8c3f4 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0xae, 0xd3, 0x30,
	0x10, 0x85, 0x31, 0x37, 0xcd, 0xbd, 0x0c, 0x6d, 0x2f, 0x32, 0x05, 0x99, 0x76, 0x53, 0x45, 0x42,
	0x44, 0x48, 0x4d, 0x25, 0x58, 0xb2, 0x8b, 0x90, 0x10, 0xac, 0x50, 0x2a, 0xb1, 0x60, 0x13, 0x35,
	0xf1, 0x94, 0x5a, 0xa1, 0x71, 0xb0, 0x9d, 0x42, 0x17, 0x3c, 0x0d, 0xbc, 0x0a, 0xef, 0x85, 0xec,
	0x24, 0xf4, 0x87, 0x20, 0x76, 0x33, 0xc7, 0xb1, 0xcf, 0x99, 0x2f, 0x03, 0xe3, 0xec, 0xb3, 0x28,
	0x79, 0x26, 0x78, 0x54, 0x29, 0x69, 0x24, 0xf5, 0x54, 0xad, 0x8b, 0xe0, 0x27, 0x81, 0xc9, 0x1b,
	0x2c, 0x51, 0xad, 0x0d, 0xae, 0x72, 0xa9, 0x30, 0xc1, 0x2f, 0x35, 0x6a, 0x43, 0x87, 0x40, 0x0a,
	0x46, 0xe6, 0x24, 0x1c, 0x26, 0xa4, 0xa0, 0x14, 0x3c, 0x8d, 0xc8, 0xd9, 0x5d, 0x27, 0xb8, 0x9a,
	0x3e, 0x06, 0x5f, 0x63, 0xae, 0xd0, 0xb0, 0x2b, 0xa7, 0xb6, 0x1d, 0x9d, 0xc0, 0x40, 0xc9, 0xba,
	0xe4, 0xcc, 0x9b, 0x93, 0x70, 0x94, 0x34, 0x8d, 0x7b, 0xc1, 0x60, 0xc5, 0x06, 0x4e, 0x74, 0x35,
	0x0d, 0xe1, 0x81, 0x28, 0x39, 0x7e, 0x4b, 0xb5, 0x91, 0x0a, 0x79, 0x9a, 0x09, 0xce, 0xfc, 0x39,
	0x09, 0xbd, 0x64, 0xec, 0xf4, 0x95, 0x93, 0x63, 0xc1, 0x83, 0xef, 0xf0, 0xe8, 0x22, 0xa5, 0xae,
	0x64, 0xa9, 0x91, 0x3e, 0x3d, 0xce, 0x95, 0x56, 0x4a, 0xca, 0x4d, 0x9b, 0x79, 0xd4, 0xa9, 0xef,
	0xad, 0x68, 0x33, 0x69, 0x7b, 0xaf, 0x1d, 0xa0, 0x69, 0xe8, 0x33, 0xb8, 0xad, 0x94, 0xdc, 0xa3,
	0x4a, 0x05, 0xc7, 0xd2, 0x08, 0x73, 0x70, 0x99, 0x87, 0xc9, 0xb8, 0x91, 0xdf, 0xb6, 0x6a, 0xf0,
	0x8b, 0x00, 0xfd, 0x80, 0x4a, 0x6c, 0x0e, 0x67, 0x8c, 0x26, 0x30, 0x38, 0xf5, 0x6c, 0x9a, 0x7f,
	0x78, 0x75, 0x04, 0xaf, 0x4e, 0x08, 0xce, 0xe0, 0xde, 0x1f, 0xff, 0xd6, 0xf9, 0xa6, 0x73, 0x3e,
	0x62, 0xb4, 0xc4, 0xfc, 0x4b, 0x8c, 0xfe, 0x7f, 0x30, 0x5e, 0xf7, 0x62, 0x5c, 0xc2, 0xc3, 0xb3,
	0x31, 0x5a, 0x88, 0x0c, 0xae, 0x75, 0x9d, 0xe7, 0xa8, 0xb5, 0x9b, 0xe4, 0x26, 0xe9, 0xda, 0x17,
	0x3f, 0x08, 0xdc, 0xc6, 0x96, 0x64, 0x2c, 0xf8, 0x0a, 0xd5, 0x5e, 0xe4, 0x48, 0xdf, 0xc1, 0xe8,
	0xec, 0x5f, 0xd0, 0x69, 0x64, 0x57, 0x29, 0xea, 0x5b, 0xa3, 0xe9, 0xac, 0xf7, 0xac, 0xf1, 0x0d,
	0xee, 0xd0, 0xd7, 0x70, 0xff, 0x24, 0x10, 0x65, 0xcd, 0xd7, 0x7f, 0xa3, 0x9e, 0x3e, 0xe9, 0x39,
	0xe9, 0x5e, 0x89, 0x9f, 0x7f, 0x0c, 0x3f, 0x09, 0xb3, 0xad, 0xb3, 0x28, 0x97, 0xbb, 0x25, 0xaf,
	0x75, 0xb1, 0x28, 0xd1, 0x7c, 0x95, 0xaa, 0x58, 0xda, 0x5b, 0x0b, 0x9d, 0x6f, 0x71, 0xb7, 0x7e,
	0x65, 0xeb, 0xcc, 0x77, 0xdb, 0xff, 0xf2, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0xa7, 0x44,
	0xef, 0x0f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlindBidServiceClient is the client API for BlindBidService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlindBidServiceClient interface {
	// Generate a proof of blind bid, and a score.
	GenerateScore(ctx context.Context, in *GenerateScoreRequest, opts ...grpc.CallOption) (*GenerateScoreResponse, error)
	// Verify a proof of blind bid, and the associated score.
	VerifyScore(ctx context.Context, in *VerifyScoreRequest, opts ...grpc.CallOption) (*VerifyScoreResponse, error)
}

type blindBidServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlindBidServiceClient(cc *grpc.ClientConn) BlindBidServiceClient {
	return &blindBidServiceClient{cc}
}

func (c *blindBidServiceClient) GenerateScore(ctx context.Context, in *GenerateScoreRequest, opts ...grpc.CallOption) (*GenerateScoreResponse, error) {
	out := new(GenerateScoreResponse)
	err := c.cc.Invoke(ctx, "/rusk.BlindBidService/GenerateScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blindBidServiceClient) VerifyScore(ctx context.Context, in *VerifyScoreRequest, opts ...grpc.CallOption) (*VerifyScoreResponse, error) {
	out := new(VerifyScoreResponse)
	err := c.cc.Invoke(ctx, "/rusk.BlindBidService/VerifyScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlindBidServiceServer is the server API for BlindBidService service.
type BlindBidServiceServer interface {
	// Generate a proof of blind bid, and a score.
	GenerateScore(context.Context, *GenerateScoreRequest) (*GenerateScoreResponse, error)
	// Verify a proof of blind bid, and the associated score.
	VerifyScore(context.Context, *VerifyScoreRequest) (*VerifyScoreResponse, error)
}

// UnimplementedBlindBidServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlindBidServiceServer struct {
}

func (*UnimplementedBlindBidServiceServer) GenerateScore(ctx context.Context, req *GenerateScoreRequest) (*GenerateScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateScore not implemented")
}
func (*UnimplementedBlindBidServiceServer) VerifyScore(ctx context.Context, req *VerifyScoreRequest) (*VerifyScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyScore not implemented")
}

func RegisterBlindBidServiceServer(s *grpc.Server, srv BlindBidServiceServer) {
	s.RegisterService(&_BlindBidService_serviceDesc, srv)
}

func _BlindBidService_GenerateScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlindBidServiceServer).GenerateScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusk.BlindBidService/GenerateScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlindBidServiceServer).GenerateScore(ctx, req.(*GenerateScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlindBidService_VerifyScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlindBidServiceServer).VerifyScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusk.BlindBidService/VerifyScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlindBidServiceServer).VerifyScore(ctx, req.(*VerifyScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlindBidService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rusk.BlindBidService",
	HandlerType: (*BlindBidServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateScore",
			Handler:    _BlindBidService_GenerateScore_Handler,
		},
		{
			MethodName: "VerifyScore",
			Handler:    _BlindBidService_VerifyScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blindbid.proto",
}
