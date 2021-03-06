// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package longGreet

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

// LongGreetServiceClient is the client API for LongGreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LongGreetServiceClient interface {
	LongGreet(ctx context.Context, opts ...grpc.CallOption) (LongGreetService_LongGreetClient, error)
}

type longGreetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLongGreetServiceClient(cc grpc.ClientConnInterface) LongGreetServiceClient {
	return &longGreetServiceClient{cc}
}

func (c *longGreetServiceClient) LongGreet(ctx context.Context, opts ...grpc.CallOption) (LongGreetService_LongGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &LongGreetService_ServiceDesc.Streams[0], "/longGreet.longGreetService/LongGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &longGreetServiceLongGreetClient{stream}
	return x, nil
}

type LongGreetService_LongGreetClient interface {
	Send(*LongGreetRequest) error
	CloseAndRecv() (*LongGreetResponse, error)
	grpc.ClientStream
}

type longGreetServiceLongGreetClient struct {
	grpc.ClientStream
}

func (x *longGreetServiceLongGreetClient) Send(m *LongGreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *longGreetServiceLongGreetClient) CloseAndRecv() (*LongGreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LongGreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LongGreetServiceServer is the server API for LongGreetService service.
// All implementations must embed UnimplementedLongGreetServiceServer
// for forward compatibility
type LongGreetServiceServer interface {
	LongGreet(LongGreetService_LongGreetServer) error
	//mustEmbedUnimplementedLongGreetServiceServer()
}

// UnimplementedLongGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLongGreetServiceServer struct {
}

func (UnimplementedLongGreetServiceServer) LongGreet(LongGreetService_LongGreetServer) error {
	return status.Errorf(codes.Unimplemented, "method LongGreet not implemented")
}
func (UnimplementedLongGreetServiceServer) mustEmbedUnimplementedLongGreetServiceServer() {}

// UnsafeLongGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LongGreetServiceServer will
// result in compilation errors.
type UnsafeLongGreetServiceServer interface {
	mustEmbedUnimplementedLongGreetServiceServer()
}

func RegisterLongGreetServiceServer(s grpc.ServiceRegistrar, srv LongGreetServiceServer) {
	s.RegisterService(&LongGreetService_ServiceDesc, srv)
}

func _LongGreetService_LongGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LongGreetServiceServer).LongGreet(&longGreetServiceLongGreetServer{stream})
}

type LongGreetService_LongGreetServer interface {
	SendAndClose(*LongGreetResponse) error
	Recv() (*LongGreetRequest, error)
	grpc.ServerStream
}

type longGreetServiceLongGreetServer struct {
	grpc.ServerStream
}

func (x *longGreetServiceLongGreetServer) SendAndClose(m *LongGreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *longGreetServiceLongGreetServer) Recv() (*LongGreetRequest, error) {
	m := new(LongGreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LongGreetService_ServiceDesc is the grpc.ServiceDesc for LongGreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LongGreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "longGreet.longGreetService",
	HandlerType: (*LongGreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LongGreet",
			Handler:       _LongGreetService_LongGreet_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "longgreet.proto",
}
