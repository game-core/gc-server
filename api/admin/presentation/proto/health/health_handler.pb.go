// ヘルスチェック

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: health/health_handler.proto

package health

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_health_health_handler_proto protoreflect.FileDescriptor

var file_health_health_handler_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x21, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x50, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x46, 0x0a, 0x05, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0xac, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x42, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x67, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0xa2, 0x02, 0x03,
	0x41, 0x41, 0x58, 0xaa, 0x02, 0x09, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xca,
	0x02, 0x09, 0x41, 0x70, 0x69, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xe2, 0x02, 0x15, 0x41, 0x70,
	0x69, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_health_health_handler_proto_goTypes = []any{
	(*HealthCheckRequest)(nil),  // 0: api.admin.HealthCheckRequest
	(*HealthCheckResponse)(nil), // 1: api.admin.HealthCheckResponse
}
var file_health_health_handler_proto_depIdxs = []int32{
	0, // 0: api.admin.Health.Check:input_type -> api.admin.HealthCheckRequest
	1, // 1: api.admin.Health.Check:output_type -> api.admin.HealthCheckResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_health_health_handler_proto_init() }
func file_health_health_handler_proto_init() {
	if File_health_health_handler_proto != nil {
		return
	}
	file_health_health_check_request_proto_init()
	file_health_health_check_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_health_health_handler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_health_handler_proto_goTypes,
		DependencyIndexes: file_health_health_handler_proto_depIdxs,
	}.Build()
	File_health_health_handler_proto = out.File
	file_health_health_handler_proto_rawDesc = nil
	file_health_health_handler_proto_goTypes = nil
	file_health_health_handler_proto_depIdxs = nil
}
