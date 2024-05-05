// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: sastoj/user/gateway/service/v1/contest.proto

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

const (
	Contest_GetContests_FullMethodName = "/api.sastoj.user.gateway.service.v1.Contest/GetContests"
	Contest_JoinContest_FullMethodName = "/api.sastoj.user.gateway.service.v1.Contest/JoinContest"
)

// ContestClient is the client API for Contest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContestClient interface {
	GetContests(ctx context.Context, in *GetContestsRequest, opts ...grpc.CallOption) (*GetContestsReply, error)
	JoinContest(ctx context.Context, in *JoinContestRequest, opts ...grpc.CallOption) (*JoinContestReply, error)
}

type contestClient struct {
	cc grpc.ClientConnInterface
}

func NewContestClient(cc grpc.ClientConnInterface) ContestClient {
	return &contestClient{cc}
}

func (c *contestClient) GetContests(ctx context.Context, in *GetContestsRequest, opts ...grpc.CallOption) (*GetContestsReply, error) {
	out := new(GetContestsReply)
	err := c.cc.Invoke(ctx, Contest_GetContests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestClient) JoinContest(ctx context.Context, in *JoinContestRequest, opts ...grpc.CallOption) (*JoinContestReply, error) {
	out := new(JoinContestReply)
	err := c.cc.Invoke(ctx, Contest_JoinContest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContestServer is the server API for Contest service.
// All implementations must embed UnimplementedContestServer
// for forward compatibility
type ContestServer interface {
	GetContests(context.Context, *GetContestsRequest) (*GetContestsReply, error)
	JoinContest(context.Context, *JoinContestRequest) (*JoinContestReply, error)
	mustEmbedUnimplementedContestServer()
}

// UnimplementedContestServer must be embedded to have forward compatible implementations.
type UnimplementedContestServer struct {
}

func (UnimplementedContestServer) GetContests(context.Context, *GetContestsRequest) (*GetContestsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContests not implemented")
}
func (UnimplementedContestServer) JoinContest(context.Context, *JoinContestRequest) (*JoinContestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinContest not implemented")
}
func (UnimplementedContestServer) mustEmbedUnimplementedContestServer() {}

// UnsafeContestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContestServer will
// result in compilation errors.
type UnsafeContestServer interface {
	mustEmbedUnimplementedContestServer()
}

func RegisterContestServer(s grpc.ServiceRegistrar, srv ContestServer) {
	s.RegisterService(&Contest_ServiceDesc, srv)
}

func _Contest_GetContests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServer).GetContests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contest_GetContests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServer).GetContests(ctx, req.(*GetContestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contest_JoinContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServer).JoinContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contest_JoinContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServer).JoinContest(ctx, req.(*JoinContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Contest_ServiceDesc is the grpc.ServiceDesc for Contest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Contest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sastoj.user.gateway.service.v1.Contest",
	HandlerType: (*ContestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContests",
			Handler:    _Contest_GetContests_Handler,
		},
		{
			MethodName: "JoinContest",
			Handler:    _Contest_JoinContest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sastoj/user/gateway/service/v1/contest.proto",
}

const (
	Problem_GetProblems_FullMethodName = "/api.sastoj.user.gateway.service.v1.Problem/GetProblems"
	Problem_GetProblem_FullMethodName  = "/api.sastoj.user.gateway.service.v1.Problem/GetProblem"
)

// ProblemClient is the client API for Problem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProblemClient interface {
	GetProblems(ctx context.Context, in *GetProblemsRequest, opts ...grpc.CallOption) (*GetProblemsReply, error)
	GetProblem(ctx context.Context, in *GetProblemRequest, opts ...grpc.CallOption) (*GetProblemReply, error)
}

type problemClient struct {
	cc grpc.ClientConnInterface
}

func NewProblemClient(cc grpc.ClientConnInterface) ProblemClient {
	return &problemClient{cc}
}

func (c *problemClient) GetProblems(ctx context.Context, in *GetProblemsRequest, opts ...grpc.CallOption) (*GetProblemsReply, error) {
	out := new(GetProblemsReply)
	err := c.cc.Invoke(ctx, Problem_GetProblems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemClient) GetProblem(ctx context.Context, in *GetProblemRequest, opts ...grpc.CallOption) (*GetProblemReply, error) {
	out := new(GetProblemReply)
	err := c.cc.Invoke(ctx, Problem_GetProblem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProblemServer is the server API for Problem service.
// All implementations must embed UnimplementedProblemServer
// for forward compatibility
type ProblemServer interface {
	GetProblems(context.Context, *GetProblemsRequest) (*GetProblemsReply, error)
	GetProblem(context.Context, *GetProblemRequest) (*GetProblemReply, error)
	mustEmbedUnimplementedProblemServer()
}

// UnimplementedProblemServer must be embedded to have forward compatible implementations.
type UnimplementedProblemServer struct {
}

func (UnimplementedProblemServer) GetProblems(context.Context, *GetProblemsRequest) (*GetProblemsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblems not implemented")
}
func (UnimplementedProblemServer) GetProblem(context.Context, *GetProblemRequest) (*GetProblemReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblem not implemented")
}
func (UnimplementedProblemServer) mustEmbedUnimplementedProblemServer() {}

// UnsafeProblemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProblemServer will
// result in compilation errors.
type UnsafeProblemServer interface {
	mustEmbedUnimplementedProblemServer()
}

func RegisterProblemServer(s grpc.ServiceRegistrar, srv ProblemServer) {
	s.RegisterService(&Problem_ServiceDesc, srv)
}

func _Problem_GetProblems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProblemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServer).GetProblems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Problem_GetProblems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServer).GetProblems(ctx, req.(*GetProblemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Problem_GetProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProblemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServer).GetProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Problem_GetProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServer).GetProblem(ctx, req.(*GetProblemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Problem_ServiceDesc is the grpc.ServiceDesc for Problem service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Problem_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sastoj.user.gateway.service.v1.Problem",
	HandlerType: (*ProblemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProblems",
			Handler:    _Problem_GetProblems_Handler,
		},
		{
			MethodName: "GetProblem",
			Handler:    _Problem_GetProblem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sastoj/user/gateway/service/v1/contest.proto",
}

const (
	Submission_Submit_FullMethodName           = "/api.sastoj.user.gateway.service.v1.Submission/Submit"
	Submission_SelfTest_FullMethodName         = "/api.sastoj.user.gateway.service.v1.Submission/SelfTest"
	Submission_GetSubmission_FullMethodName    = "/api.sastoj.user.gateway.service.v1.Submission/GetSubmission"
	Submission_ListRanking_FullMethodName      = "/api.sastoj.user.gateway.service.v1.Submission/ListRanking"
	Submission_UpdateSubmission_FullMethodName = "/api.sastoj.user.gateway.service.v1.Submission/UpdateSubmission"
	Submission_UpdateSelfTest_FullMethodName   = "/api.sastoj.user.gateway.service.v1.Submission/UpdateSelfTest"
)

// SubmissionClient is the client API for Submission service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubmissionClient interface {
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitReply, error)
	SelfTest(ctx context.Context, in *SelfTestRequest, opts ...grpc.CallOption) (*SelfTestReply, error)
	GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...grpc.CallOption) (*GetSubmissionReply, error)
	ListRanking(ctx context.Context, in *ListRankingRequest, opts ...grpc.CallOption) (*ListRankingReply, error)
	UpdateSubmission(ctx context.Context, in *UpdateSubmissionRequest, opts ...grpc.CallOption) (*UpdateSubmissionReply, error)
	UpdateSelfTest(ctx context.Context, in *UpdateSelfTestRequest, opts ...grpc.CallOption) (*UpdateSelfTestReply, error)
}

type submissionClient struct {
	cc grpc.ClientConnInterface
}

func NewSubmissionClient(cc grpc.ClientConnInterface) SubmissionClient {
	return &submissionClient{cc}
}

func (c *submissionClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitReply, error) {
	out := new(SubmitReply)
	err := c.cc.Invoke(ctx, Submission_Submit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionClient) SelfTest(ctx context.Context, in *SelfTestRequest, opts ...grpc.CallOption) (*SelfTestReply, error) {
	out := new(SelfTestReply)
	err := c.cc.Invoke(ctx, Submission_SelfTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionClient) GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...grpc.CallOption) (*GetSubmissionReply, error) {
	out := new(GetSubmissionReply)
	err := c.cc.Invoke(ctx, Submission_GetSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionClient) ListRanking(ctx context.Context, in *ListRankingRequest, opts ...grpc.CallOption) (*ListRankingReply, error) {
	out := new(ListRankingReply)
	err := c.cc.Invoke(ctx, Submission_ListRanking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionClient) UpdateSubmission(ctx context.Context, in *UpdateSubmissionRequest, opts ...grpc.CallOption) (*UpdateSubmissionReply, error) {
	out := new(UpdateSubmissionReply)
	err := c.cc.Invoke(ctx, Submission_UpdateSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *submissionClient) UpdateSelfTest(ctx context.Context, in *UpdateSelfTestRequest, opts ...grpc.CallOption) (*UpdateSelfTestReply, error) {
	out := new(UpdateSelfTestReply)
	err := c.cc.Invoke(ctx, Submission_UpdateSelfTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubmissionServer is the server API for Submission service.
// All implementations must embed UnimplementedSubmissionServer
// for forward compatibility
type SubmissionServer interface {
	Submit(context.Context, *SubmitRequest) (*SubmitReply, error)
	SelfTest(context.Context, *SelfTestRequest) (*SelfTestReply, error)
	GetSubmission(context.Context, *GetSubmissionRequest) (*GetSubmissionReply, error)
	ListRanking(context.Context, *ListRankingRequest) (*ListRankingReply, error)
	UpdateSubmission(context.Context, *UpdateSubmissionRequest) (*UpdateSubmissionReply, error)
	UpdateSelfTest(context.Context, *UpdateSelfTestRequest) (*UpdateSelfTestReply, error)
	mustEmbedUnimplementedSubmissionServer()
}

// UnimplementedSubmissionServer must be embedded to have forward compatible implementations.
type UnimplementedSubmissionServer struct {
}

func (UnimplementedSubmissionServer) Submit(context.Context, *SubmitRequest) (*SubmitReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedSubmissionServer) SelfTest(context.Context, *SelfTestRequest) (*SelfTestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelfTest not implemented")
}
func (UnimplementedSubmissionServer) GetSubmission(context.Context, *GetSubmissionRequest) (*GetSubmissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmission not implemented")
}
func (UnimplementedSubmissionServer) ListRanking(context.Context, *ListRankingRequest) (*ListRankingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRanking not implemented")
}
func (UnimplementedSubmissionServer) UpdateSubmission(context.Context, *UpdateSubmissionRequest) (*UpdateSubmissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubmission not implemented")
}
func (UnimplementedSubmissionServer) UpdateSelfTest(context.Context, *UpdateSelfTestRequest) (*UpdateSelfTestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSelfTest not implemented")
}
func (UnimplementedSubmissionServer) mustEmbedUnimplementedSubmissionServer() {}

// UnsafeSubmissionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubmissionServer will
// result in compilation errors.
type UnsafeSubmissionServer interface {
	mustEmbedUnimplementedSubmissionServer()
}

func RegisterSubmissionServer(s grpc.ServiceRegistrar, srv SubmissionServer) {
	s.RegisterService(&Submission_ServiceDesc, srv)
}

func _Submission_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Submission_Submit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Submission_SelfTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelfTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServer).SelfTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Submission_SelfTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServer).SelfTest(ctx, req.(*SelfTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Submission_GetSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServer).GetSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Submission_GetSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServer).GetSubmission(ctx, req.(*GetSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Submission_ListRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServer).ListRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Submission_ListRanking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServer).ListRanking(ctx, req.(*ListRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Submission_UpdateSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServer).UpdateSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Submission_UpdateSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServer).UpdateSubmission(ctx, req.(*UpdateSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Submission_UpdateSelfTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSelfTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServer).UpdateSelfTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Submission_UpdateSelfTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServer).UpdateSelfTest(ctx, req.(*UpdateSelfTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Submission_ServiceDesc is the grpc.ServiceDesc for Submission service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Submission_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sastoj.user.gateway.service.v1.Submission",
	HandlerType: (*SubmissionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _Submission_Submit_Handler,
		},
		{
			MethodName: "SelfTest",
			Handler:    _Submission_SelfTest_Handler,
		},
		{
			MethodName: "GetSubmission",
			Handler:    _Submission_GetSubmission_Handler,
		},
		{
			MethodName: "ListRanking",
			Handler:    _Submission_ListRanking_Handler,
		},
		{
			MethodName: "UpdateSubmission",
			Handler:    _Submission_UpdateSubmission_Handler,
		},
		{
			MethodName: "UpdateSelfTest",
			Handler:    _Submission_UpdateSelfTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sastoj/user/gateway/service/v1/contest.proto",
}
