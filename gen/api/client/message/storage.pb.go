// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: message/storage.proto

package climsg

import (
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

// Storage
type UserStorageProto struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	Items         map[int64]uint64                 `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`                                     // Items. item DataID -> amount
	Packs         map[int64]uint64                 `protobuf:"bytes,2,rep,name=packs,proto3" json:"packs,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`                                     // Packs. pack DataID -> amount
	RecoveryInfos map[int64]*ItemRecoveryInfoProto `protobuf:"bytes,3,rep,name=recovery_infos,json=recoveryInfos,proto3" json:"recovery_infos,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` // Recoverable item infos. item DataID -> properties
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserStorageProto) Reset() {
	*x = UserStorageProto{}
	mi := &file_message_storage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserStorageProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStorageProto) ProtoMessage() {}

func (x *UserStorageProto) ProtoReflect() protoreflect.Message {
	mi := &file_message_storage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStorageProto.ProtoReflect.Descriptor instead.
func (*UserStorageProto) Descriptor() ([]byte, []int) {
	return file_message_storage_proto_rawDescGZIP(), []int{0}
}

func (x *UserStorageProto) GetItems() map[int64]uint64 {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *UserStorageProto) GetPacks() map[int64]uint64 {
	if x != nil {
		return x.Packs
	}
	return nil
}

func (x *UserStorageProto) GetRecoveryInfos() map[int64]*ItemRecoveryInfoProto {
	if x != nil {
		return x.RecoveryInfos
	}
	return nil
}

// Item Recovery info
type ItemRecoveryInfoProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DataId        int64                  `protobuf:"varint,1,opt,name=data_id,json=dataId,proto3" json:"data_id,omitempty"`                         // Item DataID
	Max           int64                  `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`                                             // Max Amount
	RecoverPerSec float32                `protobuf:"fixed32,3,opt,name=recover_per_sec,json=recoverPerSec,proto3" json:"recover_per_sec,omitempty"` // Recover Per Second
	UpdatedAt     int64                  `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`                // Last updated time (Unix timestamp)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ItemRecoveryInfoProto) Reset() {
	*x = ItemRecoveryInfoProto{}
	mi := &file_message_storage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ItemRecoveryInfoProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemRecoveryInfoProto) ProtoMessage() {}

func (x *ItemRecoveryInfoProto) ProtoReflect() protoreflect.Message {
	mi := &file_message_storage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemRecoveryInfoProto.ProtoReflect.Descriptor instead.
func (*ItemRecoveryInfoProto) Descriptor() ([]byte, []int) {
	return file_message_storage_proto_rawDescGZIP(), []int{1}
}

func (x *ItemRecoveryInfoProto) GetDataId() int64 {
	if x != nil {
		return x.DataId
	}
	return 0
}

func (x *ItemRecoveryInfoProto) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *ItemRecoveryInfoProto) GetRecoverPerSec() float32 {
	if x != nil {
		return x.RecoverPerSec
	}
	return 0
}

func (x *ItemRecoveryInfoProto) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_message_storage_proto protoreflect.FileDescriptor

var file_message_storage_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0xb5, 0x03, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x3a, 0x0a, 0x05, 0x70, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x63, 0x6b,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x53, 0x0a,
	0x0e, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x38, 0x0a, 0x0a,
	0x50, 0x61, 0x63, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x60, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x89, 0x01, 0x0a, 0x15, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x61, 0x74, 0x61, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x26, 0x0a,
	0x0f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x50,
	0x65, 0x72, 0x53, 0x65, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3b, 0x63, 0x6c, 0x69, 0x6d, 0x73,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_message_storage_proto_rawDescOnce sync.Once
	file_message_storage_proto_rawDescData []byte
)

func file_message_storage_proto_rawDescGZIP() []byte {
	file_message_storage_proto_rawDescOnce.Do(func() {
		file_message_storage_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_message_storage_proto_rawDesc), len(file_message_storage_proto_rawDesc)))
	})
	return file_message_storage_proto_rawDescData
}

var file_message_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_message_storage_proto_goTypes = []any{
	(*UserStorageProto)(nil),      // 0: message.UserStorageProto
	(*ItemRecoveryInfoProto)(nil), // 1: message.ItemRecoveryInfoProto
	nil,                           // 2: message.UserStorageProto.ItemsEntry
	nil,                           // 3: message.UserStorageProto.PacksEntry
	nil,                           // 4: message.UserStorageProto.RecoveryInfosEntry
}
var file_message_storage_proto_depIdxs = []int32{
	2, // 0: message.UserStorageProto.items:type_name -> message.UserStorageProto.ItemsEntry
	3, // 1: message.UserStorageProto.packs:type_name -> message.UserStorageProto.PacksEntry
	4, // 2: message.UserStorageProto.recovery_infos:type_name -> message.UserStorageProto.RecoveryInfosEntry
	1, // 3: message.UserStorageProto.RecoveryInfosEntry.value:type_name -> message.ItemRecoveryInfoProto
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_message_storage_proto_init() }
func file_message_storage_proto_init() {
	if File_message_storage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_message_storage_proto_rawDesc), len(file_message_storage_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_storage_proto_goTypes,
		DependencyIndexes: file_message_storage_proto_depIdxs,
		MessageInfos:      file_message_storage_proto_msgTypes,
	}.Build()
	File_message_storage_proto = out.File
	file_message_storage_proto_goTypes = nil
	file_message_storage_proto_depIdxs = nil
}
