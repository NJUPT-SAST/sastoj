// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             (unknown)
// source: sastoj/user/gateway/service/v1/contest.proto

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

const OperationContestGetContests = "/api.sastoj.user.gateway.service.v1.Contest/GetContests"
const OperationContestJoinContest = "/api.sastoj.user.gateway.service.v1.Contest/JoinContest"

type ContestHTTPServer interface {
	GetContests(context.Context, *GetContestsRequest) (*GetContestsReply, error)
	JoinContest(context.Context, *JoinContestRequest) (*JoinContestReply, error)
}

func RegisterContestHTTPServer(s *http.Server, srv ContestHTTPServer) {
	r := s.Route("/")
	r.GET("/user/contests", _Contest_GetContests0_HTTP_Handler(srv))
	r.POST("/user/contests/{contest_id}", _Contest_JoinContest0_HTTP_Handler(srv))
}

func _Contest_GetContests0_HTTP_Handler(srv ContestHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetContestsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContestGetContests)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetContests(ctx, req.(*GetContestsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetContestsReply)
		return ctx.Result(200, reply)
	}
}

func _Contest_JoinContest0_HTTP_Handler(srv ContestHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in JoinContestRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationContestJoinContest)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.JoinContest(ctx, req.(*JoinContestRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*JoinContestReply)
		return ctx.Result(200, reply)
	}
}

type ContestHTTPClient interface {
	GetContests(ctx context.Context, req *GetContestsRequest, opts ...http.CallOption) (rsp *GetContestsReply, err error)
	JoinContest(ctx context.Context, req *JoinContestRequest, opts ...http.CallOption) (rsp *JoinContestReply, err error)
}

type ContestHTTPClientImpl struct {
	cc *http.Client
}

func NewContestHTTPClient(client *http.Client) ContestHTTPClient {
	return &ContestHTTPClientImpl{client}
}

func (c *ContestHTTPClientImpl) GetContests(ctx context.Context, in *GetContestsRequest, opts ...http.CallOption) (*GetContestsReply, error) {
	var out GetContestsReply
	pattern := "/user/contests"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationContestGetContests))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ContestHTTPClientImpl) JoinContest(ctx context.Context, in *JoinContestRequest, opts ...http.CallOption) (*JoinContestReply, error) {
	var out JoinContestReply
	pattern := "/user/contests/{contest_id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationContestJoinContest))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

const OperationProblemGetProblem = "/api.sastoj.user.gateway.service.v1.Problem/GetProblem"
const OperationProblemGetProblems = "/api.sastoj.user.gateway.service.v1.Problem/GetProblems"

type ProblemHTTPServer interface {
	GetProblem(context.Context, *GetProblemRequest) (*GetProblemReply, error)
	GetProblems(context.Context, *GetProblemsRequest) (*GetProblemsReply, error)
}

func RegisterProblemHTTPServer(s *http.Server, srv ProblemHTTPServer) {
	r := s.Route("/")
	r.GET("/user/contests/{contest_id}/problems", _Problem_GetProblems0_HTTP_Handler(srv))
	r.GET("/user/contests/{contest_id}/problems/{problem_id}", _Problem_GetProblem0_HTTP_Handler(srv))
}

func _Problem_GetProblems0_HTTP_Handler(srv ProblemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProblemsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProblemGetProblems)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProblems(ctx, req.(*GetProblemsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProblemsReply)
		return ctx.Result(200, reply)
	}
}

func _Problem_GetProblem0_HTTP_Handler(srv ProblemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProblemRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProblemGetProblem)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProblem(ctx, req.(*GetProblemRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProblemReply)
		return ctx.Result(200, reply)
	}
}

type ProblemHTTPClient interface {
	GetProblem(ctx context.Context, req *GetProblemRequest, opts ...http.CallOption) (rsp *GetProblemReply, err error)
	GetProblems(ctx context.Context, req *GetProblemsRequest, opts ...http.CallOption) (rsp *GetProblemsReply, err error)
}

type ProblemHTTPClientImpl struct {
	cc *http.Client
}

func NewProblemHTTPClient(client *http.Client) ProblemHTTPClient {
	return &ProblemHTTPClientImpl{client}
}

func (c *ProblemHTTPClientImpl) GetProblem(ctx context.Context, in *GetProblemRequest, opts ...http.CallOption) (*GetProblemReply, error) {
	var out GetProblemReply
	pattern := "/user/contests/{contest_id}/problems/{problem_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProblemGetProblem))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ProblemHTTPClientImpl) GetProblems(ctx context.Context, in *GetProblemsRequest, opts ...http.CallOption) (*GetProblemsReply, error) {
	var out GetProblemsReply
	pattern := "/user/contests/{contest_id}/problems"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProblemGetProblems))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

