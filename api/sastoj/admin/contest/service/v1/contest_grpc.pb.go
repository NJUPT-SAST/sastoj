// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sastoj/admin/contest/service/v1/contest.proto

package contest

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
	Contest_CreateContest_FullMethodName  = "/api.sastoj.admin.contest.Contest/CreateContest"
	Contest_UpdateContest_FullMethodName  = "/api.sastoj.admin.contest.Contest/UpdateContest"
	Contest_DeleteContest_FullMethodName  = "/api.sastoj.admin.contest.Contest/DeleteContest"
	Contest_GetContest_FullMethodName     = "/api.sastoj.admin.contest.Contest/GetContest"
	Contest_ListContest_FullMethodName    = "/api.sastoj.admin.contest.Contest/ListContest"
	Contest_AddContestants_FullMethodName = "/api.sastoj.admin.contest.Contest/AddContestants"
)

// ContestClient is the client API for Contest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContestClient interface {
	CreateContest(ctx context.Context, in *CreateContestRequest, opts ...grpc.CallOption) (*CreateContestReply, error)
	UpdateContest(ctx context.Context, in *UpdateContestRequest, opts ...grpc.CallOption) (*UpdateContestReply, error)
	DeleteContest(ctx context.Context, in *DeleteContestRequest, opts ...grpc.CallOption) (*DeleteContestReply, error)
	GetContest(ctx context.Context, in *GetContestRequest, opts ...grpc.CallOption) (*GetContestReply, error)
	ListContest(ctx context.Context, in *ListContestRequest, opts ...grpc.CallOption) (*ListContestReply, error)
	AddContestants(ctx context.Context, in *AddContestantsRequest, opts ...grpc.CallOption) (*AddContestantsReply, error)
}

type contestClient struct {
	cc grpc.ClientConnInterface
}

func NewContestClient(cc grpc.ClientConnInterface) ContestClient {
	return &contestClient{cc}
}

func (c *contestClient) CreateContest(ctx context.Context, in *CreateContestRequest, opts ...grpc.CallOption) (*CreateContestReply, error) {
	out := new(CreateContestReply)
	err := c.cc.Invoke(ctx, Contest_CreateContest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestClient) UpdateContest(ctx context.Context, in *UpdateContestRequest, opts ...grpc.CallOption) (*UpdateContestReply, error) {
	out := new(UpdateContestReply)
	err := c.cc.Invoke(ctx, Contest_UpdateContest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestClient) DeleteContest(ctx context.Context, in *DeleteContestRequest, opts ...grpc.CallOption) (*DeleteContestReply, error) {
	out := new(DeleteContestReply)
	err := c.cc.Invoke(ctx, Contest_DeleteContest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestClient) GetContest(ctx context.Context, in *GetContestRequest, opts ...grpc.CallOption) (*GetContestReply, error) {
	out := new(GetContestReply)
	err := c.cc.Invoke(ctx, Contest_GetContest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestClient) ListContest(ctx context.Context, in *ListContestRequest, opts ...grpc.CallOption) (*ListContestReply, error) {
	out := new(ListContestReply)
	err := c.cc.Invoke(ctx, Contest_ListContest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contestClient) AddContestants(ctx context.Context, in *AddContestantsRequest, opts ...grpc.CallOption) (*AddContestantsReply, error) {
	out := new(AddContestantsReply)
	err := c.cc.Invoke(ctx, Contest_AddContestants_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContestServer is the server API for Contest service.
// All implementations must embed UnimplementedContestServer
// for forward compatibility
type ContestServer interface {
	CreateContest(context.Context, *CreateContestRequest) (*CreateContestReply, error)
	UpdateContest(context.Context, *UpdateContestRequest) (*UpdateContestReply, error)
	DeleteContest(context.Context, *DeleteContestRequest) (*DeleteContestReply, error)
	GetContest(context.Context, *GetContestRequest) (*GetContestReply, error)
	ListContest(context.Context, *ListContestRequest) (*ListContestReply, error)
	AddContestants(context.Context, *AddContestantsRequest) (*AddContestantsReply, error)
	mustEmbedUnimplementedContestServer()
}

// UnimplementedContestServer must be embedded to have forward compatible implementations.
type UnimplementedContestServer struct {
}

func (UnimplementedContestServer) CreateContest(context.Context, *CreateContestRequest) (*CreateContestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContest not implemented")
}
func (UnimplementedContestServer) UpdateContest(context.Context, *UpdateContestRequest) (*UpdateContestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContest not implemented")
}
func (UnimplementedContestServer) DeleteContest(context.Context, *DeleteContestRequest) (*DeleteContestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContest not implemented")
}
func (UnimplementedContestServer) GetContest(context.Context, *GetContestRequest) (*GetContestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContest not implemented")
}
func (UnimplementedContestServer) ListContest(context.Context, *ListContestRequest) (*ListContestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContest not implemented")
}
func (UnimplementedContestServer) AddContestants(context.Context, *AddContestantsRequest) (*AddContestantsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddContestants not implemented")
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

func _Contest_CreateContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServer).CreateContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contest_CreateContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServer).CreateContest(ctx, req.(*CreateContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contest_UpdateContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServer).UpdateContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contest_UpdateContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServer).UpdateContest(ctx, req.(*UpdateContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contest_DeleteContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServer).DeleteContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contest_DeleteContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServer).DeleteContest(ctx, req.(*DeleteContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contest_GetContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServer).GetContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contest_GetContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServer).GetContest(ctx, req.(*GetContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contest_ListContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServer).ListContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contest_ListContest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServer).ListContest(ctx, req.(*ListContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contest_AddContestants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddContestantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContestServer).AddContestants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contest_AddContestants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContestServer).AddContestants(ctx, req.(*AddContestantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Contest_ServiceDesc is the grpc.ServiceDesc for Contest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Contest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sastoj.admin.contest.Contest",
	HandlerType: (*ContestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContest",
			Handler:    _Contest_CreateContest_Handler,
		},
		{
			MethodName: "UpdateContest",
			Handler:    _Contest_UpdateContest_Handler,
		},
		{
			MethodName: "DeleteContest",
			Handler:    _Contest_DeleteContest_Handler,
		},
		{
			MethodName: "GetContest",
			Handler:    _Contest_GetContest_Handler,
		},
		{
			MethodName: "ListContest",
			Handler:    _Contest_ListContest_Handler,
		},
		{
			MethodName: "AddContestants",
			Handler:    _Contest_AddContestants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sastoj/admin/contest/service/v1/contest.proto",
}
