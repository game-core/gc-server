// アカウント

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: account/account_handler.proto

package account

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

var File_account_account_handler_proto protoreflect.FileDescriptor

var file_account_account_handler_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x32, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x5f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x67, 0x65, 0x74, 0x5f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xeb, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x6c, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x72,
	0x6c, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xae,
	0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x42, 0x13, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x63,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0xa2, 0x02, 0x03, 0x41,
	0x41, 0x58, 0xaa, 0x02, 0x09, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xca, 0x02,
	0x09, 0x41, 0x70, 0x69, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xe2, 0x02, 0x15, 0x41, 0x70, 0x69,
	0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_account_account_handler_proto_goTypes = []any{
	(*AccountGetGoogleLoginUrlRequest)(nil),    // 0: api.admin.AccountGetGoogleLoginUrlRequest
	(*AccountGetGoogleLoginTokenRequest)(nil),  // 1: api.admin.AccountGetGoogleLoginTokenRequest
	(*AccountGetGoogleLoginUrlResponse)(nil),   // 2: api.admin.AccountGetGoogleLoginUrlResponse
	(*AccountGetGoogleLoginTokenResponse)(nil), // 3: api.admin.AccountGetGoogleLoginTokenResponse
}
var file_account_account_handler_proto_depIdxs = []int32{
	0, // 0: api.admin.Account.GetGoogleLoginUrl:input_type -> api.admin.AccountGetGoogleLoginUrlRequest
	1, // 1: api.admin.Account.GetGoogleLoginToken:input_type -> api.admin.AccountGetGoogleLoginTokenRequest
	2, // 2: api.admin.Account.GetGoogleLoginUrl:output_type -> api.admin.AccountGetGoogleLoginUrlResponse
	3, // 3: api.admin.Account.GetGoogleLoginToken:output_type -> api.admin.AccountGetGoogleLoginTokenResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_account_account_handler_proto_init() }
func file_account_account_handler_proto_init() {
	if File_account_account_handler_proto != nil {
		return
	}
	file_account_account_get_google_login_url_request_proto_init()
	file_account_account_get_google_login_url_response_proto_init()
	file_account_account_get_google_login_token_request_proto_init()
	file_account_account_get_google_login_token_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_account_account_handler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_account_handler_proto_goTypes,
		DependencyIndexes: file_account_account_handler_proto_depIdxs,
	}.Build()
	File_account_account_handler_proto = out.File
	file_account_account_handler_proto_rawDesc = nil
	file_account_account_handler_proto_goTypes = nil
	file_account_account_handler_proto_depIdxs = nil
}
