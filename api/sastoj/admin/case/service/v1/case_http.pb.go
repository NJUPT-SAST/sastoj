// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.26.1
// source: sastoj/admin/case/service/v1/case.proto

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

const OperationCaseServiceCreateCases = "/api.sastoj.admin.case.service.v1.CaseService/CreateCases"
const OperationCaseServiceDeleteCasesByCaseIds = "/api.sastoj.admin.case.service.v1.CaseService/DeleteCasesByCaseIds"
const OperationCaseServiceDeleteCasesByProblemId = "/api.sastoj.admin.case.service.v1.CaseService/DeleteCasesByProblemId"
const OperationCaseServiceGetCases = "/api.sastoj.admin.case.service.v1.CaseService/GetCases"
const OperationCaseServiceUpdateCase = "/api.sastoj.admin.case.service.v1.CaseService/UpdateCase"

type CaseServiceHTTPServer interface {
	CreateCases(context.Context, *CreateCasesRequest) (*CreateCasesReply, error)
	DeleteCasesByCaseIds(context.Context, *DeleteCaseByCaseIdsRequest) (*DeleteCaseByCaseIdsReply, error)
	DeleteCasesByProblemId(context.Context, *DeleteCasesByProblemIdRequest) (*DeleteCasesByProblemIdReply, error)
	GetCases(context.Context, *GetCasesRequest) (*GetCasesReply, error)
	UpdateCase(context.Context, *UpdateCaseRequest) (*UpdateCaseReply, error)
}

func RegisterCaseServiceHTTPServer(s *http.Server, srv CaseServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/case/{problem_id}", _CaseService_CreateCases0_HTTP_Handler(srv))
	r.PUT("/case", _CaseService_UpdateCase0_HTTP_Handler(srv))
	r.PUT("/case/delete", _CaseService_DeleteCasesByCaseIds0_HTTP_Handler(srv))
	r.DELETE("/case/{problem_id}", _CaseService_DeleteCasesByProblemId0_HTTP_Handler(srv))
	r.GET("/case/{problem_id}", _CaseService_GetCases0_HTTP_Handler(srv))
}

func _CaseService_CreateCases0_HTTP_Handler(srv CaseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCasesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCaseServiceCreateCases)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCases(ctx, req.(*CreateCasesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateCasesReply)
		return ctx.Result(200, reply)
	}
}

func _CaseService_UpdateCase0_HTTP_Handler(srv CaseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCaseRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCaseServiceUpdateCase)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCase(ctx, req.(*UpdateCaseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateCaseReply)
		return ctx.Result(200, reply)
	}
}

func _CaseService_DeleteCasesByCaseIds0_HTTP_Handler(srv CaseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCaseByCaseIdsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCaseServiceDeleteCasesByCaseIds)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCasesByCaseIds(ctx, req.(*DeleteCaseByCaseIdsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteCaseByCaseIdsReply)
		return ctx.Result(200, reply)
	}
}

func _CaseService_DeleteCasesByProblemId0_HTTP_Handler(srv CaseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCasesByProblemIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCaseServiceDeleteCasesByProblemId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCasesByProblemId(ctx, req.(*DeleteCasesByProblemIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteCasesByProblemIdReply)
		return ctx.Result(200, reply)
	}
}

func _CaseService_GetCases0_HTTP_Handler(srv CaseServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCasesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCaseServiceGetCases)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCases(ctx, req.(*GetCasesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCasesReply)
		return ctx.Result(200, reply)
	}
}

type CaseServiceHTTPClient interface {
	CreateCases(ctx context.Context, req *CreateCasesRequest, opts ...http.CallOption) (rsp *CreateCasesReply, err error)
	DeleteCasesByCaseIds(ctx context.Context, req *DeleteCaseByCaseIdsRequest, opts ...http.CallOption) (rsp *DeleteCaseByCaseIdsReply, err error)
	DeleteCasesByProblemId(ctx context.Context, req *DeleteCasesByProblemIdRequest, opts ...http.CallOption) (rsp *DeleteCasesByProblemIdReply, err error)
	GetCases(ctx context.Context, req *GetCasesRequest, opts ...http.CallOption) (rsp *GetCasesReply, err error)
	UpdateCase(ctx context.Context, req *UpdateCaseRequest, opts ...http.CallOption) (rsp *UpdateCaseReply, err error)
}

type CaseServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewCaseServiceHTTPClient(client *http.Client) CaseServiceHTTPClient {
	return &CaseServiceHTTPClientImpl{client}
}

func (c *CaseServiceHTTPClientImpl) CreateCases(ctx context.Context, in *CreateCasesRequest, opts ...http.CallOption) (*CreateCasesReply, error) {
	var out CreateCasesReply
	pattern := "/case/{problem_id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCaseServiceCreateCases))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CaseServiceHTTPClientImpl) DeleteCasesByCaseIds(ctx context.Context, in *DeleteCaseByCaseIdsRequest, opts ...http.CallOption) (*DeleteCaseByCaseIdsReply, error) {
	var out DeleteCaseByCaseIdsReply
	pattern := "/case/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCaseServiceDeleteCasesByCaseIds))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CaseServiceHTTPClientImpl) DeleteCasesByProblemId(ctx context.Context, in *DeleteCasesByProblemIdRequest, opts ...http.CallOption) (*DeleteCasesByProblemIdReply, error) {
	var out DeleteCasesByProblemIdReply
	pattern := "/case/{problem_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCaseServiceDeleteCasesByProblemId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CaseServiceHTTPClientImpl) GetCases(ctx context.Context, in *GetCasesRequest, opts ...http.CallOption) (*GetCasesReply, error) {
	var out GetCasesReply
	pattern := "/case/{problem_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCaseServiceGetCases))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CaseServiceHTTPClientImpl) UpdateCase(ctx context.Context, in *UpdateCaseRequest, opts ...http.CallOption) (*UpdateCaseReply, error) {
	var out UpdateCaseReply
	pattern := "/case"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCaseServiceUpdateCase))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
