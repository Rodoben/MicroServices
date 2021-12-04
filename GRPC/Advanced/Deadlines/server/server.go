package main

import (
	"Deadlines/CalculateSI"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) CalculateSiAndTotAmount(ctx context.Context, req *CalculateSI.CalculateSIRequest) (*CalculateSI.CalculateSIResponse, error) {
	p := req.GetPrinciple()
	r := req.GetRate()
	t := req.GetTime()
	if p < 0 || r < 0 || t < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Number cannot be less than zero")
	}
	si := (p * r * t) / 100
	amount := p + si
	return &CalculateSI.CalculateSIResponse{Si: si, Tot: amount}, nil
}
func (*server) CalculateSiAndTotAmountWithDeadline(ctx context.Context, req *CalculateSI.CalculateSIRequest) (*CalculateSI.CalculateSIResponse, error) {

	fmt.Println("Invoked with Deadline")
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("The Client cancelled the request!")
			return nil, status.Error(codes.Canceled, "The client cancelled the request")
		}
		time.Sleep(1 * time.Second)
	}

	p := req.GetPrinciple()
	r := req.GetRate()
	t := req.GetTime()
	if p < 0 || r < 0 || t < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Number cannot be less than zero")
	}
	si := (p * r * t) / 100
	amount := p + si
	return &CalculateSI.CalculateSIResponse{Si: si, Tot: amount}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error Listening to port:50051: %v", err)
	}
	s := grpc.NewServer()
	CalculateSI.RegisterCalculateSIServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Error while serving:%v", err)
	}
}
