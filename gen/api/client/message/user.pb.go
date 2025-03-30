// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: message/user.proto

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

// User data
type UserProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Basic         *UserBasicProto        `protobuf:"bytes,1,opt,name=basic,proto3" json:"basic,omitempty"`                       // Basic information
	Storage       *UserStorageProto      `protobuf:"bytes,2,opt,name=storage,proto3" json:"storage,omitempty"`                   // Item storage
	HeroList      *UserHeroListProto     `protobuf:"bytes,3,opt,name=hero_list,json=heroList,proto3" json:"hero_list,omitempty"` // Hero list
	Room          *UserRoomProto         `protobuf:"bytes,4,opt,name=room,proto3" json:"room,omitempty"`                         // Room
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserProto) Reset() {
	*x = UserProto{}
	mi := &file_message_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProto) ProtoMessage() {}

func (x *UserProto) ProtoReflect() protoreflect.Message {
	mi := &file_message_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProto.ProtoReflect.Descriptor instead.
func (*UserProto) Descriptor() ([]byte, []int) {
	return file_message_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserProto) GetBasic() *UserBasicProto {
	if x != nil {
		return x.Basic
	}
	return nil
}

func (x *UserProto) GetStorage() *UserStorageProto {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *UserProto) GetHeroList() *UserHeroListProto {
	if x != nil {
		return x.HeroList
	}
	return nil
}

func (x *UserProto) GetRoom() *UserRoomProto {
	if x != nil {
		return x.Room
	}
	return nil
}

// Recharge parameter
type RechargeParamProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     int64                  `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"` // Product ID in the platform
	Store         string                 `protobuf:"bytes,2,opt,name=store,proto3" json:"store,omitempty"`                           // Store
	Payload       []byte                 `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`                       // Payload
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RechargeParamProto) Reset() {
	*x = RechargeParamProto{}
	mi := &file_message_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RechargeParamProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RechargeParamProto) ProtoMessage() {}

func (x *RechargeParamProto) ProtoReflect() protoreflect.Message {
	mi := &file_message_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RechargeParamProto.ProtoReflect.Descriptor instead.
func (*RechargeParamProto) Descriptor() ([]byte, []int) {
	return file_message_user_proto_rawDescGZIP(), []int{1}
}

func (x *RechargeParamProto) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *RechargeParamProto) GetStore() string {
	if x != nil {
		return x.Store
	}
	return ""
}

func (x *RechargeParamProto) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_message_user_proto protoreflect.FileDescriptor

var file_message_user_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x6f, 0x6f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x0a, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x68, 0x65, 0x72, 0x6f,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x48, 0x65, 0x72, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x08, 0x68, 0x65, 0x72, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x63, 0x0a,
	0x12, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3b, 0x63, 0x6c, 0x69, 0x6d, 0x73, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_message_user_proto_rawDescOnce sync.Once
	file_message_user_proto_rawDescData []byte
)

func file_message_user_proto_rawDescGZIP() []byte {
	file_message_user_proto_rawDescOnce.Do(func() {
		file_message_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_message_user_proto_rawDesc), len(file_message_user_proto_rawDesc)))
	})
	return file_message_user_proto_rawDescData
}

var file_message_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_user_proto_goTypes = []any{
	(*UserProto)(nil),          // 0: message.UserProto
	(*RechargeParamProto)(nil), // 1: message.RechargeParamProto
	(*UserBasicProto)(nil),     // 2: message.UserBasicProto
	(*UserStorageProto)(nil),   // 3: message.UserStorageProto
	(*UserHeroListProto)(nil),  // 4: message.UserHeroListProto
	(*UserRoomProto)(nil),      // 5: message.UserRoomProto
}
var file_message_user_proto_depIdxs = []int32{
	2, // 0: message.UserProto.basic:type_name -> message.UserBasicProto
	3, // 1: message.UserProto.storage:type_name -> message.UserStorageProto
	4, // 2: message.UserProto.hero_list:type_name -> message.UserHeroListProto
	5, // 3: message.UserProto.room:type_name -> message.UserRoomProto
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_message_user_proto_init() }
func file_message_user_proto_init() {
	if File_message_user_proto != nil {
		return
	}
	file_message_hero_proto_init()
	file_message_user_basic_proto_init()
	file_message_storage_proto_init()
	file_message_room_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_message_user_proto_rawDesc), len(file_message_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_user_proto_goTypes,
		DependencyIndexes: file_message_user_proto_depIdxs,
		MessageInfos:      file_message_user_proto_msgTypes,
	}.Build()
	File_message_user_proto = out.File
	file_message_user_proto_goTypes = nil
	file_message_user_proto_depIdxs = nil
}
