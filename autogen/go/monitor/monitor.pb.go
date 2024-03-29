// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor.proto

package monitor

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

type Level int32

const (
	Level_WARN  Level = 0
	Level_ERROR Level = 1
	Level_FATAL Level = 2
	Level_PANIC Level = 3
)

var Level_name = map[int32]string{
	0: "WARN",
	1: "ERROR",
	2: "FATAL",
	3: "PANIC",
}

var Level_value = map[string]int32{
	"WARN":  0,
	"ERROR": 1,
	"FATAL": 2,
	"PANIC": 3,
}

func (x Level) String() string {
	return proto.EnumName(Level_name, int32(x))
}

func (Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{0}
}

type SemverRequest struct {
	Major                uint32   `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor                uint32   `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch                uint32   `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SemverRequest) Reset()         { *m = SemverRequest{} }
func (m *SemverRequest) String() string { return proto.CompactTextString(m) }
func (*SemverRequest) ProtoMessage()    {}
func (*SemverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{0}
}

func (m *SemverRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SemverRequest.Unmarshal(m, b)
}
func (m *SemverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SemverRequest.Marshal(b, m, deterministic)
}
func (m *SemverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SemverRequest.Merge(m, src)
}
func (m *SemverRequest) XXX_Size() int {
	return xxx_messageInfo_SemverRequest.Size(m)
}
func (m *SemverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SemverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SemverRequest proto.InternalMessageInfo

func (m *SemverRequest) GetMajor() uint32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *SemverRequest) GetMinor() uint32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *SemverRequest) GetPatch() uint32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

type SlowdownAlert struct {
	TimeSinceLastBlockSec uint32   `protobuf:"varint,1,opt,name=timeSinceLastBlockSec,proto3" json:"timeSinceLastBlockSec,omitempty"`
	LastKnownHeight       uint64   `protobuf:"varint,2,opt,name=lastKnownHeight,proto3" json:"lastKnownHeight,omitempty"`
	LastKnownHash         []byte   `protobuf:"bytes,3,opt,name=lastKnownHash,proto3" json:"lastKnownHash,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *SlowdownAlert) Reset()         { *m = SlowdownAlert{} }
func (m *SlowdownAlert) String() string { return proto.CompactTextString(m) }
func (*SlowdownAlert) ProtoMessage()    {}
func (*SlowdownAlert) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{1}
}

func (m *SlowdownAlert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlowdownAlert.Unmarshal(m, b)
}
func (m *SlowdownAlert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlowdownAlert.Marshal(b, m, deterministic)
}
func (m *SlowdownAlert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlowdownAlert.Merge(m, src)
}
func (m *SlowdownAlert) XXX_Size() int {
	return xxx_messageInfo_SlowdownAlert.Size(m)
}
func (m *SlowdownAlert) XXX_DiscardUnknown() {
	xxx_messageInfo_SlowdownAlert.DiscardUnknown(m)
}

var xxx_messageInfo_SlowdownAlert proto.InternalMessageInfo

func (m *SlowdownAlert) GetTimeSinceLastBlockSec() uint32 {
	if m != nil {
		return m.TimeSinceLastBlockSec
	}
	return 0
}

func (m *SlowdownAlert) GetLastKnownHeight() uint64 {
	if m != nil {
		return m.LastKnownHeight
	}
	return 0
}

func (m *SlowdownAlert) GetLastKnownHash() []byte {
	if m != nil {
		return m.LastKnownHash
	}
	return nil
}

