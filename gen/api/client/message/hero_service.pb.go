// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: message/hero_service.proto

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

type SCHeroLevelUpgrade_Code int32

const (
	SCHeroLevelUpgrade_ErrUnspecified   SCHeroLevelUpgrade_Code = 0 // Please try again later
	SCHeroLevelUpgrade_Success          SCHeroLevelUpgrade_Code = 1 // Success
	SCHeroLevelUpgrade_ErrHeroNotExist  SCHeroLevelUpgrade_Code = 2 // Hero does not exist
	SCHeroLevelUpgrade_ErrMaxLevel      SCHeroLevelUpgrade_Code = 3 // Already max level
	SCHeroLevelUpgrade_ErrCostNotEnough SCHeroLevelUpgrade_Code = 4 // Cost not enough
)

// Enum value maps for SCHeroLevelUpgrade_Code.
var (
	SCHeroLevelUpgrade_Code_name = map[int32]string{
		0: "ErrUnspecified",
		1: "Success",
		2: "ErrHeroNotExist",
		3: "ErrMaxLevel",
		4: "ErrCostNotEnough",
	}
	SCHeroLevelUpgrade_Code_value = map[string]int32{
		"ErrUnspecified":   0,
		"Success":          1,
		"ErrHeroNotExist":  2,
		"ErrMaxLevel":      3,
		"ErrCostNotEnough": 4,
	}
)

func (x SCHeroLevelUpgrade_Code) Enum() *SCHeroLevelUpgrade_Code {
	p := new(SCHeroLevelUpgrade_Code)
	*p = x
	return p
}

func (x SCHeroLevelUpgrade_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SCHeroLevelUpgrade_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_message_hero_service_proto_enumTypes[0].Descriptor()
}

func (SCHeroLevelUpgrade_Code) Type() protoreflect.EnumType {
	return &file_message_hero_service_proto_enumTypes[0]
}

