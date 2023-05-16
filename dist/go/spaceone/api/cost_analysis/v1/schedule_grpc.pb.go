// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: spaceone/api/cost_analysis/v1/schedule.proto

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
	Schedule_Create_FullMethodName  = "/spaceone.api.cost_analysis.v1.Schedule/create"
	Schedule_Update_FullMethodName  = "/spaceone.api.cost_analysis.v1.Schedule/update"
	Schedule_Enable_FullMethodName  = "/spaceone.api.cost_analysis.v1.Schedule/enable"
	Schedule_Disable_FullMethodName = "/spaceone.api.cost_analysis.v1.Schedule/disable"
	Schedule_Delete_FullMethodName  = "/spaceone.api.cost_analysis.v1.Schedule/delete"
	Schedule_Get_FullMethodName     = "/spaceone.api.cost_analysis.v1.Schedule/get"
	Schedule_List_FullMethodName    = "/spaceone.api.cost_analysis.v1.Schedule/list"
	Schedule_Stat_FullMethodName    = "/spaceone.api.cost_analysis.v1.Schedule/stat"
)

// ScheduleClient is the client API for Schedule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScheduleClient interface {
	Create(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*ScheduleInfo, error)
	Update(ctx context.Context, in *UpdateScheduleRequest, opts ...grpc.CallOption) (*ScheduleInfo, error)
	Enable(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleInfo, error)
	Disable(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleInfo, error)
	Delete(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*ScheduleInfo, error)
	List(ctx context.Context, in *ScheduleQuery, opts ...grpc.CallOption) (*SchedulesInfo, error)
	Stat(ctx context.Context, in *ScheduleStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type scheduleClient struct {
	cc grpc.ClientConnInterface
}

func NewScheduleClient(cc grpc.ClientConnInterface) ScheduleClient {
	return &scheduleClient{cc}
}

func (c *scheduleClient) Create(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*ScheduleInfo, error) {
	out := new(ScheduleInfo)
	err := c.cc.Invoke(ctx, Schedule_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) Update(ctx context.Context, in *UpdateScheduleRequest, opts ...grpc.CallOption) (*ScheduleInfo, error) {
	out := new(ScheduleInfo)
	err := c.cc.Invoke(ctx, Schedule_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) Enable(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleInfo, error) {
	out := new(ScheduleInfo)
	err := c.cc.Invoke(ctx, Schedule_Enable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) Disable(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleInfo, error) {
	out := new(ScheduleInfo)
	err := c.cc.Invoke(ctx, Schedule_Disable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) Delete(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Schedule_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) Get(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*ScheduleInfo, error) {
	out := new(ScheduleInfo)
	err := c.cc.Invoke(ctx, Schedule_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) List(ctx context.Context, in *ScheduleQuery, opts ...grpc.CallOption) (*SchedulesInfo, error) {
	out := new(SchedulesInfo)
	err := c.cc.Invoke(ctx, Schedule_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) Stat(ctx context.Context, in *ScheduleStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Schedule_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScheduleServer is the server API for Schedule service.
// All implementations must embed UnimplementedScheduleServer
// for forward compatibility
type ScheduleServer interface {
	Create(context.Context, *CreateScheduleRequest) (*ScheduleInfo, error)
	Update(context.Context, *UpdateScheduleRequest) (*ScheduleInfo, error)
	Enable(context.Context, *ScheduleRequest) (*ScheduleInfo, error)
	Disable(context.Context, *ScheduleRequest) (*ScheduleInfo, error)
	Delete(context.Context, *ScheduleRequest) (*empty.Empty, error)
	Get(context.Context, *GetScheduleRequest) (*ScheduleInfo, error)
	List(context.Context, *ScheduleQuery) (*SchedulesInfo, error)
	Stat(context.Context, *ScheduleStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedScheduleServer()
}

// UnimplementedScheduleServer must be embedded to have forward compatible implementations.
type UnimplementedScheduleServer struct {
}

func (UnimplementedScheduleServer) Create(context.Context, *CreateScheduleRequest) (*ScheduleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedScheduleServer) Update(context.Context, *UpdateScheduleRequest) (*ScheduleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedScheduleServer) Enable(context.Context, *ScheduleRequest) (*ScheduleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedScheduleServer) Disable(context.Context, *ScheduleRequest) (*ScheduleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedScheduleServer) Delete(context.Context, *ScheduleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedScheduleServer) Get(context.Context, *GetScheduleRequest) (*ScheduleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedScheduleServer) List(context.Context, *ScheduleQuery) (*SchedulesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedScheduleServer) Stat(context.Context, *ScheduleStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedScheduleServer) mustEmbedUnimplementedScheduleServer() {}

// UnsafeScheduleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScheduleServer will
// result in compilation errors.
type UnsafeScheduleServer interface {
	mustEmbedUnimplementedScheduleServer()
}

func RegisterScheduleServer(s grpc.ServiceRegistrar, srv ScheduleServer) {
	s.RegisterService(&Schedule_ServiceDesc, srv)
}

func _Schedule_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Schedule_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).Create(ctx, req.(*CreateScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Schedule_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).Update(ctx, req.(*UpdateScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Schedule_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).Enable(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Schedule_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).Disable(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Schedule_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).Delete(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Schedule_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).Get(ctx, req.(*GetScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Schedule_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).List(ctx, req.(*ScheduleQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Schedule_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).Stat(ctx, req.(*ScheduleStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Schedule_ServiceDesc is the grpc.ServiceDesc for Schedule service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Schedule_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.cost_analysis.v1.Schedule",
	HandlerType: (*ScheduleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Schedule_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Schedule_Update_Handler,
		},
		{
			MethodName: "enable",
			Handler:    _Schedule_Enable_Handler,
		},
		{
			MethodName: "disable",
			Handler:    _Schedule_Disable_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Schedule_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Schedule_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Schedule_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Schedule_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/cost_analysis/v1/schedule.proto",
}
