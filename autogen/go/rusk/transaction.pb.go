// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transaction.proto

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

type TransactionInput struct {
	Nullifier            *Nullifier `protobuf:"bytes,1,opt,name=nullifier,proto3" json:"nullifier,omitempty"`
	MerkleRoot           *Scalar    `protobuf:"bytes,2,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TransactionInput) Reset()         { *m = TransactionInput{} }
func (m *TransactionInput) String() string { return proto.CompactTextString(m) }
func (*TransactionInput) ProtoMessage()    {}
func (*TransactionInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}
func (m *TransactionInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionInput.Unmarshal(m, b)
}
func (m *TransactionInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionInput.Marshal(b, m, deterministic)
}
func (m *TransactionInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionInput.Merge(m, src)
}
func (m *TransactionInput) XXX_Size() int {
	return xxx_messageInfo_TransactionInput.Size(m)
}
func (m *TransactionInput) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionInput.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionInput proto.InternalMessageInfo

func (m *TransactionInput) GetNullifier() *Nullifier {
	if m != nil {
		return m.Nullifier
	}
	return nil
}

func (m *TransactionInput) GetMerkleRoot() *Scalar {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

type TransactionOutput struct {
	Note                 *Note      `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
	Pk                   *PublicKey `protobuf:"bytes,2,opt,name=pk,proto3" json:"pk,omitempty"`
	Value                uint64     `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	BlindingFactor       *Scalar    `protobuf:"bytes,4,opt,name=blinding_factor,json=blindingFactor,proto3" json:"blinding_factor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TransactionOutput) Reset()         { *m = TransactionOutput{} }
func (m *TransactionOutput) String() string { return proto.CompactTextString(m) }
func (*TransactionOutput) ProtoMessage()    {}
func (*TransactionOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{1}
}
func (m *TransactionOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionOutput.Unmarshal(m, b)
}
func (m *TransactionOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionOutput.Marshal(b, m, deterministic)
}
func (m *TransactionOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionOutput.Merge(m, src)
}
func (m *TransactionOutput) XXX_Size() int {
	return xxx_messageInfo_TransactionOutput.Size(m)
}
func (m *TransactionOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionOutput.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionOutput proto.InternalMessageInfo

func (m *TransactionOutput) GetNote() *Note {
	if m != nil {
		return m.Note
	}
	return nil
}

func (m *TransactionOutput) GetPk() *PublicKey {
	if m != nil {
		return m.Pk
	}
	return nil
}

func (m *TransactionOutput) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TransactionOutput) GetBlindingFactor() *Scalar {
	if m != nil {
		return m.BlindingFactor
	}
	return nil
}

type Transaction struct {
	Inputs               []*TransactionInput  `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs              []*TransactionOutput `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
	Fee                  *TransactionOutput   `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Proof                []byte               `protobuf:"bytes,4,opt,name=proof,proto3" json:"proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{2}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetInputs() []*TransactionInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Transaction) GetOutputs() []*TransactionOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *Transaction) GetFee() *TransactionOutput {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *Transaction) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionInput)(nil), "rusk.TransactionInput")
	proto.RegisterType((*TransactionOutput)(nil), "rusk.TransactionOutput")
	proto.RegisterType((*Transaction)(nil), "rusk.Transaction")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_2cc4e03d2c28c490) }