func (x SCHeroLevelUpgrade_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SCHeroLevelUpgrade_Code.Descriptor instead.
func (SCHeroLevelUpgrade_Code) EnumDescriptor() ([]byte, []int) {
	return file_message_hero_service_proto_rawDescGZIP(), []int{2, 0}
}

// @push New Hero unlocked
type SCPushHeroUnlock struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Heroes        []*HeroProto           `protobuf:"bytes,1,rep,name=heroes,proto3" json:"heroes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SCPushHeroUnlock) Reset() {
	*x = SCPushHeroUnlock{}
	mi := &file_message_hero_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SCPushHeroUnlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCPushHeroUnlock) ProtoMessage() {}

func (x *SCPushHeroUnlock) ProtoReflect() protoreflect.Message {
	mi := &file_message_hero_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCPushHeroUnlock.ProtoReflect.Descriptor instead.
func (*SCPushHeroUnlock) Descriptor() ([]byte, []int) {
	return file_message_hero_service_proto_rawDescGZIP(), []int{0}
}

func (x *SCPushHeroUnlock) GetHeroes() []*HeroProto {
	if x != nil {
		return x.Heroes
	}
	return nil
}

// Request: Hero level upgrade
type CSHeroLevelUpgrade struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HeroId        int64                  `protobuf:"varint,1,opt,name=hero_id,json=heroId,proto3" json:"hero_id,omitempty"` // Hero_id
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CSHeroLevelUpgrade) Reset() {
	*x = CSHeroLevelUpgrade{}
	mi := &file_message_hero_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CSHeroLevelUpgrade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSHeroLevelUpgrade) ProtoMessage() {}

func (x *CSHeroLevelUpgrade) ProtoReflect() protoreflect.Message {
	mi := &file_message_hero_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSHeroLevelUpgrade.ProtoReflect.Descriptor instead.
func (*CSHeroLevelUpgrade) Descriptor() ([]byte, []int) {
	return file_message_hero_service_proto_rawDescGZIP(), []int{1}
}

func (x *CSHeroLevelUpgrade) GetHeroId() int64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

// Response: Hero level upgrade
type SCHeroLevelUpgrade struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Code          SCHeroLevelUpgrade_Code `protobuf:"varint,1,opt,name=code,proto3,enum=message.SCHeroLevelUpgrade_Code" json:"code,omitempty"`                                         // Status code
	Hero          *HeroProto              `protobuf:"bytes,2,opt,name=hero,proto3" json:"hero,omitempty"`                                                                               // Hero information
	Costs         map[int64]int64         `protobuf:"bytes,3,rep,name=costs,proto3" json:"costs,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"` // Consumed materials
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SCHeroLevelUpgrade) Reset() {
	*x = SCHeroLevelUpgrade{}
	mi := &file_message_hero_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SCHeroLevelUpgrade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCHeroLevelUpgrade) ProtoMessage() {}

func (x *SCHeroLevelUpgrade) ProtoReflect() protoreflect.Message {
	mi := &file_message_hero_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCHeroLevelUpgrade.ProtoReflect.Descriptor instead.
func (*SCHeroLevelUpgrade) Descriptor() ([]byte, []int) {
	return file_message_hero_service_proto_rawDescGZIP(), []int{2}
}

func (x *SCHeroLevelUpgrade) GetCode() SCHeroLevelUpgrade_Code {
	if x != nil {
		return x.Code
	}
	return SCHeroLevelUpgrade_ErrUnspecified
}

func (x *SCHeroLevelUpgrade) GetHero() *HeroProto {
	if x != nil {
		return x.Hero
	}
	return nil
}

func (x *SCHeroLevelUpgrade) GetCosts() map[int64]int64 {
	if x != nil {
		return x.Costs
	}
	return nil
}

var File_message_hero_service_proto protoreflect.FileDescriptor

var file_message_hero_service_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x68, 0x65, 0x72,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x10, 0x53, 0x43, 0x50, 0x75, 0x73,
	0x68, 0x48, 0x65, 0x72, 0x6f, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2a, 0x0a, 0x06, 0x68,
	0x65, 0x72, 0x6f, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52,
	0x06, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x22, 0x2d, 0x0a, 0x12, 0x43, 0x53, 0x48, 0x65, 0x72,
	0x6f, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x68, 0x65, 0x72, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x68, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x22, 0xcf, 0x02, 0x0a, 0x12, 0x53, 0x43, 0x48, 0x65, 0x72,
	0x6f, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x34, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x43, 0x48, 0x65, 0x72, 0x6f, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x68, 0x65, 0x72, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x48, 0x65, 0x72, 0x6f,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x04, 0x68, 0x65, 0x72, 0x6f, 0x12, 0x3c, 0x0a, 0x05, 0x63,
	0x6f, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x43, 0x48, 0x65, 0x72, 0x6f, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x63, 0x6f, 0x73, 0x74, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x43, 0x6f, 0x73,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x63, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x45,
	0x72, 0x72, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x45, 0x72, 0x72, 0x48, 0x65, 0x72, 0x6f, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10,
	0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x4d, 0x61, 0x78, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x4e, 0x6f, 0x74,
	0x45, 0x6e, 0x6f, 0x75, 0x67, 0x68, 0x10, 0x04, 0x32, 0x7b, 0x0a, 0x0b, 0x48, 0x65, 0x72, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x10, 0x48, 0x65, 0x72, 0x6f, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1b, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x53, 0x48, 0x65, 0x72, 0x6f, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x1a, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x53, 0x43, 0x48, 0x65, 0x72, 0x6f, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x55, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a,
	0x22, 0x13, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x2f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2f, 0x75, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3b, 0x63, 0x6c, 0x69, 0x6d,
	0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_message_hero_service_proto_rawDescOnce sync.Once
	file_message_hero_service_proto_rawDescData []byte
)

func file_message_hero_service_proto_rawDescGZIP() []byte {
	file_message_hero_service_proto_rawDescOnce.Do(func() {
		file_message_hero_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_message_hero_service_proto_rawDesc), len(file_message_hero_service_proto_rawDesc)))
	})
	return file_message_hero_service_proto_rawDescData
}

var file_message_hero_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_hero_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_message_hero_service_proto_goTypes = []any{
	(SCHeroLevelUpgrade_Code)(0), // 0: message.SCHeroLevelUpgrade.Code
	(*SCPushHeroUnlock)(nil),     // 1: message.SCPushHeroUnlock
	(*CSHeroLevelUpgrade)(nil),   // 2: message.CSHeroLevelUpgrade
	(*SCHeroLevelUpgrade)(nil),   // 3: message.SCHeroLevelUpgrade
	nil,                          // 4: message.SCHeroLevelUpgrade.CostsEntry
	(*HeroProto)(nil),            // 5: message.HeroProto
}
var file_message_hero_service_proto_depIdxs = []int32{
	5, // 0: message.SCPushHeroUnlock.heroes:type_name -> message.HeroProto
	0, // 1: message.SCHeroLevelUpgrade.code:type_name -> message.SCHeroLevelUpgrade.Code
	5, // 2: message.SCHeroLevelUpgrade.hero:type_name -> message.HeroProto
	4, // 3: message.SCHeroLevelUpgrade.costs:type_name -> message.SCHeroLevelUpgrade.CostsEntry
	2, // 4: message.HeroService.HeroLevelUpgrade:input_type -> message.CSHeroLevelUpgrade
	3, // 5: message.HeroService.HeroLevelUpgrade:output_type -> message.SCHeroLevelUpgrade
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_message_hero_service_proto_init() }
func file_message_hero_service_proto_init() {
	if File_message_hero_service_proto != nil {
		return
	}
	file_message_hero_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_message_hero_service_proto_rawDesc), len(file_message_hero_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_hero_service_proto_goTypes,
		DependencyIndexes: file_message_hero_service_proto_depIdxs,
		EnumInfos:         file_message_hero_service_proto_enumTypes,
		MessageInfos:      file_message_hero_service_proto_msgTypes,
	}.Build()
	File_message_hero_service_proto = out.File
	file_message_hero_service_proto_goTypes = nil
	file_message_hero_service_proto_depIdxs = nil
}
