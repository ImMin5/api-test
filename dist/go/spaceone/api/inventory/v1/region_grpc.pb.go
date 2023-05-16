// A Region is a resource storing regional information from each cloud service provider. Regional data stored by the resource includes the latitude and longitude of the region.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: spaceone/api/inventory/v1/region.proto

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
	Region_Create_FullMethodName = "/spaceone.api.inventory.v1.Region/create"
	Region_Update_FullMethodName = "/spaceone.api.inventory.v1.Region/update"
	Region_Delete_FullMethodName = "/spaceone.api.inventory.v1.Region/delete"
	Region_Get_FullMethodName    = "/spaceone.api.inventory.v1.Region/get"
	Region_List_FullMethodName   = "/spaceone.api.inventory.v1.Region/list"
	Region_Stat_FullMethodName   = "/spaceone.api.inventory.v1.Region/stat"
)

// RegionClient is the client API for Region service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegionClient interface {
	// Creates a new Region. As the parameter `region_key`, which is automatically created when a Region is created, is in a form of `{provider}.{region_code}`, it is unique regardless of providers. You can also specify the latitude, longitude, and continent information in `tags`.
	Create(ctx context.Context, in *CreateRegionRequest, opts ...grpc.CallOption) (*RegionInfo, error)
	// Updates a specific Region. You can make changes in Region settings, including `name` and `tags`. The `tags` contain the continent, latitude, and longitude.
	Update(ctx context.Context, in *UpdateRegionRequest, opts ...grpc.CallOption) (*RegionInfo, error)
	// Deletes a specific Region. You must specify the `region_id` of the Region to delete.
	Delete(ctx context.Context, in *RegionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets a specific Region. Prints detailed information about the Region, including `name`, `region_code`, and a location coordinate.
	Get(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*RegionInfo, error)
	// Gets a list of all Regions. You can use a query to get a filtered list of Regions.
	List(ctx context.Context, in *RegionQuery, opts ...grpc.CallOption) (*RegionsInfo, error)
	Stat(ctx context.Context, in *RegionStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type regionClient struct {
	cc grpc.ClientConnInterface
}

func NewRegionClient(cc grpc.ClientConnInterface) RegionClient {
	return &regionClient{cc}
}

func (c *regionClient) Create(ctx context.Context, in *CreateRegionRequest, opts ...grpc.CallOption) (*RegionInfo, error) {
	out := new(RegionInfo)
	err := c.cc.Invoke(ctx, Region_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionClient) Update(ctx context.Context, in *UpdateRegionRequest, opts ...grpc.CallOption) (*RegionInfo, error) {
	out := new(RegionInfo)
	err := c.cc.Invoke(ctx, Region_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionClient) Delete(ctx context.Context, in *RegionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Region_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionClient) Get(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*RegionInfo, error) {
	out := new(RegionInfo)
	err := c.cc.Invoke(ctx, Region_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionClient) List(ctx context.Context, in *RegionQuery, opts ...grpc.CallOption) (*RegionsInfo, error) {
	out := new(RegionsInfo)
	err := c.cc.Invoke(ctx, Region_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionClient) Stat(ctx context.Context, in *RegionStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Region_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegionServer is the server API for Region service.
// All implementations must embed UnimplementedRegionServer
// for forward compatibility
type RegionServer interface {
	// Creates a new Region. As the parameter `region_key`, which is automatically created when a Region is created, is in a form of `{provider}.{region_code}`, it is unique regardless of providers. You can also specify the latitude, longitude, and continent information in `tags`.
	Create(context.Context, *CreateRegionRequest) (*RegionInfo, error)
	// Updates a specific Region. You can make changes in Region settings, including `name` and `tags`. The `tags` contain the continent, latitude, and longitude.
	Update(context.Context, *UpdateRegionRequest) (*RegionInfo, error)
	// Deletes a specific Region. You must specify the `region_id` of the Region to delete.
	Delete(context.Context, *RegionRequest) (*empty.Empty, error)
	// Gets a specific Region. Prints detailed information about the Region, including `name`, `region_code`, and a location coordinate.
	Get(context.Context, *GetRegionRequest) (*RegionInfo, error)
	// Gets a list of all Regions. You can use a query to get a filtered list of Regions.
	List(context.Context, *RegionQuery) (*RegionsInfo, error)
	Stat(context.Context, *RegionStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedRegionServer()
}

// UnimplementedRegionServer must be embedded to have forward compatible implementations.
type UnimplementedRegionServer struct {
}

func (UnimplementedRegionServer) Create(context.Context, *CreateRegionRequest) (*RegionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRegionServer) Update(context.Context, *UpdateRegionRequest) (*RegionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRegionServer) Delete(context.Context, *RegionRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRegionServer) Get(context.Context, *GetRegionRequest) (*RegionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRegionServer) List(context.Context, *RegionQuery) (*RegionsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRegionServer) Stat(context.Context, *RegionStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedRegionServer) mustEmbedUnimplementedRegionServer() {}

// UnsafeRegionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegionServer will
// result in compilation errors.
type UnsafeRegionServer interface {
	mustEmbedUnimplementedRegionServer()
}

func RegisterRegionServer(s grpc.ServiceRegistrar, srv RegionServer) {
	s.RegisterService(&Region_ServiceDesc, srv)
}

func _Region_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Region_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServer).Create(ctx, req.(*CreateRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Region_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Region_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServer).Update(ctx, req.(*UpdateRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Region_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Region_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServer).Delete(ctx, req.(*RegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Region_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Region_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServer).Get(ctx, req.(*GetRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Region_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegionQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Region_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServer).List(ctx, req.(*RegionQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Region_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegionStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Region_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServer).Stat(ctx, req.(*RegionStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Region_ServiceDesc is the grpc.ServiceDesc for Region service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Region_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.inventory.v1.Region",
	HandlerType: (*RegionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Region_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Region_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Region_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Region_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Region_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Region_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/inventory/v1/region.proto",
}
