package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"ex2/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Println("invoked with", req)
	firstname := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstname + " number" + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)

	}
	return nil
}

func main() {
	fmt.Println("Hi!!! I am Server Streaming Grpc Server")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("coudnt coonect %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic("failed to serve:")
	}
}
