// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: one.proto

package proto

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

type MessageOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MessageOne) Reset() {
	*x = MessageOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_one_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOne) ProtoMessage() {}

func (x *MessageOne) ProtoReflect() protoreflect.Message {
	mi := &file_one_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOne.ProtoReflect.Descriptor instead.
func (*MessageOne) Descriptor() ([]byte, []int) {
	return file_one_proto_rawDescGZIP(), []int{0}
}

func (x *MessageOne) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_one_proto protoreflect.FileDescriptor

var file_one_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x20, 0x0a, 0x0a, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x2b, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6d, 0x65,
	0x73, 0x68, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2d,
	0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_one_proto_rawDescOnce sync.Once
	file_one_proto_rawDescData = file_one_proto_rawDesc
)

func file_one_proto_rawDescGZIP() []byte {
	file_one_proto_rawDescOnce.Do(func() {
		file_one_proto_rawDescData = protoimpl.X.CompressGZIP(file_one_proto_rawDescData)
	})
	return file_one_proto_rawDescData
}

var file_one_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_one_proto_goTypes = []interface{}{
	(*MessageOne)(nil), // 0: example.simple.MessageOne
}
var file_one_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_one_proto_init() }
func file_one_proto_init() {
	if File_one_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_one_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOne); i {
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
			RawDescriptor: file_one_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_one_proto_goTypes,
		DependencyIndexes: file_one_proto_depIdxs,
		MessageInfos:      file_one_proto_msgTypes,
	}.Build()
	File_one_proto = out.File
	file_one_proto_rawDesc = nil
	file_one_proto_goTypes = nil
	file_one_proto_depIdxs = nil
}