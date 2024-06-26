// ヘルスチェック

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: health/masterHealth/master_health.proto

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

type MasterHealth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HealthId         int64            `protobuf:"varint,1,opt,name=health_id,json=healthId,proto3" json:"health_id,omitempty"`
	Name             string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MasterHealthEnum MasterHealthEnum `protobuf:"varint,3,opt,name=master_health_enum,json=masterHealthEnum,proto3,enum=api.game.MasterHealthEnum" json:"master_health_enum,omitempty"`
}

func (x *MasterHealth) Reset() {
	*x = MasterHealth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_masterHealth_master_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterHealth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterHealth) ProtoMessage() {}

func (x *MasterHealth) ProtoReflect() protoreflect.Message {
	mi := &file_health_masterHealth_master_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterHealth.ProtoReflect.Descriptor instead.
func (*MasterHealth) Descriptor() ([]byte, []int) {
	return file_health_masterHealth_master_health_proto_rawDescGZIP(), []int{0}
}

func (x *MasterHealth) GetHealthId() int64 {
	if x != nil {
		return x.HealthId
	}
	return 0
}

func (x *MasterHealth) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MasterHealth) GetMasterHealthEnum() MasterHealthEnum {
	if x != nil {
		return x.MasterHealthEnum
	}
	return MasterHealthEnum_MasterNone
}

var File_health_masterHealth_master_health_proto protoreflect.FileDescriptor

var file_health_masterHealth_master_health_proto_rawDesc = []byte{
	0x0a, 0x27, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x69, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x1a, 0x2c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x89, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x12, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x10, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x45, 0x6e, 0x75, 0x6d, 0x42, 0xb2, 0x01,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x42, 0x11,
	0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x63, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0xa2, 0x02, 0x03, 0x41, 0x47, 0x58, 0xaa, 0x02, 0x08, 0x41, 0x70, 0x69, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0xca, 0x02, 0x08, 0x41, 0x70, 0x69, 0x5c, 0x47, 0x61, 0x6d, 0x65, 0xe2,
	0x02, 0x14, 0x41, 0x70, 0x69, 0x5c, 0x47, 0x61, 0x6d, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x47, 0x61,
	0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_health_masterHealth_master_health_proto_rawDescOnce sync.Once
	file_health_masterHealth_master_health_proto_rawDescData = file_health_masterHealth_master_health_proto_rawDesc
)

func file_health_masterHealth_master_health_proto_rawDescGZIP() []byte {
	file_health_masterHealth_master_health_proto_rawDescOnce.Do(func() {
		file_health_masterHealth_master_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_masterHealth_master_health_proto_rawDescData)
	})
	return file_health_masterHealth_master_health_proto_rawDescData
}

var file_health_masterHealth_master_health_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_health_masterHealth_master_health_proto_goTypes = []any{
	(*MasterHealth)(nil),  // 0: api.game.MasterHealth
	(MasterHealthEnum)(0), // 1: api.game.MasterHealthEnum
}
var file_health_masterHealth_master_health_proto_depIdxs = []int32{
	1, // 0: api.game.MasterHealth.master_health_enum:type_name -> api.game.MasterHealthEnum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_health_masterHealth_master_health_proto_init() }
func file_health_masterHealth_master_health_proto_init() {
	if File_health_masterHealth_master_health_proto != nil {
		return
	}
	file_health_masterHealth_master_health_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_health_masterHealth_master_health_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MasterHealth); i {
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
			RawDescriptor: file_health_masterHealth_master_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_health_masterHealth_master_health_proto_goTypes,
		DependencyIndexes: file_health_masterHealth_master_health_proto_depIdxs,
		MessageInfos:      file_health_masterHealth_master_health_proto_msgTypes,
	}.Build()
	File_health_masterHealth_master_health_proto = out.File
	file_health_masterHealth_master_health_proto_rawDesc = nil
	file_health_masterHealth_master_health_proto_goTypes = nil
	file_health_masterHealth_master_health_proto_depIdxs = nil
}
