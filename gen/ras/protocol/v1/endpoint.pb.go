// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: ras/protocol/v1/endpoint.proto

package protocolv1

import (
	_ "github.com/v8platform/protos/gen/ras/encoding"
	v1 "github.com/v8platform/protos/gen/ras/messages/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EndpointOpen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service  string            `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Encoding string            `protobuf:"bytes,2,opt,name=encoding,proto3" json:"encoding,omitempty"`
	Version  string            `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
	Params   map[string]*Param `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EndpointOpen) Reset() {
	*x = EndpointOpen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointOpen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointOpen) ProtoMessage() {}

func (x *EndpointOpen) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointOpen.ProtoReflect.Descriptor instead.
func (*EndpointOpen) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *EndpointOpen) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *EndpointOpen) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

func (x *EndpointOpen) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *EndpointOpen) GetParams() map[string]*Param {
	if x != nil {
		return x.Params
	}
	return nil
}

type EndpointOpenAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId  string            `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Version    string            `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"`
	EnpoiontId int32             `protobuf:"varint,3,opt,name=enpoiont_id,json=enpoiontId,proto3" json:"enpoiont_id,omitempty"`
	Params     map[string]*Param `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EndpointOpenAck) Reset() {
	*x = EndpointOpenAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointOpenAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointOpenAck) ProtoMessage() {}

func (x *EndpointOpenAck) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointOpenAck.ProtoReflect.Descriptor instead.
func (*EndpointOpenAck) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_endpoint_proto_rawDescGZIP(), []int{1}
}

func (x *EndpointOpenAck) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *EndpointOpenAck) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *EndpointOpenAck) GetEnpoiontId() int32 {
	if x != nil {
		return x.EnpoiontId
	}
	return 0
}

func (x *EndpointOpenAck) GetParams() map[string]*Param {
	if x != nil {
		return x.Params
	}
	return nil
}

type EndpointMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId int32                `protobuf:"varint,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	Format     int32                `protobuf:"varint,2,opt,name=format,proto3" json:"format,omitempty"`
	Type       EndpointDataType     `protobuf:"varint,3,opt,name=type,proto3,enum=ras.protocol.v1.EndpointDataType" json:"type,omitempty"`
	Data       *EndpointDataMessage `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EndpointMessage) Reset() {
	*x = EndpointMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointMessage) ProtoMessage() {}

func (x *EndpointMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointMessage.ProtoReflect.Descriptor instead.
func (*EndpointMessage) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_endpoint_proto_rawDescGZIP(), []int{2}
}

func (x *EndpointMessage) GetEndpointId() int32 {
	if x != nil {
		return x.EndpointId
	}
	return 0
}

func (x *EndpointMessage) GetFormat() int32 {
	if x != nil {
		return x.Format
	}
	return 0
}

func (x *EndpointMessage) GetType() EndpointDataType {
	if x != nil {
		return x.Type
	}
	return EndpointDataType_ENDPOINT_DATA_TYPE_VOID_MESSAGE
}

func (x *EndpointMessage) GetData() *EndpointDataMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

type EndpointMessageAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId int32            `protobuf:"varint,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	Format     int32            `protobuf:"varint,2,opt,name=format,proto3" json:"format,omitempty"`
	Type       EndpointDataType `protobuf:"varint,3,opt,name=type,proto3,enum=ras.protocol.v1.EndpointDataType" json:"type,omitempty"`
	// Types that are assignable to Data:
	//	*EndpointMessageAck_VoidMessage
	//	*EndpointMessageAck_Message
	//	*EndpointMessageAck_Failure
	Data isEndpointMessageAck_Data `protobuf_oneof:"data"`
}

func (x *EndpointMessageAck) Reset() {
	*x = EndpointMessageAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointMessageAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointMessageAck) ProtoMessage() {}

func (x *EndpointMessageAck) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointMessageAck.ProtoReflect.Descriptor instead.
func (*EndpointMessageAck) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_endpoint_proto_rawDescGZIP(), []int{3}
}

func (x *EndpointMessageAck) GetEndpointId() int32 {
	if x != nil {
		return x.EndpointId
	}
	return 0
}

func (x *EndpointMessageAck) GetFormat() int32 {
	if x != nil {
		return x.Format
	}
	return 0
}

func (x *EndpointMessageAck) GetType() EndpointDataType {
	if x != nil {
		return x.Type
	}
	return EndpointDataType_ENDPOINT_DATA_TYPE_VOID_MESSAGE
}

func (m *EndpointMessageAck) GetData() isEndpointMessageAck_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *EndpointMessageAck) GetVoidMessage() *EndpointDataVoidMessage {
	if x, ok := x.GetData().(*EndpointMessageAck_VoidMessage); ok {
		return x.VoidMessage
	}
	return nil
}

