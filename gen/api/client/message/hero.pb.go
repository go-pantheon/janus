// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: message/hero.proto

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

// Hero list
type UserHeroListProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Heroes        map[int64]*HeroProto   `protobuf:"bytes,1,rep,name=heroes,proto3" json:"heroes,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` // All heroes
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserHeroListProto) Reset() {
	*x = UserHeroListProto{}
	mi := &file_message_hero_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserHeroListProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserHeroListProto) ProtoMessage() {}

func (x *UserHeroListProto) ProtoReflect() protoreflect.Message {
	mi := &file_message_hero_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserHeroListProto.ProtoReflect.Descriptor instead.
func (*UserHeroListProto) Descriptor() ([]byte, []int) {
	return file_message_hero_proto_rawDescGZIP(), []int{0}
}

func (x *UserHeroListProto) GetHeroes() map[int64]*HeroProto {
	if x != nil {
		return x.Heroes
	}
	return nil
}

// Hero
type HeroProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                       // Hero ID
	DataId        int64                  `protobuf:"varint,2,opt,name=data_id,json=dataId,proto3" json:"data_id,omitempty"` // Hero DataID
	Level         uint64                 `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`                 // Hero level
	Equips        []int64                `protobuf:"varint,4,rep,packed,name=equips,proto3" json:"equips,omitempty"`        // Equip IDs
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HeroProto) Reset() {
	*x = HeroProto{}
	mi := &file_message_hero_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeroProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeroProto) ProtoMessage() {}

func (x *HeroProto) ProtoReflect() protoreflect.Message {
	mi := &file_message_hero_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeroProto.ProtoReflect.Descriptor instead.
func (*HeroProto) Descriptor() ([]byte, []int) {
	return file_message_hero_proto_rawDescGZIP(), []int{1}
}

func (x *HeroProto) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HeroProto) GetDataId() int64 {
	if x != nil {
		return x.DataId
	}
	return 0
}

func (x *HeroProto) GetLevel() uint64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *HeroProto) GetEquips() []int64 {
	if x != nil {
		return x.Equips
	}
	return nil
}

var File_message_hero_proto protoreflect.FileDescriptor

var file_message_hero_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa2, 0x01,
	0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x48, 0x65, 0x72, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x3e, 0x0a, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x48, 0x65, 0x72, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x48, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x72,
	0x6f, 0x65, 0x73, 0x1a, 0x4d, 0x0a, 0x0b, 0x48, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x48, 0x65,
	0x72, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x62, 0x0a, 0x09, 0x48, 0x65, 0x72, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x64, 0x61, 0x74, 0x61, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x71, 0x75, 0x69, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06,
	0x65, 0x71, 0x75, 0x69, 0x70, 0x73, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3b, 0x63, 0x6c, 0x69,
	0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_message_hero_proto_rawDescOnce sync.Once
	file_message_hero_proto_rawDescData []byte
)

func file_message_hero_proto_rawDescGZIP() []byte {
	file_message_hero_proto_rawDescOnce.Do(func() {
		file_message_hero_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_message_hero_proto_rawDesc), len(file_message_hero_proto_rawDesc)))
	})
	return file_message_hero_proto_rawDescData
}

var file_message_hero_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_message_hero_proto_goTypes = []any{
	(*UserHeroListProto)(nil), // 0: message.UserHeroListProto
	(*HeroProto)(nil),         // 1: message.HeroProto
	nil,                       // 2: message.UserHeroListProto.HeroesEntry
}
var file_message_hero_proto_depIdxs = []int32{
	2, // 0: message.UserHeroListProto.heroes:type_name -> message.UserHeroListProto.HeroesEntry
	1, // 1: message.UserHeroListProto.HeroesEntry.value:type_name -> message.HeroProto
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_hero_proto_init() }
func file_message_hero_proto_init() {
	if File_message_hero_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_message_hero_proto_rawDesc), len(file_message_hero_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_hero_proto_goTypes,
		DependencyIndexes: file_message_hero_proto_depIdxs,
		MessageInfos:      file_message_hero_proto_msgTypes,
	}.Build()
	File_message_hero_proto = out.File
	file_message_hero_proto_goTypes = nil
	file_message_hero_proto_depIdxs = nil
}
