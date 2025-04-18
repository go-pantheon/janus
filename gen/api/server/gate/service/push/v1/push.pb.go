// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: gate/service/push/v1/push.proto

package servicev1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type PushRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           int64                  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Bodies        []*PushBody            `protobuf:"bytes,2,rep,name=bodies,proto3" json:"bodies,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	mi := &file_gate_service_push_v1_push_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gate_service_push_v1_push_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_gate_service_push_v1_push_proto_rawDescGZIP(), []int{0}
}

func (x *PushRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PushRequest) GetBodies() []*PushBody {
	if x != nil {
		return x.Bodies
	}
	return nil
}

type PushResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PushResponse) Reset() {
	*x = PushResponse{}
	mi := &file_gate_service_push_v1_push_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResponse) ProtoMessage() {}

func (x *PushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gate_service_push_v1_push_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResponse.ProtoReflect.Descriptor instead.
func (*PushResponse) Descriptor() ([]byte, []int) {
	return file_gate_service_push_v1_push_proto_rawDescGZIP(), []int{1}
}

type MulticastRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           []int64                `protobuf:"varint,1,rep,packed,name=uid,proto3" json:"uid,omitempty"`
	Bodies        []*PushBody            `protobuf:"bytes,2,rep,name=bodies,proto3" json:"bodies,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MulticastRequest) Reset() {
	*x = MulticastRequest{}
	mi := &file_gate_service_push_v1_push_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MulticastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MulticastRequest) ProtoMessage() {}

func (x *MulticastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gate_service_push_v1_push_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MulticastRequest.ProtoReflect.Descriptor instead.
func (*MulticastRequest) Descriptor() ([]byte, []int) {
	return file_gate_service_push_v1_push_proto_rawDescGZIP(), []int{2}
}

func (x *MulticastRequest) GetUid() []int64 {
	if x != nil {
		return x.Uid
	}
	return nil
}

func (x *MulticastRequest) GetBodies() []*PushBody {
	if x != nil {
		return x.Bodies
	}
	return nil
}

type MulticastResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MulticastResponse) Reset() {
	*x = MulticastResponse{}
	mi := &file_gate_service_push_v1_push_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MulticastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MulticastResponse) ProtoMessage() {}

func (x *MulticastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gate_service_push_v1_push_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MulticastResponse.ProtoReflect.Descriptor instead.
func (*MulticastResponse) Descriptor() ([]byte, []int) {
	return file_gate_service_push_v1_push_proto_rawDescGZIP(), []int{3}
}

type BroadcastRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bodies        []*PushBody            `protobuf:"bytes,1,rep,name=bodies,proto3" json:"bodies,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BroadcastRequest) Reset() {
	*x = BroadcastRequest{}
	mi := &file_gate_service_push_v1_push_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRequest) ProtoMessage() {}

func (x *BroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gate_service_push_v1_push_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRequest.ProtoReflect.Descriptor instead.
func (*BroadcastRequest) Descriptor() ([]byte, []int) {
	return file_gate_service_push_v1_push_proto_rawDescGZIP(), []int{4}
}

func (x *BroadcastRequest) GetBodies() []*PushBody {
	if x != nil {
		return x.Bodies
	}
	return nil
}

type BroadcastResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BroadcastResponse) Reset() {
	*x = BroadcastResponse{}
	mi := &file_gate_service_push_v1_push_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BroadcastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastResponse) ProtoMessage() {}

func (x *BroadcastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gate_service_push_v1_push_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastResponse.ProtoReflect.Descriptor instead.
func (*BroadcastResponse) Descriptor() ([]byte, []int) {
	return file_gate_service_push_v1_push_proto_rawDescGZIP(), []int{5}
}

type PushBody struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"` // Message body proto bytes array
	Obj           int64                  `protobuf:"varint,2,opt,name=obj,proto3" json:"obj,omitempty"`  // Object ID, according to the business agreement
	Mod           int32                  `protobuf:"varint,3,opt,name=mod,proto3" json:"mod,omitempty"`  // Module ID, globally unique
	Seq           int32                  `protobuf:"varint,4,opt,name=seq,proto3" json:"seq,omitempty"`  // Message ID within the module, unique within the module
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PushBody) Reset() {
	*x = PushBody{}
	mi := &file_gate_service_push_v1_push_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushBody) ProtoMessage() {}

func (x *PushBody) ProtoReflect() protoreflect.Message {
	mi := &file_gate_service_push_v1_push_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushBody.ProtoReflect.Descriptor instead.
func (*PushBody) Descriptor() ([]byte, []int) {
	return file_gate_service_push_v1_push_proto_rawDescGZIP(), []int{6}
}

func (x *PushBody) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PushBody) GetObj() int64 {
	if x != nil {
		return x.Obj
	}
	return 0
}

