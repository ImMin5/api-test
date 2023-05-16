// A Budget is a planned amount of cost expenditure for reduction and prediction of infrastructure costs.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: spaceone/api/cost_analysis/v1/budget.proto

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
	Budget_Create_FullMethodName          = "/spaceone.api.cost_analysis.v1.Budget/create"
	Budget_Update_FullMethodName          = "/spaceone.api.cost_analysis.v1.Budget/update"
	Budget_SetNotification_FullMethodName = "/spaceone.api.cost_analysis.v1.Budget/set_notification"
	Budget_Delete_FullMethodName          = "/spaceone.api.cost_analysis.v1.Budget/delete"
	Budget_Get_FullMethodName             = "/spaceone.api.cost_analysis.v1.Budget/get"
	Budget_List_FullMethodName            = "/spaceone.api.cost_analysis.v1.Budget/list"
	Budget_Stat_FullMethodName            = "/spaceone.api.cost_analysis.v1.Budget/stat"
)

// BudgetClient is the client API for Budget service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BudgetClient interface {
	// Creates a new Budget. When creating a Budget, it should be set for a specific ProjectGroup or Project. The budgeted amount and date of the `planned_limits` should be specified on a monthly or yearly basis.
	Create(ctx context.Context, in *CreateBudgetRequest, opts ...grpc.CallOption) (*BudgetInfo, error)
	// Updates a specific Budget. You can make changes in the budgeted amount of the time period specified while creating the resource.
	Update(ctx context.Context, in *UpdateBudgetRequest, opts ...grpc.CallOption) (*BudgetInfo, error)
	// Sets a notification on a specific Budget. Sets a threshold on the budget, and if the cost exceeds the threshold, a notification is raised.
	SetNotification(ctx context.Context, in *SetBudgetNotificationRequest, opts ...grpc.CallOption) (*BudgetInfo, error)
	// Deletes a specific Budget. You must specify the `budget_id` of the Budget to delete.
	Delete(ctx context.Context, in *BudgetRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets a specific Budget. Prints detailed information about the Budget, including `planned_limits` of the project group or project for the pre-defined period.
	Get(ctx context.Context, in *GetBudgetRequest, opts ...grpc.CallOption) (*BudgetInfo, error)
	// Gets a list of all Budgets. You can use a query to get a filtered list of Budgets.
	List(ctx context.Context, in *BudgetQuery, opts ...grpc.CallOption) (*BudgetsInfo, error)
	Stat(ctx context.Context, in *BudgetStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type budgetClient struct {
	cc grpc.ClientConnInterface
}

func NewBudgetClient(cc grpc.ClientConnInterface) BudgetClient {
	return &budgetClient{cc}
}

func (c *budgetClient) Create(ctx context.Context, in *CreateBudgetRequest, opts ...grpc.CallOption) (*BudgetInfo, error) {
	out := new(BudgetInfo)
	err := c.cc.Invoke(ctx, Budget_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetClient) Update(ctx context.Context, in *UpdateBudgetRequest, opts ...grpc.CallOption) (*BudgetInfo, error) {
	out := new(BudgetInfo)
	err := c.cc.Invoke(ctx, Budget_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetClient) SetNotification(ctx context.Context, in *SetBudgetNotificationRequest, opts ...grpc.CallOption) (*BudgetInfo, error) {
	out := new(BudgetInfo)
	err := c.cc.Invoke(ctx, Budget_SetNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetClient) Delete(ctx context.Context, in *BudgetRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Budget_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetClient) Get(ctx context.Context, in *GetBudgetRequest, opts ...grpc.CallOption) (*BudgetInfo, error) {
	out := new(BudgetInfo)
	err := c.cc.Invoke(ctx, Budget_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetClient) List(ctx context.Context, in *BudgetQuery, opts ...grpc.CallOption) (*BudgetsInfo, error) {
	out := new(BudgetsInfo)
	err := c.cc.Invoke(ctx, Budget_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetClient) Stat(ctx context.Context, in *BudgetStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Budget_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BudgetServer is the server API for Budget service.
// All implementations must embed UnimplementedBudgetServer
// for forward compatibility
type BudgetServer interface {
	// Creates a new Budget. When creating a Budget, it should be set for a specific ProjectGroup or Project. The budgeted amount and date of the `planned_limits` should be specified on a monthly or yearly basis.
	Create(context.Context, *CreateBudgetRequest) (*BudgetInfo, error)
	// Updates a specific Budget. You can make changes in the budgeted amount of the time period specified while creating the resource.
	Update(context.Context, *UpdateBudgetRequest) (*BudgetInfo, error)
	// Sets a notification on a specific Budget. Sets a threshold on the budget, and if the cost exceeds the threshold, a notification is raised.
	SetNotification(context.Context, *SetBudgetNotificationRequest) (*BudgetInfo, error)
	// Deletes a specific Budget. You must specify the `budget_id` of the Budget to delete.
	Delete(context.Context, *BudgetRequest) (*empty.Empty, error)
	// Gets a specific Budget. Prints detailed information about the Budget, including `planned_limits` of the project group or project for the pre-defined period.
	Get(context.Context, *GetBudgetRequest) (*BudgetInfo, error)
	// Gets a list of all Budgets. You can use a query to get a filtered list of Budgets.
	List(context.Context, *BudgetQuery) (*BudgetsInfo, error)
	Stat(context.Context, *BudgetStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedBudgetServer()
}

// UnimplementedBudgetServer must be embedded to have forward compatible implementations.
type UnimplementedBudgetServer struct {
}

func (UnimplementedBudgetServer) Create(context.Context, *CreateBudgetRequest) (*BudgetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBudgetServer) Update(context.Context, *UpdateBudgetRequest) (*BudgetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBudgetServer) SetNotification(context.Context, *SetBudgetNotificationRequest) (*BudgetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNotification not implemented")
}
func (UnimplementedBudgetServer) Delete(context.Context, *BudgetRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBudgetServer) Get(context.Context, *GetBudgetRequest) (*BudgetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBudgetServer) List(context.Context, *BudgetQuery) (*BudgetsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBudgetServer) Stat(context.Context, *BudgetStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedBudgetServer) mustEmbedUnimplementedBudgetServer() {}

// UnsafeBudgetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BudgetServer will
// result in compilation errors.
type UnsafeBudgetServer interface {
	mustEmbedUnimplementedBudgetServer()
}

func RegisterBudgetServer(s grpc.ServiceRegistrar, srv BudgetServer) {
	s.RegisterService(&Budget_ServiceDesc, srv)
}

func _Budget_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Budget_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServer).Create(ctx, req.(*CreateBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Budget_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Budget_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServer).Update(ctx, req.(*UpdateBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Budget_SetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBudgetNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServer).SetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Budget_SetNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServer).SetNotification(ctx, req.(*SetBudgetNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Budget_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Budget_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServer).Delete(ctx, req.(*BudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Budget_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Budget_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServer).Get(ctx, req.(*GetBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Budget_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BudgetQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Budget_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServer).List(ctx, req.(*BudgetQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Budget_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BudgetStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Budget_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServer).Stat(ctx, req.(*BudgetStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Budget_ServiceDesc is the grpc.ServiceDesc for Budget service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Budget_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.cost_analysis.v1.Budget",
	HandlerType: (*BudgetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Budget_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Budget_Update_Handler,
		},
		{
			MethodName: "set_notification",
			Handler:    _Budget_SetNotification_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Budget_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Budget_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Budget_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Budget_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/cost_analysis/v1/budget.proto",
}
