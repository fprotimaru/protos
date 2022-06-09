// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: protos/parser_service.proto

package parser_pb

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_parser_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_parser_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protos_parser_service_proto_rawDescGZIP(), []int{0}
}

type ParseDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFinished bool `protobuf:"varint,1,opt,name=is_finished,json=isFinished,proto3" json:"is_finished,omitempty"`
}

func (x *ParseDataResponse) Reset() {
	*x = ParseDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_parser_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseDataResponse) ProtoMessage() {}

func (x *ParseDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_parser_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseDataResponse.ProtoReflect.Descriptor instead.
func (*ParseDataResponse) Descriptor() ([]byte, []int) {
	return file_protos_parser_service_proto_rawDescGZIP(), []int{1}
}

func (x *ParseDataResponse) GetIsFinished() bool {
	if x != nil {
		return x.IsFinished
	}
	return false
}

var File_protos_parser_service_proto protoreflect.FileDescriptor

var file_protos_parser_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x34, 0x0a, 0x11, 0x50, 0x61,
	0x72, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x32, 0x42, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x73, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72,
	0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_parser_service_proto_rawDescOnce sync.Once
	file_protos_parser_service_proto_rawDescData = file_protos_parser_service_proto_rawDesc
)

func file_protos_parser_service_proto_rawDescGZIP() []byte {
	file_protos_parser_service_proto_rawDescOnce.Do(func() {
		file_protos_parser_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_parser_service_proto_rawDescData)
	})
	return file_protos_parser_service_proto_rawDescData
}

var file_protos_parser_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_parser_service_proto_goTypes = []interface{}{
	(*Empty)(nil),             // 0: pb.Empty
	(*ParseDataResponse)(nil), // 1: pb.ParseDataResponse
}
var file_protos_parser_service_proto_depIdxs = []int32{
	0, // 0: pb.PostParserService.ParseData:input_type -> pb.Empty
	1, // 1: pb.PostParserService.ParseData:output_type -> pb.ParseDataResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_parser_service_proto_init() }
func file_protos_parser_service_proto_init() {
	if File_protos_parser_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_parser_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_protos_parser_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseDataResponse); i {
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
			RawDescriptor: file_protos_parser_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_parser_service_proto_goTypes,
		DependencyIndexes: file_protos_parser_service_proto_depIdxs,
		MessageInfos:      file_protos_parser_service_proto_msgTypes,
	}.Build()
	File_protos_parser_service_proto = out.File
	file_protos_parser_service_proto_rawDesc = nil
	file_protos_parser_service_proto_goTypes = nil
	file_protos_parser_service_proto_depIdxs = nil
}