func (x *PushBody) GetMod() int32 {
	if x != nil {
		return x.Mod
	}
	return 0
}

func (x *PushBody) GetSeq() int32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

var File_gate_service_push_v1_push_proto protoreflect.FileDescriptor

var file_gate_service_push_v1_push_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x75, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x62, 0x6f, 0x64, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x06, 0x62, 0x6f, 0x64, 0x69, 0x65, 0x73, 0x22, 0x0e,
	0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c,
	0x0a, 0x10, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x62, 0x6f, 0x64, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68,
	0x42, 0x6f, 0x64, 0x79, 0x52, 0x06, 0x62, 0x6f, 0x64, 0x69, 0x65, 0x73, 0x22, 0x13, 0x0a, 0x11,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4a, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x62, 0x6f, 0x64, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73,
	0x68, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x06, 0x62, 0x6f, 0x64, 0x69, 0x65, 0x73, 0x22, 0x13, 0x0a,
	0x11, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x54, 0x0a, 0x08, 0x50, 0x75, 0x73, 0x68, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6f, 0x62, 0x6a, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6d, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x71, 0x32, 0xd8, 0x02, 0x0a, 0x0b, 0x50, 0x75, 0x73,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68,
	0x12, 0x21, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x3a,
	0x01, 0x2a, 0x22, 0x05, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x12, 0x73, 0x0a, 0x09, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x75,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a,
	0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x12, 0x73,
	0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x75, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_gate_service_push_v1_push_proto_rawDescOnce sync.Once
	file_gate_service_push_v1_push_proto_rawDescData []byte
)

func file_gate_service_push_v1_push_proto_rawDescGZIP() []byte {
	file_gate_service_push_v1_push_proto_rawDescOnce.Do(func() {
		file_gate_service_push_v1_push_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_gate_service_push_v1_push_proto_rawDesc), len(file_gate_service_push_v1_push_proto_rawDesc)))
	})
	return file_gate_service_push_v1_push_proto_rawDescData
}

var file_gate_service_push_v1_push_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_gate_service_push_v1_push_proto_goTypes = []any{
	(*PushRequest)(nil),       // 0: gate.service.push.v1.PushRequest
	(*PushResponse)(nil),      // 1: gate.service.push.v1.PushResponse
	(*MulticastRequest)(nil),  // 2: gate.service.push.v1.MulticastRequest
	(*MulticastResponse)(nil), // 3: gate.service.push.v1.MulticastResponse
	(*BroadcastRequest)(nil),  // 4: gate.service.push.v1.BroadcastRequest
	(*BroadcastResponse)(nil), // 5: gate.service.push.v1.BroadcastResponse
	(*PushBody)(nil),          // 6: gate.service.push.v1.PushBody
}
var file_gate_service_push_v1_push_proto_depIdxs = []int32{
	6, // 0: gate.service.push.v1.PushRequest.bodies:type_name -> gate.service.push.v1.PushBody
	6, // 1: gate.service.push.v1.MulticastRequest.bodies:type_name -> gate.service.push.v1.PushBody
	6, // 2: gate.service.push.v1.BroadcastRequest.bodies:type_name -> gate.service.push.v1.PushBody
	0, // 3: gate.service.push.v1.PushService.Push:input_type -> gate.service.push.v1.PushRequest
	2, // 4: gate.service.push.v1.PushService.Multicast:input_type -> gate.service.push.v1.MulticastRequest
	4, // 5: gate.service.push.v1.PushService.Broadcast:input_type -> gate.service.push.v1.BroadcastRequest
	1, // 6: gate.service.push.v1.PushService.Push:output_type -> gate.service.push.v1.PushResponse
	3, // 7: gate.service.push.v1.PushService.Multicast:output_type -> gate.service.push.v1.MulticastResponse
	5, // 8: gate.service.push.v1.PushService.Broadcast:output_type -> gate.service.push.v1.BroadcastResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gate_service_push_v1_push_proto_init() }
func file_gate_service_push_v1_push_proto_init() {
	if File_gate_service_push_v1_push_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_gate_service_push_v1_push_proto_rawDesc), len(file_gate_service_push_v1_push_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gate_service_push_v1_push_proto_goTypes,
		DependencyIndexes: file_gate_service_push_v1_push_proto_depIdxs,
		MessageInfos:      file_gate_service_push_v1_push_proto_msgTypes,
	}.Build()
	File_gate_service_push_v1_push_proto = out.File
	file_gate_service_push_v1_push_proto_goTypes = nil
	file_gate_service_push_v1_push_proto_depIdxs = nil
}
