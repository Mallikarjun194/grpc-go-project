// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: maxsum.proto

package proto

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

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SumServiceClient interface {
	MaxService(ctx context.Context, opts ...grpc.CallOption) (SumService_MaxServiceClient, error)
}

type sumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSumServiceClient(cc grpc.ClientConnInterface) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) MaxService(ctx context.Context, opts ...grpc.CallOption) (SumService_MaxServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &SumService_ServiceDesc.Streams[0], "/maxsum.sumService/MaxService", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServiceMaxServiceClient{stream}
	return x, nil
}

type SumService_MaxServiceClient interface {
	Send(*Request) error
	Recv() (*Resp, error)
	grpc.ClientStream
}

type sumServiceMaxServiceClient struct {
	grpc.ClientStream
}

func (x *sumServiceMaxServiceClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sumServiceMaxServiceClient) Recv() (*Resp, error) {
	m := new(Resp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SumServiceServer is the server API for SumService service.
// All implementations must embed UnimplementedSumServiceServer
// for forward compatibility
type SumServiceServer interface {
	MaxService(SumService_MaxServiceServer) error
	mustEmbedUnimplementedSumServiceServer()
}

// UnimplementedSumServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSumServiceServer struct {
}

func (UnimplementedSumServiceServer) MaxService(SumService_MaxServiceServer) error {
	return status.Errorf(codes.Unimplemented, "method MaxService not implemented")
}
func (UnimplementedSumServiceServer) mustEmbedUnimplementedSumServiceServer() {}

// UnsafeSumServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SumServiceServer will
// result in compilation errors.
type UnsafeSumServiceServer interface {
	mustEmbedUnimplementedSumServiceServer()
}

func RegisterSumServiceServer(s grpc.ServiceRegistrar, srv SumServiceServer) {
	s.RegisterService(&SumService_ServiceDesc, srv)
}

func _SumService_MaxService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SumServiceServer).MaxService(&sumServiceMaxServiceServer{stream})
}

type SumService_MaxServiceServer interface {
	Send(*Resp) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type sumServiceMaxServiceServer struct {
	grpc.ServerStream
}

func (x *sumServiceMaxServiceServer) Send(m *Resp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumServiceMaxServiceServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SumService_ServiceDesc is the grpc.ServiceDesc for SumService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SumService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "maxsum.sumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MaxService",
			Handler:       _SumService_MaxService_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "maxsum.proto",
}
