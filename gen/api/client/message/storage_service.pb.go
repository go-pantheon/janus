// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: message/storage_service.proto

package climsg

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

type SCUsePack_Code int32

const (
	SCUsePack_ErrUnspecified  SCUsePack_Code = 0 // Please try again later
	SCUsePack_Success         SCUsePack_Code = 1 // Success
	SCUsePack_ErrPackNotExist SCUsePack_Code = 2 // Pack does not exist
	SCUsePack_ErrPackLimit    SCUsePack_Code = 3 // Usage restricted
)

// Enum value maps for SCUsePack_Code.
var (
	SCUsePack_Code_name = map[int32]string{
		0: "ErrUnspecified",
		1: "Success",
		2: "ErrPackNotExist",
		3: "ErrPackLimit",
	}
	SCUsePack_Code_value = map[string]int32{
		"ErrUnspecified":  0,
		"Success":         1,
		"ErrPackNotExist": 2,
		"ErrPackLimit":    3,
	}
)

func (x SCUsePack_Code) Enum() *SCUsePack_Code {
	p := new(SCUsePack_Code)
	*p = x
	return p
}

func (x SCUsePack_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SCUsePack_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_message_storage_service_proto_enumTypes[0].Descriptor()
}

func (SCUsePack_Code) Type() protoreflect.EnumType {
	return &file_message_storage_service_proto_enumTypes[0]
}

func (x SCUsePack_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SCUsePack_Code.Descriptor instead.
func (SCUsePack_Code) EnumDescriptor() ([]byte, []int) {
	return file_message_storage_service_proto_rawDescGZIP(), []int{2, 0}
}

// @push Item has been updated
type SCPushItemUpdated struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         map[int64]uint64       `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"` // Items with changed amounts. item DataID -> current amount. value is 0 means the item has been deleted
	Packs         map[int64]uint64       `protobuf:"bytes,2,rep,name=packs,proto3" json:"packs,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"` // Packs with changed amounts. pack DataID -> current amount. value is 0 means the pack has been deleted
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SCPushItemUpdated) Reset() {
	*x = SCPushItemUpdated{}
	mi := &file_message_storage_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SCPushItemUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCPushItemUpdated) ProtoMessage() {}

func (x *SCPushItemUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_message_storage_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCPushItemUpdated.ProtoReflect.Descriptor instead.
func (*SCPushItemUpdated) Descriptor() ([]byte, []int) {
	return file_message_storage_service_proto_rawDescGZIP(), []int{0}
}

func (x *SCPushItemUpdated) GetItems() map[int64]uint64 {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SCPushItemUpdated) GetPacks() map[int64]uint64 {
	if x != nil {
		return x.Packs
	}
	return nil
}

// Request: Use pack
type CSUsePack struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // Pack ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CSUsePack) Reset() {
	*x = CSUsePack{}
	mi := &file_message_storage_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CSUsePack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSUsePack) ProtoMessage() {}

