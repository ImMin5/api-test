// description of dashboard

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: spaceone/api/dashboard/v1/custom_widget.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CustomWidget_Create_FullMethodName = "/spaceone.api.dashboard.v1.CustomWidget/create"
	CustomWidget_Update_FullMethodName = "/spaceone.api.dashboard.v1.CustomWidget/update"
	CustomWidget_Delete_FullMethodName = "/spaceone.api.dashboard.v1.CustomWidget/delete"
	CustomWidget_Get_FullMethodName    = "/spaceone.api.dashboard.v1.CustomWidget/get"
	CustomWidget_List_FullMethodName   = "/spaceone.api.dashboard.v1.CustomWidget/list"
	CustomWidget_Stat_FullMethodName   = "/spaceone.api.dashboard.v1.CustomWidget/stat"
)

// CustomWidgetClient is the client API for CustomWidget service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomWidgetClient interface {
	Create(ctx context.Context, in *CreateCustomWidgetRequest, opts ...grpc.CallOption) (*CustomWidgetInfo, error)
	Update(ctx context.Context, in *UpdateCustomWidgetRequest, opts ...grpc.CallOption) (*CustomWidgetInfo, error)
	Delete(ctx context.Context, in *CustomWidgetRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *GetCustomWidgetRequest, opts ...grpc.CallOption) (*CustomWidgetInfo, error)
	List(ctx context.Context, in *CustomWidgetQuery, opts ...grpc.CallOption) (*CustomWidgetsInfo, error)
	Stat(ctx context.Context, in *CustomWidgetStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type customWidgetClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomWidgetClient(cc grpc.ClientConnInterface) CustomWidgetClient {
	return &customWidgetClient{cc}
}

func (c *customWidgetClient) Create(ctx context.Context, in *CreateCustomWidgetRequest, opts ...grpc.CallOption) (*CustomWidgetInfo, error) {
	out := new(CustomWidgetInfo)
	err := c.cc.Invoke(ctx, CustomWidget_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customWidgetClient) Update(ctx context.Context, in *UpdateCustomWidgetRequest, opts ...grpc.CallOption) (*CustomWidgetInfo, error) {
	out := new(CustomWidgetInfo)
	err := c.cc.Invoke(ctx, CustomWidget_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customWidgetClient) Delete(ctx context.Context, in *CustomWidgetRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, CustomWidget_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customWidgetClient) Get(ctx context.Context, in *GetCustomWidgetRequest, opts ...grpc.CallOption) (*CustomWidgetInfo, error) {
	out := new(CustomWidgetInfo)
	err := c.cc.Invoke(ctx, CustomWidget_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customWidgetClient) List(ctx context.Context, in *CustomWidgetQuery, opts ...grpc.CallOption) (*CustomWidgetsInfo, error) {
	out := new(CustomWidgetsInfo)
	err := c.cc.Invoke(ctx, CustomWidget_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customWidgetClient) Stat(ctx context.Context, in *CustomWidgetStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, CustomWidget_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomWidgetServer is the server API for CustomWidget service.
// All implementations must embed UnimplementedCustomWidgetServer
// for forward compatibility
type CustomWidgetServer interface {
	Create(context.Context, *CreateCustomWidgetRequest) (*CustomWidgetInfo, error)
	Update(context.Context, *UpdateCustomWidgetRequest) (*CustomWidgetInfo, error)
	Delete(context.Context, *CustomWidgetRequest) (*empty.Empty, error)
	Get(context.Context, *GetCustomWidgetRequest) (*CustomWidgetInfo, error)
	List(context.Context, *CustomWidgetQuery) (*CustomWidgetsInfo, error)
	Stat(context.Context, *CustomWidgetStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedCustomWidgetServer()
}

// UnimplementedCustomWidgetServer must be embedded to have forward compatible implementations.
type UnimplementedCustomWidgetServer struct {
}

func (UnimplementedCustomWidgetServer) Create(context.Context, *CreateCustomWidgetRequest) (*CustomWidgetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCustomWidgetServer) Update(context.Context, *UpdateCustomWidgetRequest) (*CustomWidgetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCustomWidgetServer) Delete(context.Context, *CustomWidgetRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCustomWidgetServer) Get(context.Context, *GetCustomWidgetRequest) (*CustomWidgetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCustomWidgetServer) List(context.Context, *CustomWidgetQuery) (*CustomWidgetsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCustomWidgetServer) Stat(context.Context, *CustomWidgetStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedCustomWidgetServer) mustEmbedUnimplementedCustomWidgetServer() {}

// UnsafeCustomWidgetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomWidgetServer will
// result in compilation errors.
type UnsafeCustomWidgetServer interface {
	mustEmbedUnimplementedCustomWidgetServer()
}

func RegisterCustomWidgetServer(s grpc.ServiceRegistrar, srv CustomWidgetServer) {
	s.RegisterService(&CustomWidget_ServiceDesc, srv)
}

func _CustomWidget_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomWidgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomWidgetServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomWidget_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomWidgetServer).Create(ctx, req.(*CreateCustomWidgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomWidget_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomWidgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomWidgetServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomWidget_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomWidgetServer).Update(ctx, req.(*UpdateCustomWidgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomWidget_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomWidgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomWidgetServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomWidget_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomWidgetServer).Delete(ctx, req.(*CustomWidgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomWidget_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomWidgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomWidgetServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomWidget_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomWidgetServer).Get(ctx, req.(*GetCustomWidgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomWidget_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomWidgetQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomWidgetServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomWidget_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomWidgetServer).List(ctx, req.(*CustomWidgetQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomWidget_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomWidgetStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomWidgetServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomWidget_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomWidgetServer).Stat(ctx, req.(*CustomWidgetStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomWidget_ServiceDesc is the grpc.ServiceDesc for CustomWidget service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomWidget_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.dashboard.v1.CustomWidget",
	HandlerType: (*CustomWidgetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _CustomWidget_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _CustomWidget_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _CustomWidget_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _CustomWidget_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _CustomWidget_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _CustomWidget_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/dashboard/v1/custom_widget.proto",
}