var fileDescriptor_2cc4e03d2c28c490 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0x4a, 0xe3, 0x40,
	0x14, 0x86, 0x49, 0xda, 0xed, 0xb2, 0x27, 0x65, 0xbb, 0x1d, 0x96, 0xdd, 0xd0, 0x0b, 0x2d, 0xbd,
	0x8a, 0x42, 0x13, 0xac, 0xf8, 0x02, 0x5e, 0x08, 0x22, 0xa8, 0x8c, 0x5e, 0x79, 0x53, 0x92, 0x74,
	0x12, 0x87, 0x4c, 0xe7, 0x84, 0xc9, 0x8c, 0xd0, 0xc7, 0xf1, 0x11, 0x7c, 0x43, 0xc9, 0x4c, 0x42,
	0x8b, 0x8a, 0x77, 0x33, 0x27, 0xdf, 0x7f, 0xfe, 0xff, 0x9f, 0xc0, 0x54, 0xab, 0x54, 0x36, 0x69,
	0xae, 0x39, 0xca, 0xb8, 0x56, 0xa8, 0x91, 0x0c, 0x95, 0x69, 0xaa, 0x59, 0x50, 0x70, 0x26, 0x36,
	0x6e, 0x34, 0x83, 0x8a, 0xed, 0x9a, 0xfe, 0x2c, 0x51, 0x33, 0x77, 0x5e, 0xd4, 0xf0, 0xe7, 0x71,
	0xaf, 0xbf, 0x96, 0xb5, 0xd1, 0x64, 0x09, 0xbf, 0xa4, 0x11, 0x82, 0x17, 0x9c, 0xa9, 0xd0, 0x9b,
	0x7b, 0x51, 0xb0, 0x9a, 0xc4, 0xed, 0xca, 0xf8, 0xb6, 0x1f, 0xd3, 0x3d, 0x41, 0x96, 0x10, 0x6c,
	0x99, 0xaa, 0x04, 0x5b, 0x2b, 0x44, 0x1d, 0xfa, 0x56, 0x30, 0x76, 0x82, 0x87, 0x3c, 0x15, 0xa9,
	0xa2, 0xe0, 0x00, 0x8a, 0xa8, 0x17, 0xaf, 0x1e, 0x4c, 0x0f, 0x2c, 0xef, 0x8c, 0x6e, 0x3d, 0x8f,
	0x60, 0xd8, 0xa6, 0xea, 0xec, 0xa0, 0xb3, 0x43, 0xcd, 0xa8, 0x9d, 0x93, 0x63, 0xf0, 0xeb, 0xaa,
	0xdb, 0xdd, 0x85, 0xb9, 0x37, 0x99, 0xe0, 0xf9, 0x0d, 0xdb, 0x51, 0xbf, 0xae, 0xc8, 0x5f, 0xf8,
	0xf1, 0x92, 0x0a, 0xc3, 0xc2, 0xc1, 0xdc, 0x8b, 0x46, 0xd4, 0x5d, 0xc8, 0x05, 0x4c, 0x32, 0xc1,
	0xe5, 0x86, 0xcb, 0x72, 0x5d, 0xa4, 0xb9, 0x46, 0x15, 0x0e, 0xbf, 0xc8, 0xf7, 0xbb, 0x87, 0xae,
	0x2c, 0xb3, 0x78, 0xf3, 0x20, 0x38, 0xc8, 0x48, 0x62, 0x18, 0xf1, 0xf6, 0x69, 0x9a, 0xd0, 0x9b,
	0x0f, 0xa2, 0x60, 0xf5, 0xcf, 0xa9, 0x3f, 0xbe, 0x1c, 0xed, 0x28, 0x72, 0x06, 0x3f, 0xd1, 0xf6,
	0x6a, 0x42, 0xdf, 0x0a, 0xfe, 0x7f, 0x12, 0xb8, 0xde, 0xb4, 0xe7, 0xc8, 0x09, 0x0c, 0x0a, 0xe6,
	0xd2, 0x7f, 0x83, 0xb7, 0x4c, 0x5b, 0xb5, 0x56, 0x88, 0x85, 0xad, 0x32, 0xa6, 0xee, 0x72, 0x79,
	0xfa, 0x14, 0x95, 0x5c, 0x3f, 0x9b, 0x2c, 0xce, 0x71, 0x9b, 0x6c, 0x4c, 0x53, 0x2d, 0xed, 0x2f,
	0xce, 0x4c, 0x91, 0xa4, 0x46, 0x63, 0xc9, 0x64, 0x52, 0x62, 0xd2, 0x2e, 0xce, 0x46, 0xf6, 0xcb,
	0xf9, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x6f, 0x73, 0x31, 0x3c, 0x02, 0x00, 0x00,
}
