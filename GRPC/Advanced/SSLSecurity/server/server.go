package main

import (
	"SSLSecurity/MultiplicationPb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) MultiplicationTableWithDeadline(req *MultiplicationPb.MultiplicationRequest, stream MultiplicationPb.MultiplicationService_MultiplicationTableWithDeadlineServer) error {

	// fmt.Println("Invoked with Deadline")
	// for i := 0; i < 3; i++ {
	// 	if ctx.Err() == context.DeadlineExceeded {
	// 		fmt.Println("The Client cancelled the request!")
	// 		return status.Error(codes.Canceled, "The client cancelled the request")
	// 	}
	// 	time.Sleep(1 * time.Second)
	// }

	num := req.GetNum()

	if num < 0 {
		return status.Errorf(codes.InvalidArgument, "Received Negative Number")
	}
	var i int32
	for i = 1; i <= 10; i++ {
		result := i * num
		res := &MultiplicationPb.MultiplicationResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error listening on the port 50051: %v", err)
	}
	s := grpc.NewServer()
	MultiplicationPb.RegisterMultiplicationServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Error while serving:", err)
	}
}
