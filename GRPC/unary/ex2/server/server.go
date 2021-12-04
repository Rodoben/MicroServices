package main

import (
	"context"
	"ex2/sumpb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	fmt.Println("Invoked With", req)
	num1 := req.GetNumbers().GetNum1()
	num2 := req.GetNumbers().GetNum2()

	result := num1 + num2
	res := sumpb.SumResponse{
		Result: result,
	}
	return &res, nil
}
func main() {

	fmt.Println("Hi i am Unary server and will add numbers")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error in connection %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic("failed to serve:")
	}
}
