// 管理者アカウントのGoogleToken

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: account/adminAccountGoogleToken/admin_account_google_token.proto

package adminAccountGoogleToken

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AdminAccountGoogleToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiredAt    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
}

func (x *AdminAccountGoogleToken) Reset() {
	*x = AdminAccountGoogleToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_adminAccountGoogleToken_admin_account_google_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminAccountGoogleToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminAccountGoogleToken) ProtoMessage() {}

func (x *AdminAccountGoogleToken) ProtoReflect() protoreflect.Message {
	mi := &file_account_adminAccountGoogleToken_admin_account_google_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminAccountGoogleToken.ProtoReflect.Descriptor instead.
func (*AdminAccountGoogleToken) Descriptor() ([]byte, []int) {
	return file_account_adminAccountGoogleToken_admin_account_google_token_proto_rawDescGZIP(), []int{0}
}

func (x *AdminAccountGoogleToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AdminAccountGoogleToken) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *AdminAccountGoogleToken) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

var File_account_adminAccountGoogleToken_admin_account_google_token_proto protoreflect.FileDescriptor

var file_account_adminAccountGoogleToken_admin_account_google_token_proto_rawDesc = []byte{
	0x0a, 0x40, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c,
	0x01, 0x0a, 0x17, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x42, 0xcf, 0x01,
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42,
	0x1c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x5b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6d, 0x65,
	0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0xa2, 0x02, 0x03, 0x41,
	0x41, 0x58, 0xaa, 0x02, 0x09, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xca, 0x02,
	0x09, 0x41, 0x70, 0x69, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xe2, 0x02, 0x15, 0x41, 0x70, 0x69,
	0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_adminAccountGoogleToken_admin_account_google_token_proto_rawDescOnce sync.Once
	file_account_adminAccountGoogleToken_admin_account_google_token_proto_rawDescData = file_account_adminAccountGoogleToken_admin_account_google_token_proto_rawDesc
)

func file_account_adminAccountGoogleToken_admin_account_google_token_proto_rawDescGZIP() []byte {
	file_account_adminAccountGoogleToken_admin_account_google_token_proto_rawDescOnce.Do(func() {
		file_account_adminAccountGoogleToken_admin_account_google_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_adminAccountGoogleToken_admin_account_google_token_proto_rawDescData)
	})
	return file_account_adminAccountGoogleToken_admin_account_google_token_proto_rawDescData
}

var file_account_adminAccountGoogleToken_admin_account_google_token_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_account_adminAccountGoogleToken_admin_account_google_token_proto_goTypes = []any{
	(*AdminAccountGoogleToken)(nil), // 0: api.admin.AdminAccountGoogleToken
	(*timestamppb.Timestamp)(nil),   // 1: google.protobuf.Timestamp
}
var file_account_adminAccountGoogleToken_admin_account_google_token_proto_depIdxs = []int32{
	1, // 0: api.admin.AdminAccountGoogleToken.expired_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_account_adminAccountGoogleToken_admin_account_google_token_proto_init() }
func file_account_adminAccountGoogleToken_admin_account_google_token_proto_init() {
	if File_account_adminAccountGoogleToken_admin_account_google_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_adminAccountGoogleToken_admin_account_google_token_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AdminAccountGoogleToken); i {
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
			RawDescriptor: file_account_adminAccountGoogleToken_admin_account_google_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_adminAccountGoogleToken_admin_account_google_token_proto_goTypes,
		DependencyIndexes: file_account_adminAccountGoogleToken_admin_account_google_token_proto_depIdxs,
		MessageInfos:      file_account_adminAccountGoogleToken_admin_account_google_token_proto_msgTypes,
	}.Build()
	File_account_adminAccountGoogleToken_admin_account_google_token_proto = out.File
	file_account_adminAccountGoogleToken_admin_account_google_token_proto_rawDesc = nil
	file_account_adminAccountGoogleToken_admin_account_google_token_proto_goTypes = nil
	file_account_adminAccountGoogleToken_admin_account_google_token_proto_depIdxs = nil
}