type Field struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{2}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *Field) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ErrorAlert struct {
	Level                Level    `protobuf:"varint,1,opt,name=level,proto3,enum=monitor.Level" json:"level,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	TimestampMillis      string   `protobuf:"bytes,3,opt,name=timestampMillis,proto3" json:"timestampMillis,omitempty"`
	File                 string   `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
	Line                 uint32   `protobuf:"varint,5,opt,name=line,proto3" json:"line,omitempty"`
	Function             string   `protobuf:"bytes,6,opt,name=function,proto3" json:"function,omitempty"`
	Fields               []*Field `protobuf:"bytes,7,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorAlert) Reset()         { *m = ErrorAlert{} }
func (m *ErrorAlert) String() string { return proto.CompactTextString(m) }
func (*ErrorAlert) ProtoMessage()    {}
func (*ErrorAlert) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{3}
}

func (m *ErrorAlert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorAlert.Unmarshal(m, b)
}
func (m *ErrorAlert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorAlert.Marshal(b, m, deterministic)
}
func (m *ErrorAlert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorAlert.Merge(m, src)
}
func (m *ErrorAlert) XXX_Size() int {
	return xxx_messageInfo_ErrorAlert.Size(m)
}
func (m *ErrorAlert) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorAlert.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorAlert proto.InternalMessageInfo

func (m *ErrorAlert) GetLevel() Level {
	if m != nil {
		return m.Level
	}
	return Level_WARN
}

func (m *ErrorAlert) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ErrorAlert) GetTimestampMillis() string {
	if m != nil {
		return m.TimestampMillis
	}
	return ""
}

func (m *ErrorAlert) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *ErrorAlert) GetLine() uint32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *ErrorAlert) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *ErrorAlert) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{4}
}

func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (m *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(m, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{5}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type BlockUpdate struct {
	Height               uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TxAmount             uint32   `protobuf:"varint,4,opt,name=txAmount,proto3" json:"txAmount,omitempty"`
	BlockTimeSec         uint32   `protobuf:"varint,5,opt,name=blockTimeSec,proto3" json:"blockTimeSec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockUpdate) Reset()         { *m = BlockUpdate{} }
func (m *BlockUpdate) String() string { return proto.CompactTextString(m) }
func (*BlockUpdate) ProtoMessage()    {}
func (*BlockUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{6}
}

func (m *BlockUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockUpdate.Unmarshal(m, b)
}
func (m *BlockUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockUpdate.Marshal(b, m, deterministic)
}
func (m *BlockUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockUpdate.Merge(m, src)
}
func (m *BlockUpdate) XXX_Size() int {
	return xxx_messageInfo_BlockUpdate.Size(m)
}
func (m *BlockUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_BlockUpdate proto.InternalMessageInfo

func (m *BlockUpdate) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockUpdate) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockUpdate) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockUpdate) GetTxAmount() uint32 {
	if m != nil {
		return m.TxAmount
	}
	return 0
}

func (m *BlockUpdate) GetBlockTimeSec() uint32 {
	if m != nil {
		return m.BlockTimeSec
	}
	return 0
}

func init() {
	proto.RegisterEnum("monitor.Level", Level_name, Level_value)
	proto.RegisterType((*SemverRequest)(nil), "monitor.SemverRequest")
	proto.RegisterType((*SlowdownAlert)(nil), "monitor.SlowdownAlert")
	proto.RegisterType((*Field)(nil), "monitor.Field")
	proto.RegisterType((*ErrorAlert)(nil), "monitor.ErrorAlert")
	proto.RegisterType((*EmptyRequest)(nil), "monitor.EmptyRequest")
	proto.RegisterType((*EmptyResponse)(nil), "monitor.EmptyResponse")
	proto.RegisterType((*BlockUpdate)(nil), "monitor.BlockUpdate")
}

func init() { proto.RegisterFile("monitor.proto", fileDescriptor_44174b7b2a306b71) }