func (x *EndpointMessageAck) GetMessage() *EndpointDataMessage {
	if x, ok := x.GetData().(*EndpointMessageAck_Message); ok {
		return x.Message
	}
	return nil
}

func (x *EndpointMessageAck) GetFailure() *EndpointFailureMessage {
	if x, ok := x.GetData().(*EndpointMessageAck_Failure); ok {
		return x.Failure
	}
	return nil
}

type isEndpointMessageAck_Data interface {
	isEndpointMessageAck_Data()
}

type EndpointMessageAck_VoidMessage struct {
	VoidMessage *EndpointDataVoidMessage `protobuf:"bytes,4,opt,name=void_message,json=voidMessage,proto3,oneof"`
}

type EndpointMessageAck_Message struct {
	Message *EndpointDataMessage `protobuf:"bytes,5,opt,name=message,proto3,oneof"`
}

type EndpointMessageAck_Failure struct {
	Failure *EndpointFailureMessage `protobuf:"bytes,6,opt,name=failure,proto3,oneof"`
}

func (*EndpointMessageAck_VoidMessage) isEndpointMessageAck_Data() {}

func (*EndpointMessageAck_Message) isEndpointMessageAck_Data() {}

func (*EndpointMessageAck_Failure) isEndpointMessageAck_Data() {}

type EndpointFailureAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId  string      `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Version    string      `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	EndpointId int32       `protobuf:"varint,4,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	ClassCause string      `protobuf:"bytes,5,opt,name=class_cause,json=classCause,proto3" json:"class_cause,omitempty"`
	Message    string      `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	Trace      []string    `protobuf:"bytes,7,rep,name=trace,proto3" json:"trace,omitempty"`
	Cause      *CauseError `protobuf:"bytes,8,opt,name=cause,proto3,oneof" json:"cause,omitempty"`
}

func (x *EndpointFailureAck) Reset() {
	*x = EndpointFailureAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointFailureAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointFailureAck) ProtoMessage() {}

func (x *EndpointFailureAck) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointFailureAck.ProtoReflect.Descriptor instead.
func (*EndpointFailureAck) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_endpoint_proto_rawDescGZIP(), []int{4}
}

func (x *EndpointFailureAck) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *EndpointFailureAck) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *EndpointFailureAck) GetEndpointId() int32 {
	if x != nil {
		return x.EndpointId
	}
	return 0
}

func (x *EndpointFailureAck) GetClassCause() string {
	if x != nil {
		return x.ClassCause
	}
	return ""
}

func (x *EndpointFailureAck) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EndpointFailureAck) GetTrace() []string {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *EndpointFailureAck) GetCause() *CauseError {
	if x != nil {
		return x.Cause
	}
	return nil
}

type CauseError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId string      `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Message   string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Cause     *CauseError `protobuf:"bytes,3,opt,name=cause,proto3,oneof" json:"cause,omitempty"`
}

func (x *CauseError) Reset() {
	*x = CauseError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CauseError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CauseError) ProtoMessage() {}

func (x *CauseError) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CauseError.ProtoReflect.Descriptor instead.
func (*CauseError) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_endpoint_proto_rawDescGZIP(), []int{5}
}

func (x *CauseError) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *CauseError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CauseError) GetCause() *CauseError {
	if x != nil {
		return x.Cause
	}
	return nil
}

type EndpointDataVoidMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EndpointDataVoidMessage) Reset() {
	*x = EndpointDataVoidMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointDataVoidMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointDataVoidMessage) ProtoMessage() {}

func (x *EndpointDataVoidMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointDataVoidMessage.ProtoReflect.Descriptor instead.
func (*EndpointDataVoidMessage) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_endpoint_proto_rawDescGZIP(), []int{6}
}

type EndpointDataMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type v1.EndpointMessageType `protobuf:"varint,1,opt,name=type,proto3,enum=ras.messages.v1.EndpointMessageType" json:"type,omitempty"`
	Data []byte                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EndpointDataMessage) Reset() {
	*x = EndpointDataMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointDataMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointDataMessage) ProtoMessage() {}

func (x *EndpointDataMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointDataMessage.ProtoReflect.Descriptor instead.
func (*EndpointDataMessage) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_endpoint_proto_rawDescGZIP(), []int{7}
}

func (x *EndpointDataMessage) GetType() v1.EndpointMessageType {
	if x != nil {
		return x.Type
	}
	return v1.EndpointMessageType_GET_AGENT_ADMINS_REQUEST
}

func (x *EndpointDataMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type EndpointFailureMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId string      `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Message   string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Cause     *CauseError `protobuf:"bytes,3,opt,name=cause,proto3,oneof" json:"cause,omitempty"`
}

