// A Supervisor is a resource managing the lifecycle of the plugin instances to deploy. A Supervisor manages the deployment of plugin instances by deploying or deleting the `pod` of the plugin instances.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: spaceone/api/plugin/v1/supervisor.proto

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
	Supervisor_Publish_FullMethodName       = "/spaceone.api.plugin.v1.Supervisor/publish"
	Supervisor_Register_FullMethodName      = "/spaceone.api.plugin.v1.Supervisor/register"
	Supervisor_Update_FullMethodName        = "/spaceone.api.plugin.v1.Supervisor/update"
	Supervisor_Deregister_FullMethodName    = "/spaceone.api.plugin.v1.Supervisor/deregister"
	Supervisor_Enable_FullMethodName        = "/spaceone.api.plugin.v1.Supervisor/enable"
	Supervisor_Disable_FullMethodName       = "/spaceone.api.plugin.v1.Supervisor/disable"
	Supervisor_RecoverPlugin_FullMethodName = "/spaceone.api.plugin.v1.Supervisor/recover_plugin"
	Supervisor_Get_FullMethodName           = "/spaceone.api.plugin.v1.Supervisor/get"
	Supervisor_List_FullMethodName          = "/spaceone.api.plugin.v1.Supervisor/list"
	Supervisor_Stat_FullMethodName          = "/spaceone.api.plugin.v1.Supervisor/stat"
	Supervisor_ListPlugins_FullMethodName   = "/spaceone.api.plugin.v1.Supervisor/list_plugins"
)

