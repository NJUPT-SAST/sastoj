// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: api/sastoj/gojudge/judger/service/v1/judge.proto

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
	Judge_CreateOne_FullMethodName    = "/sastoj.gojudge.judger.service.v1.Judge/CreateOne"
	Judge_FetchStatus_FullMethodName  = "/sastoj.gojudge.judger.service.v1.Judge/FetchStatus"
	Judge_ModifyConfig_FullMethodName = "/sastoj.gojudge.judger.service.v1.Judge/ModifyConfig"
	Judge_DeleteOne_FullMethodName    = "/sastoj.gojudge.judger.service.v1.Judge/DeleteOne"
	Judge_ListAll_FullMethodName      = "/sastoj.gojudge.judger.service.v1.Judge/ListAll"
)

// JudgeClient is the client API for Judge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JudgeClient interface {
	// CreateOne create a judge with config
	CreateOne(ctx context.Context, in *CreateOneRequest, opts ...grpc.CallOption) (*CreateOneReply, error)
	// Fetch status for one specific judge by uuid or name
	FetchStatus(ctx context.Context, in *FetchStatusRequest, opts ...grpc.CallOption) (*FetchStatusReply, error)
	// Modify config for one specific judge by uuid or name
	ModifyConfig(ctx context.Context, in *ModifyConfigRequest, opts ...grpc.CallOption) (*ModifyConfigReply, error)
	// Delete one judge
	DeleteOne(ctx context.Context, in *DeleteOneRequest, opts ...grpc.CallOption) (*DeleteOneReply, error)
	// List all judges and check status
	ListAll(ctx context.Context, in *ListAllRequest, opts ...grpc.CallOption) (*ListAllReply, error)
}

type judgeClient struct {
	cc grpc.ClientConnInterface
}

func NewJudgeClient(cc grpc.ClientConnInterface) JudgeClient {
	return &judgeClient{cc}
}

func (c *judgeClient) CreateOne(ctx context.Context, in *CreateOneRequest, opts ...grpc.CallOption) (*CreateOneReply, error) {
	out := new(CreateOneReply)
	err := c.cc.Invoke(ctx, Judge_CreateOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeClient) FetchStatus(ctx context.Context, in *FetchStatusRequest, opts ...grpc.CallOption) (*FetchStatusReply, error) {
	out := new(FetchStatusReply)
	err := c.cc.Invoke(ctx, Judge_FetchStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeClient) ModifyConfig(ctx context.Context, in *ModifyConfigRequest, opts ...grpc.CallOption) (*ModifyConfigReply, error) {
	out := new(ModifyConfigReply)
	err := c.cc.Invoke(ctx, Judge_ModifyConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeClient) DeleteOne(ctx context.Context, in *DeleteOneRequest, opts ...grpc.CallOption) (*DeleteOneReply, error) {
	out := new(DeleteOneReply)
	err := c.cc.Invoke(ctx, Judge_DeleteOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeClient) ListAll(ctx context.Context, in *ListAllRequest, opts ...grpc.CallOption) (*ListAllReply, error) {
	out := new(ListAllReply)
	err := c.cc.Invoke(ctx, Judge_ListAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JudgeServer is the server API for Judge service.
// All implementations must embed UnimplementedJudgeServer
// for forward compatibility
type JudgeServer interface {
	// CreateOne create a judge with config
	CreateOne(context.Context, *CreateOneRequest) (*CreateOneReply, error)
	// Fetch status for one specific judge by uuid or name
	FetchStatus(context.Context, *FetchStatusRequest) (*FetchStatusReply, error)
	// Modify config for one specific judge by uuid or name
	ModifyConfig(context.Context, *ModifyConfigRequest) (*ModifyConfigReply, error)
	// Delete one judge
	DeleteOne(context.Context, *DeleteOneRequest) (*DeleteOneReply, error)
	// List all judges and check status
	ListAll(context.Context, *ListAllRequest) (*ListAllReply, error)
	mustEmbedUnimplementedJudgeServer()
}

// UnimplementedJudgeServer must be embedded to have forward compatible implementations.
type UnimplementedJudgeServer struct {
}

func (UnimplementedJudgeServer) CreateOne(context.Context, *CreateOneRequest) (*CreateOneReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOne not implemented")
}
func (UnimplementedJudgeServer) FetchStatus(context.Context, *FetchStatusRequest) (*FetchStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchStatus not implemented")
}
func (UnimplementedJudgeServer) ModifyConfig(context.Context, *ModifyConfigRequest) (*ModifyConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyConfig not implemented")
}
func (UnimplementedJudgeServer) DeleteOne(context.Context, *DeleteOneRequest) (*DeleteOneReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOne not implemented")
}
func (UnimplementedJudgeServer) ListAll(context.Context, *ListAllRequest) (*ListAllReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAll not implemented")
}
func (UnimplementedJudgeServer) mustEmbedUnimplementedJudgeServer() {}

// UnsafeJudgeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JudgeServer will
// result in compilation errors.
type UnsafeJudgeServer interface {
	mustEmbedUnimplementedJudgeServer()
}

func RegisterJudgeServer(s grpc.ServiceRegistrar, srv JudgeServer) {
	s.RegisterService(&Judge_ServiceDesc, srv)
}

func _Judge_CreateOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).CreateOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Judge_CreateOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).CreateOne(ctx, req.(*CreateOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judge_FetchStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).FetchStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Judge_FetchStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).FetchStatus(ctx, req.(*FetchStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judge_ModifyConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).ModifyConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Judge_ModifyConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).ModifyConfig(ctx, req.(*ModifyConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judge_DeleteOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).DeleteOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Judge_DeleteOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).DeleteOne(ctx, req.(*DeleteOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judge_ListAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).ListAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Judge_ListAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).ListAll(ctx, req.(*ListAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Judge_ServiceDesc is the grpc.ServiceDesc for Judge service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Judge_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sastoj.gojudge.judger.service.v1.Judge",
	HandlerType: (*JudgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOne",
			Handler:    _Judge_CreateOne_Handler,
		},
		{
			MethodName: "FetchStatus",
			Handler:    _Judge_FetchStatus_Handler,
		},
		{
			MethodName: "ModifyConfig",
			Handler:    _Judge_ModifyConfig_Handler,
		},
		{
			MethodName: "DeleteOne",
			Handler:    _Judge_DeleteOne_Handler,
		},
		{
			MethodName: "ListAll",
			Handler:    _Judge_ListAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/sastoj/gojudge/judger/service/v1/judge.proto",
}
