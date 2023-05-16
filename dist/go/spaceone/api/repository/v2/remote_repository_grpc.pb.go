// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: spaceone/api/repository/v2/remote_repository.proto

package v2

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
	RemoteRepository_Get_FullMethodName  = "/spaceone.api.repository.v2.RemoteRepository/get"
	RemoteRepository_List_FullMethodName = "/spaceone.api.repository.v2.RemoteRepository/list"
)

// RemoteRepositoryClient is the client API for RemoteRepository service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteRepositoryClient interface {
	Get(ctx context.Context, in *GetRemoteRepository, opts ...grpc.CallOption) (*RemoteRepositoryInfo, error)
	List(ctx context.Context, in *RemoteRepositoryQuery, opts ...grpc.CallOption) (*RemoteRepositoriesInfo, error)
}

type remoteRepositoryClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteRepositoryClient(cc grpc.ClientConnInterface) RemoteRepositoryClient {
	return &remoteRepositoryClient{cc}
}

func (c *remoteRepositoryClient) Get(ctx context.Context, in *GetRemoteRepository, opts ...grpc.CallOption) (*RemoteRepositoryInfo, error) {
	out := new(RemoteRepositoryInfo)
	err := c.cc.Invoke(ctx, RemoteRepository_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteRepositoryClient) List(ctx context.Context, in *RemoteRepositoryQuery, opts ...grpc.CallOption) (*RemoteRepositoriesInfo, error) {
	out := new(RemoteRepositoriesInfo)
	err := c.cc.Invoke(ctx, RemoteRepository_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteRepositoryServer is the server API for RemoteRepository service.
// All implementations must embed UnimplementedRemoteRepositoryServer
// for forward compatibility
type RemoteRepositoryServer interface {
	Get(context.Context, *GetRemoteRepository) (*RemoteRepositoryInfo, error)
	List(context.Context, *RemoteRepositoryQuery) (*RemoteRepositoriesInfo, error)
	mustEmbedUnimplementedRemoteRepositoryServer()
}

// UnimplementedRemoteRepositoryServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteRepositoryServer struct {
}

func (UnimplementedRemoteRepositoryServer) Get(context.Context, *GetRemoteRepository) (*RemoteRepositoryInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRemoteRepositoryServer) List(context.Context, *RemoteRepositoryQuery) (*RemoteRepositoriesInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRemoteRepositoryServer) mustEmbedUnimplementedRemoteRepositoryServer() {}

// UnsafeRemoteRepositoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteRepositoryServer will
// result in compilation errors.
type UnsafeRemoteRepositoryServer interface {
	mustEmbedUnimplementedRemoteRepositoryServer()
}

func RegisterRemoteRepositoryServer(s grpc.ServiceRegistrar, srv RemoteRepositoryServer) {
	s.RegisterService(&RemoteRepository_ServiceDesc, srv)
}

func _RemoteRepository_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRemoteRepository)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteRepositoryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteRepository_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteRepositoryServer).Get(ctx, req.(*GetRemoteRepository))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteRepository_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteRepositoryQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteRepositoryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteRepository_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteRepositoryServer).List(ctx, req.(*RemoteRepositoryQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteRepository_ServiceDesc is the grpc.ServiceDesc for RemoteRepository service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteRepository_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.repository.v2.RemoteRepository",
	HandlerType: (*RemoteRepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _RemoteRepository_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _RemoteRepository_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/repository/v2/remote_repository.proto",
}
