// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             (unknown)
// source: sastoj/admin/judger/service/v1/judger.proto

package v1

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

const OperationJudgerGetJudger = "/api.sastoj.admin.judger.service.v1.Judger/GetJudger"
const OperationJudgerUpdateJudger = "/api.sastoj.admin.judger.service.v1.Judger/UpdateJudger"

type JudgerHTTPServer interface {
	GetJudger(context.Context, *GetJudgerRequest) (*GetJudgerReply, error)
	UpdateJudger(context.Context, *UpdateJudgerRequest) (*UpdateJudgerReply, error)
}

func RegisterJudgerHTTPServer(s *http.Server, srv JudgerHTTPServer) {
	r := s.Route("/")
	r.POST("/judger/{problem_id}", _Judger_UpdateJudger0_HTTP_Handler(srv))
	r.GET("/judger/{problem_id}", _Judger_GetJudger0_HTTP_Handler(srv))
}

func _Judger_UpdateJudger0_HTTP_Handler(srv JudgerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateJudgerRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationJudgerUpdateJudger)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateJudger(ctx, req.(*UpdateJudgerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateJudgerReply)
		return ctx.Result(200, reply)
	}
}

func _Judger_GetJudger0_HTTP_Handler(srv JudgerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetJudgerRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationJudgerGetJudger)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetJudger(ctx, req.(*GetJudgerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetJudgerReply)
		return ctx.Result(200, reply)
	}
}

type JudgerHTTPClient interface {
	GetJudger(ctx context.Context, req *GetJudgerRequest, opts ...http.CallOption) (rsp *GetJudgerReply, err error)
	UpdateJudger(ctx context.Context, req *UpdateJudgerRequest, opts ...http.CallOption) (rsp *UpdateJudgerReply, err error)
}

type JudgerHTTPClientImpl struct {
	cc *http.Client
}

func NewJudgerHTTPClient(client *http.Client) JudgerHTTPClient {
	return &JudgerHTTPClientImpl{client}
}

func (c *JudgerHTTPClientImpl) GetJudger(ctx context.Context, in *GetJudgerRequest, opts ...http.CallOption) (*GetJudgerReply, error) {
	var out GetJudgerReply
	pattern := "/judger/{problem_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationJudgerGetJudger))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *JudgerHTTPClientImpl) UpdateJudger(ctx context.Context, in *UpdateJudgerRequest, opts ...http.CallOption) (*UpdateJudgerReply, error) {
	var out UpdateJudgerReply
	pattern := "/judger/{problem_id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationJudgerUpdateJudger))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
