// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: module/modules.proto

package climod

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

// Module ID
type ModuleID int32

const (
	ModuleID_ModuleUnspecified ModuleID = 0
	// System module, like handshake, heartbeat, etc.
	ModuleID_System ModuleID = 1
	// Dev module, like gm command execution, etc.
	ModuleID_Dev ModuleID = 2
	// User module, like login, user data sync, some user basic data operations, etc.
	ModuleID_User    ModuleID = 3
	ModuleID_Storage ModuleID = 4
	ModuleID_Hero    ModuleID = 5
	ModuleID_Room    ModuleID = 6
)

// Enum value maps for ModuleID.
var (
	ModuleID_name = map[int32]string{
		0: "ModuleUnspecified",
		1: "System",
		2: "Dev",
		3: "User",
		4: "Storage",
		5: "Hero",
		6: "Room",
	}
	ModuleID_value = map[string]int32{
		"ModuleUnspecified": 0,
		"System":            1,
		"Dev":               2,
		"User":              3,
		"Storage":           4,
		"Hero":              5,
		"Room":              6,
	}
)

func (x ModuleID) Enum() *ModuleID {
	p := new(ModuleID)
	*p = x
	return p
}

func (x ModuleID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ModuleID) Descriptor() protoreflect.EnumDescriptor {
	return file_module_modules_proto_enumTypes[0].Descriptor()
}

func (ModuleID) Type() protoreflect.EnumType {
	return &file_module_modules_proto_enumTypes[0]
}

func (x ModuleID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ModuleID.Descriptor instead.
func (ModuleID) EnumDescriptor() ([]byte, []int) {
	return file_module_modules_proto_rawDescGZIP(), []int{0}
}

var File_module_modules_proto protoreflect.FileDescriptor

var file_module_modules_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2a, 0x61,
	0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x44, 0x65, 0x76, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x10, 0x04, 0x12, 0x08, 0x0a,
	0x04, 0x48, 0x65, 0x72, 0x6f, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x10,
	0x06, 0x42, 0x1a, 0x5a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x3b, 0x63, 0x6c, 0x69, 0x6d, 0x6f, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_module_modules_proto_rawDescOnce sync.Once
	file_module_modules_proto_rawDescData []byte
)

func file_module_modules_proto_rawDescGZIP() []byte {
	file_module_modules_proto_rawDescOnce.Do(func() {
		file_module_modules_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_module_modules_proto_rawDesc), len(file_module_modules_proto_rawDesc)))
	})
	return file_module_modules_proto_rawDescData
}

var file_module_modules_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_module_modules_proto_goTypes = []any{
	(ModuleID)(0), // 0: module.ModuleID
}
var file_module_modules_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_module_modules_proto_init() }
func file_module_modules_proto_init() {
	if File_module_modules_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_module_modules_proto_rawDesc), len(file_module_modules_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_module_modules_proto_goTypes,
		DependencyIndexes: file_module_modules_proto_depIdxs,
		EnumInfos:         file_module_modules_proto_enumTypes,
	}.Build()
	File_module_modules_proto = out.File
	file_module_modules_proto_goTypes = nil
	file_module_modules_proto_depIdxs = nil
}
