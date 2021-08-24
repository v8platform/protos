// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: ras/api/v1/cluster_admin.proto

package apiv1

import (
	v1 "github.com/v8platform/protos/gen/ras/messages/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AdminsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *AdminsRequest) Reset() {
	*x = AdminsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_api_v1_cluster_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminsRequest) ProtoMessage() {}

func (x *AdminsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_api_v1_cluster_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminsRequest.ProtoReflect.Descriptor instead.
func (*AdminsRequest) Descriptor() ([]byte, []int) {
	return file_ras_api_v1_cluster_admin_proto_rawDescGZIP(), []int{0}
}

func (x *AdminsRequest) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type AdminsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*v1.AdminInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *AdminsResponse) Reset() {
	*x = AdminsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_api_v1_cluster_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminsResponse) ProtoMessage() {}

func (x *AdminsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ras_api_v1_cluster_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminsResponse.ProtoReflect.Descriptor instead.
func (*AdminsResponse) Descriptor() ([]byte, []int) {
	return file_ras_api_v1_cluster_admin_proto_rawDescGZIP(), []int{1}
}

func (x *AdminsResponse) GetItems() []*v1.AdminInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster   *Cluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	AdminName string   `protobuf:"bytes,2,opt,name=admin_name,json=adminName,proto3" json:"admin_name,omitempty"`
}

func (x *DeleteAdminRequest) Reset() {
	*x = DeleteAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_api_v1_cluster_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdminRequest) ProtoMessage() {}

func (x *DeleteAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_api_v1_cluster_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAdminRequest.ProtoReflect.Descriptor instead.
func (*DeleteAdminRequest) Descriptor() ([]byte, []int) {
	return file_ras_api_v1_cluster_admin_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteAdminRequest) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *DeleteAdminRequest) GetAdminName() string {
	if x != nil {
		return x.AdminName
	}
	return ""
}

type AddAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster   *Cluster      `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	AdminInfo *v1.AdminInfo `protobuf:"bytes,2,opt,name=admin_info,json=adminInfo,proto3" json:"admin_info,omitempty"`
}

func (x *AddAdminRequest) Reset() {
	*x = AddAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_api_v1_cluster_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAdminRequest) ProtoMessage() {}

func (x *AddAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_api_v1_cluster_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAdminRequest.ProtoReflect.Descriptor instead.
func (*AddAdminRequest) Descriptor() ([]byte, []int) {
	return file_ras_api_v1_cluster_admin_proto_rawDescGZIP(), []int{3}
}

func (x *AddAdminRequest) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *AddAdminRequest) GetAdminInfo() *v1.AdminInfo {
	if x != nil {
		return x.AdminInfo
	}
	return nil
}

var File_ras_api_v1_cluster_admin_proto protoreflect.FileDescriptor

var file_ras_api_v1_cluster_admin_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x72, 0x61,
	0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x72, 0x61, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0d, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x0e, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x61, 0x73,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x62, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x7b, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xe4,
	0x01, 0x0a, 0x13, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x73,
	0x12, 0x19, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x08, 0x41, 0x64, 0x64,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x72, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0xa0, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67,
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
	file_ras_api_v1_cluster_admin_proto_rawDescOnce sync.Once
	file_ras_api_v1_cluster_admin_proto_rawDescData = file_ras_api_v1_cluster_admin_proto_rawDesc
)

func file_ras_api_v1_cluster_admin_proto_rawDescGZIP() []byte {
	file_ras_api_v1_cluster_admin_proto_rawDescOnce.Do(func() {
		file_ras_api_v1_cluster_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_ras_api_v1_cluster_admin_proto_rawDescData)
	})
	return file_ras_api_v1_cluster_admin_proto_rawDescData
}

var file_ras_api_v1_cluster_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ras_api_v1_cluster_admin_proto_goTypes = []interface{}{
	(*AdminsRequest)(nil),      // 0: ras.api.v1.AdminsRequest
	(*AdminsResponse)(nil),     // 1: ras.api.v1.AdminsResponse
	(*DeleteAdminRequest)(nil), // 2: ras.api.v1.DeleteAdminRequest
	(*AddAdminRequest)(nil),    // 3: ras.api.v1.AddAdminRequest
	(*Cluster)(nil),            // 4: ras.api.v1.Cluster
	(*v1.AdminInfo)(nil),       // 5: ras.messages.v1.AdminInfo
	(*emptypb.Empty)(nil),      // 6: google.protobuf.Empty
}
var file_ras_api_v1_cluster_admin_proto_depIdxs = []int32{
	4, // 0: ras.api.v1.AdminsRequest.cluster:type_name -> ras.api.v1.Cluster
	5, // 1: ras.api.v1.AdminsResponse.items:type_name -> ras.messages.v1.AdminInfo
	4, // 2: ras.api.v1.DeleteAdminRequest.cluster:type_name -> ras.api.v1.Cluster
	4, // 3: ras.api.v1.AddAdminRequest.cluster:type_name -> ras.api.v1.Cluster
	5, // 4: ras.api.v1.AddAdminRequest.admin_info:type_name -> ras.messages.v1.AdminInfo
	0, // 5: ras.api.v1.ClusterAdminService.Admins:input_type -> ras.api.v1.AdminsRequest
	3, // 6: ras.api.v1.ClusterAdminService.AddAdmin:input_type -> ras.api.v1.AddAdminRequest
	2, // 7: ras.api.v1.ClusterAdminService.DeleteAdmin:input_type -> ras.api.v1.DeleteAdminRequest
	1, // 8: ras.api.v1.ClusterAdminService.Admins:output_type -> ras.api.v1.AdminsResponse
	6, // 9: ras.api.v1.ClusterAdminService.AddAdmin:output_type -> google.protobuf.Empty
	6, // 10: ras.api.v1.ClusterAdminService.DeleteAdmin:output_type -> google.protobuf.Empty
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ras_api_v1_cluster_admin_proto_init() }
func file_ras_api_v1_cluster_admin_proto_init() {
	if File_ras_api_v1_cluster_admin_proto != nil {
		return
	}
	file_ras_api_v1_helpers_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ras_api_v1_cluster_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminsRequest); i {
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
		file_ras_api_v1_cluster_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminsResponse); i {
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
		file_ras_api_v1_cluster_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAdminRequest); i {
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
		file_ras_api_v1_cluster_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAdminRequest); i {
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
			RawDescriptor: file_ras_api_v1_cluster_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ras_api_v1_cluster_admin_proto_goTypes,
		DependencyIndexes: file_ras_api_v1_cluster_admin_proto_depIdxs,
		MessageInfos:      file_ras_api_v1_cluster_admin_proto_msgTypes,
	}.Build()
	File_ras_api_v1_cluster_admin_proto = out.File
	file_ras_api_v1_cluster_admin_proto_rawDesc = nil
	file_ras_api_v1_cluster_admin_proto_goTypes = nil
	file_ras_api_v1_cluster_admin_proto_depIdxs = nil
}
