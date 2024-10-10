// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: v1/favorite.proto

package v1

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

// FavoriteServiceClient is the client API for FavoriteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavoriteServiceClient interface {
	AddFavorite(ctx context.Context, in *AddFavoriteRequest, opts ...grpc.CallOption) (*AddFavoriteResponse, error)
	RemoveFavorite(ctx context.Context, in *RemoveFavoriteRequest, opts ...grpc.CallOption) (*RemoveFavoriteResponse, error)
	ListFavorite(ctx context.Context, in *ListFavoriteRequest, opts ...grpc.CallOption) (*ListFavoriteResponse, error)
	CountFavorite(ctx context.Context, in *CountFavoriteRequest, opts ...grpc.CallOption) (*CountFavoriteResponse, error)
	IsFavorite(ctx context.Context, in *IsFavoriteRequest, opts ...grpc.CallOption) (*IsFavoriteResponse, error)
}

type favoriteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteServiceClient(cc grpc.ClientConnInterface) FavoriteServiceClient {
	return &favoriteServiceClient{cc}
}

func (c *favoriteServiceClient) AddFavorite(ctx context.Context, in *AddFavoriteRequest, opts ...grpc.CallOption) (*AddFavoriteResponse, error) {
	out := new(AddFavoriteResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.FavoriteService/AddFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteServiceClient) RemoveFavorite(ctx context.Context, in *RemoveFavoriteRequest, opts ...grpc.CallOption) (*RemoveFavoriteResponse, error) {
	out := new(RemoveFavoriteResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.FavoriteService/RemoveFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteServiceClient) ListFavorite(ctx context.Context, in *ListFavoriteRequest, opts ...grpc.CallOption) (*ListFavoriteResponse, error) {
	out := new(ListFavoriteResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.FavoriteService/ListFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteServiceClient) CountFavorite(ctx context.Context, in *CountFavoriteRequest, opts ...grpc.CallOption) (*CountFavoriteResponse, error) {
	out := new(CountFavoriteResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.FavoriteService/CountFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteServiceClient) IsFavorite(ctx context.Context, in *IsFavoriteRequest, opts ...grpc.CallOption) (*IsFavoriteResponse, error) {
	out := new(IsFavoriteResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.FavoriteService/IsFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteServiceServer is the server API for FavoriteService service.
// All implementations must embed UnimplementedFavoriteServiceServer
// for forward compatibility
type FavoriteServiceServer interface {
	AddFavorite(context.Context, *AddFavoriteRequest) (*AddFavoriteResponse, error)
	RemoveFavorite(context.Context, *RemoveFavoriteRequest) (*RemoveFavoriteResponse, error)
	ListFavorite(context.Context, *ListFavoriteRequest) (*ListFavoriteResponse, error)
	CountFavorite(context.Context, *CountFavoriteRequest) (*CountFavoriteResponse, error)
	IsFavorite(context.Context, *IsFavoriteRequest) (*IsFavoriteResponse, error)
	mustEmbedUnimplementedFavoriteServiceServer()
}

// UnimplementedFavoriteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFavoriteServiceServer struct {
}

func (UnimplementedFavoriteServiceServer) AddFavorite(context.Context, *AddFavoriteRequest) (*AddFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFavorite not implemented")
}
func (UnimplementedFavoriteServiceServer) RemoveFavorite(context.Context, *RemoveFavoriteRequest) (*RemoveFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFavorite not implemented")
}
func (UnimplementedFavoriteServiceServer) ListFavorite(context.Context, *ListFavoriteRequest) (*ListFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFavorite not implemented")
}
func (UnimplementedFavoriteServiceServer) CountFavorite(context.Context, *CountFavoriteRequest) (*CountFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountFavorite not implemented")
}
func (UnimplementedFavoriteServiceServer) IsFavorite(context.Context, *IsFavoriteRequest) (*IsFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFavorite not implemented")
}
func (UnimplementedFavoriteServiceServer) mustEmbedUnimplementedFavoriteServiceServer() {}

// UnsafeFavoriteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavoriteServiceServer will
// result in compilation errors.
type UnsafeFavoriteServiceServer interface {
	mustEmbedUnimplementedFavoriteServiceServer()
}

func RegisterFavoriteServiceServer(s grpc.ServiceRegistrar, srv FavoriteServiceServer) {
	s.RegisterService(&FavoriteService_ServiceDesc, srv)
}

func _FavoriteService_AddFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).AddFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.FavoriteService/AddFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).AddFavorite(ctx, req.(*AddFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteService_RemoveFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).RemoveFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.FavoriteService/RemoveFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).RemoveFavorite(ctx, req.(*RemoveFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteService_ListFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).ListFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.FavoriteService/ListFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).ListFavorite(ctx, req.(*ListFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteService_CountFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).CountFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.FavoriteService/CountFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).CountFavorite(ctx, req.(*CountFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteService_IsFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).IsFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.FavoriteService/IsFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).IsFavorite(ctx, req.(*IsFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FavoriteService_ServiceDesc is the grpc.ServiceDesc for FavoriteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FavoriteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shortVideoCoreService.api.v1.FavoriteService",
	HandlerType: (*FavoriteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFavorite",
			Handler:    _FavoriteService_AddFavorite_Handler,
		},
		{
			MethodName: "RemoveFavorite",
			Handler:    _FavoriteService_RemoveFavorite_Handler,
		},
		{
			MethodName: "ListFavorite",
			Handler:    _FavoriteService_ListFavorite_Handler,
		},
		{
			MethodName: "CountFavorite",
			Handler:    _FavoriteService_CountFavorite_Handler,
		},
		{
			MethodName: "IsFavorite",
			Handler:    _FavoriteService_IsFavorite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/favorite.proto",
}