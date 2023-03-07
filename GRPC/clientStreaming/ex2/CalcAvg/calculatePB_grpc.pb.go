// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package CalcAvg

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

// CalculateServiceClient is the client API for CalculateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculateServiceClient interface {
	CalcAvg(ctx context.Context, opts ...grpc.CallOption) (CalculateService_CalcAvgClient, error)
}

type calculateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateServiceClient(cc grpc.ClientConnInterface) CalculateServiceClient {
	return &calculateServiceClient{cc}
}

func (c *calculateServiceClient) CalcAvg(ctx context.Context, opts ...grpc.CallOption) (CalculateService_CalcAvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculateService_ServiceDesc.Streams[0], "/calcAvg.CalculateService/CalcAvg", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceCalcAvgClient{stream}
	return x, nil
}

type CalculateService_CalcAvgClient interface {
	Send(*CalcAvgRequest) error
	CloseAndRecv() (*CalcAvgResponse, error)
	grpc.ClientStream
}

type calculateServiceCalcAvgClient struct {
	grpc.ClientStream
}

func (x *calculateServiceCalcAvgClient) Send(m *CalcAvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceCalcAvgClient) CloseAndRecv() (*CalcAvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CalcAvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculateServiceServer is the server API for CalculateService service.
// All implementations must embed UnimplementedCalculateServiceServer
// for forward compatibility
type CalculateServiceServer interface {
	CalcAvg(CalculateService_CalcAvgServer) error
	//mustEmbedUnimplementedCalculateServiceServer()
}

// UnimplementedCalculateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculateServiceServer struct {
}

func (UnimplementedCalculateServiceServer) CalcAvg(CalculateService_CalcAvgServer) error {
	return status.Errorf(codes.Unimplemented, "method CalcAvg not implemented")
}
func (UnimplementedCalculateServiceServer) mustEmbedUnimplementedCalculateServiceServer() {}

// UnsafeCalculateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculateServiceServer will
// result in compilation errors.
type UnsafeCalculateServiceServer interface {
	mustEmbedUnimplementedCalculateServiceServer()
}

func RegisterCalculateServiceServer(s grpc.ServiceRegistrar, srv CalculateServiceServer) {
	s.RegisterService(&CalculateService_ServiceDesc, srv)
}

func _CalculateService_CalcAvg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).CalcAvg(&calculateServiceCalcAvgServer{stream})
}

type CalculateService_CalcAvgServer interface {
	SendAndClose(*CalcAvgResponse) error
	Recv() (*CalcAvgRequest, error)
	grpc.ServerStream
}

type calculateServiceCalcAvgServer struct {
	grpc.ServerStream
}

func (x *calculateServiceCalcAvgServer) SendAndClose(m *CalcAvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceCalcAvgServer) Recv() (*CalcAvgRequest, error) {
	m := new(CalcAvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculateService_ServiceDesc is the grpc.ServiceDesc for CalculateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calcAvg.CalculateService",
	HandlerType: (*CalculateServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CalcAvg",
			Handler:       _CalculateService_CalcAvg_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculatePB.proto",
}