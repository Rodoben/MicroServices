package main

import (
	"ex1/longGreet"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) LongGreet(stream longGreet.LongGreetService_LongGreetServer) error {
	fmt.Println("Long Greet function invoked")
	result := "Hello "
	for {
		msg, err := stream.Recv()
		fmt.Println(msg)
		if err == io.EOF {
			// read client stream finished
			return stream.SendAndClose(&longGreet.LongGreetResponse{
				Result: result,
			})

		}
		if err != nil {
			log.Fatalf("Error while reading %v", err)
		}
		firstname := msg.GetFirstName()
		result += "Hello " + firstname + "! "
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Coudnt Listen on Port: %v", err)
	}
	s := grpc.NewServer()
	longGreet.RegisterLongGreetServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Error While Serving")
	}
}
