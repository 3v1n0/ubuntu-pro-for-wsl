// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: agentapi.proto

package agentapi

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
	UI_ApplyProToken_FullMethodName     = "/agentapi.UI/ApplyProToken"
	UI_Ping_FullMethodName              = "/agentapi.UI/Ping"
	UI_SubscpriptionInfo_FullMethodName = "/agentapi.UI/SubscpriptionInfo"
)

// UIClient is the client API for UI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UIClient interface {
	ApplyProToken(ctx context.Context, in *ProAttachInfo, opts ...grpc.CallOption) (*Empty, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	SubscpriptionInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SubscriptionInfo, error)
}

type uIClient struct {
	cc grpc.ClientConnInterface
}

func NewUIClient(cc grpc.ClientConnInterface) UIClient {
	return &uIClient{cc}
}

func (c *uIClient) ApplyProToken(ctx context.Context, in *ProAttachInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UI_ApplyProToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uIClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UI_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uIClient) SubscpriptionInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SubscriptionInfo, error) {
	out := new(SubscriptionInfo)
	err := c.cc.Invoke(ctx, UI_SubscpriptionInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UIServer is the server API for UI service.
// All implementations must embed UnimplementedUIServer
// for forward compatibility
type UIServer interface {
	ApplyProToken(context.Context, *ProAttachInfo) (*Empty, error)
	Ping(context.Context, *Empty) (*Empty, error)
	SubscpriptionInfo(context.Context, *Empty) (*SubscriptionInfo, error)
	mustEmbedUnimplementedUIServer()
}

// UnimplementedUIServer must be embedded to have forward compatible implementations.
type UnimplementedUIServer struct {
}

func (UnimplementedUIServer) ApplyProToken(context.Context, *ProAttachInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyProToken not implemented")
}
func (UnimplementedUIServer) Ping(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedUIServer) SubscpriptionInfo(context.Context, *Empty) (*SubscriptionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscpriptionInfo not implemented")
}
func (UnimplementedUIServer) mustEmbedUnimplementedUIServer() {}

// UnsafeUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UIServer will
// result in compilation errors.
type UnsafeUIServer interface {
	mustEmbedUnimplementedUIServer()
}

func RegisterUIServer(s grpc.ServiceRegistrar, srv UIServer) {
	s.RegisterService(&UI_ServiceDesc, srv)
}

func _UI_ApplyProToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProAttachInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).ApplyProToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UI_ApplyProToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).ApplyProToken(ctx, req.(*ProAttachInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _UI_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UI_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UI_SubscpriptionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).SubscpriptionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UI_SubscpriptionInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).SubscpriptionInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// UI_ServiceDesc is the grpc.ServiceDesc for UI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agentapi.UI",
	HandlerType: (*UIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyProToken",
			Handler:    _UI_ApplyProToken_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _UI_Ping_Handler,
		},
		{
			MethodName: "SubscpriptionInfo",
			Handler:    _UI_SubscpriptionInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agentapi.proto",
}

const (
	WSLInstance_Connected_FullMethodName = "/agentapi.WSLInstance/Connected"
)

// WSLInstanceClient is the client API for WSLInstance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WSLInstanceClient interface {
	Connected(ctx context.Context, opts ...grpc.CallOption) (WSLInstance_ConnectedClient, error)
}

type wSLInstanceClient struct {
	cc grpc.ClientConnInterface
}

func NewWSLInstanceClient(cc grpc.ClientConnInterface) WSLInstanceClient {
	return &wSLInstanceClient{cc}
}

func (c *wSLInstanceClient) Connected(ctx context.Context, opts ...grpc.CallOption) (WSLInstance_ConnectedClient, error) {
	stream, err := c.cc.NewStream(ctx, &WSLInstance_ServiceDesc.Streams[0], WSLInstance_Connected_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &wSLInstanceConnectedClient{stream}
	return x, nil
}

type WSLInstance_ConnectedClient interface {
	Send(*DistroInfo) error
	Recv() (*Port, error)
	grpc.ClientStream
}

type wSLInstanceConnectedClient struct {
	grpc.ClientStream
}

func (x *wSLInstanceConnectedClient) Send(m *DistroInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *wSLInstanceConnectedClient) Recv() (*Port, error) {
	m := new(Port)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WSLInstanceServer is the server API for WSLInstance service.
// All implementations must embed UnimplementedWSLInstanceServer
// for forward compatibility
type WSLInstanceServer interface {
	Connected(WSLInstance_ConnectedServer) error
	mustEmbedUnimplementedWSLInstanceServer()
}

// UnimplementedWSLInstanceServer must be embedded to have forward compatible implementations.
type UnimplementedWSLInstanceServer struct {
}

func (UnimplementedWSLInstanceServer) Connected(WSLInstance_ConnectedServer) error {
	return status.Errorf(codes.Unimplemented, "method Connected not implemented")
}
func (UnimplementedWSLInstanceServer) mustEmbedUnimplementedWSLInstanceServer() {}

// UnsafeWSLInstanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WSLInstanceServer will
// result in compilation errors.
type UnsafeWSLInstanceServer interface {
	mustEmbedUnimplementedWSLInstanceServer()
}

func RegisterWSLInstanceServer(s grpc.ServiceRegistrar, srv WSLInstanceServer) {
	s.RegisterService(&WSLInstance_ServiceDesc, srv)
}

func _WSLInstance_Connected_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WSLInstanceServer).Connected(&wSLInstanceConnectedServer{stream})
}

type WSLInstance_ConnectedServer interface {
	Send(*Port) error
	Recv() (*DistroInfo, error)
	grpc.ServerStream
}

type wSLInstanceConnectedServer struct {
	grpc.ServerStream
}

func (x *wSLInstanceConnectedServer) Send(m *Port) error {
	return x.ServerStream.SendMsg(m)
}

func (x *wSLInstanceConnectedServer) Recv() (*DistroInfo, error) {
	m := new(DistroInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WSLInstance_ServiceDesc is the grpc.ServiceDesc for WSLInstance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WSLInstance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agentapi.WSLInstance",
	HandlerType: (*WSLInstanceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connected",
			Handler:       _WSLInstance_Connected_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "agentapi.proto",
}
