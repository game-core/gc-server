// ログインボーナス

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: loginBonus/masterLoginBonus/master_login_bonus.proto

package masterLoginBonus

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MasterLoginBonus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MasterLoginBonusId int64  `protobuf:"varint,1,opt,name=master_login_bonus_id,json=masterLoginBonusId,proto3" json:"master_login_bonus_id,omitempty"`
	MasterEventId      int64  `protobuf:"varint,2,opt,name=master_event_id,json=masterEventId,proto3" json:"master_event_id,omitempty"`
	Name               string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MasterLoginBonus) Reset() {
	*x = MasterLoginBonus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loginBonus_masterLoginBonus_master_login_bonus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterLoginBonus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterLoginBonus) ProtoMessage() {}

func (x *MasterLoginBonus) ProtoReflect() protoreflect.Message {
	mi := &file_loginBonus_masterLoginBonus_master_login_bonus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterLoginBonus.ProtoReflect.Descriptor instead.
func (*MasterLoginBonus) Descriptor() ([]byte, []int) {
	return file_loginBonus_masterLoginBonus_master_login_bonus_proto_rawDescGZIP(), []int{0}
}

func (x *MasterLoginBonus) GetMasterLoginBonusId() int64 {
	if x != nil {
		return x.MasterLoginBonusId
	}
	return 0
}

func (x *MasterLoginBonus) GetMasterEventId() int64 {
	if x != nil {
		return x.MasterEventId
	}
	return 0
}

func (x *MasterLoginBonus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_loginBonus_masterLoginBonus_master_login_bonus_proto protoreflect.FileDescriptor

var file_loginBonus_masterLoginBonus_master_login_bonus_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x2f, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x2f, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x22, 0x81, 0x01, 0x0a, 0x10, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x42, 0x6f, 0x6e, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0xbe, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x42, 0x15, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x56,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2d,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x42, 0x6f, 0x6e, 0x75, 0x73, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x47, 0x58, 0xaa, 0x02, 0x08, 0x41,
	0x70, 0x69, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0xca, 0x02, 0x08, 0x41, 0x70, 0x69, 0x5c, 0x47, 0x61,
	0x6d, 0x65, 0xe2, 0x02, 0x14, 0x41, 0x70, 0x69, 0x5c, 0x47, 0x61, 0x6d, 0x65, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x47, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loginBonus_masterLoginBonus_master_login_bonus_proto_rawDescOnce sync.Once
	file_loginBonus_masterLoginBonus_master_login_bonus_proto_rawDescData = file_loginBonus_masterLoginBonus_master_login_bonus_proto_rawDesc
)

func file_loginBonus_masterLoginBonus_master_login_bonus_proto_rawDescGZIP() []byte {
	file_loginBonus_masterLoginBonus_master_login_bonus_proto_rawDescOnce.Do(func() {
		file_loginBonus_masterLoginBonus_master_login_bonus_proto_rawDescData = protoimpl.X.CompressGZIP(file_loginBonus_masterLoginBonus_master_login_bonus_proto_rawDescData)
	})
	return file_loginBonus_masterLoginBonus_master_login_bonus_proto_rawDescData
}

var file_loginBonus_masterLoginBonus_master_login_bonus_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_loginBonus_masterLoginBonus_master_login_bonus_proto_goTypes = []any{
	(*MasterLoginBonus)(nil), // 0: api.game.MasterLoginBonus
}
var file_loginBonus_masterLoginBonus_master_login_bonus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loginBonus_masterLoginBonus_master_login_bonus_proto_init() }
func file_loginBonus_masterLoginBonus_master_login_bonus_proto_init() {
	if File_loginBonus_masterLoginBonus_master_login_bonus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loginBonus_masterLoginBonus_master_login_bonus_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MasterLoginBonus); i {
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
			RawDescriptor: file_loginBonus_masterLoginBonus_master_login_bonus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_loginBonus_masterLoginBonus_master_login_bonus_proto_goTypes,
		DependencyIndexes: file_loginBonus_masterLoginBonus_master_login_bonus_proto_depIdxs,
		MessageInfos:      file_loginBonus_masterLoginBonus_master_login_bonus_proto_msgTypes,
	}.Build()
	File_loginBonus_masterLoginBonus_master_login_bonus_proto = out.File
	file_loginBonus_masterLoginBonus_master_login_bonus_proto_rawDesc = nil
	file_loginBonus_masterLoginBonus_master_login_bonus_proto_goTypes = nil
	file_loginBonus_masterLoginBonus_master_login_bonus_proto_depIdxs = nil
}