func (x *EndpointFailureMessage) Reset() {
	*x = EndpointFailureMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointFailureMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointFailureMessage) ProtoMessage() {}

func (x *EndpointFailureMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_endpoint_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointFailureMessage.ProtoReflect.Descriptor instead.
func (*EndpointFailureMessage) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_endpoint_proto_rawDescGZIP(), []int{8}
}

func (x *EndpointFailureMessage) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *EndpointFailureMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EndpointFailureMessage) GetCause() *CauseError {
	if x != nil {
		return x.Cause
	}
	return nil
}

var File_ras_protocol_v1_endpoint_proto protoreflect.FileDescriptor

var file_ras_protocol_v1_endpoint_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16,
	0x72, 0x61, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x61, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x61,
	0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x0c, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x72, 0x61, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x51, 0x0a,
	0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x3a, 0x06, 0x88, 0xf5, 0xea, 0x94, 0x0e, 0x0b, 0x22, 0x8c, 0x02, 0x0a, 0x0f, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x70, 0x6f, 0x69, 0x6f, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x6e, 0x70, 0x6f,
	0x69, 0x6f, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x63, 0x6b, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x51, 0x0a, 0x0b,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72,
	0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a,
	0x06, 0x88, 0xf5, 0xea, 0x94, 0x0e, 0x0c, 0x22, 0xc3, 0x01, 0x0a, 0x0f, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x61, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x06, 0x88, 0xf5, 0xea, 0x94, 0x0e, 0x0e, 0x22, 0xea, 0x02,
	0x0a, 0x12, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x41, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x35, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x72, 0x61,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x76, 0x6f, 0x69, 0x64, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x72, 0x61, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x56, 0x6f, 0x69, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x76, 0x6f, 0x69, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x3a, 0x06, 0x88, 0xf5, 0xea, 0x94,
	0x0e, 0x0e, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x89, 0x02, 0x0a, 0x12, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x41, 0x63,
	0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x05,
	0x63, 0x61, 0x75, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x61,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x75, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x63, 0x61, 0x75, 0x73,
	0x65, 0x88, 0x01, 0x01, 0x3a, 0x06, 0x88, 0xf5, 0xea, 0x94, 0x0e, 0x0f, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x0a, 0x43, 0x61, 0x75, 0x73, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36,
	0x0a, 0x05, 0x63, 0x61, 0x75, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x75, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x63, 0x61,
	0x75, 0x73, 0x65, 0x88, 0x01, 0x01, 0x3a, 0x06, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x00, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x22, 0x21, 0x0a, 0x17, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x56, 0x6f, 0x69, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x3a, 0x06, 0x90, 0xf5, 0xea, 0x94, 0x0e, 0x00, 0x22, 0x6b, 0x0a, 0x13, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x24, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x3a, 0x06, 0x90, 0xf5, 0xea, 0x94, 0x0e, 0x01, 0x22, 0x9c, 0x01, 0x0a, 0x16, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x05,
	0x63, 0x61, 0x75, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x61,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x75, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x63, 0x61, 0x75, 0x73,
	0x65, 0x88, 0x01, 0x01, 0x3a, 0x07, 0x90, 0xf5, 0xea, 0x94, 0x0e, 0xff, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x42, 0xcb, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x42,
	0x0d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x52, 0x50, 0x58, 0xaa, 0x02, 0x0f, 0x52, 0x61, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x52, 0x61, 0x73, 0x5c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x52, 0x61, 0x73, 0x5c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x52, 0x61, 0x73, 0x3a, 0x3a, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x06, 0x08,
	0x01, 0x10, 0x01, 0x18, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ras_protocol_v1_endpoint_proto_rawDescOnce sync.Once
	file_ras_protocol_v1_endpoint_proto_rawDescData = file_ras_protocol_v1_endpoint_proto_rawDesc
)

func file_ras_protocol_v1_endpoint_proto_rawDescGZIP() []byte {
	file_ras_protocol_v1_endpoint_proto_rawDescOnce.Do(func() {
		file_ras_protocol_v1_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_ras_protocol_v1_endpoint_proto_rawDescData)
	})
	return file_ras_protocol_v1_endpoint_proto_rawDescData
}

