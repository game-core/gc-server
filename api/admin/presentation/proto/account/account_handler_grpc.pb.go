// アカウント

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: account/account_handler.proto

package account

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Account_GetGoogleUrl_FullMethodName   = "/api.admin.Account/GetGoogleUrl"
	Account_GetGoogleToken_FullMethodName = "/api.admin.Account/GetGoogleToken"
)

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	GetGoogleUrl(ctx context.Context, in *AccountGetGoogleUrlRequest, opts ...grpc.CallOption) (*AccountGetGoogleUrlResponse, error)
	GetGoogleToken(ctx context.Context, in *AccountGetGoogleTokenRequest, opts ...grpc.CallOption) (*AccountGetGoogleTokenResponse, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) GetGoogleUrl(ctx context.Context, in *AccountGetGoogleUrlRequest, opts ...grpc.CallOption) (*AccountGetGoogleUrlResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountGetGoogleUrlResponse)
	err := c.cc.Invoke(ctx, Account_GetGoogleUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetGoogleToken(ctx context.Context, in *AccountGetGoogleTokenRequest, opts ...grpc.CallOption) (*AccountGetGoogleTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccountGetGoogleTokenResponse)
	err := c.cc.Invoke(ctx, Account_GetGoogleToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations should embed UnimplementedAccountServer
// for forward compatibility
type AccountServer interface {
	GetGoogleUrl(context.Context, *AccountGetGoogleUrlRequest) (*AccountGetGoogleUrlResponse, error)
	GetGoogleToken(context.Context, *AccountGetGoogleTokenRequest) (*AccountGetGoogleTokenResponse, error)
}

// UnimplementedAccountServer should be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (UnimplementedAccountServer) GetGoogleUrl(context.Context, *AccountGetGoogleUrlRequest) (*AccountGetGoogleUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoogleUrl not implemented")
}
func (UnimplementedAccountServer) GetGoogleToken(context.Context, *AccountGetGoogleTokenRequest) (*AccountGetGoogleTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoogleToken not implemented")
}

// UnsafeAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServer will
// result in compilation errors.
type UnsafeAccountServer interface {
	mustEmbedUnimplementedAccountServer()
}

func RegisterAccountServer(s grpc.ServiceRegistrar, srv AccountServer) {
	s.RegisterService(&Account_ServiceDesc, srv)
}

func _Account_GetGoogleUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountGetGoogleUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetGoogleUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_GetGoogleUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetGoogleUrl(ctx, req.(*AccountGetGoogleUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetGoogleToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountGetGoogleTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetGoogleToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_GetGoogleToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetGoogleToken(ctx, req.(*AccountGetGoogleTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Account_ServiceDesc is the grpc.ServiceDesc for Account service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Account_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.admin.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGoogleUrl",
			Handler:    _Account_GetGoogleUrl_Handler,
		},
		{
			MethodName: "GetGoogleToken",
			Handler:    _Account_GetGoogleToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/account_handler.proto",
}