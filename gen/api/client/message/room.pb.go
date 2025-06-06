// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: message/room.proto

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

type UserRoomProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        int64                  `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"` // Room ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRoomProto) Reset() {
	*x = UserRoomProto{}
	mi := &file_message_room_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRoomProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRoomProto) ProtoMessage() {}

func (x *UserRoomProto) ProtoReflect() protoreflect.Message {
	mi := &file_message_room_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRoomProto.ProtoReflect.Descriptor instead.
func (*UserRoomProto) Descriptor() ([]byte, []int) {
	return file_message_room_proto_rawDescGZIP(), []int{0}
}

func (x *UserRoomProto) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

// Room data
type RoomProto struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Basic         *RoomBasicProto            `protobuf:"bytes,1,opt,name=basic,proto3" json:"basic,omitempty"`                                                                                // Room basic data
	Members       map[int64]*RoomMemberProto `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` // user ID -> room member data
	Invitees      []int64                    `protobuf:"varint,3,rep,packed,name=invitees,proto3" json:"invitees,omitempty"`                                                                  // Invitees
	Requests      []int64                    `protobuf:"varint,4,rep,packed,name=requests,proto3" json:"requests,omitempty"`                                                                  // Requests
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoomProto) Reset() {
	*x = RoomProto{}
	mi := &file_message_room_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomProto) ProtoMessage() {}

func (x *RoomProto) ProtoReflect() protoreflect.Message {
	mi := &file_message_room_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomProto.ProtoReflect.Descriptor instead.
func (*RoomProto) Descriptor() ([]byte, []int) {
	return file_message_room_proto_rawDescGZIP(), []int{1}
}

func (x *RoomProto) GetBasic() *RoomBasicProto {
	if x != nil {
		return x.Basic
	}
	return nil
}

func (x *RoomProto) GetMembers() map[int64]*RoomMemberProto {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *RoomProto) GetInvitees() []int64 {
	if x != nil {
		return x.Invitees
	}
	return nil
}

func (x *RoomProto) GetRequests() []int64 {
	if x != nil {
		return x.Requests
	}
	return nil
}

// Room basic data
type RoomBasicProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                // Room ID
	RoomType      int64                  `protobuf:"varint,2,opt,name=room_type,json=roomType,proto3" json:"room_type,omitempty"`    // room type
	Creator       *UserBasicProto        `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`                       // Creator
	CreatedAt     int64                  `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // Created at
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoomBasicProto) Reset() {
	*x = RoomBasicProto{}
	mi := &file_message_room_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomBasicProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomBasicProto) ProtoMessage() {}

func (x *RoomBasicProto) ProtoReflect() protoreflect.Message {
	mi := &file_message_room_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomBasicProto.ProtoReflect.Descriptor instead.
func (*RoomBasicProto) Descriptor() ([]byte, []int) {
	return file_message_room_proto_rawDescGZIP(), []int{2}
}

func (x *RoomBasicProto) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoomBasicProto) GetRoomType() int64 {
	if x != nil {
		return x.RoomType
	}
	return 0
}

func (x *RoomBasicProto) GetCreator() *UserBasicProto {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *RoomBasicProto) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

// Room member
type RoomMemberProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *UserBasicProto        `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`                          // User basic data
	JoinedAt      int64                  `protobuf:"varint,2,opt,name=joined_at,json=joinedAt,proto3" json:"joined_at,omitempty"` // Joined at
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoomMemberProto) Reset() {
	*x = RoomMemberProto{}
	mi := &file_message_room_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomMemberProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomMemberProto) ProtoMessage() {}

func (x *RoomMemberProto) ProtoReflect() protoreflect.Message {
	mi := &file_message_room_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomMemberProto.ProtoReflect.Descriptor instead.
func (*RoomMemberProto) Descriptor() ([]byte, []int) {
	return file_message_room_proto_rawDescGZIP(), []int{3}
}

func (x *RoomMemberProto) GetUser() *UserBasicProto {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *RoomMemberProto) GetJoinedAt() int64 {
	if x != nil {
		return x.JoinedAt
	}
	return 0
}

var File_message_room_proto protoreflect.FileDescriptor

var file_message_room_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x18, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49,
	0x64, 0x22, 0x83, 0x02, 0x0a, 0x09, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x2d, 0x0a, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x12, 0x39,
	0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x1a, 0x54, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8f, 0x01, 0x0a, 0x0e, 0x52, 0x6f, 0x6f, 0x6d,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72,
	0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5b, 0x0a, 0x0f, 0x52, 0x6f, 0x6f,
	0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f, 0x69,
	0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6a, 0x6f,
	0x69, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3b, 0x63, 0x6c, 0x69,
	0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_message_room_proto_rawDescOnce sync.Once
	file_message_room_proto_rawDescData []byte
)

func file_message_room_proto_rawDescGZIP() []byte {
	file_message_room_proto_rawDescOnce.Do(func() {
		file_message_room_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_message_room_proto_rawDesc), len(file_message_room_proto_rawDesc)))
	})
	return file_message_room_proto_rawDescData
}

var file_message_room_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_message_room_proto_goTypes = []any{
	(*UserRoomProto)(nil),   // 0: message.UserRoomProto
	(*RoomProto)(nil),       // 1: message.RoomProto
	(*RoomBasicProto)(nil),  // 2: message.RoomBasicProto
	(*RoomMemberProto)(nil), // 3: message.RoomMemberProto
	nil,                     // 4: message.RoomProto.MembersEntry
	(*UserBasicProto)(nil),  // 5: message.UserBasicProto
}
var file_message_room_proto_depIdxs = []int32{
	2, // 0: message.RoomProto.basic:type_name -> message.RoomBasicProto
	4, // 1: message.RoomProto.members:type_name -> message.RoomProto.MembersEntry
	5, // 2: message.RoomBasicProto.creator:type_name -> message.UserBasicProto
	5, // 3: message.RoomMemberProto.user:type_name -> message.UserBasicProto
	3, // 4: message.RoomProto.MembersEntry.value:type_name -> message.RoomMemberProto
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_message_room_proto_init() }
func file_message_room_proto_init() {
	if File_message_room_proto != nil {
		return
	}
	file_message_user_basic_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_message_room_proto_rawDesc), len(file_message_room_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_room_proto_goTypes,
		DependencyIndexes: file_message_room_proto_depIdxs,
		MessageInfos:      file_message_room_proto_msgTypes,
	}.Build()
	File_message_room_proto = out.File
	file_message_room_proto_goTypes = nil
	file_message_room_proto_depIdxs = nil
}
