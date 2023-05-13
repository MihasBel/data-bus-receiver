// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: v1/bus/bus.proto

package bus

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bus_bus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bus_bus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_v1_bus_bus_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_v1_bus_bus_proto protoreflect.FileDescriptor

var file_v1_bus_bus_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x76, 0x31, 0x2e, 0x62, 0x75, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x47, 0x0a, 0x0a, 0x42, 0x75,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e, 0x76, 0x31, 0x2e,
	0x62, 0x75, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x08, 0x5a, 0x06, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_bus_bus_proto_rawDescOnce sync.Once
	file_v1_bus_bus_proto_rawDescData = file_v1_bus_bus_proto_rawDesc
)

func file_v1_bus_bus_proto_rawDescGZIP() []byte {
	file_v1_bus_bus_proto_rawDescOnce.Do(func() {
		file_v1_bus_bus_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_bus_bus_proto_rawDescData)
	})
	return file_v1_bus_bus_proto_rawDescData
}

var file_v1_bus_bus_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v1_bus_bus_proto_goTypes = []interface{}{
	(*Message)(nil),       // 0: v1.bus.Message
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_v1_bus_bus_proto_depIdxs = []int32{
	0, // 0: v1.bus.BusService.ReceiveMessage:input_type -> v1.bus.Message
	1, // 1: v1.bus.BusService.ReceiveMessage:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_bus_bus_proto_init() }
func file_v1_bus_bus_proto_init() {
	if File_v1_bus_bus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_bus_bus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_v1_bus_bus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_bus_bus_proto_goTypes,
		DependencyIndexes: file_v1_bus_bus_proto_depIdxs,
		MessageInfos:      file_v1_bus_bus_proto_msgTypes,
	}.Build()
	File_v1_bus_bus_proto = out.File
	file_v1_bus_bus_proto_rawDesc = nil
	file_v1_bus_bus_proto_goTypes = nil
	file_v1_bus_bus_proto_depIdxs = nil
}
