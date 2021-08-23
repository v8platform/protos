// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: ras/api/v1/helpers.proto

package apiv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Usr *string `protobuf:"bytes,2,opt,name=usr,proto3,oneof" json:"usr,omitempty"`
	Pwd *string `protobuf:"bytes,3,opt,name=pwd,proto3,oneof" json:"pwd,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_api_v1_helpers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_ras_api_v1_helpers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_ras_api_v1_helpers_proto_rawDescGZIP(), []int{0}
}

func (x *Cluster) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cluster) GetUsr() string {
	if x != nil && x.Usr != nil {
		return *x.Usr
	}
	return ""
}

func (x *Cluster) GetPwd() string {
	if x != nil && x.Pwd != nil {
		return *x.Pwd
	}
	return ""
}

var File_ras_api_v1_helpers_proto protoreflect.FileDescriptor

var file_ras_api_v1_helpers_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x61, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x57, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x03,
	0x75, 0x73, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x73, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x70, 0x77, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x03, 0x70, 0x77, 0x64, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75,
	0x73, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x70, 0x77, 0x64, 0x42, 0x9b, 0x01, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x48,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x72, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x52, 0x41, 0x58, 0xaa, 0x02, 0x0a, 0x52, 0x61, 0x73, 0x2e, 0x41, 0x70, 0x69,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x52, 0x61, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x16, 0x52, 0x61, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x52, 0x61, 0x73, 0x3a,
	0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ras_api_v1_helpers_proto_rawDescOnce sync.Once
	file_ras_api_v1_helpers_proto_rawDescData = file_ras_api_v1_helpers_proto_rawDesc
)

func file_ras_api_v1_helpers_proto_rawDescGZIP() []byte {
	file_ras_api_v1_helpers_proto_rawDescOnce.Do(func() {
		file_ras_api_v1_helpers_proto_rawDescData = protoimpl.X.CompressGZIP(file_ras_api_v1_helpers_proto_rawDescData)
	})
	return file_ras_api_v1_helpers_proto_rawDescData
}

var file_ras_api_v1_helpers_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ras_api_v1_helpers_proto_goTypes = []interface{}{
	(*Cluster)(nil), // 0: ras.api.v1.Cluster
}
var file_ras_api_v1_helpers_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ras_api_v1_helpers_proto_init() }
func file_ras_api_v1_helpers_proto_init() {
	if File_ras_api_v1_helpers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ras_api_v1_helpers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
	file_ras_api_v1_helpers_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ras_api_v1_helpers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ras_api_v1_helpers_proto_goTypes,
		DependencyIndexes: file_ras_api_v1_helpers_proto_depIdxs,
		MessageInfos:      file_ras_api_v1_helpers_proto_msgTypes,
	}.Build()
	File_ras_api_v1_helpers_proto = out.File
	file_ras_api_v1_helpers_proto_rawDesc = nil
	file_ras_api_v1_helpers_proto_goTypes = nil
	file_ras_api_v1_helpers_proto_depIdxs = nil
}
