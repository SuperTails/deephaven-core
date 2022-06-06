// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: deephaven/proto/console.proto

package console

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

// ConsoleServiceClient is the client API for ConsoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsoleServiceClient interface {
	GetConsoleTypes(ctx context.Context, in *GetConsoleTypesRequest, opts ...grpc.CallOption) (*GetConsoleTypesResponse, error)
	StartConsole(ctx context.Context, in *StartConsoleRequest, opts ...grpc.CallOption) (*StartConsoleResponse, error)
	SubscribeToLogs(ctx context.Context, in *LogSubscriptionRequest, opts ...grpc.CallOption) (ConsoleService_SubscribeToLogsClient, error)
	ExecuteCommand(ctx context.Context, in *ExecuteCommandRequest, opts ...grpc.CallOption) (*ExecuteCommandResponse, error)
	CancelCommand(ctx context.Context, in *CancelCommandRequest, opts ...grpc.CallOption) (*CancelCommandResponse, error)
	BindTableToVariable(ctx context.Context, in *BindTableToVariableRequest, opts ...grpc.CallOption) (*BindTableToVariableResponse, error)
	//
	// Starts a stream for autocomplete on the current session. More than one console,
	// more than one document can be edited at a time using this, and they can separately
	// be closed as well. A given document should only be edited within one stream at a
	// time.
	AutoCompleteStream(ctx context.Context, opts ...grpc.CallOption) (ConsoleService_AutoCompleteStreamClient, error)
	//
	// Half of the browser-based (browser's can't do bidirectional streams without websockets)
	// implementation for AutoCompleteStream.
	OpenAutoCompleteStream(ctx context.Context, in *AutoCompleteRequest, opts ...grpc.CallOption) (ConsoleService_OpenAutoCompleteStreamClient, error)
	//
	// Other half of the browser-based implementation for AutoCompleteStream.
	NextAutoCompleteStream(ctx context.Context, in *AutoCompleteRequest, opts ...grpc.CallOption) (*BrowserNextResponse, error)
}

type consoleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsoleServiceClient(cc grpc.ClientConnInterface) ConsoleServiceClient {
	return &consoleServiceClient{cc}
}

