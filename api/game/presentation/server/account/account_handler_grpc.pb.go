// アカウント

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Account_Get_FullMethodName    = "/api.game.Account/Get"
	Account_Create_FullMethodName = "/api.game.Account/Create"
	Account_Login_FullMethodName  = "/api.game.Account/Login"
)

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	Get(ctx context.Context, in *AccountGetRequest, opts ...grpc.CallOption) (*AccountGetResponse, error)
	Create(ctx context.Context, in *AccountCreateRequest, opts ...grpc.CallOption) (*AccountCreateResponse, error)
	Login(ctx context.Context, in *AccountLoginRequest, opts ...grpc.CallOption) (*AccountLoginResponse, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Get(ctx context.Context, in *AccountGetRequest, opts ...grpc.CallOption) (*AccountGetResponse, error) {
	out := new(AccountGetResponse)
	err := c.cc.Invoke(ctx, Account_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Create(ctx context.Context, in *AccountCreateRequest, opts ...grpc.CallOption) (*AccountCreateResponse, error) {
	out := new(AccountCreateResponse)
	err := c.cc.Invoke(ctx, Account_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Login(ctx context.Context, in *AccountLoginRequest, opts ...grpc.CallOption) (*AccountLoginResponse, error) {
	out := new(AccountLoginResponse)
	err := c.cc.Invoke(ctx, Account_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations should embed UnimplementedAccountServer
// for forward compatibility
type AccountServer interface {
	Get(context.Context, *AccountGetRequest) (*AccountGetResponse, error)
	Create(context.Context, *AccountCreateRequest) (*AccountCreateResponse, error)
	Login(context.Context, *AccountLoginRequest) (*AccountLoginResponse, error)
}

// UnimplementedAccountServer should be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (UnimplementedAccountServer) Get(context.Context, *AccountGetRequest) (*AccountGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAccountServer) Create(context.Context, *AccountCreateRequest) (*AccountCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccountServer) Login(context.Context, *AccountLoginRequest) (*AccountLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
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

func _Account_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Get(ctx, req.(*AccountGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Create(ctx, req.(*AccountCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Login(ctx, req.(*AccountLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Account_ServiceDesc is the grpc.ServiceDesc for Account service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Account_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.game.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Account_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Account_Create_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Account_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/account_handler.proto",
}