var fileDescriptor_44174b7b2a306b71 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0xe3, 0x38, 0x6d, 0xa6, 0x75, 0x1a, 0x2d, 0x6d, 0x65, 0x55, 0x3c, 0x54, 0x56, 0x85,
	0x2a, 0x50, 0x1b, 0xa9, 0x45, 0x48, 0x48, 0xf0, 0xe0, 0xa0, 0x56, 0x45, 0xb4, 0x05, 0x36, 0x45,
	0x48, 0xbc, 0x39, 0xce, 0x26, 0x59, 0xba, 0xde, 0x35, 0xf6, 0xba, 0x21, 0x1f, 0x82, 0xf8, 0x1a,
	0x7e, 0x84, 0xaf, 0x41, 0x3b, 0x76, 0x9c, 0x8b, 0xa0, 0xbc, 0xcd, 0x39, 0xb3, 0xe3, 0x99, 0x73,
	0x66, 0xd7, 0xe0, 0xc6, 0x4a, 0x72, 0xad, 0xd2, 0x93, 0x24, 0x55, 0x5a, 0x91, 0xf5, 0x12, 0xfa,
	0x1f, 0xc1, 0xed, 0xb1, 0xf8, 0x9e, 0xa5, 0x94, 0x7d, 0xcb, 0x59, 0xa6, 0xc9, 0x0e, 0x38, 0x71,
	0xf8, 0x55, 0xa5, 0x9e, 0x75, 0x60, 0x1d, 0xb9, 0xb4, 0x00, 0xc8, 0x72, 0xa9, 0x52, 0xaf, 0x56,
	0xb2, 0x06, 0x18, 0x36, 0x09, 0x75, 0x34, 0xf6, 0xec, 0x82, 0x45, 0xe0, 0xff, 0xb0, 0xc0, 0xed,
	0x09, 0x35, 0x19, 0xa8, 0x89, 0x0c, 0x04, 0x4b, 0x35, 0x79, 0x0e, 0xbb, 0x9a, 0xc7, 0xac, 0xc7,
	0x65, 0xc4, 0xae, 0xc2, 0x4c, 0x77, 0x85, 0x8a, 0xee, 0x7a, 0x2c, 0x2a, 0x7b, 0xfc, 0x3d, 0x49,
	0x8e, 0x60, 0x5b, 0x84, 0x99, 0x7e, 0x27, 0xd5, 0x44, 0x5e, 0x32, 0x3e, 0x1a, 0x6b, 0xec, 0x5e,
	0xa7, 0xab, 0x34, 0x39, 0x04, 0x77, 0x4e, 0x85, 0x59, 0x31, 0xcf, 0x16, 0x5d, 0x26, 0xfd, 0x33,
	0x70, 0x2e, 0x38, 0x13, 0x03, 0x33, 0xf6, 0xd0, 0x04, 0xd8, 0xbe, 0x49, 0x0b, 0x60, 0xd8, 0xfb,
	0x50, 0xe4, 0x0c, 0x9b, 0x34, 0x69, 0x01, 0xfc, 0xdf, 0x16, 0xc0, 0x79, 0x9a, 0xaa, 0xb4, 0x50,
	0x72, 0x08, 0x8e, 0x60, 0xf7, 0x4c, 0x60, 0x69, 0xeb, 0xb4, 0x75, 0x32, 0xb3, 0xf5, 0xca, 0xb0,
	0xb4, 0x48, 0x92, 0x36, 0xd8, 0x71, 0x36, 0x2a, 0x3f, 0x64, 0x42, 0xa3, 0xc5, 0x88, 0xcc, 0x74,
	0x18, 0x27, 0xd7, 0x5c, 0x08, 0x9e, 0xe1, 0x8c, 0x4d, 0xba, 0x4a, 0x13, 0x02, 0xf5, 0x21, 0x17,
	0xcc, 0xab, 0x63, 0x1a, 0x63, 0xc3, 0x09, 0x2e, 0x99, 0xe7, 0xa0, 0x5d, 0x18, 0x93, 0x7d, 0xd8,
	0x18, 0xe6, 0x32, 0xd2, 0x5c, 0x49, 0xaf, 0x81, 0x67, 0x2b, 0x4c, 0x9e, 0x40, 0x03, 0x35, 0x65,
	0xde, 0xfa, 0x81, 0x7d, 0xb4, 0xb9, 0x30, 0x26, 0x1a, 0x40, 0xcb, 0xac, 0xdf, 0x82, 0xad, 0xf3,
	0x38, 0xd1, 0xd3, 0x72, 0xf7, 0xfe, 0x36, 0xb8, 0x25, 0xce, 0x12, 0x25, 0x33, 0xe6, 0xff, 0xb4,
	0x60, 0x13, 0xf7, 0xf1, 0x29, 0x19, 0x84, 0x9a, 0x91, 0x3d, 0x68, 0x8c, 0x8b, 0x4d, 0x58, 0xb8,
	0x89, 0x12, 0x99, 0x01, 0xc7, 0xc6, 0xf7, 0x1a, 0xfa, 0x8e, 0x31, 0x79, 0x0c, 0xcd, 0x4a, 0x1b,
	0x8a, 0xb5, 0xe9, 0x9c, 0x30, 0xe3, 0xeb, 0xef, 0x41, 0xac, 0x72, 0xa9, 0x51, 0xaa, 0x4b, 0x2b,
	0x4c, 0x7c, 0xd8, 0xea, 0x9b, 0xa6, 0xb7, 0xe6, 0x5a, 0xb0, 0xa8, 0x94, 0xbd, 0xc4, 0x3d, 0x3d,
	0x05, 0x07, 0x2d, 0x27, 0x1b, 0x50, 0xff, 0x1c, 0xd0, 0x9b, 0xf6, 0x1a, 0x69, 0x82, 0x73, 0x4e,
	0xe9, 0x7b, 0xda, 0xb6, 0x4c, 0x78, 0x11, 0xdc, 0x06, 0x57, 0xed, 0x9a, 0x09, 0x3f, 0x04, 0x37,
	0x6f, 0xdf, 0xb4, 0xed, 0xd3, 0x5f, 0x35, 0x58, 0xbf, 0x2e, 0x8c, 0x20, 0x2f, 0xc1, 0xb9, 0x64,
	0x42, 0x28, 0xb2, 0x57, 0x79, 0xb3, 0xf4, 0x0e, 0xf6, 0xe7, 0xfc, 0xb2, 0x25, 0x6b, 0xe4, 0x05,
	0xd8, 0xdd, 0x29, 0x23, 0xbb, 0xab, 0x07, 0xfe, 0x57, 0xf7, 0x1a, 0x36, 0x6f, 0x94, 0xe6, 0xc3,
	0x29, 0x3a, 0x4a, 0x76, 0xaa, 0x83, 0x0b, 0x0e, 0x3f, 0x50, 0xde, 0x85, 0x56, 0x51, 0x3e, 0x7b,
	0x5b, 0x8b, 0xa3, 0x2f, 0x3e, 0xb7, 0x07, 0xbe, 0xf1, 0x6a, 0x36, 0x02, 0x5e, 0x69, 0xf2, 0x68,
	0x7e, 0xb0, 0xba, 0xe2, 0xff, 0xae, 0xee, 0x1e, 0x7f, 0x79, 0x36, 0xe2, 0x7a, 0x9c, 0xf7, 0x4f,
	0x22, 0x15, 0x77, 0x06, 0x79, 0x76, 0x77, 0x8c, 0x7f, 0x93, 0x7e, 0x3e, 0xec, 0x84, 0xb9, 0x56,
	0x23, 0x26, 0x3b, 0x23, 0xd5, 0x29, 0xcb, 0xfb, 0x0d, 0x4c, 0x9e, 0xfd, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0xe1, 0x95, 0x0d, 0x7b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MonitorClient is the client API for Monitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitorClient interface {
	Hello(ctx context.Context, in *SemverRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	Bye(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	NotifyBlock(ctx context.Context, in *BlockUpdate, opts ...grpc.CallOption) (*EmptyResponse, error)
	NotifySlowdown(ctx context.Context, in *SlowdownAlert, opts ...grpc.CallOption) (*EmptyResponse, error)
	NotifyError(ctx context.Context, in *ErrorAlert, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type monitorClient struct {
	cc *grpc.ClientConn
}

func NewMonitorClient(cc *grpc.ClientConn) MonitorClient {
	return &monitorClient{cc}
}

func (c *monitorClient) Hello(ctx context.Context, in *SemverRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/monitor.Monitor/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorClient) Bye(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/monitor.Monitor/Bye", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorClient) NotifyBlock(ctx context.Context, in *BlockUpdate, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/monitor.Monitor/NotifyBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorClient) NotifySlowdown(ctx context.Context, in *SlowdownAlert, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/monitor.Monitor/NotifySlowdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorClient) NotifyError(ctx context.Context, in *ErrorAlert, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/monitor.Monitor/NotifyError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorServer is the server API for Monitor service.
type MonitorServer interface {
	Hello(context.Context, *SemverRequest) (*EmptyResponse, error)
	Bye(context.Context, *EmptyRequest) (*EmptyResponse, error)
	NotifyBlock(context.Context, *BlockUpdate) (*EmptyResponse, error)
	NotifySlowdown(context.Context, *SlowdownAlert) (*EmptyResponse, error)
	NotifyError(context.Context, *ErrorAlert) (*EmptyResponse, error)
}

// UnimplementedMonitorServer can be embedded to have forward compatible implementations.
type UnimplementedMonitorServer struct {
}

func (*UnimplementedMonitorServer) Hello(ctx context.Context, req *SemverRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedMonitorServer) Bye(ctx context.Context, req *EmptyRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bye not implemented")
}
func (*UnimplementedMonitorServer) NotifyBlock(ctx context.Context, req *BlockUpdate) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyBlock not implemented")
}
func (*UnimplementedMonitorServer) NotifySlowdown(ctx context.Context, req *SlowdownAlert) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifySlowdown not implemented")
}
func (*UnimplementedMonitorServer) NotifyError(ctx context.Context, req *ErrorAlert) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyError not implemented")
}

func RegisterMonitorServer(s *grpc.Server, srv MonitorServer) {
	s.RegisterService(&_Monitor_serviceDesc, srv)
}

func _Monitor_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SemverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.Monitor/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).Hello(ctx, req.(*SemverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitor_Bye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).Bye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.Monitor/Bye",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).Bye(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitor_NotifyBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).NotifyBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.Monitor/NotifyBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).NotifyBlock(ctx, req.(*BlockUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitor_NotifySlowdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlowdownAlert)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).NotifySlowdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.Monitor/NotifySlowdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).NotifySlowdown(ctx, req.(*SlowdownAlert))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitor_NotifyError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ErrorAlert)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).NotifyError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.Monitor/NotifyError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).NotifyError(ctx, req.(*ErrorAlert))
	}
	return interceptor(ctx, in, info, handler)
}

var _Monitor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monitor.Monitor",
	HandlerType: (*MonitorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Monitor_Hello_Handler,
		},
		{
			MethodName: "Bye",
			Handler:    _Monitor_Bye_Handler,
		},
		{
			MethodName: "NotifyBlock",
			Handler:    _Monitor_NotifyBlock_Handler,
		},
		{
			MethodName: "NotifySlowdown",
			Handler:    _Monitor_NotifySlowdown_Handler,
		},
		{
			MethodName: "NotifyError",
			Handler:    _Monitor_NotifyError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitor.proto",
}
