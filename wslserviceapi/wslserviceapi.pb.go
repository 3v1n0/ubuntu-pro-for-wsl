// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.12.4
// source: wslserviceapi.proto

package wslserviceapi

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

type AttachInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AttachInfo) Reset() {
	*x = AttachInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wslserviceapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachInfo) ProtoMessage() {}

func (x *AttachInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wslserviceapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachInfo.ProtoReflect.Descriptor instead.
func (*AttachInfo) Descriptor() ([]byte, []int) {
	return file_wslserviceapi_proto_rawDescGZIP(), []int{0}
}

func (x *AttachInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wslserviceapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_wslserviceapi_proto_msgTypes[1]
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
	return file_wslserviceapi_proto_rawDescGZIP(), []int{1}
}

var File_wslserviceapi_proto protoreflect.FileDescriptor

var file_wslserviceapi_proto_rawDesc = []byte{
	0x0a, 0x13, 0x77, 0x73, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x77, 0x73, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x61, 0x70, 0x69, 0x22, 0x22, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x32, 0x7b, 0x0a, 0x03, 0x57, 0x53, 0x4c, 0x12, 0x3e, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x77, 0x73, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x14, 0x2e, 0x77, 0x73, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67,
	0x12, 0x14, 0x2e, 0x77, 0x73, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x77, 0x73, 0x6c, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x3b,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x6e,
	0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x2f, 0x75, 0x62, 0x75, 0x6e, 0x74, 0x75, 0x2d, 0x70, 0x72,
	0x6f, 0x2d, 0x66, 0x6f, 0x72, 0x2d, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x2f, 0x77, 0x73,
	0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_wslserviceapi_proto_rawDescOnce sync.Once
	file_wslserviceapi_proto_rawDescData = file_wslserviceapi_proto_rawDesc
)

func file_wslserviceapi_proto_rawDescGZIP() []byte {
	file_wslserviceapi_proto_rawDescOnce.Do(func() {
		file_wslserviceapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_wslserviceapi_proto_rawDescData)
	})
	return file_wslserviceapi_proto_rawDescData
}

var file_wslserviceapi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wslserviceapi_proto_goTypes = []interface{}{
	(*AttachInfo)(nil), // 0: wslserviceapi.AttachInfo
	(*Empty)(nil),      // 1: wslserviceapi.Empty
}
var file_wslserviceapi_proto_depIdxs = []int32{
	0, // 0: wslserviceapi.WSL.ProAttach:input_type -> wslserviceapi.AttachInfo
	1, // 1: wslserviceapi.WSL.Ping:input_type -> wslserviceapi.Empty
	1, // 2: wslserviceapi.WSL.ProAttach:output_type -> wslserviceapi.Empty
	1, // 3: wslserviceapi.WSL.Ping:output_type -> wslserviceapi.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wslserviceapi_proto_init() }
func file_wslserviceapi_proto_init() {
	if File_wslserviceapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wslserviceapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachInfo); i {
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
		file_wslserviceapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wslserviceapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wslserviceapi_proto_goTypes,
		DependencyIndexes: file_wslserviceapi_proto_depIdxs,
		MessageInfos:      file_wslserviceapi_proto_msgTypes,
	}.Build()
	File_wslserviceapi_proto = out.File
	file_wslserviceapi_proto_rawDesc = nil
	file_wslserviceapi_proto_goTypes = nil
	file_wslserviceapi_proto_depIdxs = nil
}
