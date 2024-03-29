// A Quota is a limit on protocol usage for a day or a month. You can manage the use of the protocol with a Quota.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: spaceone/api/notification/v1/quota.proto

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
	Quota_Create_FullMethodName = "/spaceone.api.notification.v1.Quota/create"
	Quota_Update_FullMethodName = "/spaceone.api.notification.v1.Quota/update"
	Quota_Delete_FullMethodName = "/spaceone.api.notification.v1.Quota/delete"
	Quota_Get_FullMethodName    = "/spaceone.api.notification.v1.Quota/get"
	Quota_List_FullMethodName   = "/spaceone.api.notification.v1.Quota/list"
	Quota_Stat_FullMethodName   = "/spaceone.api.notification.v1.Quota/stat"
)

// QuotaClient is the client API for Quota service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuotaClient interface {
	// Creates a new Quota limiting the use of a selected Protocol for a day or a month. If the parameter `limit` has no value, it will be deemed unlimited. If a Protocol has not set a Quota, the default Quota set in the Config will be applied.
	Create(ctx context.Context, in *CreateQuotaRequest, opts ...grpc.CallOption) (*QuotaInfo, error)
	// Updates a specific Quota. You can make changes in Quota `limit`, managing the use of the Protocol.
	Update(ctx context.Context, in *UpdateQuotaRequest, opts ...grpc.CallOption) (*QuotaInfo, error)
	// Deletes a specific Quota. The default Quota set in the Config will be applied to the Protocol you deleted the Quota of.
	Delete(ctx context.Context, in *QuotaRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets a specific Quota. Prints detailed information about the Quota, including the `limit` and the Protocol limited by the Quota.
	Get(ctx context.Context, in *QuotaRequest, opts ...grpc.CallOption) (*QuotaInfo, error)
	// Gets a list of all Quotas. You can use a query to get a filtered list of Quotas.
	List(ctx context.Context, in *QuotaQuery, opts ...grpc.CallOption) (*QuotasInfo, error)
	Stat(ctx context.Context, in *QuotaStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type quotaClient struct {
	cc grpc.ClientConnInterface
}

func NewQuotaClient(cc grpc.ClientConnInterface) QuotaClient {
	return &quotaClient{cc}
}

func (c *quotaClient) Create(ctx context.Context, in *CreateQuotaRequest, opts ...grpc.CallOption) (*QuotaInfo, error) {
	out := new(QuotaInfo)
	err := c.cc.Invoke(ctx, Quota_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaClient) Update(ctx context.Context, in *UpdateQuotaRequest, opts ...grpc.CallOption) (*QuotaInfo, error) {
	out := new(QuotaInfo)
	err := c.cc.Invoke(ctx, Quota_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaClient) Delete(ctx context.Context, in *QuotaRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Quota_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaClient) Get(ctx context.Context, in *QuotaRequest, opts ...grpc.CallOption) (*QuotaInfo, error) {
	out := new(QuotaInfo)
	err := c.cc.Invoke(ctx, Quota_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaClient) List(ctx context.Context, in *QuotaQuery, opts ...grpc.CallOption) (*QuotasInfo, error) {
	out := new(QuotasInfo)
	err := c.cc.Invoke(ctx, Quota_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quotaClient) Stat(ctx context.Context, in *QuotaStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Quota_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuotaServer is the server API for Quota service.
// All implementations must embed UnimplementedQuotaServer
// for forward compatibility
type QuotaServer interface {
	// Creates a new Quota limiting the use of a selected Protocol for a day or a month. If the parameter `limit` has no value, it will be deemed unlimited. If a Protocol has not set a Quota, the default Quota set in the Config will be applied.
	Create(context.Context, *CreateQuotaRequest) (*QuotaInfo, error)
	// Updates a specific Quota. You can make changes in Quota `limit`, managing the use of the Protocol.
	Update(context.Context, *UpdateQuotaRequest) (*QuotaInfo, error)
	// Deletes a specific Quota. The default Quota set in the Config will be applied to the Protocol you deleted the Quota of.
	Delete(context.Context, *QuotaRequest) (*empty.Empty, error)
	// Gets a specific Quota. Prints detailed information about the Quota, including the `limit` and the Protocol limited by the Quota.
	Get(context.Context, *QuotaRequest) (*QuotaInfo, error)
	// Gets a list of all Quotas. You can use a query to get a filtered list of Quotas.
	List(context.Context, *QuotaQuery) (*QuotasInfo, error)
	Stat(context.Context, *QuotaStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedQuotaServer()
}

// UnimplementedQuotaServer must be embedded to have forward compatible implementations.
type UnimplementedQuotaServer struct {
}

func (UnimplementedQuotaServer) Create(context.Context, *CreateQuotaRequest) (*QuotaInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedQuotaServer) Update(context.Context, *UpdateQuotaRequest) (*QuotaInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedQuotaServer) Delete(context.Context, *QuotaRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedQuotaServer) Get(context.Context, *QuotaRequest) (*QuotaInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedQuotaServer) List(context.Context, *QuotaQuery) (*QuotasInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedQuotaServer) Stat(context.Context, *QuotaStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedQuotaServer) mustEmbedUnimplementedQuotaServer() {}

// UnsafeQuotaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuotaServer will
// result in compilation errors.
type UnsafeQuotaServer interface {
	mustEmbedUnimplementedQuotaServer()
}

func RegisterQuotaServer(s grpc.ServiceRegistrar, srv QuotaServer) {
	s.RegisterService(&Quota_ServiceDesc, srv)
}

func _Quota_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Quota_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServer).Create(ctx, req.(*CreateQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quota_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Quota_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServer).Update(ctx, req.(*UpdateQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quota_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Quota_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServer).Delete(ctx, req.(*QuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quota_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Quota_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServer).Get(ctx, req.(*QuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quota_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotaQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Quota_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServer).List(ctx, req.(*QuotaQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quota_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuotaStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Quota_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServer).Stat(ctx, req.(*QuotaStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Quota_ServiceDesc is the grpc.ServiceDesc for Quota service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Quota_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.notification.v1.Quota",
	HandlerType: (*QuotaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Quota_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Quota_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Quota_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Quota_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Quota_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Quota_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/notification/v1/quota.proto",
}
