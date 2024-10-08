// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v3.12.4
// source: svapi/follow.proto

package svapi

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationFollowServiceAddFollow = "/svapi.FollowService/AddFollow"
const OperationFollowServiceListFollowing = "/svapi.FollowService/ListFollowing"
const OperationFollowServiceRemoveFollow = "/svapi.FollowService/RemoveFollow"

type FollowServiceHTTPServer interface {
	AddFollow(context.Context, *AddFollowRequest) (*AddFollowResponse, error)
	ListFollowing(context.Context, *ListFollowingRequest) (*ListFollowingResponse, error)
	RemoveFollow(context.Context, *RemoveFollowRequest) (*RemoveFollowResponse, error)
}

func RegisterFollowServiceHTTPServer(s *http.Server, srv FollowServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/follow", _FollowService_AddFollow0_HTTP_Handler(srv))
	r.DELETE("/follow", _FollowService_RemoveFollow0_HTTP_Handler(srv))
	r.GET("/follow", _FollowService_ListFollowing0_HTTP_Handler(srv))
}

func _FollowService_AddFollow0_HTTP_Handler(srv FollowServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddFollowRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFollowServiceAddFollow)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddFollow(ctx, req.(*AddFollowRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return ctx.Result(200, out)
	}
}

func _FollowService_RemoveFollow0_HTTP_Handler(srv FollowServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RemoveFollowRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFollowServiceRemoveFollow)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RemoveFollow(ctx, req.(*RemoveFollowRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return ctx.Result(200, out)
	}
}

func _FollowService_ListFollowing0_HTTP_Handler(srv FollowServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListFollowingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFollowServiceListFollowing)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListFollowing(ctx, req.(*ListFollowingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return ctx.Result(200, out)
	}
}

type FollowServiceHTTPClient interface {
	AddFollow(ctx context.Context, req *AddFollowRequest, opts ...http.CallOption) (rsp *AddFollowResponse, err error)
	ListFollowing(ctx context.Context, req *ListFollowingRequest, opts ...http.CallOption) (rsp *ListFollowingResponse, err error)
	RemoveFollow(ctx context.Context, req *RemoveFollowRequest, opts ...http.CallOption) (rsp *RemoveFollowResponse, err error)
}

type FollowServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewFollowServiceHTTPClient(client *http.Client) FollowServiceHTTPClient {
	return &FollowServiceHTTPClientImpl{client}
}

func (c *FollowServiceHTTPClientImpl) AddFollow(ctx context.Context, in *AddFollowRequest, opts ...http.CallOption) (*AddFollowResponse, error) {
	var out AddFollowResponse
	pattern := "/follow"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFollowServiceAddFollow))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FollowServiceHTTPClientImpl) ListFollowing(ctx context.Context, in *ListFollowingRequest, opts ...http.CallOption) (*ListFollowingResponse, error) {
	var out ListFollowingResponse
	pattern := "/follow"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationFollowServiceListFollowing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FollowServiceHTTPClientImpl) RemoveFollow(ctx context.Context, in *RemoveFollowRequest, opts ...http.CallOption) (*RemoveFollowResponse, error) {
	var out RemoveFollowResponse
	pattern := "/follow"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFollowServiceRemoveFollow))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
