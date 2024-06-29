// 管理者アカウントのGoogleToken情報

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: account/adminAccountGoogleTokenInfo/admin_account_google_token_info.proto

package adminAccountGoogleTokenInfo

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

type AdminAccountGoogleTokenInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email         string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	VerifiedEmail bool   `protobuf:"varint,3,opt,name=verified_email,json=verifiedEmail,proto3" json:"verified_email,omitempty"`
	ExpiresIn     int64  `protobuf:"varint,4,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	IssuedTo      string `protobuf:"bytes,5,opt,name=issued_to,json=issuedTo,proto3" json:"issued_to,omitempty"`
	Scope         string `protobuf:"bytes,6,opt,name=scope,proto3" json:"scope,omitempty"`
}

func (x *AdminAccountGoogleTokenInfo) Reset() {
	*x = AdminAccountGoogleTokenInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminAccountGoogleTokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminAccountGoogleTokenInfo) ProtoMessage() {}

func (x *AdminAccountGoogleTokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminAccountGoogleTokenInfo.ProtoReflect.Descriptor instead.
func (*AdminAccountGoogleTokenInfo) Descriptor() ([]byte, []int) {
	return file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_rawDescGZIP(), []int{0}
}

func (x *AdminAccountGoogleTokenInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AdminAccountGoogleTokenInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AdminAccountGoogleTokenInfo) GetVerifiedEmail() bool {
	if x != nil {
		return x.VerifiedEmail
	}
	return false
}

func (x *AdminAccountGoogleTokenInfo) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *AdminAccountGoogleTokenInfo) GetIssuedTo() string {
	if x != nil {
		return x.IssuedTo
	}
	return ""
}

func (x *AdminAccountGoogleTokenInfo) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

var File_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto protoreflect.FileDescriptor

var file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_rawDesc = []byte{
	0x0a, 0x49, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0xc5, 0x01, 0x0a, 0x1b, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x54, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x42, 0xd7,
	0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x42, 0x20, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x63, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0xa2, 0x02, 0x03, 0x41, 0x41, 0x58, 0xaa, 0x02, 0x09, 0x41, 0x70,
	0x69, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xca, 0x02, 0x09, 0x41, 0x70, 0x69, 0x5c, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0xe2, 0x02, 0x15, 0x41, 0x70, 0x69, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x41, 0x70,
	0x69, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_rawDescOnce sync.Once
	file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_rawDescData = file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_rawDesc
)

func file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_rawDescGZIP() []byte {
	file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_rawDescOnce.Do(func() {
		file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_rawDescData)
	})
	return file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_rawDescData
}

var file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_goTypes = []any{
	(*AdminAccountGoogleTokenInfo)(nil), // 0: api.admin.AdminAccountGoogleTokenInfo
}
var file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_init() }
func file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_init() {
	if File_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AdminAccountGoogleTokenInfo); i {
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
			RawDescriptor: file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_goTypes,
		DependencyIndexes: file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_depIdxs,
		MessageInfos:      file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_msgTypes,
	}.Build()
	File_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto = out.File
	file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_rawDesc = nil
	file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_goTypes = nil
	file_account_adminAccountGoogleTokenInfo_admin_account_google_token_info_proto_depIdxs = nil
}