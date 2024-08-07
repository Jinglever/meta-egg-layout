// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.22.2
// source: meta_egg_layout.proto

package meta_egg_layout

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MetaEggLayout_CreateUser_FullMethodName    = "/metaegglayout.MetaEggLayout/CreateUser"
	MetaEggLayout_GetUserDetail_FullMethodName = "/metaegglayout.MetaEggLayout/GetUserDetail"
	MetaEggLayout_GetUserList_FullMethodName   = "/metaegglayout.MetaEggLayout/GetUserList"
	MetaEggLayout_UpdateUser_FullMethodName    = "/metaegglayout.MetaEggLayout/UpdateUser"
	MetaEggLayout_DeleteUser_FullMethodName    = "/metaegglayout.MetaEggLayout/DeleteUser"
)

// MetaEggLayoutClient is the client API for MetaEggLayout service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetaEggLayoutClient interface {
	// 创建用户
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserDetail, error)
	// 获取用户详情
	GetUserDetail(ctx context.Context, in *GetUserDetailRequest, opts ...grpc.CallOption) (*UserDetail, error)
	// 获取用户列表
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error)
	// 更新用户
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除用户
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type metaEggLayoutClient struct {
	cc grpc.ClientConnInterface
}

func NewMetaEggLayoutClient(cc grpc.ClientConnInterface) MetaEggLayoutClient {
	return &metaEggLayoutClient{cc}
}

func (c *metaEggLayoutClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserDetail, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserDetail)
	err := c.cc.Invoke(ctx, MetaEggLayout_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaEggLayoutClient) GetUserDetail(ctx context.Context, in *GetUserDetailRequest, opts ...grpc.CallOption) (*UserDetail, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserDetail)
	err := c.cc.Invoke(ctx, MetaEggLayout_GetUserDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaEggLayoutClient) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserListResponse)
	err := c.cc.Invoke(ctx, MetaEggLayout_GetUserList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaEggLayoutClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MetaEggLayout_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaEggLayoutClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MetaEggLayout_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetaEggLayoutServer is the server API for MetaEggLayout service.
// All implementations must embed UnimplementedMetaEggLayoutServer
// for forward compatibility
type MetaEggLayoutServer interface {
	// 创建用户
	CreateUser(context.Context, *CreateUserRequest) (*UserDetail, error)
	// 获取用户详情
	GetUserDetail(context.Context, *GetUserDetailRequest) (*UserDetail, error)
	// 获取用户列表
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error)
	// 更新用户
	UpdateUser(context.Context, *UpdateUserRequest) (*emptypb.Empty, error)
	// 删除用户
	DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMetaEggLayoutServer()
}

// UnimplementedMetaEggLayoutServer must be embedded to have forward compatible implementations.
type UnimplementedMetaEggLayoutServer struct {
}

func (UnimplementedMetaEggLayoutServer) CreateUser(context.Context, *CreateUserRequest) (*UserDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMetaEggLayoutServer) GetUserDetail(context.Context, *GetUserDetailRequest) (*UserDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetail not implemented")
}
func (UnimplementedMetaEggLayoutServer) GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedMetaEggLayoutServer) UpdateUser(context.Context, *UpdateUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedMetaEggLayoutServer) DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedMetaEggLayoutServer) mustEmbedUnimplementedMetaEggLayoutServer() {}

// UnsafeMetaEggLayoutServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetaEggLayoutServer will
// result in compilation errors.
type UnsafeMetaEggLayoutServer interface {
	mustEmbedUnimplementedMetaEggLayoutServer()
}

func RegisterMetaEggLayoutServer(s grpc.ServiceRegistrar, srv MetaEggLayoutServer) {
	s.RegisterService(&MetaEggLayout_ServiceDesc, srv)
}

func _MetaEggLayout_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaEggLayoutServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetaEggLayout_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaEggLayoutServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaEggLayout_GetUserDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaEggLayoutServer).GetUserDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetaEggLayout_GetUserDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaEggLayoutServer).GetUserDetail(ctx, req.(*GetUserDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaEggLayout_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaEggLayoutServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetaEggLayout_GetUserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaEggLayoutServer).GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaEggLayout_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaEggLayoutServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetaEggLayout_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaEggLayoutServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaEggLayout_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaEggLayoutServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetaEggLayout_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaEggLayoutServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetaEggLayout_ServiceDesc is the grpc.ServiceDesc for MetaEggLayout service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetaEggLayout_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metaegglayout.MetaEggLayout",
	HandlerType: (*MetaEggLayoutServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _MetaEggLayout_CreateUser_Handler,
		},
		{
			MethodName: "GetUserDetail",
			Handler:    _MetaEggLayout_GetUserDetail_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _MetaEggLayout_GetUserList_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _MetaEggLayout_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _MetaEggLayout_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meta_egg_layout.proto",
}
