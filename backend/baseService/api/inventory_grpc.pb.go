// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: inventory.proto

package api

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

// InventoryServiceClient is the client API for InventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryServiceClient interface {
	CreateGoodInventory(ctx context.Context, in *CreateGoodRequest, opts ...grpc.CallOption) (*CreateGoodResponse, error)
	QueryGoodsInventory(ctx context.Context, in *QueryGoodsRequest, opts ...grpc.CallOption) (*QueryGoodsResponse, error)
	UpdateGoodsInventory(ctx context.Context, in *UpdateGoodsRequest, opts ...grpc.CallOption) (*UpdateGoodsResponse, error)
	Operate(ctx context.Context, in *OperateRequest, opts ...grpc.CallOption) (*OperateResponse, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) CreateGoodInventory(ctx context.Context, in *CreateGoodRequest, opts ...grpc.CallOption) (*CreateGoodResponse, error) {
	out := new(CreateGoodResponse)
	err := c.cc.Invoke(ctx, "/api.InventoryService/CreateGoodInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) QueryGoodsInventory(ctx context.Context, in *QueryGoodsRequest, opts ...grpc.CallOption) (*QueryGoodsResponse, error) {
	out := new(QueryGoodsResponse)
	err := c.cc.Invoke(ctx, "/api.InventoryService/QueryGoodsInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) UpdateGoodsInventory(ctx context.Context, in *UpdateGoodsRequest, opts ...grpc.CallOption) (*UpdateGoodsResponse, error) {
	out := new(UpdateGoodsResponse)
	err := c.cc.Invoke(ctx, "/api.InventoryService/UpdateGoodsInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) Operate(ctx context.Context, in *OperateRequest, opts ...grpc.CallOption) (*OperateResponse, error) {
	out := new(OperateResponse)
	err := c.cc.Invoke(ctx, "/api.InventoryService/Operate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServiceServer is the server API for InventoryService service.
// All implementations should embed UnimplementedInventoryServiceServer
// for forward compatibility
type InventoryServiceServer interface {
	CreateGoodInventory(context.Context, *CreateGoodRequest) (*CreateGoodResponse, error)
	QueryGoodsInventory(context.Context, *QueryGoodsRequest) (*QueryGoodsResponse, error)
	UpdateGoodsInventory(context.Context, *UpdateGoodsRequest) (*UpdateGoodsResponse, error)
	Operate(context.Context, *OperateRequest) (*OperateResponse, error)
}

// UnimplementedInventoryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedInventoryServiceServer struct {
}

func (UnimplementedInventoryServiceServer) CreateGoodInventory(context.Context, *CreateGoodRequest) (*CreateGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoodInventory not implemented")
}
func (UnimplementedInventoryServiceServer) QueryGoodsInventory(context.Context, *QueryGoodsRequest) (*QueryGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryGoodsInventory not implemented")
}
func (UnimplementedInventoryServiceServer) UpdateGoodsInventory(context.Context, *UpdateGoodsRequest) (*UpdateGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoodsInventory not implemented")
}
func (UnimplementedInventoryServiceServer) Operate(context.Context, *OperateRequest) (*OperateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Operate not implemented")
}

// UnsafeInventoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServiceServer will
// result in compilation errors.
type UnsafeInventoryServiceServer interface {
	mustEmbedUnimplementedInventoryServiceServer()
}

func RegisterInventoryServiceServer(s grpc.ServiceRegistrar, srv InventoryServiceServer) {
	s.RegisterService(&InventoryService_ServiceDesc, srv)
}

func _InventoryService_CreateGoodInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).CreateGoodInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.InventoryService/CreateGoodInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).CreateGoodInventory(ctx, req.(*CreateGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_QueryGoodsInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).QueryGoodsInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.InventoryService/QueryGoodsInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).QueryGoodsInventory(ctx, req.(*QueryGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_UpdateGoodsInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).UpdateGoodsInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.InventoryService/UpdateGoodsInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).UpdateGoodsInventory(ctx, req.(*UpdateGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_Operate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).Operate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.InventoryService/Operate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).Operate(ctx, req.(*OperateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryService_ServiceDesc is the grpc.ServiceDesc for InventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGoodInventory",
			Handler:    _InventoryService_CreateGoodInventory_Handler,
		},
		{
			MethodName: "QueryGoodsInventory",
			Handler:    _InventoryService_QueryGoodsInventory_Handler,
		},
		{
			MethodName: "UpdateGoodsInventory",
			Handler:    _InventoryService_UpdateGoodsInventory_Handler,
		},
		{
			MethodName: "Operate",
			Handler:    _InventoryService_Operate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory.proto",
}