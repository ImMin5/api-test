// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: spaceone/api/inventory/v1/resource_group.proto

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

// ResourceGroupClient is the client API for ResourceGroup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceGroupClient interface {
	// desc: Creates a new ResourceGroup. You can integrate `resource`s from different `provider`s by specifying the cloud resources to be grouped in the `resources` parameter.
	// request_example: >-
	// {
	// "name": "azure-group-1",
	// "resources": [
	// {
	// "resource_type": "inventory.Server?provider=azure&cloud_service_group=Compute&cloud_service_type=VirtualMachine",
	// "filter": [
	// {"o": "eq", "k": "provider", "v": "azure"},
	// {"o": "eq", "k": "cloud_service_group", "v": "Compute"},
	// {"o": "eq", "k": "cloud_service_type", "v": "VirtualMachine"}
	// ]
	// }
	// ],
	// "options": {
	// "raw_filter": []
	// },
	// "tags": {
	// "a": "b"
	// }
	// }
	// response_example: >-
	// {
	// "resource_group_id": "rsc-grp-7d46a1fc7738",
	// "name": "azure-group-1",
	// "resources": [
	// {
	// "resource_type": "inventory.Server?provider=azure&cloud_service_group=Compute&cloud_service_type=VirtualMachine",
	// "filter": [
	// {
	// "k": "provider",
	// "v": "azure",
	// "o": "eq"
	// },
	// {
	// "k": "cloud_service_group",
	// "v": "Compute",
	// "o": "eq"
	// },
	// {
	// "k": "cloud_service_type",
	// "v": "VirtualMachine",
	// "o": "eq"
	// }
	// ]
	// }
	// ],
	// "options": {
	// "raw_filter": []
	// },
	// "tags": {
	// "a": "b"
	// },
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-06-23T01:50:33.152Z"
	// }
	Create(ctx context.Context, in *CreateResourceGroupRequest, opts ...grpc.CallOption) (*ResourceGroupInfo, error)
	// desc: Updates a specific ResourceGroup. You can make changes in ResourceGroup settings, including `name`, `tags`, and grouped resources in the ResourceGroup.
	// request_example: >-
	// {
	// "resource_group_id": "rsc-grp-7d46a1fc7738",
	// "name": "azure-grp-test2",
	// "resources": [
	// {
	// "resource_type": "inventory.Server?provider=azure&cloud_service_group=Compute&cloud_service_type=VirtualMachine",
	// "filter": [
	// {"k": "provider", "v": "azure", "o": "eq"},
	// {"o": "eq", "k": "cloud_service_group", "v": "Compute"},
	// {"v": "VirtualMachine", "k": "cloud_service_type", "o": "eq"}
	// ]
	// }
	// ],
	// "options": {},
	// "tags": {
	// "b": "c"
	// }
	// }
	// response_example: >-
	// {
	// "resource_group_id": "rsc-grp-7d46a1fc7738",
	// "name": "azure-grp-test2",
	// "resources": [
	// {
	// "resource_type": "inventory.Server?provider=azure&cloud_service_group=Compute&cloud_service_type=VirtualMachine",
	// "filter": [
	// {
	// "k": "provider",
	// "v": "azure",
	// "o": "eq"
	// },
	// {
	// "k": "cloud_service_group",
	// "v": "Compute",
	// "o": "eq"
	// },
	// {
	// "k": "cloud_service_type",
	// "v": "VirtualMachine",
	// "o": "eq"
	// }
	// ]
	// }
	// ],
	// "options": {
	// "raw_filter": []
	// },
	// "tags": {
	// "a": "b"
	// },
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-06-23T01:50:33.152Z"
	// }
	Update(ctx context.Context, in *UpdateResourceGroupRequest, opts ...grpc.CallOption) (*ResourceGroupInfo, error)
	// desc: Deletes a specific ResourceGroup. You must specify the `resource_group_id` of the ResourceGroup to delete.
	// request_example: >-
	// {
	// "resource_group_id": "rsc-grp-aa3c4ca465b2"
	// }
	Delete(ctx context.Context, in *ResourceGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// desc: Gets a specific ResourceGroup. Prints detailed information about the ResourceGroup, including the information of the grouped cloud resources.
	// request_example: >-
	// {
	// "resource_group_id": "rsc-grp-aa3c4ca465b2"
	// }
	// response_example: >-
	// {
	// "resource_group_id": "rsc-grp-aa3c4ca465b2",
	// "name": "stset",
	// "resources": [
	// {
	// "resource_type": "inventory.Server?provider=aws&cloud_service_group=EC2&cloud_service_type=Instance",
	// "filter": [
	// {
	// "k": "provider",
	// "o": "eq",
	// "v": "aws"
	// },
	// {
	// "v": "EC2",
	// "k": "cloud_service_group",
	// "o": "eq"
	// },
	// {
	// "o": "eq",
	// "k": "cloud_service_type",
	// "v": "Instance"
	// }
	// ],
	// "keyword": "test"
	// }
	// ],
	// "options": {
	// "raw_filter": []
	// },
	// "tags": {},
	// "project_id": "project-18655561c535",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2021-06-01T10:23:20.537Z"
	// }
	Get(ctx context.Context, in *GetResourceGroupRequest, opts ...grpc.CallOption) (*ResourceGroupInfo, error)
	// desc: Gets a list of all ResourceGroups. You can use a query to get a filtered list of ResourceGroups.
	// request_example: >-
	// {
	// "query": {
	// "filter": [
	// {
	// "key": "name",
	// "value": [
	// "azure-vmss-group",
	// "stset"
	// ],
	// "operator": "in"
	// }
	// ]
	// }
	// }
	// response_example: >-
	// {
	// "results": [
	// {
	// "resource_group_id": "rsc-grp-4c86e071e0f0",
	// "name": "azure-vmss-group",
	// "resources": [
	// {
	// "resource_type": "inventory.CloudService?provider=azure&cloud_service_group=Compute&cloud_service_type=VmScaleSet",
	// "filter": [
	// {
	// "k": "provider",
	// "v": "azure",
	// "o": "eq"
	// },
	// {
	// "v": "Compute",
	// "k": "cloud_service_group",
	// "o": "eq"
	// },
	// {
	// "k": "cloud_service_type",
	// "v": "VmScaleSet",
	// "o": "eq"
	// }
	// ]
	// }
	// ],
	// "options": {
	// "raw_filter": []
	// },
	// "tags": {},
	// "project_id": "project-9074eea97d7e",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2021-04-23T03:23:40.037Z"
	// },
	// {
	// "resource_group_id": "rsc-grp-aa3c4ca465b2",
	// "name": "stset",
	// "resources": [
	// {
	// "resource_type": "inventory.Server?provider=aws&cloud_service_group=EC2&cloud_service_type=Instance",
	// "filter": [
	// {
	// "k": "provider",
	// "v": "aws",
	// "o": "eq"
	// },
	// {
	// "v": "EC2",
	// "k": "cloud_service_group",
	// "o": "eq"
	// },
	// {
	// "k": "cloud_service_type",
	// "v": "Instance",
	// "o": "eq"
	// }
	// ],
	// "keyword": "test"
	// }
	// ],
	// "options": {
	// "raw_filter": [
	// [
	// "test"
	// ]
	// ]
	// },
	// "tags": {},
	// "project_id": "project-18655561c535",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2021-06-01T10:23:20.537Z"
	// }
	// ],
	// "total_count": 2
	// }
	List(ctx context.Context, in *ResourceGroupQuery, opts ...grpc.CallOption) (*ResourceGroupsInfo, error)
	Stat(ctx context.Context, in *ResourceGroupStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type resourceGroupClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceGroupClient(cc grpc.ClientConnInterface) ResourceGroupClient {
	return &resourceGroupClient{cc}
}

func (c *resourceGroupClient) Create(ctx context.Context, in *CreateResourceGroupRequest, opts ...grpc.CallOption) (*ResourceGroupInfo, error) {
	out := new(ResourceGroupInfo)
	err := c.cc.Invoke(ctx, "/spaceone.api.inventory.v1.ResourceGroup/create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceGroupClient) Update(ctx context.Context, in *UpdateResourceGroupRequest, opts ...grpc.CallOption) (*ResourceGroupInfo, error) {
	out := new(ResourceGroupInfo)
	err := c.cc.Invoke(ctx, "/spaceone.api.inventory.v1.ResourceGroup/update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceGroupClient) Delete(ctx context.Context, in *ResourceGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/spaceone.api.inventory.v1.ResourceGroup/delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceGroupClient) Get(ctx context.Context, in *GetResourceGroupRequest, opts ...grpc.CallOption) (*ResourceGroupInfo, error) {
	out := new(ResourceGroupInfo)
	err := c.cc.Invoke(ctx, "/spaceone.api.inventory.v1.ResourceGroup/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceGroupClient) List(ctx context.Context, in *ResourceGroupQuery, opts ...grpc.CallOption) (*ResourceGroupsInfo, error) {
	out := new(ResourceGroupsInfo)
	err := c.cc.Invoke(ctx, "/spaceone.api.inventory.v1.ResourceGroup/list", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceGroupClient) Stat(ctx context.Context, in *ResourceGroupStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, "/spaceone.api.inventory.v1.ResourceGroup/stat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceGroupServer is the server API for ResourceGroup service.
// All implementations must embed UnimplementedResourceGroupServer
// for forward compatibility
type ResourceGroupServer interface {
	// desc: Creates a new ResourceGroup. You can integrate `resource`s from different `provider`s by specifying the cloud resources to be grouped in the `resources` parameter.
	// request_example: >-
	// {
	// "name": "azure-group-1",
	// "resources": [
	// {
	// "resource_type": "inventory.Server?provider=azure&cloud_service_group=Compute&cloud_service_type=VirtualMachine",
	// "filter": [
	// {"o": "eq", "k": "provider", "v": "azure"},
	// {"o": "eq", "k": "cloud_service_group", "v": "Compute"},
	// {"o": "eq", "k": "cloud_service_type", "v": "VirtualMachine"}
	// ]
	// }
	// ],
	// "options": {
	// "raw_filter": []
	// },
	// "tags": {
	// "a": "b"
	// }
	// }
	// response_example: >-
	// {
	// "resource_group_id": "rsc-grp-7d46a1fc7738",
	// "name": "azure-group-1",
	// "resources": [
	// {
	// "resource_type": "inventory.Server?provider=azure&cloud_service_group=Compute&cloud_service_type=VirtualMachine",
	// "filter": [
	// {
	// "k": "provider",
	// "v": "azure",
	// "o": "eq"
	// },
	// {
	// "k": "cloud_service_group",
	// "v": "Compute",
	// "o": "eq"
	// },
	// {
	// "k": "cloud_service_type",
	// "v": "VirtualMachine",
	// "o": "eq"
	// }
	// ]
	// }
	// ],
	// "options": {
	// "raw_filter": []
	// },
	// "tags": {
	// "a": "b"
	// },
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-06-23T01:50:33.152Z"
	// }
	Create(context.Context, *CreateResourceGroupRequest) (*ResourceGroupInfo, error)
	// desc: Updates a specific ResourceGroup. You can make changes in ResourceGroup settings, including `name`, `tags`, and grouped resources in the ResourceGroup.
	// request_example: >-
	// {
	// "resource_group_id": "rsc-grp-7d46a1fc7738",
	// "name": "azure-grp-test2",
	// "resources": [
	// {
	// "resource_type": "inventory.Server?provider=azure&cloud_service_group=Compute&cloud_service_type=VirtualMachine",
	// "filter": [
	// {"k": "provider", "v": "azure", "o": "eq"},
	// {"o": "eq", "k": "cloud_service_group", "v": "Compute"},
	// {"v": "VirtualMachine", "k": "cloud_service_type", "o": "eq"}
	// ]
	// }
	// ],
	// "options": {},
	// "tags": {
	// "b": "c"
	// }
	// }
	// response_example: >-
	// {
	// "resource_group_id": "rsc-grp-7d46a1fc7738",
	// "name": "azure-grp-test2",
	// "resources": [
	// {
	// "resource_type": "inventory.Server?provider=azure&cloud_service_group=Compute&cloud_service_type=VirtualMachine",
	// "filter": [
	// {
	// "k": "provider",
	// "v": "azure",
	// "o": "eq"
	// },
	// {
	// "k": "cloud_service_group",
	// "v": "Compute",
	// "o": "eq"
	// },
	// {
	// "k": "cloud_service_type",
	// "v": "VirtualMachine",
	// "o": "eq"
	// }
	// ]
	// }
	// ],
	// "options": {
	// "raw_filter": []
	// },
	// "tags": {
	// "a": "b"
	// },
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2022-06-23T01:50:33.152Z"
	// }
	Update(context.Context, *UpdateResourceGroupRequest) (*ResourceGroupInfo, error)
	// desc: Deletes a specific ResourceGroup. You must specify the `resource_group_id` of the ResourceGroup to delete.
	// request_example: >-
	// {
	// "resource_group_id": "rsc-grp-aa3c4ca465b2"
	// }
	Delete(context.Context, *ResourceGroupRequest) (*empty.Empty, error)
	// desc: Gets a specific ResourceGroup. Prints detailed information about the ResourceGroup, including the information of the grouped cloud resources.
	// request_example: >-
	// {
	// "resource_group_id": "rsc-grp-aa3c4ca465b2"
	// }
	// response_example: >-
	// {
	// "resource_group_id": "rsc-grp-aa3c4ca465b2",
	// "name": "stset",
	// "resources": [
	// {
	// "resource_type": "inventory.Server?provider=aws&cloud_service_group=EC2&cloud_service_type=Instance",
	// "filter": [
	// {
	// "k": "provider",
	// "o": "eq",
	// "v": "aws"
	// },
	// {
	// "v": "EC2",
	// "k": "cloud_service_group",
	// "o": "eq"
	// },
	// {
	// "o": "eq",
	// "k": "cloud_service_type",
	// "v": "Instance"
	// }
	// ],
	// "keyword": "test"
	// }
	// ],
	// "options": {
	// "raw_filter": []
	// },
	// "tags": {},
	// "project_id": "project-18655561c535",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2021-06-01T10:23:20.537Z"
	// }
	Get(context.Context, *GetResourceGroupRequest) (*ResourceGroupInfo, error)
	// desc: Gets a list of all ResourceGroups. You can use a query to get a filtered list of ResourceGroups.
	// request_example: >-
	// {
	// "query": {
	// "filter": [
	// {
	// "key": "name",
	// "value": [
	// "azure-vmss-group",
	// "stset"
	// ],
	// "operator": "in"
	// }
	// ]
	// }
	// }
	// response_example: >-
	// {
	// "results": [
	// {
	// "resource_group_id": "rsc-grp-4c86e071e0f0",
	// "name": "azure-vmss-group",
	// "resources": [
	// {
	// "resource_type": "inventory.CloudService?provider=azure&cloud_service_group=Compute&cloud_service_type=VmScaleSet",
	// "filter": [
	// {
	// "k": "provider",
	// "v": "azure",
	// "o": "eq"
	// },
	// {
	// "v": "Compute",
	// "k": "cloud_service_group",
	// "o": "eq"
	// },
	// {
	// "k": "cloud_service_type",
	// "v": "VmScaleSet",
	// "o": "eq"
	// }
	// ]
	// }
	// ],
	// "options": {
	// "raw_filter": []
	// },
	// "tags": {},
	// "project_id": "project-9074eea97d7e",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2021-04-23T03:23:40.037Z"
	// },
	// {
	// "resource_group_id": "rsc-grp-aa3c4ca465b2",
	// "name": "stset",
	// "resources": [
	// {
	// "resource_type": "inventory.Server?provider=aws&cloud_service_group=EC2&cloud_service_type=Instance",
	// "filter": [
	// {
	// "k": "provider",
	// "v": "aws",
	// "o": "eq"
	// },
	// {
	// "v": "EC2",
	// "k": "cloud_service_group",
	// "o": "eq"
	// },
	// {
	// "k": "cloud_service_type",
	// "v": "Instance",
	// "o": "eq"
	// }
	// ],
	// "keyword": "test"
	// }
	// ],
	// "options": {
	// "raw_filter": [
	// [
	// "test"
	// ]
	// ]
	// },
	// "tags": {},
	// "project_id": "project-18655561c535",
	// "domain_id": "domain-58010aa2e451",
	// "created_at": "2021-06-01T10:23:20.537Z"
	// }
	// ],
	// "total_count": 2
	// }
	List(context.Context, *ResourceGroupQuery) (*ResourceGroupsInfo, error)
	Stat(context.Context, *ResourceGroupStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedResourceGroupServer()
}

// UnimplementedResourceGroupServer must be embedded to have forward compatible implementations.
type UnimplementedResourceGroupServer struct {
}

func (UnimplementedResourceGroupServer) Create(context.Context, *CreateResourceGroupRequest) (*ResourceGroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedResourceGroupServer) Update(context.Context, *UpdateResourceGroupRequest) (*ResourceGroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedResourceGroupServer) Delete(context.Context, *ResourceGroupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedResourceGroupServer) Get(context.Context, *GetResourceGroupRequest) (*ResourceGroupInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedResourceGroupServer) List(context.Context, *ResourceGroupQuery) (*ResourceGroupsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedResourceGroupServer) Stat(context.Context, *ResourceGroupStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedResourceGroupServer) mustEmbedUnimplementedResourceGroupServer() {}

// UnsafeResourceGroupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceGroupServer will
// result in compilation errors.
type UnsafeResourceGroupServer interface {
	mustEmbedUnimplementedResourceGroupServer()
}

func RegisterResourceGroupServer(s grpc.ServiceRegistrar, srv ResourceGroupServer) {
	s.RegisterService(&ResourceGroup_ServiceDesc, srv)
}

func _ResourceGroup_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResourceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceGroupServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spaceone.api.inventory.v1.ResourceGroup/create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceGroupServer).Create(ctx, req.(*CreateResourceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceGroup_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResourceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceGroupServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spaceone.api.inventory.v1.ResourceGroup/update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceGroupServer).Update(ctx, req.(*UpdateResourceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceGroup_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceGroupServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spaceone.api.inventory.v1.ResourceGroup/delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceGroupServer).Delete(ctx, req.(*ResourceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceGroup_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceGroupServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spaceone.api.inventory.v1.ResourceGroup/get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceGroupServer).Get(ctx, req.(*GetResourceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceGroup_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceGroupQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceGroupServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spaceone.api.inventory.v1.ResourceGroup/list",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceGroupServer).List(ctx, req.(*ResourceGroupQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceGroup_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceGroupStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceGroupServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spaceone.api.inventory.v1.ResourceGroup/stat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceGroupServer).Stat(ctx, req.(*ResourceGroupStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceGroup_ServiceDesc is the grpc.ServiceDesc for ResourceGroup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceGroup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.inventory.v1.ResourceGroup",
	HandlerType: (*ResourceGroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _ResourceGroup_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _ResourceGroup_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _ResourceGroup_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _ResourceGroup_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _ResourceGroup_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _ResourceGroup_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/inventory/v1/resource_group.proto",
}
