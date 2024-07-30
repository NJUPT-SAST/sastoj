// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sastoj/user/contest/service/v1/contest.proto

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
	ContestService_GetContests_FullMethodName     = "/api.sastoj.user.contest.service.v1.ContestService/GetContests"
	ContestService_JoinContest_FullMethodName     = "/api.sastoj.user.contest.service.v1.ContestService/JoinContest"
	ContestService_GetProblems_FullMethodName     = "/api.sastoj.user.contest.service.v1.ContestService/GetProblems"
	ContestService_GetProblem_FullMethodName      = "/api.sastoj.user.contest.service.v1.ContestService/GetProblem"
	ContestService_Submit_FullMethodName          = "/api.sastoj.user.contest.service.v1.ContestService/Submit"
	ContestService_SelfTest_FullMethodName        = "/api.sastoj.user.contest.service.v1.ContestService/SelfTest"
	ContestService_GetSubmission_FullMethodName   = "/api.sastoj.user.contest.service.v1.ContestService/GetSubmission"
	ContestService_ListRanking_FullMethodName     = "/api.sastoj.user.contest.service.v1.ContestService/ListRanking"
	ContestService_RegisterGateway_FullMethodName = "/api.sastoj.user.contest.service.v1.ContestService/RegisterGateway"
)

// ContestServiceClient is the client API for ContestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContestServiceClient interface {
	GetContests(ctx context.Context, in *GetContestsRequest, opts ...grpc.CallOption) (*GetContestsReply, error)
	JoinContest(ctx context.Context, in *JoinContestRequest, opts ...grpc.CallOption) (*JoinContestReply, error)
	GetProblems(ctx context.Context, in *GetProblemsRequest, opts ...grpc.CallOption) (*GetProblemsReply, error)
	GetProblem(ctx context.Context, in *GetProblemRequest, opts ...grpc.CallOption) (*GetProblemReply, error)
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitReply, error)
	SelfTest(ctx context.Context, in *SelfTestRequest, opts ...grpc.CallOption) (*SelfTestReply, error)
	GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...grpc.CallOption) (*GetSubmissionReply, error)
	ListRanking(ctx context.Context, in *ListRankingRequest, opts ...grpc.CallOption) (*ListRankingReply, error)
	RegisterGateway(ctx context.Context, in *RegisterGatewayRequest, opts ...grpc.CallOption) (*RegisterGatewayReply, error)
}

type contestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContestServiceClient(cc grpc.ClientConnInterface) ContestServiceClient {
	return &contestServiceClient{cc}
}

