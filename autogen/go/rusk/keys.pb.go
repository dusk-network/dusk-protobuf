// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: keys.proto

package rusk

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

type SecretKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A *Scalar `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B *Scalar `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
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

func (x *SecretKey) GetA() *Scalar {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *SecretKey) GetB() *Scalar {
	if x != nil {
		return x.B
	}
	return nil
}

type ViewKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A  *Scalar          `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	BG *CompressedPoint `protobuf:"bytes,2,opt,name=b_g,json=bG,proto3" json:"b_g,omitempty"`
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

func (x *ViewKey) GetA() *Scalar {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *ViewKey) GetBG() *CompressedPoint {
	if x != nil {
		return x.BG
	}
	return nil
}

type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AG *CompressedPoint `protobuf:"bytes,1,opt,name=a_g,json=aG,proto3" json:"a_g,omitempty"`
	BG *CompressedPoint `protobuf:"bytes,2,opt,name=b_g,json=bG,proto3" json:"b_g,omitempty"`
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

func (x *PublicKey) GetAG() *CompressedPoint {
	if x != nil {
		return x.AG
	}
	return nil
}

func (x *PublicKey) GetBG() *CompressedPoint {
	if x != nil {
		return x.BG
	}
	return nil
}

var File_keys_proto protoreflect.FileDescriptor

var file_keys_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x75,
	0x73, 0x6b, 0x1a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x43, 0x0a, 0x09, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x01,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x53,
	0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x01, 0x61, 0x12, 0x1a, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61,
	0x72, 0x52, 0x01, 0x62, 0x22, 0x4d, 0x0a, 0x07, 0x56, 0x69, 0x65, 0x77, 0x4b, 0x65, 0x79, 0x12,
	0x1a, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x75, 0x73,
	0x6b, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x01, 0x61, 0x12, 0x26, 0x0a, 0x03, 0x62,
	0x5f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x02, 0x62, 0x47, 0x22, 0x5b, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x26, 0x0a, 0x03, 0x61, 0x5f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x72, 0x75, 0x73, 0x6b, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02, 0x61, 0x47, 0x12, 0x26, 0x0a, 0x03, 0x62, 0x5f, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x75, 0x73, 0x6b, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02, 0x62, 0x47,
	0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x75, 0x73, 0x6b, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x75, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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

var file_keys_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_keys_proto_goTypes = []interface{}{
	(*SecretKey)(nil),       // 0: rusk.SecretKey
	(*ViewKey)(nil),         // 1: rusk.ViewKey
	(*PublicKey)(nil),       // 2: rusk.PublicKey
	(*Scalar)(nil),          // 3: rusk.Scalar
	(*CompressedPoint)(nil), // 4: rusk.CompressedPoint
}
var file_keys_proto_depIdxs = []int32{
	3, // 0: rusk.SecretKey.a:type_name -> rusk.Scalar
	3, // 1: rusk.SecretKey.b:type_name -> rusk.Scalar
	3, // 2: rusk.ViewKey.a:type_name -> rusk.Scalar
	4, // 3: rusk.ViewKey.b_g:type_name -> rusk.CompressedPoint
	4, // 4: rusk.PublicKey.a_g:type_name -> rusk.CompressedPoint
	4, // 5: rusk.PublicKey.b_g:type_name -> rusk.CompressedPoint
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_keys_proto_init() }
func file_keys_proto_init() {
	if File_keys_proto != nil {
		return
	}
	file_field_proto_init()
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_keys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
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