var file_ras_protocol_v1_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_ras_protocol_v1_endpoint_proto_goTypes = []interface{}{
	(*EndpointOpen)(nil),            // 0: ras.protocol.v1.EndpointOpen
	(*EndpointOpenAck)(nil),         // 1: ras.protocol.v1.EndpointOpenAck
	(*EndpointMessage)(nil),         // 2: ras.protocol.v1.EndpointMessage
	(*EndpointMessageAck)(nil),      // 3: ras.protocol.v1.EndpointMessageAck
	(*EndpointFailureAck)(nil),      // 4: ras.protocol.v1.EndpointFailureAck
	(*CauseError)(nil),              // 5: ras.protocol.v1.CauseError
	(*EndpointDataVoidMessage)(nil), // 6: ras.protocol.v1.EndpointDataVoidMessage
	(*EndpointDataMessage)(nil),     // 7: ras.protocol.v1.EndpointDataMessage
	(*EndpointFailureMessage)(nil),  // 8: ras.protocol.v1.EndpointFailureMessage
	nil,                             // 9: ras.protocol.v1.EndpointOpen.ParamsEntry
	nil,                             // 10: ras.protocol.v1.EndpointOpenAck.ParamsEntry
	(EndpointDataType)(0),           // 11: ras.protocol.v1.EndpointDataType
	(v1.EndpointMessageType)(0),     // 12: ras.messages.v1.EndpointMessageType
	(*Param)(nil),                   // 13: ras.protocol.v1.Param
}
var file_ras_protocol_v1_endpoint_proto_depIdxs = []int32{
	9,  // 0: ras.protocol.v1.EndpointOpen.params:type_name -> ras.protocol.v1.EndpointOpen.ParamsEntry
	10, // 1: ras.protocol.v1.EndpointOpenAck.params:type_name -> ras.protocol.v1.EndpointOpenAck.ParamsEntry
	11, // 2: ras.protocol.v1.EndpointMessage.type:type_name -> ras.protocol.v1.EndpointDataType
	7,  // 3: ras.protocol.v1.EndpointMessage.data:type_name -> ras.protocol.v1.EndpointDataMessage
	11, // 4: ras.protocol.v1.EndpointMessageAck.type:type_name -> ras.protocol.v1.EndpointDataType
	6,  // 5: ras.protocol.v1.EndpointMessageAck.void_message:type_name -> ras.protocol.v1.EndpointDataVoidMessage
	7,  // 6: ras.protocol.v1.EndpointMessageAck.message:type_name -> ras.protocol.v1.EndpointDataMessage
	8,  // 7: ras.protocol.v1.EndpointMessageAck.failure:type_name -> ras.protocol.v1.EndpointFailureMessage
	5,  // 8: ras.protocol.v1.EndpointFailureAck.cause:type_name -> ras.protocol.v1.CauseError
	5,  // 9: ras.protocol.v1.CauseError.cause:type_name -> ras.protocol.v1.CauseError
	12, // 10: ras.protocol.v1.EndpointDataMessage.type:type_name -> ras.messages.v1.EndpointMessageType
	5,  // 11: ras.protocol.v1.EndpointFailureMessage.cause:type_name -> ras.protocol.v1.CauseError
	13, // 12: ras.protocol.v1.EndpointOpen.ParamsEntry.value:type_name -> ras.protocol.v1.Param
	13, // 13: ras.protocol.v1.EndpointOpenAck.ParamsEntry.value:type_name -> ras.protocol.v1.Param
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_ras_protocol_v1_endpoint_proto_init() }
func file_ras_protocol_v1_endpoint_proto_init() {
	if File_ras_protocol_v1_endpoint_proto != nil {
		return
	}
	file_ras_protocol_v1_param_proto_init()
	file_ras_protocol_v1_types_proto_init()
	file_ras_protocol_v1_packet_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ras_protocol_v1_endpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointOpen); i {
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
		file_ras_protocol_v1_endpoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointOpenAck); i {
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
		file_ras_protocol_v1_endpoint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointMessage); i {
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
		file_ras_protocol_v1_endpoint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointMessageAck); i {
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
		file_ras_protocol_v1_endpoint_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointFailureAck); i {
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
		file_ras_protocol_v1_endpoint_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CauseError); i {
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
		file_ras_protocol_v1_endpoint_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointDataVoidMessage); i {
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
		file_ras_protocol_v1_endpoint_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointDataMessage); i {
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
		file_ras_protocol_v1_endpoint_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointFailureMessage); i {
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
	file_ras_protocol_v1_endpoint_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*EndpointMessageAck_VoidMessage)(nil),
		(*EndpointMessageAck_Message)(nil),
		(*EndpointMessageAck_Failure)(nil),
	}
	file_ras_protocol_v1_endpoint_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_ras_protocol_v1_endpoint_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_ras_protocol_v1_endpoint_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ras_protocol_v1_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ras_protocol_v1_endpoint_proto_goTypes,
		DependencyIndexes: file_ras_protocol_v1_endpoint_proto_depIdxs,
		MessageInfos:      file_ras_protocol_v1_endpoint_proto_msgTypes,
	}.Build()
	File_ras_protocol_v1_endpoint_proto = out.File
	file_ras_protocol_v1_endpoint_proto_rawDesc = nil
	file_ras_protocol_v1_endpoint_proto_goTypes = nil
	file_ras_protocol_v1_endpoint_proto_depIdxs = nil
}