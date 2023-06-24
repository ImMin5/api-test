// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: spaceone/api/inventory/v1/change_history.proto

package v1

import (
	context "context"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that ˚`this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ChangeHistory_List_FullMethodName = "/spaceone.api.inventory.v1.ChangeHistory/list"
	ChangeHistory_Stat_FullMethodName = "/spaceone.api.inventory.v1.ChangeHistory/stat"
)

// ChangeHistoryClient is the client API for ChangeHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChangeHistoryClient interface {
	List(ctx context.Context, in *ChangeHistoryQuery, opts ...grpc.CallOption) (*ChangeHistoryInfo, error)
	Stat(ctx context.Context, in *ChangeHistoryStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type changeHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewChangeHistoryClient(cc grpc.ClientConnInterface) ChangeHistoryClient {
	return &changeHistoryClient{cc}
}

func (c *changeHistoryClient) List(ctx context.Context, in *ChangeHistoryQuery, opts ...grpc.CallOption) (*ChangeHistoryInfo, error) {
	out := new(ChangeHistoryInfo)
	err := c.cc.Invoke(ctx, ChangeHistory_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *changeHistoryClient) Stat(ctx context.Context, in *ChangeHistoryStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, ChangeHistory_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChangeHistoryServer is the server API for ChangeHistory service.
// All implementations must embed UnimplementedChangeHistoryServer
// for forward compatibility
type ChangeHistoryServer interface {
	List(context.Context, *ChangeHistoryQuery) (*ChangeHistoryInfo, error)
	Stat(context.Context, *ChangeHistoryStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedChangeHistoryServer()
}

// UnimplementedChangeHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedChangeHistoryServer struct {
}

func (UnimplementedChangeHistoryServer) List(context.Context, *ChangeHistoryQuery) (*ChangeHistoryInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedChangeHistoryServer) Stat(context.Context, *ChangeHistoryStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedChangeHistoryServer) mustEmbedUnimplementedChangeHistoryServer() {}

// UnsafeChangeHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChangeHistoryServer will
// result in compilation errors.
type UnsafeChangeHistoryServer interface {
	mustEmbedUnimplementedChangeHistoryServer()
}

func RegisterChangeHistoryServer(s grpc.ServiceRegistrar, srv ChangeHistoryServer) {
	s.RegisterService(&ChangeHistory_ServiceDesc, srv)
}

func _ChangeHistory_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeHistoryQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChangeHistoryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChangeHistory_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChangeHistoryServer).List(ctx, req.(*ChangeHistoryQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChangeHistory_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeHistoryStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChangeHistoryServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChangeHistory_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChangeHistoryServer).Stat(ctx, req.(*ChangeHistoryStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// ChangeHistory_ServiceDesc is the grpc.ServiceDesc for ChangeHistory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChangeHistory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.inventory.v1.ChangeHistory",
	HandlerType: (*ChangeHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list",
			Handler:    _ChangeHistory_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _ChangeHistory_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/inventory/v1/change_history.proto",
}
