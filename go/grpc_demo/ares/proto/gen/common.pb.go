// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: common.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Envelope struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TypeId        uint32                 `protobuf:"varint,1,opt,name=typeId,proto3" json:"typeId,omitempty"`
	PvId          uint32                 `protobuf:"varint,2,opt,name=pvId,proto3" json:"pvId,omitempty"`
	Payload       []byte                 `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	mi := &file_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *Envelope) GetTypeId() uint32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *Envelope) GetPvId() uint32 {
	if x != nil {
		return x.PvId
	}
	return 0
}

func (x *Envelope) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var file_common_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*uint32)(nil),
		Field:         1000,
		Name:          "proto.type_id",
		Tag:           "varint,1000,opt,name=type_id",
		Filename:      "common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1001,
		Name:          "proto.msg_package",
		Tag:           "bytes,1001,opt,name=msg_package",
		Filename:      "common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1002,
		Name:          "proto.msg_registry",
		Tag:           "bytes,1002,opt,name=msg_registry",
		Filename:      "common.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional uint32 type_id = 1000;
	E_TypeId = &file_common_proto_extTypes[0]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// optional string msg_package = 1001;
	E_MsgPackage = &file_common_proto_extTypes[1]
	// optional string msg_registry = 1002;
	E_MsgRegistry = &file_common_proto_extTypes[2]
)

var File_common_proto protoreflect.FileDescriptor

const file_common_proto_rawDesc = "" +
	"\n" +
	"\fcommon.proto\x12\x05proto\x1a google/protobuf/descriptor.proto\"P\n" +
	"\bEnvelope\x12\x16\n" +
	"\x06typeId\x18\x01 \x01(\rR\x06typeId\x12\x12\n" +
	"\x04pvId\x18\x02 \x01(\rR\x04pvId\x12\x18\n" +
	"\apayload\x18\x03 \x01(\fR\apayload:9\n" +
	"\atype_id\x12\x1f.google.protobuf.MessageOptions\x18\xe8\a \x01(\rR\x06typeId:>\n" +
	"\vmsg_package\x12\x1c.google.protobuf.FileOptions\x18\xe9\a \x01(\tR\n" +
	"msgPackage:@\n" +
	"\fmsg_registry\x12\x1c.google.protobuf.FileOptions\x18\xea\a \x01(\tR\vmsgRegistryB\x06Z\x04./pbb\x06proto3"

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData []byte
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_proto_rawDesc), len(file_common_proto_rawDesc)))
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_proto_goTypes = []any{
	(*Envelope)(nil),                    // 0: proto.Envelope
	(*descriptorpb.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
	(*descriptorpb.FileOptions)(nil),    // 2: google.protobuf.FileOptions
}
var file_common_proto_depIdxs = []int32{
	1, // 0: proto.type_id:extendee -> google.protobuf.MessageOptions
	2, // 1: proto.msg_package:extendee -> google.protobuf.FileOptions
	2, // 2: proto.msg_registry:extendee -> google.protobuf.FileOptions
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_proto_rawDesc), len(file_common_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
		ExtensionInfos:    file_common_proto_extTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