func (x *CSUsePack) ProtoReflect() protoreflect.Message {
	mi := &file_message_storage_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSUsePack.ProtoReflect.Descriptor instead.
func (*CSUsePack) Descriptor() ([]byte, []int) {
	return file_message_storage_service_proto_rawDescGZIP(), []int{1}
}

func (x *CSUsePack) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Response: Use pack
type SCUsePack struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          SCUsePack_Code         `protobuf:"varint,1,opt,name=code,proto3,enum=message.SCUsePack_Code" json:"code,omitempty"`
	Id            int64                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`                                                                                    // Pack ID
	Prizes        map[int64]int64        `protobuf:"bytes,3,rep,name=prizes,proto3" json:"prizes,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"` // the prizes that the player has received
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SCUsePack) Reset() {
	*x = SCUsePack{}
	mi := &file_message_storage_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SCUsePack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCUsePack) ProtoMessage() {}

func (x *SCUsePack) ProtoReflect() protoreflect.Message {
	mi := &file_message_storage_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCUsePack.ProtoReflect.Descriptor instead.
func (*SCUsePack) Descriptor() ([]byte, []int) {
	return file_message_storage_service_proto_rawDescGZIP(), []int{2}
}

func (x *SCUsePack) GetCode() SCUsePack_Code {
	if x != nil {
		return x.Code
	}
	return SCUsePack_ErrUnspecified
}

func (x *SCUsePack) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SCUsePack) GetPrizes() map[int64]int64 {
	if x != nil {
		return x.Prizes
	}
	return nil
}

var File_message_storage_service_proto protoreflect.FileDescriptor

var file_message_storage_service_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x02,
	0x0a, 0x11, 0x53, 0x43, 0x50, 0x75, 0x73, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x43, 0x50,
	0x75, 0x73, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x3b, 0x0a, 0x05, 0x70, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x43, 0x50, 0x75, 0x73, 0x68,
	0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x50, 0x61, 0x63, 0x6b,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x61, 0x63, 0x6b, 0x73, 0x1a, 0x38, 0x0a,
	0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x38, 0x0a, 0x0a, 0x50, 0x61, 0x63, 0x6b, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x1b, 0x0a, 0x09, 0x43, 0x53, 0x55, 0x73, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8b,
	0x02, 0x0a, 0x09, 0x53, 0x43, 0x55, 0x73, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x12, 0x2b, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x43, 0x55, 0x73, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x70, 0x72, 0x69,
	0x7a, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x53, 0x43, 0x55, 0x73, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x2e, 0x50, 0x72,
	0x69, 0x7a, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x72, 0x69, 0x7a, 0x65,
	0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x7a, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4e, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x72, 0x72, 0x55, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x72, 0x72, 0x50, 0x61, 0x63, 0x6b,
	0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x72,
	0x72, 0x50, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x10, 0x03, 0x32, 0x61, 0x0a, 0x0e,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f,
	0x0a, 0x07, 0x55, 0x73, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x12, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x43, 0x53, 0x55, 0x73, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x1a, 0x12, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x43, 0x55, 0x73, 0x65, 0x50, 0x61, 0x63,
	0x6b, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x42,
	0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x3b, 0x63, 0x6c, 0x69, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_message_storage_service_proto_rawDescOnce sync.Once
	file_message_storage_service_proto_rawDescData []byte
)

func file_message_storage_service_proto_rawDescGZIP() []byte {
	file_message_storage_service_proto_rawDescOnce.Do(func() {
		file_message_storage_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_message_storage_service_proto_rawDesc), len(file_message_storage_service_proto_rawDesc)))
	})
	return file_message_storage_service_proto_rawDescData
}

var file_message_storage_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_storage_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_message_storage_service_proto_goTypes = []any{
	(SCUsePack_Code)(0),       // 0: message.SCUsePack.Code
	(*SCPushItemUpdated)(nil), // 1: message.SCPushItemUpdated
	(*CSUsePack)(nil),         // 2: message.CSUsePack
	(*SCUsePack)(nil),         // 3: message.SCUsePack
	nil,                       // 4: message.SCPushItemUpdated.ItemsEntry
	nil,                       // 5: message.SCPushItemUpdated.PacksEntry
	nil,                       // 6: message.SCUsePack.PrizesEntry
}
var file_message_storage_service_proto_depIdxs = []int32{
	4, // 0: message.SCPushItemUpdated.items:type_name -> message.SCPushItemUpdated.ItemsEntry
	5, // 1: message.SCPushItemUpdated.packs:type_name -> message.SCPushItemUpdated.PacksEntry
	0, // 2: message.SCUsePack.code:type_name -> message.SCUsePack.Code
	6, // 3: message.SCUsePack.prizes:type_name -> message.SCUsePack.PrizesEntry
	2, // 4: message.StorageService.UsePack:input_type -> message.CSUsePack
	3, // 5: message.StorageService.UsePack:output_type -> message.SCUsePack
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_message_storage_service_proto_init() }
func file_message_storage_service_proto_init() {
	if File_message_storage_service_proto != nil {
		return
	}
	file_message_storage_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_message_storage_service_proto_rawDesc), len(file_message_storage_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_storage_service_proto_goTypes,
		DependencyIndexes: file_message_storage_service_proto_depIdxs,
		EnumInfos:         file_message_storage_service_proto_enumTypes,
		MessageInfos:      file_message_storage_service_proto_msgTypes,
	}.Build()
	File_message_storage_service_proto = out.File
	file_message_storage_service_proto_goTypes = nil
	file_message_storage_service_proto_depIdxs = nil
}
