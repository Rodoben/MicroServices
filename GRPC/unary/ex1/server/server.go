package main

import (
	"context"
	"fmt"
	"net"

	"ex1/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("invvoked with", req)
	first := req.GetGreeting().GetFirstName()
	last := req.GetGreeting().GetLastName()
	result := "Hello " + first + " " + last
	res := greetpb.GreetResponse{
		Result: result,
	}
	return &res, nil
}

func main() {
	fmt.Println("Hi!!! I am Unary Grpc Server")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic("failed to serve:")
	}
}
