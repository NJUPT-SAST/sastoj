// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: sastoj/gojudge/judger/gojudge/v1/gojudge.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Executor_Exec_FullMethodName       = "/pb.Executor/Exec"
	Executor_ExecStream_FullMethodName = "/pb.Executor/ExecStream"
	Executor_FileList_FullMethodName   = "/pb.Executor/FileList"
	Executor_FileGet_FullMethodName    = "/pb.Executor/FileGet"
	Executor_FileAdd_FullMethodName    = "/pb.Executor/FileAdd"
	Executor_FileDelete_FullMethodName = "/pb.Executor/FileDelete"
)

// ExecutorClient is the client API for Executor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecutorClient interface {
	// Exec defines unary RPC to run a program with resource limitations
	Exec(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// ExecStream defines streaming RPC to run a program with real-time input &
	// output. The first request must be execRequest and the following request
	// must be execInput. The last response must be execResponse and the others
	// are execOutput. TTY attribute will create single pty for the program thus
	// stdout & stderr should have same name
	ExecStream(ctx context.Context, opts ...grpc.CallOption) (Executor_ExecStreamClient, error)
	// FileList lists all files available in the file store
	FileList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FileListType, error)
	// FileGet download the file from the file store
	FileGet(ctx context.Context, in *FileID, opts ...grpc.CallOption) (*FileContent, error)
	// FileAdd create a file into the file store
	FileAdd(ctx context.Context, in *FileContent, opts ...grpc.CallOption) (*FileID, error)
	// FileDelete deletes a file from the file store
	FileDelete(ctx context.Context, in *FileID, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type executorClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutorClient(cc grpc.ClientConnInterface) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) Exec(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Executor_Exec_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) ExecStream(ctx context.Context, opts ...grpc.CallOption) (Executor_ExecStreamClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Executor_ServiceDesc.Streams[0], Executor_ExecStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &executorExecStreamClient{ClientStream: stream}
	return x, nil
}

type Executor_ExecStreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type executorExecStreamClient struct {
	grpc.ClientStream
}

func (x *executorExecStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *executorExecStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *executorClient) FileList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FileListType, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileListType)
	err := c.cc.Invoke(ctx, Executor_FileList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) FileGet(ctx context.Context, in *FileID, opts ...grpc.CallOption) (*FileContent, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileContent)
	err := c.cc.Invoke(ctx, Executor_FileGet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) FileAdd(ctx context.Context, in *FileContent, opts ...grpc.CallOption) (*FileID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileID)
	err := c.cc.Invoke(ctx, Executor_FileAdd_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) FileDelete(ctx context.Context, in *FileID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Executor_FileDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutorServer is the server API for Executor service.
// All implementations must embed UnimplementedExecutorServer
// for forward compatibility
type ExecutorServer interface {
	// Exec defines unary RPC to run a program with resource limitations
	Exec(context.Context, *Request) (*Response, error)
	// ExecStream defines streaming RPC to run a program with real-time input &
	// output. The first request must be execRequest and the following request
	// must be execInput. The last response must be execResponse and the others
	// are execOutput. TTY attribute will create single pty for the program thus
	// stdout & stderr should have same name
	ExecStream(Executor_ExecStreamServer) error
	// FileList lists all files available in the file store
	FileList(context.Context, *emptypb.Empty) (*FileListType, error)
	// FileGet download the file from the file store
	FileGet(context.Context, *FileID) (*FileContent, error)
	// FileAdd create a file into the file store
	FileAdd(context.Context, *FileContent) (*FileID, error)
	// FileDelete deletes a file from the file store
	FileDelete(context.Context, *FileID) (*emptypb.Empty, error)
	mustEmbedUnimplementedExecutorServer()
}

// UnimplementedExecutorServer must be embedded to have forward compatible implementations.
type UnimplementedExecutorServer struct {
}

func (UnimplementedExecutorServer) Exec(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}
func (UnimplementedExecutorServer) ExecStream(Executor_ExecStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecStream not implemented")
}
func (UnimplementedExecutorServer) FileList(context.Context, *emptypb.Empty) (*FileListType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileList not implemented")
}
func (UnimplementedExecutorServer) FileGet(context.Context, *FileID) (*FileContent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileGet not implemented")
}
func (UnimplementedExecutorServer) FileAdd(context.Context, *FileContent) (*FileID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileAdd not implemented")
}
func (UnimplementedExecutorServer) FileDelete(context.Context, *FileID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileDelete not implemented")
}
func (UnimplementedExecutorServer) mustEmbedUnimplementedExecutorServer() {}

// UnsafeExecutorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecutorServer will
// result in compilation errors.
type UnsafeExecutorServer interface {
	mustEmbedUnimplementedExecutorServer()
}

func RegisterExecutorServer(s grpc.ServiceRegistrar, srv ExecutorServer) {
	s.RegisterService(&Executor_ServiceDesc, srv)
}

func _Executor_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Executor_Exec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).Exec(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_ExecStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExecutorServer).ExecStream(&executorExecStreamServer{ServerStream: stream})
}

type Executor_ExecStreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type executorExecStreamServer struct {
	grpc.ServerStream
}

func (x *executorExecStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *executorExecStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Executor_FileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).FileList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Executor_FileList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).FileList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_FileGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).FileGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Executor_FileGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).FileGet(ctx, req.(*FileID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_FileAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileContent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).FileAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Executor_FileAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).FileAdd(ctx, req.(*FileContent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_FileDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).FileDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Executor_FileDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).FileDelete(ctx, req.(*FileID))
	}
	return interceptor(ctx, in, info, handler)
}

// Executor_ServiceDesc is the grpc.ServiceDesc for Executor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Executor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exec",
			Handler:    _Executor_Exec_Handler,
		},
		{
			MethodName: "FileList",
			Handler:    _Executor_FileList_Handler,
		},
		{
			MethodName: "FileGet",
			Handler:    _Executor_FileGet_Handler,
		},
		{
			MethodName: "FileAdd",
			Handler:    _Executor_FileAdd_Handler,
		},
		{
			MethodName: "FileDelete",
			Handler:    _Executor_FileDelete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecStream",
			Handler:       _Executor_ExecStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "sastoj/gojudge/judger/gojudge/v1/gojudge.proto",
}
