// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: provisioner.proto

package rusk

import (
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

type Provisioner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKeyBls    []byte   `protobuf:"bytes,1,opt,name=public_key_bls,json=publicKeyBls,proto3" json:"public_key_bls,omitempty"`
	RawPublicKeyBls []byte   `protobuf:"bytes,2,opt,name=raw_public_key_bls,json=rawPublicKeyBls,proto3" json:"raw_public_key_bls,omitempty"`
	Stakes          []*Stake `protobuf:"bytes,3,rep,name=stakes,proto3" json:"stakes,omitempty"`
}

func (x *Provisioner) Reset() {
	*x = Provisioner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provisioner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provisioner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provisioner) ProtoMessage() {}

func (x *Provisioner) ProtoReflect() protoreflect.Message {
	mi := &file_provisioner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provisioner.ProtoReflect.Descriptor instead.
func (*Provisioner) Descriptor() ([]byte, []int) {
	return file_provisioner_proto_rawDescGZIP(), []int{0}
}

func (x *Provisioner) GetPublicKeyBls() []byte {
	if x != nil {
		return x.PublicKeyBls
	}
	return nil
}

func (x *Provisioner) GetRawPublicKeyBls() []byte {
	if x != nil {
		return x.RawPublicKeyBls
	}
	return nil
}

func (x *Provisioner) GetStakes() []*Stake {
	if x != nil {
		return x.Stakes
	}
	return nil
}

var File_provisioner_proto protoreflect.FileDescriptor

var file_provisioner_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x75, 0x73, 0x6b, 0x1a, 0x0b, 0x73, 0x74, 0x61, 0x6b, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x62, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x6c, 0x73, 0x12, 0x2b, 0x0a, 0x12,
	0x72, 0x61, 0x77, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x62,
	0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x72, 0x61, 0x77, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x6c, 0x73, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x6b, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x75, 0x73, 0x6b,
	0x2e, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_provisioner_proto_rawDescOnce sync.Once
	file_provisioner_proto_rawDescData = file_provisioner_proto_rawDesc
)

func file_provisioner_proto_rawDescGZIP() []byte {
	file_provisioner_proto_rawDescOnce.Do(func() {
		file_provisioner_proto_rawDescData = protoimpl.X.CompressGZIP(file_provisioner_proto_rawDescData)
	})
	return file_provisioner_proto_rawDescData
}

var file_provisioner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_provisioner_proto_goTypes = []interface{}{
	(*Provisioner)(nil), // 0: rusk.Provisioner
	(*Stake)(nil),       // 1: rusk.Stake
}
var file_provisioner_proto_depIdxs = []int32{
	1, // 0: rusk.Provisioner.stakes:type_name -> rusk.Stake
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_provisioner_proto_init() }
func file_provisioner_proto_init() {
	if File_provisioner_proto != nil {
		return
	}
	file_stake_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_provisioner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provisioner); i {
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
			RawDescriptor: file_provisioner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_provisioner_proto_goTypes,
		DependencyIndexes: file_provisioner_proto_depIdxs,
		MessageInfos:      file_provisioner_proto_msgTypes,
	}.Build()
	File_provisioner_proto = out.File
	file_provisioner_proto_rawDesc = nil
	file_provisioner_proto_goTypes = nil
	file_provisioner_proto_depIdxs = nil
}