// SupervisorClient is the client API for Supervisor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupervisorClient interface {
	// Creates a new Supervisor. Only Users with the `MANAGED` permission can set the Supervisor `public`. The Supervisor manages the lifecycle of plugin instances by the Supervisor's state. When a Supervisor is created, the state of the resource is `PENDING`. If the state remains the same for 5 minutes, the state is changed to `DISCONNECTED`.
	Publish(ctx context.Context, in *PublishSupervisorRequest, opts ...grpc.CallOption) (*SupervisorInfo, error)
	// Registers a specific Supervisor. You must specify the `supervisor_id` of the Supervisor to register. The `state` of the Supervisor changes from `PENDING` to `ENABLED`.
	Register(ctx context.Context, in *RegisterSupervisorRequest, opts ...grpc.CallOption) (*SupervisorInfo, error)
	// Updates a specific Supervisor. You can make changes in Supervisor settings, including `labels`, `tags`, and the `bool` type parameter `is_public`.
	Update(ctx context.Context, in *RegisterSupervisorRequest, opts ...grpc.CallOption) (*SupervisorInfo, error)
	// Deregisters and deletes a specific Supervisor. You must specify the `supervisor_id` of the Supervisor to deregister.
	Deregister(ctx context.Context, in *SupervisorRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Enables a specific Supervisor. By changing the `state` parameter to `ENABLED`, the Supervisor can deploy or delete the `pod` of the plugin instance.
	Enable(ctx context.Context, in *SupervisorRequest, opts ...grpc.CallOption) (*SupervisorInfo, error)
	// Disables a specific Supervisor. By changing the `state` parameter to `DISABLED`, the Supervisor cannot deploy or delete the `pod` of the plugin instance.
	Disable(ctx context.Context, in *SupervisorRequest, opts ...grpc.CallOption) (*SupervisorInfo, error)
	// Recovers a specific plugin instance in a specific Supervisor. Changes the `state` of the Supervisor to `RE-PROVISIONING`.
	RecoverPlugin(ctx context.Context, in *RecoverPluginRequest, opts ...grpc.CallOption) (*PluginInfo, error)
	Get(ctx context.Context, in *GetSupervisorRequest, opts ...grpc.CallOption) (*SupervisorInfo, error)
	// Gets a list of all Supervisors. You can use a query to get a filtered list of Supervisors.
	List(ctx context.Context, in *SupervisorQuery, opts ...grpc.CallOption) (*SupervisorsInfo, error)
	Stat(ctx context.Context, in *SupervisorStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
	// Gets a list of all plugin instances regardless of Supervisors. Prints detailed information about the plugin instances, including `version`, `state`, and the relevant Supervisor.
	ListPlugins(ctx context.Context, in *PluginQuery, opts ...grpc.CallOption) (*PluginsInfo, error)
}

type supervisorClient struct {
	cc grpc.ClientConnInterface
}

func NewSupervisorClient(cc grpc.ClientConnInterface) SupervisorClient {
	return &supervisorClient{cc}
}

func (c *supervisorClient) Publish(ctx context.Context, in *PublishSupervisorRequest, opts ...grpc.CallOption) (*SupervisorInfo, error) {
	out := new(SupervisorInfo)
	err := c.cc.Invoke(ctx, Supervisor_Publish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Register(ctx context.Context, in *RegisterSupervisorRequest, opts ...grpc.CallOption) (*SupervisorInfo, error) {
	out := new(SupervisorInfo)
	err := c.cc.Invoke(ctx, Supervisor_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Update(ctx context.Context, in *RegisterSupervisorRequest, opts ...grpc.CallOption) (*SupervisorInfo, error) {
	out := new(SupervisorInfo)
	err := c.cc.Invoke(ctx, Supervisor_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Deregister(ctx context.Context, in *SupervisorRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Supervisor_Deregister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Enable(ctx context.Context, in *SupervisorRequest, opts ...grpc.CallOption) (*SupervisorInfo, error) {
	out := new(SupervisorInfo)
	err := c.cc.Invoke(ctx, Supervisor_Enable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Disable(ctx context.Context, in *SupervisorRequest, opts ...grpc.CallOption) (*SupervisorInfo, error) {
	out := new(SupervisorInfo)
	err := c.cc.Invoke(ctx, Supervisor_Disable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) RecoverPlugin(ctx context.Context, in *RecoverPluginRequest, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := c.cc.Invoke(ctx, Supervisor_RecoverPlugin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Get(ctx context.Context, in *GetSupervisorRequest, opts ...grpc.CallOption) (*SupervisorInfo, error) {
	out := new(SupervisorInfo)
	err := c.cc.Invoke(ctx, Supervisor_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) List(ctx context.Context, in *SupervisorQuery, opts ...grpc.CallOption) (*SupervisorsInfo, error) {
	out := new(SupervisorsInfo)
	err := c.cc.Invoke(ctx, Supervisor_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Stat(ctx context.Context, in *SupervisorStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Supervisor_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) ListPlugins(ctx context.Context, in *PluginQuery, opts ...grpc.CallOption) (*PluginsInfo, error) {
	out := new(PluginsInfo)
	err := c.cc.Invoke(ctx, Supervisor_ListPlugins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupervisorServer is the server API for Supervisor service.
// All implementations must embed UnimplementedSupervisorServer
// for forward compatibility
type SupervisorServer interface {
	// Creates a new Supervisor. Only Users with the `MANAGED` permission can set the Supervisor `public`. The Supervisor manages the lifecycle of plugin instances by the Supervisor's state. When a Supervisor is created, the state of the resource is `PENDING`. If the state remains the same for 5 minutes, the state is changed to `DISCONNECTED`.
	Publish(context.Context, *PublishSupervisorRequest) (*SupervisorInfo, error)
	// Registers a specific Supervisor. You must specify the `supervisor_id` of the Supervisor to register. The `state` of the Supervisor changes from `PENDING` to `ENABLED`.
	Register(context.Context, *RegisterSupervisorRequest) (*SupervisorInfo, error)
	// Updates a specific Supervisor. You can make changes in Supervisor settings, including `labels`, `tags`, and the `bool` type parameter `is_public`.
	Update(context.Context, *RegisterSupervisorRequest) (*SupervisorInfo, error)
	// Deregisters and deletes a specific Supervisor. You must specify the `supervisor_id` of the Supervisor to deregister.
	Deregister(context.Context, *SupervisorRequest) (*empty.Empty, error)
	// Enables a specific Supervisor. By changing the `state` parameter to `ENABLED`, the Supervisor can deploy or delete the `pod` of the plugin instance.
	Enable(context.Context, *SupervisorRequest) (*SupervisorInfo, error)
	// Disables a specific Supervisor. By changing the `state` parameter to `DISABLED`, the Supervisor cannot deploy or delete the `pod` of the plugin instance.
	Disable(context.Context, *SupervisorRequest) (*SupervisorInfo, error)
	// Recovers a specific plugin instance in a specific Supervisor. Changes the `state` of the Supervisor to `RE-PROVISIONING`.
	RecoverPlugin(context.Context, *RecoverPluginRequest) (*PluginInfo, error)
	Get(context.Context, *GetSupervisorRequest) (*SupervisorInfo, error)
	// Gets a list of all Supervisors. You can use a query to get a filtered list of Supervisors.
	List(context.Context, *SupervisorQuery) (*SupervisorsInfo, error)
	Stat(context.Context, *SupervisorStatQuery) (*_struct.Struct, error)
	// Gets a list of all plugin instances regardless of Supervisors. Prints detailed information about the plugin instances, including `version`, `state`, and the relevant Supervisor.
	ListPlugins(context.Context, *PluginQuery) (*PluginsInfo, error)
	mustEmbedUnimplementedSupervisorServer()
}

// UnimplementedSupervisorServer must be embedded to have forward compatible implementations.
type UnimplementedSupervisorServer struct {
}

func (UnimplementedSupervisorServer) Publish(context.Context, *PublishSupervisorRequest) (*SupervisorInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedSupervisorServer) Register(context.Context, *RegisterSupervisorRequest) (*SupervisorInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedSupervisorServer) Update(context.Context, *RegisterSupervisorRequest) (*SupervisorInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSupervisorServer) Deregister(context.Context, *SupervisorRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deregister not implemented")
}
func (UnimplementedSupervisorServer) Enable(context.Context, *SupervisorRequest) (*SupervisorInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedSupervisorServer) Disable(context.Context, *SupervisorRequest) (*SupervisorInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedSupervisorServer) RecoverPlugin(context.Context, *RecoverPluginRequest) (*PluginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverPlugin not implemented")
}
func (UnimplementedSupervisorServer) Get(context.Context, *GetSupervisorRequest) (*SupervisorInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSupervisorServer) List(context.Context, *SupervisorQuery) (*SupervisorsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSupervisorServer) Stat(context.Context, *SupervisorStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedSupervisorServer) ListPlugins(context.Context, *PluginQuery) (*PluginsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlugins not implemented")
}
func (UnimplementedSupervisorServer) mustEmbedUnimplementedSupervisorServer() {}

// UnsafeSupervisorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupervisorServer will
// result in compilation errors.
type UnsafeSupervisorServer interface {
	mustEmbedUnimplementedSupervisorServer()
}

func RegisterSupervisorServer(s grpc.ServiceRegistrar, srv SupervisorServer) {
	s.RegisterService(&Supervisor_ServiceDesc, srv)
}

func _Supervisor_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishSupervisorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Supervisor_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Publish(ctx, req.(*PublishSupervisorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterSupervisorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Supervisor_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Register(ctx, req.(*RegisterSupervisorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterSupervisorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Supervisor_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Update(ctx, req.(*RegisterSupervisorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupervisorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Supervisor_Deregister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Deregister(ctx, req.(*SupervisorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupervisorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Supervisor_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Enable(ctx, req.(*SupervisorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupervisorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Supervisor_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Disable(ctx, req.(*SupervisorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_RecoverPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).RecoverPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Supervisor_RecoverPlugin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).RecoverPlugin(ctx, req.(*RecoverPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSupervisorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Supervisor_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Get(ctx, req.(*GetSupervisorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupervisorQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Supervisor_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).List(ctx, req.(*SupervisorQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupervisorStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Supervisor_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Stat(ctx, req.(*SupervisorStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_ListPlugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).ListPlugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Supervisor_ListPlugins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).ListPlugins(ctx, req.(*PluginQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Supervisor_ServiceDesc is the grpc.ServiceDesc for Supervisor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Supervisor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.plugin.v1.Supervisor",
	HandlerType: (*SupervisorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "publish",
			Handler:    _Supervisor_Publish_Handler,
		},
		{
			MethodName: "register",
			Handler:    _Supervisor_Register_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Supervisor_Update_Handler,
		},
		{
			MethodName: "deregister",
			Handler:    _Supervisor_Deregister_Handler,
		},
		{
			MethodName: "enable",
			Handler:    _Supervisor_Enable_Handler,
		},
		{
			MethodName: "disable",
			Handler:    _Supervisor_Disable_Handler,
		},
		{
			MethodName: "recover_plugin",
			Handler:    _Supervisor_RecoverPlugin_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Supervisor_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Supervisor_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Supervisor_Stat_Handler,
		},
		{
			MethodName: "list_plugins",
			Handler:    _Supervisor_ListPlugins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/plugin/v1/supervisor.proto",
}