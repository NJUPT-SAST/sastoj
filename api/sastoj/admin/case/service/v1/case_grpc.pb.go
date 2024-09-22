// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: admin/case/service/v1/case.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CaseService_DeleteCasesByProblemId_FullMethodName = "/api.sastoj.admin.case.service.v1.CaseService/DeleteCasesByProblemId"
)

// CaseServiceClient is the client API for CaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaseServiceClient interface {
	DeleteCasesByProblemId(ctx context.Context, in *DeleteCasesByProblemIdRequest, opts ...grpc.CallOption) (*DeleteCasesByProblemIdReply, error)
}

type caseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCaseServiceClient(cc grpc.ClientConnInterface) CaseServiceClient {
	return &caseServiceClient{cc}
}

func (c *caseServiceClient) DeleteCasesByProblemId(ctx context.Context, in *DeleteCasesByProblemIdRequest, opts ...grpc.CallOption) (*DeleteCasesByProblemIdReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCasesByProblemIdReply)
	err := c.cc.Invoke(ctx, CaseService_DeleteCasesByProblemId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaseServiceServer is the server API for CaseService service.
// All implementations must embed UnimplementedCaseServiceServer
// for forward compatibility.
type CaseServiceServer interface {
	DeleteCasesByProblemId(context.Context, *DeleteCasesByProblemIdRequest) (*DeleteCasesByProblemIdReply, error)
	mustEmbedUnimplementedCaseServiceServer()
}

// UnimplementedCaseServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCaseServiceServer struct{}

func (UnimplementedCaseServiceServer) DeleteCasesByProblemId(context.Context, *DeleteCasesByProblemIdRequest) (*DeleteCasesByProblemIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCasesByProblemId not implemented")
}
func (UnimplementedCaseServiceServer) mustEmbedUnimplementedCaseServiceServer() {}
func (UnimplementedCaseServiceServer) testEmbeddedByValue()                     {}

// UnsafeCaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaseServiceServer will
// result in compilation errors.
type UnsafeCaseServiceServer interface {
	mustEmbedUnimplementedCaseServiceServer()
}

func RegisterCaseServiceServer(s grpc.ServiceRegistrar, srv CaseServiceServer) {
	// If the following call pancis, it indicates UnimplementedCaseServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CaseService_ServiceDesc, srv)
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

// CaseService_ServiceDesc is the grpc.ServiceDesc for CaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sastoj.admin.case.service.v1.CaseService",
	HandlerType: (*CaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteCasesByProblemId",
			Handler:    _CaseService_DeleteCasesByProblemId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/case/service/v1/case.proto",
}
