// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: basic_fields.proto

package rusk

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type BlsScalar struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlsScalar) Reset()         { *m = BlsScalar{} }
func (m *BlsScalar) String() string { return proto.CompactTextString(m) }
func (*BlsScalar) ProtoMessage()    {}
func (*BlsScalar) Descriptor() ([]byte, []int) {
	return fileDescriptor_d26fe2b44342a40e, []int{0}
}
func (m *BlsScalar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlsScalar.Unmarshal(m, b)
}
func (m *BlsScalar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlsScalar.Marshal(b, m, deterministic)
}
func (m *BlsScalar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlsScalar.Merge(m, src)
}
func (m *BlsScalar) XXX_Size() int {
	return xxx_messageInfo_BlsScalar.Size(m)
}
func (m *BlsScalar) XXX_DiscardUnknown() {
	xxx_messageInfo_BlsScalar.DiscardUnknown(m)
}

var xxx_messageInfo_BlsScalar proto.InternalMessageInfo

func (m *BlsScalar) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type JubJubScalar struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JubJubScalar) Reset()         { *m = JubJubScalar{} }
func (m *JubJubScalar) String() string { return proto.CompactTextString(m) }
func (*JubJubScalar) ProtoMessage()    {}
func (*JubJubScalar) Descriptor() ([]byte, []int) {
	return fileDescriptor_d26fe2b44342a40e, []int{1}
}
func (m *JubJubScalar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JubJubScalar.Unmarshal(m, b)
}
func (m *JubJubScalar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JubJubScalar.Marshal(b, m, deterministic)
}
func (m *JubJubScalar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JubJubScalar.Merge(m, src)
}
func (m *JubJubScalar) XXX_Size() int {
	return xxx_messageInfo_JubJubScalar.Size(m)
}
func (m *JubJubScalar) XXX_DiscardUnknown() {
	xxx_messageInfo_JubJubScalar.DiscardUnknown(m)
}

var xxx_messageInfo_JubJubScalar proto.InternalMessageInfo

func (m *JubJubScalar) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type JubJubCompressed struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JubJubCompressed) Reset()         { *m = JubJubCompressed{} }
func (m *JubJubCompressed) String() string { return proto.CompactTextString(m) }
func (*JubJubCompressed) ProtoMessage()    {}
func (*JubJubCompressed) Descriptor() ([]byte, []int) {
	return fileDescriptor_d26fe2b44342a40e, []int{2}
}
func (m *JubJubCompressed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JubJubCompressed.Unmarshal(m, b)
}
func (m *JubJubCompressed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JubJubCompressed.Marshal(b, m, deterministic)
}
func (m *JubJubCompressed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JubJubCompressed.Merge(m, src)
}
func (m *JubJubCompressed) XXX_Size() int {
	return xxx_messageInfo_JubJubCompressed.Size(m)
}
func (m *JubJubCompressed) XXX_DiscardUnknown() {
	xxx_messageInfo_JubJubCompressed.DiscardUnknown(m)
}

var xxx_messageInfo_JubJubCompressed proto.InternalMessageInfo

func (m *JubJubCompressed) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PoseidonCipher struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoseidonCipher) Reset()         { *m = PoseidonCipher{} }
func (m *PoseidonCipher) String() string { return proto.CompactTextString(m) }
func (*PoseidonCipher) ProtoMessage()    {}
func (*PoseidonCipher) Descriptor() ([]byte, []int) {
	return fileDescriptor_d26fe2b44342a40e, []int{3}
}
func (m *PoseidonCipher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoseidonCipher.Unmarshal(m, b)
}
func (m *PoseidonCipher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoseidonCipher.Marshal(b, m, deterministic)
}
func (m *PoseidonCipher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoseidonCipher.Merge(m, src)
}
func (m *PoseidonCipher) XXX_Size() int {
	return xxx_messageInfo_PoseidonCipher.Size(m)
}
func (m *PoseidonCipher) XXX_DiscardUnknown() {
	xxx_messageInfo_PoseidonCipher.DiscardUnknown(m)
}

var xxx_messageInfo_PoseidonCipher proto.InternalMessageInfo

func (m *PoseidonCipher) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Proof struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_d26fe2b44342a40e, []int{4}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proof.Unmarshal(m, b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return xxx_messageInfo_Proof.Size(m)
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func (m *Proof) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*BlsScalar)(nil), "rusk.BlsScalar")
	proto.RegisterType((*JubJubScalar)(nil), "rusk.JubJubScalar")
	proto.RegisterType((*JubJubCompressed)(nil), "rusk.JubJubCompressed")
	proto.RegisterType((*PoseidonCipher)(nil), "rusk.PoseidonCipher")
	proto.RegisterType((*Proof)(nil), "rusk.Proof")
}

func init() { proto.RegisterFile("basic_fields.proto", fileDescriptor_d26fe2b44342a40e) }

var fileDescriptor_d26fe2b44342a40e = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x4a, 0x2c, 0xce,
	0x4c, 0x8e, 0x4f, 0xcb, 0x4c, 0xcd, 0x49, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x29, 0x2a, 0x2d, 0xce, 0x56, 0x92, 0xe7, 0xe2, 0x74, 0xca, 0x29, 0x0e, 0x4e, 0x4e, 0xcc, 0x49,
	0x2c, 0x12, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09,
	0x02, 0xb3, 0x95, 0x94, 0xb8, 0x78, 0xbc, 0x4a, 0x93, 0xbc, 0x4a, 0x93, 0xf0, 0xa8, 0x51, 0xe3,
	0x12, 0x80, 0xa8, 0x71, 0xce, 0xcf, 0x2d, 0x28, 0x4a, 0x2d, 0x2e, 0x4e, 0x4d, 0xc1, 0xaa, 0x4e,
	0x85, 0x8b, 0x2f, 0x20, 0xbf, 0x38, 0x35, 0x33, 0x25, 0x3f, 0xcf, 0x39, 0xb3, 0x20, 0x23, 0x15,
	0xbb, 0x69, 0xd2, 0x5c, 0xac, 0x01, 0x45, 0xf9, 0xf9, 0x69, 0xd8, 0x24, 0x9d, 0xb4, 0xa2, 0x34,
	0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x53, 0x4a, 0x8b, 0xb3, 0x75,
	0xf3, 0x52, 0x4b, 0xca, 0xf3, 0x8b, 0xb2, 0xf5, 0x41, 0xfe, 0xd1, 0x2d, 0x4e, 0xce, 0x48, 0xcd,
	0x4d, 0xb4, 0x06, 0xb1, 0x93, 0xd8, 0xc0, 0x1e, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x85,
	0x2b, 0x20, 0x3d, 0xfe, 0x00, 0x00, 0x00,
}
