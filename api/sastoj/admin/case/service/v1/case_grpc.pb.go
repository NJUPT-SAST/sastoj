// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: api/sastoj/admin/case/service/v1/case.proto

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
	CaseService_CreateCases_FullMethodName            = "/api.sastoj.admin.case.service.v1.CaseService/CreateCases"
	CaseService_UpdateCase_FullMethodName             = "/api.sastoj.admin.case.service.v1.CaseService/UpdateCase"
	CaseService_DeleteCasesByCaseIds_FullMethodName   = "/api.sastoj.admin.case.service.v1.CaseService/DeleteCasesByCaseIds"
	CaseService_DeleteCasesByProblemId_FullMethodName = "/api.sastoj.admin.case.service.v1.CaseService/DeleteCasesByProblemId"
	CaseService_GetCases_FullMethodName               = "/api.sastoj.admin.case.service.v1.CaseService/GetCases"
)

// CaseServiceClient is the client API for CaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaseServiceClient interface {
	CreateCases(ctx context.Context, in *CreateCasesRequest, opts ...grpc.CallOption) (*CreateCasesReply, error)
	UpdateCase(ctx context.Context, in *UpdateCaseRequest, opts ...grpc.CallOption) (*UpdateCaseReply, error)
	DeleteCasesByCaseIds(ctx context.Context, in *DeleteCaseByCaseIdsRequest, opts ...grpc.CallOption) (*DeleteCaseByCaseIdsReply, error)
	DeleteCasesByProblemId(ctx context.Context, in *DeleteCasesByProblemIdRequest, opts ...grpc.CallOption) (*DeleteCasesByProblemIdReply, error)
	GetCases(ctx context.Context, in *GetCasesRequest, opts ...grpc.CallOption) (*GetCasesReply, error)
}

type caseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCaseServiceClient(cc grpc.ClientConnInterface) CaseServiceClient {
	return &caseServiceClient{cc}
}

func (c *caseServiceClient) CreateCases(ctx context.Context, in *CreateCasesRequest, opts ...grpc.CallOption) (*CreateCasesReply, error) {
	out := new(CreateCasesReply)
	err := c.cc.Invoke(ctx, CaseService_CreateCases_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseServiceClient) UpdateCase(ctx context.Context, in *UpdateCaseRequest, opts ...grpc.CallOption) (*UpdateCaseReply, error) {
	out := new(UpdateCaseReply)
	err := c.cc.Invoke(ctx, CaseService_UpdateCase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseServiceClient) DeleteCasesByCaseIds(ctx context.Context, in *DeleteCaseByCaseIdsRequest, opts ...grpc.CallOption) (*DeleteCaseByCaseIdsReply, error) {
	out := new(DeleteCaseByCaseIdsReply)
	err := c.cc.Invoke(ctx, CaseService_DeleteCasesByCaseIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseServiceClient) DeleteCasesByProblemId(ctx context.Context, in *DeleteCasesByProblemIdRequest, opts ...grpc.CallOption) (*DeleteCasesByProblemIdReply, error) {
	out := new(DeleteCasesByProblemIdReply)
	err := c.cc.Invoke(ctx, CaseService_DeleteCasesByProblemId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caseServiceClient) GetCases(ctx context.Context, in *GetCasesRequest, opts ...grpc.CallOption) (*GetCasesReply, error) {
	out := new(GetCasesReply)
	err := c.cc.Invoke(ctx, CaseService_GetCases_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaseServiceServer is the server API for CaseService service.
// All implementations must embed UnimplementedCaseServiceServer
// for forward compatibility
type CaseServiceServer interface {
	CreateCases(context.Context, *CreateCasesRequest) (*CreateCasesReply, error)
	UpdateCase(context.Context, *UpdateCaseRequest) (*UpdateCaseReply, error)
	DeleteCasesByCaseIds(context.Context, *DeleteCaseByCaseIdsRequest) (*DeleteCaseByCaseIdsReply, error)
	DeleteCasesByProblemId(context.Context, *DeleteCasesByProblemIdRequest) (*DeleteCasesByProblemIdReply, error)
	GetCases(context.Context, *GetCasesRequest) (*GetCasesReply, error)
	mustEmbedUnimplementedCaseServiceServer()
}

// UnimplementedCaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCaseServiceServer struct {
}

func (UnimplementedCaseServiceServer) CreateCases(context.Context, *CreateCasesRequest) (*CreateCasesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCases not implemented")
}
func (UnimplementedCaseServiceServer) UpdateCase(context.Context, *UpdateCaseRequest) (*UpdateCaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCase not implemented")
}
func (UnimplementedCaseServiceServer) DeleteCasesByCaseIds(context.Context, *DeleteCaseByCaseIdsRequest) (*DeleteCaseByCaseIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCasesByCaseIds not implemented")
}
func (UnimplementedCaseServiceServer) DeleteCasesByProblemId(context.Context, *DeleteCasesByProblemIdRequest) (*DeleteCasesByProblemIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCasesByProblemId not implemented")
}
func (UnimplementedCaseServiceServer) GetCases(context.Context, *GetCasesRequest) (*GetCasesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCases not implemented")
}
func (UnimplementedCaseServiceServer) mustEmbedUnimplementedCaseServiceServer() {}

// UnsafeCaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaseServiceServer will
// result in compilation errors.
type UnsafeCaseServiceServer interface {
	mustEmbedUnimplementedCaseServiceServer()
}

func RegisterCaseServiceServer(s grpc.ServiceRegistrar, srv CaseServiceServer) {
	s.RegisterService(&CaseService_ServiceDesc, srv)
}

func _CaseService_CreateCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).CreateCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaseService_CreateCases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).CreateCases(ctx, req.(*CreateCasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseService_UpdateCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).UpdateCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaseService_UpdateCase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).UpdateCase(ctx, req.(*UpdateCaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseService_DeleteCasesByCaseIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCaseByCaseIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).DeleteCasesByCaseIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaseService_DeleteCasesByCaseIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).DeleteCasesByCaseIds(ctx, req.(*DeleteCaseByCaseIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseService_DeleteCasesByProblemId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCasesByProblemIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).DeleteCasesByProblemId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaseService_DeleteCasesByProblemId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).DeleteCasesByProblemId(ctx, req.(*DeleteCasesByProblemIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaseService_GetCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaseServiceServer).GetCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaseService_GetCases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaseServiceServer).GetCases(ctx, req.(*GetCasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CaseService_ServiceDesc is the grpc.ServiceDesc for CaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sastoj.admin.case.service.v1.CaseService",
	HandlerType: (*CaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCases",
			Handler:    _CaseService_CreateCases_Handler,
		},
		{
			MethodName: "UpdateCase",
			Handler:    _CaseService_UpdateCase_Handler,
		},
		{
			MethodName: "DeleteCasesByCaseIds",
			Handler:    _CaseService_DeleteCasesByCaseIds_Handler,
		},
		{
			MethodName: "DeleteCasesByProblemId",
			Handler:    _CaseService_DeleteCasesByProblemId_Handler,
		},
		{
			MethodName: "GetCases",
			Handler:    _CaseService_GetCases_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/sastoj/admin/case/service/v1/case.proto",
}