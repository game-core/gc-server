// 交換受け取りレスポンス

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: exchange/exchange_receive_response.proto

package exchange

import (
	reflect "reflect"
	sync "sync"

	userExchange "github.com/game-core/gc-server/api/game/presentation/proto/exchange/userExchange"
	userExchangeItem "github.com/game-core/gc-server/api/game/presentation/proto/exchange/userExchangeItem"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExchangeReceiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserExchange     *userExchange.UserExchange         `protobuf:"bytes,1,opt,name=user_exchange,json=userExchange,proto3,oneof" json:"user_exchange,omitempty"`
	UserExchangeItem *userExchangeItem.UserExchangeItem `protobuf:"bytes,2,opt,name=user_exchange_item,json=userExchangeItem,proto3,oneof" json:"user_exchange_item,omitempty"`
}

func (x *ExchangeReceiveResponse) Reset() {
	*x = ExchangeReceiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_exchange_receive_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeReceiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeReceiveResponse) ProtoMessage() {}

func (x *ExchangeReceiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_exchange_receive_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeReceiveResponse.ProtoReflect.Descriptor instead.
func (*ExchangeReceiveResponse) Descriptor() ([]byte, []int) {
	return file_exchange_exchange_receive_response_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeReceiveResponse) GetUserExchange() *userExchange.UserExchange {
	if x != nil {
		return x.UserExchange
	}
	return nil
}

func (x *ExchangeReceiveResponse) GetUserExchangeItem() *userExchangeItem.UserExchangeItem {
	if x != nil {
		return x.UserExchangeItem
	}
	return nil
}

var File_exchange_exchange_receive_response_proto protoreflect.FileDescriptor

var file_exchange_exchange_receive_response_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70, 0x69, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x1a, 0x29, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x32, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x17, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x40, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00,
	0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x4d, 0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x01, 0x52, 0x10, 0x75, 0x73, 0x65,
	0x72, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x88, 0x01, 0x01,
	0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x42, 0xb2, 0x01, 0x0a, 0x0c, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x42, 0x1c, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x67, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0xa2,
	0x02, 0x03, 0x41, 0x47, 0x58, 0xaa, 0x02, 0x08, 0x41, 0x70, 0x69, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0xca, 0x02, 0x08, 0x41, 0x70, 0x69, 0x5c, 0x47, 0x61, 0x6d, 0x65, 0xe2, 0x02, 0x14, 0x41, 0x70,
	0x69, 0x5c, 0x47, 0x61, 0x6d, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x09, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x47, 0x61, 0x6d, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exchange_exchange_receive_response_proto_rawDescOnce sync.Once
	file_exchange_exchange_receive_response_proto_rawDescData = file_exchange_exchange_receive_response_proto_rawDesc
)

func file_exchange_exchange_receive_response_proto_rawDescGZIP() []byte {
	file_exchange_exchange_receive_response_proto_rawDescOnce.Do(func() {
		file_exchange_exchange_receive_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_exchange_exchange_receive_response_proto_rawDescData)
	})
	return file_exchange_exchange_receive_response_proto_rawDescData
}

var file_exchange_exchange_receive_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_exchange_exchange_receive_response_proto_goTypes = []any{
	(*ExchangeReceiveResponse)(nil),           // 0: api.game.ExchangeReceiveResponse
	(*userExchange.UserExchange)(nil),         // 1: api.game.UserExchange
	(*userExchangeItem.UserExchangeItem)(nil), // 2: api.game.UserExchangeItem
}
var file_exchange_exchange_receive_response_proto_depIdxs = []int32{
	1, // 0: api.game.ExchangeReceiveResponse.user_exchange:type_name -> api.game.UserExchange
	2, // 1: api.game.ExchangeReceiveResponse.user_exchange_item:type_name -> api.game.UserExchangeItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_exchange_exchange_receive_response_proto_init() }
func file_exchange_exchange_receive_response_proto_init() {
	if File_exchange_exchange_receive_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exchange_exchange_receive_response_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ExchangeReceiveResponse); i {
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
	file_exchange_exchange_receive_response_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_exchange_exchange_receive_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_exchange_exchange_receive_response_proto_goTypes,
		DependencyIndexes: file_exchange_exchange_receive_response_proto_depIdxs,
		MessageInfos:      file_exchange_exchange_receive_response_proto_msgTypes,
	}.Build()
	File_exchange_exchange_receive_response_proto = out.File
	file_exchange_exchange_receive_response_proto_rawDesc = nil
	file_exchange_exchange_receive_response_proto_goTypes = nil
	file_exchange_exchange_receive_response_proto_depIdxs = nil
}