const OperationSubmissionGetSubmission = "/api.sastoj.user.gateway.service.v1.Submission/GetSubmission"
const OperationSubmissionListRanking = "/api.sastoj.user.gateway.service.v1.Submission/ListRanking"
const OperationSubmissionSelfTest = "/api.sastoj.user.gateway.service.v1.Submission/SelfTest"
const OperationSubmissionSubmit = "/api.sastoj.user.gateway.service.v1.Submission/Submit"

type SubmissionHTTPServer interface {
	GetSubmission(context.Context, *GetSubmissionRequest) (*GetSubmissionReply, error)
	ListRanking(context.Context, *ListRankingRequest) (*ListRankingReply, error)
	SelfTest(context.Context, *SelfTestRequest) (*SelfTestReply, error)
	Submit(context.Context, *SubmitRequest) (*SubmitReply, error)
}

func RegisterSubmissionHTTPServer(s *http.Server, srv SubmissionHTTPServer) {
	r := s.Route("/")
	r.POST("/user/contests/{contest_id}/problems/{problem_id}/submission", _Submission_Submit0_HTTP_Handler(srv))
	r.POST("/user/contests/{contest_id}/problems/{problem_id}/test", _Submission_SelfTest0_HTTP_Handler(srv))
	r.GET("/user/contests/{contest_id}/problems/{problem_id}/submissions/{submission_id}", _Submission_GetSubmission0_HTTP_Handler(srv))
	r.GET("/user/contests/{contest_id}/ranking", _Submission_ListRanking0_HTTP_Handler(srv))
}

func _Submission_Submit0_HTTP_Handler(srv SubmissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SubmitRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubmissionSubmit)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Submit(ctx, req.(*SubmitRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SubmitReply)
		return ctx.Result(200, reply)
	}
}

func _Submission_SelfTest0_HTTP_Handler(srv SubmissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SelfTestRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubmissionSelfTest)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SelfTest(ctx, req.(*SelfTestRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SelfTestReply)
		return ctx.Result(200, reply)
	}
}

func _Submission_GetSubmission0_HTTP_Handler(srv SubmissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSubmissionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubmissionGetSubmission)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSubmission(ctx, req.(*GetSubmissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSubmissionReply)
		return ctx.Result(200, reply)
	}
}

func _Submission_ListRanking0_HTTP_Handler(srv SubmissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRankingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSubmissionListRanking)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListRanking(ctx, req.(*ListRankingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListRankingReply)
		return ctx.Result(200, reply)
	}
}

type SubmissionHTTPClient interface {
	GetSubmission(ctx context.Context, req *GetSubmissionRequest, opts ...http.CallOption) (rsp *GetSubmissionReply, err error)
	ListRanking(ctx context.Context, req *ListRankingRequest, opts ...http.CallOption) (rsp *ListRankingReply, err error)
	SelfTest(ctx context.Context, req *SelfTestRequest, opts ...http.CallOption) (rsp *SelfTestReply, err error)
	Submit(ctx context.Context, req *SubmitRequest, opts ...http.CallOption) (rsp *SubmitReply, err error)
}

type SubmissionHTTPClientImpl struct {
	cc *http.Client
}

func NewSubmissionHTTPClient(client *http.Client) SubmissionHTTPClient {
	return &SubmissionHTTPClientImpl{client}
}

func (c *SubmissionHTTPClientImpl) GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...http.CallOption) (*GetSubmissionReply, error) {
	var out GetSubmissionReply
	pattern := "/user/contests/{contest_id}/problems/{problem_id}/submissions/{submission_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSubmissionGetSubmission))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SubmissionHTTPClientImpl) ListRanking(ctx context.Context, in *ListRankingRequest, opts ...http.CallOption) (*ListRankingReply, error) {
	var out ListRankingReply
	pattern := "/user/contests/{contest_id}/ranking"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSubmissionListRanking))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SubmissionHTTPClientImpl) SelfTest(ctx context.Context, in *SelfTestRequest, opts ...http.CallOption) (*SelfTestReply, error) {
	var out SelfTestReply
	pattern := "/user/contests/{contest_id}/problems/{problem_id}/test"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSubmissionSelfTest))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SubmissionHTTPClientImpl) Submit(ctx context.Context, in *SubmitRequest, opts ...http.CallOption) (*SubmitReply, error) {
	var out SubmitReply
	pattern := "/user/contests/{contest_id}/problems/{problem_id}/submission"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSubmissionSubmit))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