func (c *contestServiceClient) GetContests(ctx context.Context, in *GetContestsRequest, opts ...grpc.CallOption) (*GetContestsReply, error) {
	out := new(GetContestsReply)
	err := c.cc.Invoke(ctx, ContestService_GetContests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) JoinContest(ctx context.Context, in *JoinContestRequest, opts ...grpc.CallOption) (*JoinContestReply, error) {
	out := new(JoinContestReply)
	err := c.cc.Invoke(ctx, ContestService_JoinContest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) GetProblems(ctx context.Context, in *GetProblemsRequest, opts ...grpc.CallOption) (*GetProblemsReply, error) {
	out := new(GetProblemsReply)
	err := c.cc.Invoke(ctx, ContestService_GetProblems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) GetProblem(ctx context.Context, in *GetProblemRequest, opts ...grpc.CallOption) (*GetProblemReply, error) {
	out := new(GetProblemReply)
	err := c.cc.Invoke(ctx, ContestService_GetProblem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitReply, error) {
	out := new(SubmitReply)
	err := c.cc.Invoke(ctx, ContestService_Submit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) SelfTest(ctx context.Context, in *SelfTestRequest, opts ...grpc.CallOption) (*SelfTestReply, error) {
	out := new(SelfTestReply)
	err := c.cc.Invoke(ctx, ContestService_SelfTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...grpc.CallOption) (*GetSubmissionReply, error) {
	out := new(GetSubmissionReply)
	err := c.cc.Invoke(ctx, ContestService_GetSubmission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) ListRanking(ctx context.Context, in *ListRankingRequest, opts ...grpc.CallOption) (*ListRankingReply, error) {
	out := new(ListRankingReply)
	err := c.cc.Invoke(ctx, ContestService_ListRanking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestServiceClient) RegisterGateway(ctx context.Context, in *RegisterGatewayRequest, opts ...grpc.CallOption) (*RegisterGatewayReply, error) {
	out := new(RegisterGatewayReply)
	err := c.cc.Invoke(ctx, ContestService_RegisterGateway_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContestServiceServer is the server API for ContestService service.
// All implementations must embed UnimplementedContestServiceServer
// for forward compatibility
type ContestServiceServer interface {
	GetContests(context.Context, *GetContestsRequest) (*GetContestsReply, error)
	JoinContest(context.Context, *JoinContestRequest) (*JoinContestReply, error)
	GetProblems(context.Context, *GetProblemsRequest) (*GetProblemsReply, error)
	GetProblem(context.Context, *GetProblemRequest) (*GetProblemReply, error)
	Submit(context.Context, *SubmitRequest) (*SubmitReply, error)
	SelfTest(context.Context, *SelfTestRequest) (*SelfTestReply, error)
	GetSubmission(context.Context, *GetSubmissionRequest) (*GetSubmissionReply, error)
	ListRanking(context.Context, *ListRankingRequest) (*ListRankingReply, error)
	RegisterGateway(context.Context, *RegisterGatewayRequest) (*RegisterGatewayReply, error)
	mustEmbedUnimplementedContestServiceServer()
}

// UnimplementedContestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContestServiceServer struct {
}

func (UnimplementedContestServiceServer) GetContests(context.Context, *GetContestsRequest) (*GetContestsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContests not implemented")
}
func (UnimplementedContestServiceServer) JoinContest(context.Context, *JoinContestRequest) (*JoinContestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinContest not implemented")
}
func (UnimplementedContestServiceServer) GetProblems(context.Context, *GetProblemsRequest) (*GetProblemsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblems not implemented")
}
func (UnimplementedContestServiceServer) GetProblem(context.Context, *GetProblemRequest) (*GetProblemReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblem not implemented")
}
func (UnimplementedContestServiceServer) Submit(context.Context, *SubmitRequest) (*SubmitReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedContestServiceServer) SelfTest(context.Context, *SelfTestRequest) (*SelfTestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelfTest not implemented")
}
func (UnimplementedContestServiceServer) GetSubmission(context.Context, *GetSubmissionRequest) (*GetSubmissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmission not implemented")
}
func (UnimplementedContestServiceServer) ListRanking(context.Context, *ListRankingRequest) (*ListRankingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRanking not implemented")
}
func (UnimplementedContestServiceServer) RegisterGateway(context.Context, *RegisterGatewayRequest) (*RegisterGatewayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterGateway not implemented")
}
func (UnimplementedContestServiceServer) mustEmbedUnimplementedContestServiceServer() {}

// UnsafeContestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContestServiceServer will
// result in compilation errors.
type UnsafeContestServiceServer interface {
	mustEmbedUnimplementedContestServiceServer()
}

func RegisterContestServiceServer(s grpc.ServiceRegistrar, srv ContestServiceServer) {
	s.RegisterService(&ContestService_ServiceDesc, srv)
}

func _ContestService_GetContests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).GetContests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_GetContests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).GetContests(ctx, req.(*GetContestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_JoinContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).JoinContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_JoinContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).JoinContest(ctx, req.(*JoinContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_GetProblems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProblemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).GetProblems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_GetProblems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).GetProblems(ctx, req.(*GetProblemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_GetProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProblemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).GetProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_GetProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).GetProblem(ctx, req.(*GetProblemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_Submit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_SelfTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelfTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).SelfTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_SelfTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).SelfTest(ctx, req.(*SelfTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_GetSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).GetSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_GetSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).GetSubmission(ctx, req.(*GetSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_ListRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).ListRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_ListRanking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).ListRanking(ctx, req.(*ListRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContestService_RegisterGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServiceServer).RegisterGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContestService_RegisterGateway_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServiceServer).RegisterGateway(ctx, req.(*RegisterGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContestService_ServiceDesc is the grpc.ServiceDesc for ContestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sastoj.user.contest.service.v1.ContestService",
	HandlerType: (*ContestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContests",
			Handler:    _ContestService_GetContests_Handler,
		},
		{
			MethodName: "JoinContest",
			Handler:    _ContestService_JoinContest_Handler,
		},
		{
			MethodName: "GetProblems",
			Handler:    _ContestService_GetProblems_Handler,
		},
		{
			MethodName: "GetProblem",
			Handler:    _ContestService_GetProblem_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _ContestService_Submit_Handler,
		},
		{
			MethodName: "SelfTest",
			Handler:    _ContestService_SelfTest_Handler,
		},
		{
			MethodName: "GetSubmission",
			Handler:    _ContestService_GetSubmission_Handler,
		},
		{
			MethodName: "ListRanking",
			Handler:    _ContestService_ListRanking_Handler,
		},
		{
			MethodName: "RegisterGateway",
			Handler:    _ContestService_RegisterGateway_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sastoj/user/contest/service/v1/contest.proto",
}
