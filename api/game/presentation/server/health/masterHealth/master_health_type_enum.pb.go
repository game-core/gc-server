// ヘルスチェックタイプ

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: health/masterHealth/master_health_type_enum.proto

package masterHealth

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

type MasterHealthType int32

const (
	MasterHealthType_MasterNone    MasterHealthType = 0
	MasterHealthType_MasterSuccess MasterHealthType = 1
)

// Enum value maps for MasterHealthType.
var (
	MasterHealthType_name = map[int32]string{
		0: "MasterNone",
		1: "MasterSuccess",
	}
	MasterHealthType_value = map[string]int32{
		"MasterNone":    0,
		"MasterSuccess": 1,
	}
)

func (x MasterHealthType) Enum() *MasterHealthType {
	p := new(MasterHealthType)
	*p = x
	return p
}

func (x MasterHealthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MasterHealthType) Descriptor() protoreflect.EnumDescriptor {
	return file_health_masterHealth_master_health_type_enum_proto_enumTypes[0].Descriptor()
}

func (MasterHealthType) Type() protoreflect.EnumType {
	return &file_health_masterHealth_master_health_type_enum_proto_enumTypes[0]
}

func (x MasterHealthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MasterHealthType.Descriptor instead.
func (MasterHealthType) EnumDescriptor() ([]byte, []int) {
	return file_health_masterHealth_master_health_type_enum_proto_rawDescGZIP(), []int{0}
}

var File_health_masterHealth_master_health_type_enum_proto protoreflect.FileDescriptor

var file_health_masterHealth_master_health_type_enum_proto_rawDesc = []byte{
	0x0a, 0x31, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2a, 0x35, 0x0a,
	0x10, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x6e, 0x65, 0x10,
	0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x10, 0x01, 0x42, 0xbb, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x42, 0x19, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x61, 0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0xa2, 0x02, 0x03, 0x41, 0x47, 0x58, 0xaa, 0x02, 0x08, 0x41, 0x70, 0x69, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0xca, 0x02, 0x08, 0x41, 0x70, 0x69, 0x5c, 0x47, 0x61, 0x6d, 0x65, 0xe2,
	0x02, 0x14, 0x41, 0x70, 0x69, 0x5c, 0x47, 0x61, 0x6d, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x47, 0x61,
	0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_health_masterHealth_master_health_type_enum_proto_rawDescOnce sync.Once
	file_health_masterHealth_master_health_type_enum_proto_rawDescData = file_health_masterHealth_master_health_type_enum_proto_rawDesc
)

func file_health_masterHealth_master_health_type_enum_proto_rawDescGZIP() []byte {
	file_health_masterHealth_master_health_type_enum_proto_rawDescOnce.Do(func() {
		file_health_masterHealth_master_health_type_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_masterHealth_master_health_type_enum_proto_rawDescData)
	})
	return file_health_masterHealth_master_health_type_enum_proto_rawDescData
}

var file_health_masterHealth_master_health_type_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_health_masterHealth_master_health_type_enum_proto_goTypes = []interface{}{
	(MasterHealthType)(0), // 0: api.game.MasterHealthType
}
var file_health_masterHealth_master_health_type_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_health_masterHealth_master_health_type_enum_proto_init() }
func file_health_masterHealth_master_health_type_enum_proto_init() {
	if File_health_masterHealth_master_health_type_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_health_masterHealth_master_health_type_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_health_masterHealth_master_health_type_enum_proto_goTypes,
		DependencyIndexes: file_health_masterHealth_master_health_type_enum_proto_depIdxs,
		EnumInfos:         file_health_masterHealth_master_health_type_enum_proto_enumTypes,
	}.Build()
	File_health_masterHealth_master_health_type_enum_proto = out.File
	file_health_masterHealth_master_health_type_enum_proto_rawDesc = nil
	file_health_masterHealth_master_health_type_enum_proto_goTypes = nil
	file_health_masterHealth_master_health_type_enum_proto_depIdxs = nil
}