func (c *consoleServiceClient) GetConsoleTypes(ctx context.Context, in *GetConsoleTypesRequest, opts ...grpc.CallOption) (*GetConsoleTypesResponse, error) {
	out := new(GetConsoleTypesResponse)
	err := c.cc.Invoke(ctx, "/io.deephaven.proto.backplane.script.grpc.ConsoleService/GetConsoleTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consoleServiceClient) StartConsole(ctx context.Context, in *StartConsoleRequest, opts ...grpc.CallOption) (*StartConsoleResponse, error) {
	out := new(StartConsoleResponse)
	err := c.cc.Invoke(ctx, "/io.deephaven.proto.backplane.script.grpc.ConsoleService/StartConsole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consoleServiceClient) SubscribeToLogs(ctx context.Context, in *LogSubscriptionRequest, opts ...grpc.CallOption) (ConsoleService_SubscribeToLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConsoleService_ServiceDesc.Streams[0], "/io.deephaven.proto.backplane.script.grpc.ConsoleService/SubscribeToLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &consoleServiceSubscribeToLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConsoleService_SubscribeToLogsClient interface {
	Recv() (*LogSubscriptionData, error)
	grpc.ClientStream
}

type consoleServiceSubscribeToLogsClient struct {
	grpc.ClientStream
}

func (x *consoleServiceSubscribeToLogsClient) Recv() (*LogSubscriptionData, error) {
	m := new(LogSubscriptionData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *consoleServiceClient) ExecuteCommand(ctx context.Context, in *ExecuteCommandRequest, opts ...grpc.CallOption) (*ExecuteCommandResponse, error) {
	out := new(ExecuteCommandResponse)
	err := c.cc.Invoke(ctx, "/io.deephaven.proto.backplane.script.grpc.ConsoleService/ExecuteCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consoleServiceClient) CancelCommand(ctx context.Context, in *CancelCommandRequest, opts ...grpc.CallOption) (*CancelCommandResponse, error) {
	out := new(CancelCommandResponse)
	err := c.cc.Invoke(ctx, "/io.deephaven.proto.backplane.script.grpc.ConsoleService/CancelCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consoleServiceClient) BindTableToVariable(ctx context.Context, in *BindTableToVariableRequest, opts ...grpc.CallOption) (*BindTableToVariableResponse, error) {
	out := new(BindTableToVariableResponse)
	err := c.cc.Invoke(ctx, "/io.deephaven.proto.backplane.script.grpc.ConsoleService/BindTableToVariable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consoleServiceClient) AutoCompleteStream(ctx context.Context, opts ...grpc.CallOption) (ConsoleService_AutoCompleteStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConsoleService_ServiceDesc.Streams[1], "/io.deephaven.proto.backplane.script.grpc.ConsoleService/AutoCompleteStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &consoleServiceAutoCompleteStreamClient{stream}
	return x, nil
}

type ConsoleService_AutoCompleteStreamClient interface {
	Send(*AutoCompleteRequest) error
	Recv() (*AutoCompleteResponse, error)
	grpc.ClientStream
}

type consoleServiceAutoCompleteStreamClient struct {
	grpc.ClientStream
}

func (x *consoleServiceAutoCompleteStreamClient) Send(m *AutoCompleteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *consoleServiceAutoCompleteStreamClient) Recv() (*AutoCompleteResponse, error) {
	m := new(AutoCompleteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *consoleServiceClient) OpenAutoCompleteStream(ctx context.Context, in *AutoCompleteRequest, opts ...grpc.CallOption) (ConsoleService_OpenAutoCompleteStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConsoleService_ServiceDesc.Streams[2], "/io.deephaven.proto.backplane.script.grpc.ConsoleService/OpenAutoCompleteStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &consoleServiceOpenAutoCompleteStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConsoleService_OpenAutoCompleteStreamClient interface {
	Recv() (*AutoCompleteResponse, error)
	grpc.ClientStream
}

type consoleServiceOpenAutoCompleteStreamClient struct {
	grpc.ClientStream
}

func (x *consoleServiceOpenAutoCompleteStreamClient) Recv() (*AutoCompleteResponse, error) {
	m := new(AutoCompleteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *consoleServiceClient) NextAutoCompleteStream(ctx context.Context, in *AutoCompleteRequest, opts ...grpc.CallOption) (*BrowserNextResponse, error) {
	out := new(BrowserNextResponse)
	err := c.cc.Invoke(ctx, "/io.deephaven.proto.backplane.script.grpc.ConsoleService/NextAutoCompleteStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsoleServiceServer is the server API for ConsoleService service.
// All implementations must embed UnimplementedConsoleServiceServer
// for forward compatibility
type ConsoleServiceServer interface {
	GetConsoleTypes(context.Context, *GetConsoleTypesRequest) (*GetConsoleTypesResponse, error)
	StartConsole(context.Context, *StartConsoleRequest) (*StartConsoleResponse, error)
	SubscribeToLogs(*LogSubscriptionRequest, ConsoleService_SubscribeToLogsServer) error
	ExecuteCommand(context.Context, *ExecuteCommandRequest) (*ExecuteCommandResponse, error)
	CancelCommand(context.Context, *CancelCommandRequest) (*CancelCommandResponse, error)
	BindTableToVariable(context.Context, *BindTableToVariableRequest) (*BindTableToVariableResponse, error)
	//
	// Starts a stream for autocomplete on the current session. More than one console,
	// more than one document can be edited at a time using this, and they can separately
	// be closed as well. A given document should only be edited within one stream at a
	// time.
	AutoCompleteStream(ConsoleService_AutoCompleteStreamServer) error
	//
	// Half of the browser-based (browser's can't do bidirectional streams without websockets)
	// implementation for AutoCompleteStream.
	OpenAutoCompleteStream(*AutoCompleteRequest, ConsoleService_OpenAutoCompleteStreamServer) error
	//
	// Other half of the browser-based implementation for AutoCompleteStream.
	NextAutoCompleteStream(context.Context, *AutoCompleteRequest) (*BrowserNextResponse, error)
	mustEmbedUnimplementedConsoleServiceServer()
}

// UnimplementedConsoleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConsoleServiceServer struct {
}

func (UnimplementedConsoleServiceServer) GetConsoleTypes(context.Context, *GetConsoleTypesRequest) (*GetConsoleTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsoleTypes not implemented")
}
func (UnimplementedConsoleServiceServer) StartConsole(context.Context, *StartConsoleRequest) (*StartConsoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartConsole not implemented")
}
func (UnimplementedConsoleServiceServer) SubscribeToLogs(*LogSubscriptionRequest, ConsoleService_SubscribeToLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToLogs not implemented")
}
func (UnimplementedConsoleServiceServer) ExecuteCommand(context.Context, *ExecuteCommandRequest) (*ExecuteCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCommand not implemented")
}
func (UnimplementedConsoleServiceServer) CancelCommand(context.Context, *CancelCommandRequest) (*CancelCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelCommand not implemented")
}
func (UnimplementedConsoleServiceServer) BindTableToVariable(context.Context, *BindTableToVariableRequest) (*BindTableToVariableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindTableToVariable not implemented")
}
func (UnimplementedConsoleServiceServer) AutoCompleteStream(ConsoleService_AutoCompleteStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AutoCompleteStream not implemented")
}
func (UnimplementedConsoleServiceServer) OpenAutoCompleteStream(*AutoCompleteRequest, ConsoleService_OpenAutoCompleteStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method OpenAutoCompleteStream not implemented")
}
func (UnimplementedConsoleServiceServer) NextAutoCompleteStream(context.Context, *AutoCompleteRequest) (*BrowserNextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextAutoCompleteStream not implemented")
}
func (UnimplementedConsoleServiceServer) mustEmbedUnimplementedConsoleServiceServer() {}

// UnsafeConsoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsoleServiceServer will
// result in compilation errors.
type UnsafeConsoleServiceServer interface {
	mustEmbedUnimplementedConsoleServiceServer()
}

func RegisterConsoleServiceServer(s grpc.ServiceRegistrar, srv ConsoleServiceServer) {
	s.RegisterService(&ConsoleService_ServiceDesc, srv)
}

func _ConsoleService_GetConsoleTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConsoleTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsoleServiceServer).GetConsoleTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.deephaven.proto.backplane.script.grpc.ConsoleService/GetConsoleTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsoleServiceServer).GetConsoleTypes(ctx, req.(*GetConsoleTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsoleService_StartConsole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartConsoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsoleServiceServer).StartConsole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.deephaven.proto.backplane.script.grpc.ConsoleService/StartConsole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsoleServiceServer).StartConsole(ctx, req.(*StartConsoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsoleService_SubscribeToLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogSubscriptionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConsoleServiceServer).SubscribeToLogs(m, &consoleServiceSubscribeToLogsServer{stream})
}

type ConsoleService_SubscribeToLogsServer interface {
	Send(*LogSubscriptionData) error
	grpc.ServerStream
}

type consoleServiceSubscribeToLogsServer struct {
	grpc.ServerStream
}

func (x *consoleServiceSubscribeToLogsServer) Send(m *LogSubscriptionData) error {
	return x.ServerStream.SendMsg(m)
}

func _ConsoleService_ExecuteCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsoleServiceServer).ExecuteCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.deephaven.proto.backplane.script.grpc.ConsoleService/ExecuteCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsoleServiceServer).ExecuteCommand(ctx, req.(*ExecuteCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsoleService_CancelCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsoleServiceServer).CancelCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.deephaven.proto.backplane.script.grpc.ConsoleService/CancelCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsoleServiceServer).CancelCommand(ctx, req.(*CancelCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsoleService_BindTableToVariable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindTableToVariableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsoleServiceServer).BindTableToVariable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.deephaven.proto.backplane.script.grpc.ConsoleService/BindTableToVariable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsoleServiceServer).BindTableToVariable(ctx, req.(*BindTableToVariableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsoleService_AutoCompleteStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConsoleServiceServer).AutoCompleteStream(&consoleServiceAutoCompleteStreamServer{stream})
}

type ConsoleService_AutoCompleteStreamServer interface {
	Send(*AutoCompleteResponse) error
	Recv() (*AutoCompleteRequest, error)
	grpc.ServerStream
}

type consoleServiceAutoCompleteStreamServer struct {
	grpc.ServerStream
}

func (x *consoleServiceAutoCompleteStreamServer) Send(m *AutoCompleteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *consoleServiceAutoCompleteStreamServer) Recv() (*AutoCompleteRequest, error) {
	m := new(AutoCompleteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ConsoleService_OpenAutoCompleteStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AutoCompleteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConsoleServiceServer).OpenAutoCompleteStream(m, &consoleServiceOpenAutoCompleteStreamServer{stream})
}

type ConsoleService_OpenAutoCompleteStreamServer interface {
	Send(*AutoCompleteResponse) error
	grpc.ServerStream
}

type consoleServiceOpenAutoCompleteStreamServer struct {
	grpc.ServerStream
}

func (x *consoleServiceOpenAutoCompleteStreamServer) Send(m *AutoCompleteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ConsoleService_NextAutoCompleteStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AutoCompleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsoleServiceServer).NextAutoCompleteStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.deephaven.proto.backplane.script.grpc.ConsoleService/NextAutoCompleteStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsoleServiceServer).NextAutoCompleteStream(ctx, req.(*AutoCompleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConsoleService_ServiceDesc is the grpc.ServiceDesc for ConsoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConsoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.deephaven.proto.backplane.script.grpc.ConsoleService",
	HandlerType: (*ConsoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConsoleTypes",
			Handler:    _ConsoleService_GetConsoleTypes_Handler,
		},
		{
			MethodName: "StartConsole",
			Handler:    _ConsoleService_StartConsole_Handler,
		},
		{
			MethodName: "ExecuteCommand",
			Handler:    _ConsoleService_ExecuteCommand_Handler,
		},
		{
			MethodName: "CancelCommand",
			Handler:    _ConsoleService_CancelCommand_Handler,
		},
		{
			MethodName: "BindTableToVariable",
			Handler:    _ConsoleService_BindTableToVariable_Handler,
		},
		{
			MethodName: "NextAutoCompleteStream",
			Handler:    _ConsoleService_NextAutoCompleteStream_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToLogs",
			Handler:       _ConsoleService_SubscribeToLogs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AutoCompleteStream",
			Handler:       _ConsoleService_AutoCompleteStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "OpenAutoCompleteStream",
			Handler:       _ConsoleService_OpenAutoCompleteStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "deephaven/proto/console.proto",
}
