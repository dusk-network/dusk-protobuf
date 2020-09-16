// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: transaction.proto

package _

import (
	proto "github.com/golang/protobuf/proto"
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

type Crossover struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValueComm    *JubJubCompressed `protobuf:"bytes,1,opt,name=value_comm,json=valueComm,proto3" json:"value_comm,omitempty"`
	Nonce        *BlsScalar        `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	EncyptedData *PoseidonCipher   `protobuf:"bytes,3,opt,name=encypted_data,json=encyptedData,proto3" json:"encypted_data,omitempty"`
}

func (x *Crossover) Reset() {
	*x = Crossover{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crossover) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crossover) ProtoMessage() {}

func (x *Crossover) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crossover.ProtoReflect.Descriptor instead.
func (*Crossover) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Crossover) GetValueComm() *JubJubCompressed {
	if x != nil {
		return x.ValueComm
	}
	return nil
}

func (x *Crossover) GetNonce() *BlsScalar {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *Crossover) GetEncyptedData() *PoseidonCipher {
	if x != nil {
		return x.EncyptedData
	}
	return nil
}

type Fee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GasLimit uint64            `protobuf:"varint,1,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasPrice uint64            `protobuf:"varint,2,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	R        *JubJubCompressed `protobuf:"bytes,3,opt,name=R,proto3" json:"R,omitempty"`
	PkR      *JubJubCompressed `protobuf:"bytes,4,opt,name=pk_r,json=pkR,proto3" json:"pk_r,omitempty"`
}

func (x *Fee) Reset() {
	*x = Fee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fee) ProtoMessage() {}

func (x *Fee) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fee.ProtoReflect.Descriptor instead.
func (*Fee) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *Fee) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *Fee) GetGasPrice() uint64 {
	if x != nil {
		return x.GasPrice
	}
	return 0
}

func (x *Fee) GetR() *JubJubCompressed {
	if x != nil {
		return x.R
	}
	return nil
}

func (x *Fee) GetPkR() *JubJubCompressed {
	if x != nil {
		return x.PkR
	}
	return nil
}

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Randomness   *JubJubCompressed `protobuf:"bytes,1,opt,name=randomness,proto3" json:"randomness,omitempty"`
	PkR          *JubJubCompressed `protobuf:"bytes,2,opt,name=pk_r,json=pkR,proto3" json:"pk_r,omitempty"`
	Commitment   *JubJubCompressed `protobuf:"bytes,3,opt,name=commitment,proto3" json:"commitment,omitempty"`
	Nonce        *BlsScalar        `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	EncyptedData *PoseidonCipher   `protobuf:"bytes,5,opt,name=encypted_data,json=encyptedData,proto3" json:"encypted_data,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *Note) GetRandomness() *JubJubCompressed {
	if x != nil {
		return x.Randomness
	}
	return nil
}

func (x *Note) GetPkR() *JubJubCompressed {
	if x != nil {
		return x.PkR
	}
	return nil
}

func (x *Note) GetCommitment() *JubJubCompressed {
	if x != nil {
		return x.Commitment
	}
	return nil
}

func (x *Note) GetNonce() *BlsScalar {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *Note) GetEncyptedData() *PoseidonCipher {
	if x != nil {
		return x.EncyptedData
	}
	return nil
}

type TransactionPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Anchor        *BlsScalar   `protobuf:"bytes,1,opt,name=anchor,proto3" json:"anchor,omitempty"`
	Nullifier     []*BlsScalar `protobuf:"bytes,2,rep,name=nullifier,proto3" json:"nullifier,omitempty"`
	Crossover     *Crossover   `protobuf:"bytes,3,opt,name=crossover,proto3" json:"crossover,omitempty"`
	Notes         []*Note      `protobuf:"bytes,4,rep,name=notes,proto3" json:"notes,omitempty"`
	Fee           *Fee         `protobuf:"bytes,5,opt,name=fee,proto3" json:"fee,omitempty"`
	SpendingProof *Proof       `protobuf:"bytes,6,opt,name=spending_proof,json=spendingProof,proto3" json:"spending_proof,omitempty"`
	CallData      []byte       `protobuf:"bytes,7,opt,name=call_data,json=callData,proto3" json:"call_data,omitempty"`
}

