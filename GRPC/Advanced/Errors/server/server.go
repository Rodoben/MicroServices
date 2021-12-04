package main

import (
	"Errors/SquareRoot"
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) SquareRoot(ctx context.Context, req *SquareRoot.SquareRootRequest) (*SquareRoot.SquareRootResponse, error) {
	var num float32 = req.GetNum()

	fmt.Println(num)
	if num < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Received negative number")
	}

	return &SquareRoot.SquareRootResponse{
		Result: float32(math.Sqrt(float64(num))),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error Listening to port 50051: %v", err)
	}

	s := grpc.NewServer()
	SquareRoot.RegisterSquareRootServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Error While Serving on port", err)
	}

}
