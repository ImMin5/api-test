// A ProjectAlertConfig is a resource to set the alert policies for a Project.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: spaceone/api/monitoring/v1/project_alert_config.proto

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
	ProjectAlertConfig_Create_FullMethodName = "/spaceone.api.monitoring.v1.ProjectAlertConfig/create"
	ProjectAlertConfig_Update_FullMethodName = "/spaceone.api.monitoring.v1.ProjectAlertConfig/update"
	ProjectAlertConfig_Delete_FullMethodName = "/spaceone.api.monitoring.v1.ProjectAlertConfig/delete"
	ProjectAlertConfig_Get_FullMethodName    = "/spaceone.api.monitoring.v1.ProjectAlertConfig/get"
	ProjectAlertConfig_List_FullMethodName   = "/spaceone.api.monitoring.v1.ProjectAlertConfig/list"
	ProjectAlertConfig_Stat_FullMethodName   = "/spaceone.api.monitoring.v1.ProjectAlertConfig/stat"
)

// ProjectAlertConfigClient is the client API for ProjectAlertConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectAlertConfigClient interface {
	// Creates a new ProjectAlertConfig in a specific Project. When creating a ProjectAlertConfig, validation of the Project is preceded. After the validation is done, ProjectAlertConfig enables EscalationPolicy to be set in the Project, or enables `enum` type `recovery_mode` and `notification_urgency` to be set through the `options` parameter.  The parameter `recovery_mode` is for changing the state of the Alert to `resolved` if the external monitoring solution sends the resolved Alert. The parameter `notification_urgency` is used to choose whether you will get all Alerts or only urgent ones.
	Create(ctx context.Context, in *CreateProjectAlertConfigRequest, opts ...grpc.CallOption) (*ProjectAlertConfigInfo, error)
	// Updates a specific ProjectAlertConfig. You can make changes in ProjectAlertConfig settings, including the EscalationPolicy to apply. You can also change `notification_urgency` and `recovery_mode` by modifying the `options` parameter.
	Update(ctx context.Context, in *UpdateProjectAlertConfigRequest, opts ...grpc.CallOption) (*ProjectAlertConfigInfo, error)
	// Deletes a specific ProjectAlertConfig. Deletes alert configuration data in a Project.
	Delete(ctx context.Context, in *ProjectAlertConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets a specific ProjectAlertConfig. Prints detailed information about the ProjectAlertConfig, including EscalationPolicy, recovery mode, and notification urgency.
	Get(ctx context.Context, in *GetProjectAlertConfigRequest, opts ...grpc.CallOption) (*ProjectAlertConfigInfo, error)
	// Gets a list of all ProjectAlertConfigs from all projects configured in the same domain. You can use a query to get a filtered list of ProjectAlertConfigs.
	List(ctx context.Context, in *ProjectAlertConfigQuery, opts ...grpc.CallOption) (*ProjectAlertConfigsInfo, error)
	Stat(ctx context.Context, in *ProjectAlertConfigStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type projectAlertConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectAlertConfigClient(cc grpc.ClientConnInterface) ProjectAlertConfigClient {
	return &projectAlertConfigClient{cc}
}

func (c *projectAlertConfigClient) Create(ctx context.Context, in *CreateProjectAlertConfigRequest, opts ...grpc.CallOption) (*ProjectAlertConfigInfo, error) {
	out := new(ProjectAlertConfigInfo)
	err := c.cc.Invoke(ctx, ProjectAlertConfig_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAlertConfigClient) Update(ctx context.Context, in *UpdateProjectAlertConfigRequest, opts ...grpc.CallOption) (*ProjectAlertConfigInfo, error) {
	out := new(ProjectAlertConfigInfo)
	err := c.cc.Invoke(ctx, ProjectAlertConfig_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAlertConfigClient) Delete(ctx context.Context, in *ProjectAlertConfigRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectAlertConfig_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAlertConfigClient) Get(ctx context.Context, in *GetProjectAlertConfigRequest, opts ...grpc.CallOption) (*ProjectAlertConfigInfo, error) {
	out := new(ProjectAlertConfigInfo)
	err := c.cc.Invoke(ctx, ProjectAlertConfig_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAlertConfigClient) List(ctx context.Context, in *ProjectAlertConfigQuery, opts ...grpc.CallOption) (*ProjectAlertConfigsInfo, error) {
	out := new(ProjectAlertConfigsInfo)
	err := c.cc.Invoke(ctx, ProjectAlertConfig_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAlertConfigClient) Stat(ctx context.Context, in *ProjectAlertConfigStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, ProjectAlertConfig_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectAlertConfigServer is the server API for ProjectAlertConfig service.
// All implementations must embed UnimplementedProjectAlertConfigServer
// for forward compatibility
type ProjectAlertConfigServer interface {
	// Creates a new ProjectAlertConfig in a specific Project. When creating a ProjectAlertConfig, validation of the Project is preceded. After the validation is done, ProjectAlertConfig enables EscalationPolicy to be set in the Project, or enables `enum` type `recovery_mode` and `notification_urgency` to be set through the `options` parameter.  The parameter `recovery_mode` is for changing the state of the Alert to `resolved` if the external monitoring solution sends the resolved Alert. The parameter `notification_urgency` is used to choose whether you will get all Alerts or only urgent ones.
	Create(context.Context, *CreateProjectAlertConfigRequest) (*ProjectAlertConfigInfo, error)
	// Updates a specific ProjectAlertConfig. You can make changes in ProjectAlertConfig settings, including the EscalationPolicy to apply. You can also change `notification_urgency` and `recovery_mode` by modifying the `options` parameter.
	Update(context.Context, *UpdateProjectAlertConfigRequest) (*ProjectAlertConfigInfo, error)
	// Deletes a specific ProjectAlertConfig. Deletes alert configuration data in a Project.
	Delete(context.Context, *ProjectAlertConfigRequest) (*empty.Empty, error)
	// Gets a specific ProjectAlertConfig. Prints detailed information about the ProjectAlertConfig, including EscalationPolicy, recovery mode, and notification urgency.
	Get(context.Context, *GetProjectAlertConfigRequest) (*ProjectAlertConfigInfo, error)
	// Gets a list of all ProjectAlertConfigs from all projects configured in the same domain. You can use a query to get a filtered list of ProjectAlertConfigs.
	List(context.Context, *ProjectAlertConfigQuery) (*ProjectAlertConfigsInfo, error)
	Stat(context.Context, *ProjectAlertConfigStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedProjectAlertConfigServer()
}

// UnimplementedProjectAlertConfigServer must be embedded to have forward compatible implementations.
type UnimplementedProjectAlertConfigServer struct {
}

func (UnimplementedProjectAlertConfigServer) Create(context.Context, *CreateProjectAlertConfigRequest) (*ProjectAlertConfigInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProjectAlertConfigServer) Update(context.Context, *UpdateProjectAlertConfigRequest) (*ProjectAlertConfigInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProjectAlertConfigServer) Delete(context.Context, *ProjectAlertConfigRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProjectAlertConfigServer) Get(context.Context, *GetProjectAlertConfigRequest) (*ProjectAlertConfigInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedProjectAlertConfigServer) List(context.Context, *ProjectAlertConfigQuery) (*ProjectAlertConfigsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProjectAlertConfigServer) Stat(context.Context, *ProjectAlertConfigStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedProjectAlertConfigServer) mustEmbedUnimplementedProjectAlertConfigServer() {}

// UnsafeProjectAlertConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectAlertConfigServer will
// result in compilation errors.
type UnsafeProjectAlertConfigServer interface {
	mustEmbedUnimplementedProjectAlertConfigServer()
}

func RegisterProjectAlertConfigServer(s grpc.ServiceRegistrar, srv ProjectAlertConfigServer) {
	s.RegisterService(&ProjectAlertConfig_ServiceDesc, srv)
}

func _ProjectAlertConfig_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectAlertConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAlertConfigServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAlertConfig_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAlertConfigServer).Create(ctx, req.(*CreateProjectAlertConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAlertConfig_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectAlertConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAlertConfigServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAlertConfig_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAlertConfigServer).Update(ctx, req.(*UpdateProjectAlertConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAlertConfig_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectAlertConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAlertConfigServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAlertConfig_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAlertConfigServer).Delete(ctx, req.(*ProjectAlertConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAlertConfig_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectAlertConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAlertConfigServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAlertConfig_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAlertConfigServer).Get(ctx, req.(*GetProjectAlertConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAlertConfig_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectAlertConfigQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAlertConfigServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAlertConfig_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAlertConfigServer).List(ctx, req.(*ProjectAlertConfigQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAlertConfig_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectAlertConfigStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAlertConfigServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAlertConfig_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAlertConfigServer).Stat(ctx, req.(*ProjectAlertConfigStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectAlertConfig_ServiceDesc is the grpc.ServiceDesc for ProjectAlertConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectAlertConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.monitoring.v1.ProjectAlertConfig",
	HandlerType: (*ProjectAlertConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _ProjectAlertConfig_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _ProjectAlertConfig_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _ProjectAlertConfig_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _ProjectAlertConfig_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _ProjectAlertConfig_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _ProjectAlertConfig_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/monitoring/v1/project_alert_config.proto",
}