func (x *TransactionPayload) Reset() {
	*x = TransactionPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionPayload) ProtoMessage() {}

func (x *TransactionPayload) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionPayload.ProtoReflect.Descriptor instead.
func (*TransactionPayload) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionPayload) GetAnchor() *BlsScalar {
	if x != nil {
		return x.Anchor
	}
	return nil
}

func (x *TransactionPayload) GetNullifier() []*BlsScalar {
	if x != nil {
		return x.Nullifier
	}
	return nil
}

func (x *TransactionPayload) GetCrossover() *Crossover {
	if x != nil {
		return x.Crossover
	}
	return nil
}

func (x *TransactionPayload) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *TransactionPayload) GetFee() *Fee {
	if x != nil {
		return x.Fee
	}
	return nil
}

func (x *TransactionPayload) GetSpendingProof() *Proof {
	if x != nil {
		return x.SpendingProof
	}
	return nil
}

func (x *TransactionPayload) GetCallData() []byte {
	if x != nil {
		return x.CallData
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint32              `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Type      uint32              `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	TxPayload *TransactionPayload `protobuf:"bytes,3,opt,name=tx_payload,json=txPayload,proto3" json:"tx_payload,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *Transaction) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Transaction) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Transaction) GetTxPayload() *TransactionPayload {
	if x != nil {
		return x.TxPayload
	}
	return nil
}

var File_transaction_proto protoreflect.FileDescriptor

var file_transaction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x75, 0x73, 0x6b, 0x1a, 0x12, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01,
	0x0a, 0x09, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0a, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x4a, 0x75, 0x62, 0x4a, 0x75, 0x62, 0x43, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x12, 0x25, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x42, 0x6c, 0x73, 0x53, 0x63, 0x61, 0x6c,
	0x61, 0x72, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x65, 0x6e, 0x63,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x50, 0x6f, 0x73, 0x65, 0x69, 0x64, 0x6f, 0x6e,
	0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x52, 0x0c, 0x65, 0x6e, 0x63, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x90, 0x01, 0x0a, 0x03, 0x46, 0x65, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61,
	0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x01, 0x52, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x4a, 0x75, 0x62, 0x4a, 0x75, 0x62, 0x43,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52, 0x01, 0x52, 0x12, 0x29, 0x0a, 0x04,
	0x70, 0x6b, 0x5f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x75, 0x73,
	0x6b, 0x2e, 0x4a, 0x75, 0x62, 0x4a, 0x75, 0x62, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x64, 0x52, 0x03, 0x70, 0x6b, 0x52, 0x22, 0x83, 0x02, 0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65,
	0x12, 0x36, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x4a, 0x75, 0x62, 0x4a,
	0x75, 0x62, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52, 0x0a, 0x72, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x70, 0x6b, 0x5f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x4a, 0x75,
	0x62, 0x4a, 0x75, 0x62, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52, 0x03,
	0x70, 0x6b, 0x52, 0x12, 0x36, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x4a,
	0x75, 0x62, 0x4a, 0x75, 0x62, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52,
	0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x75, 0x73,
	0x6b, 0x2e, 0x42, 0x6c, 0x73, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x65, 0x6e, 0x63, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x75, 0x73, 0x6b,
	0x2e, 0x50, 0x6f, 0x73, 0x65, 0x69, 0x64, 0x6f, 0x6e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x52,
	0x0c, 0x65, 0x6e, 0x63, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x22, 0xab, 0x02,
	0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x27, 0x0a, 0x06, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x42, 0x6c, 0x73, 0x53,
	0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x06, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x12, 0x2d, 0x0a,
	0x09, 0x6e, 0x75, 0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x42, 0x6c, 0x73, 0x53, 0x63, 0x61, 0x6c, 0x61,
	0x72, 0x52, 0x09, 0x6e, 0x75, 0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x09,
	0x63, 0x72, 0x6f, 0x73, 0x73, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x6f, 0x76, 0x65, 0x72,
	0x52, 0x09, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x05, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x75, 0x73,
	0x6b, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x1b, 0x0a,
	0x03, 0x66, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x75, 0x73,
	0x6b, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x32, 0x0a, 0x0e, 0x73, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52,
	0x0d, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x22, 0x74, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x74, 0x78, 0x5f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72,
	0x75, 0x73, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x09, 0x74, 0x78, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x75, 0x73, 0x6b, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x75, 0x73,
	0x6b, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_transaction_proto_rawDescOnce sync.Once
	file_transaction_proto_rawDescData = file_transaction_proto_rawDesc
)

func file_transaction_proto_rawDescGZIP() []byte {
	file_transaction_proto_rawDescOnce.Do(func() {
		file_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_proto_rawDescData)
	})
	return file_transaction_proto_rawDescData
}

var file_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_transaction_proto_goTypes = []interface{}{
	(*Crossover)(nil),          // 0: rusk.Crossover
	(*Fee)(nil),                // 1: rusk.Fee
	(*Note)(nil),               // 2: rusk.Note
	(*TransactionPayload)(nil), // 3: rusk.TransactionPayload
	(*Transaction)(nil),        // 4: rusk.Transaction
	(*JubJubCompressed)(nil),   // 5: rusk.JubJubCompressed
	(*BlsScalar)(nil),          // 6: rusk.BlsScalar
	(*PoseidonCipher)(nil),     // 7: rusk.PoseidonCipher
	(*Proof)(nil),              // 8: rusk.Proof
}
var file_transaction_proto_depIdxs = []int32{
	5,  // 0: rusk.Crossover.value_comm:type_name -> rusk.JubJubCompressed
	6,  // 1: rusk.Crossover.nonce:type_name -> rusk.BlsScalar
	7,  // 2: rusk.Crossover.encypted_data:type_name -> rusk.PoseidonCipher
	5,  // 3: rusk.Fee.R:type_name -> rusk.JubJubCompressed
	5,  // 4: rusk.Fee.pk_r:type_name -> rusk.JubJubCompressed
	5,  // 5: rusk.Note.randomness:type_name -> rusk.JubJubCompressed
	5,  // 6: rusk.Note.pk_r:type_name -> rusk.JubJubCompressed
	5,  // 7: rusk.Note.commitment:type_name -> rusk.JubJubCompressed
	6,  // 8: rusk.Note.nonce:type_name -> rusk.BlsScalar
	7,  // 9: rusk.Note.encypted_data:type_name -> rusk.PoseidonCipher
	6,  // 10: rusk.TransactionPayload.anchor:type_name -> rusk.BlsScalar
	6,  // 11: rusk.TransactionPayload.nullifier:type_name -> rusk.BlsScalar
	0,  // 12: rusk.TransactionPayload.crossover:type_name -> rusk.Crossover
	2,  // 13: rusk.TransactionPayload.notes:type_name -> rusk.Note
	1,  // 14: rusk.TransactionPayload.fee:type_name -> rusk.Fee
	8,  // 15: rusk.TransactionPayload.spending_proof:type_name -> rusk.Proof
	3,  // 16: rusk.Transaction.tx_payload:type_name -> rusk.TransactionPayload
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_transaction_proto_init() }
func file_transaction_proto_init() {
	if File_transaction_proto != nil {
		return
	}
	file_basic_fields_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crossover); i {
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
		file_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fee); i {
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
		file_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
		file_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionPayload); i {
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
		file_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
			RawDescriptor: file_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_proto_goTypes,
		DependencyIndexes: file_transaction_proto_depIdxs,
		MessageInfos:      file_transaction_proto_msgTypes,
	}.Build()
	File_transaction_proto = out.File
	file_transaction_proto_rawDesc = nil
	file_transaction_proto_goTypes = nil
	file_transaction_proto_depIdxs = nil
}
