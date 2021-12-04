package main

import (
	"ex1/GreetPb"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetEveryone(stream GreetPb.GreetService_GreetEveryoneServer) error {
	fmt.Println("invoked GreetEveryone")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream:%v", err)
			return err
		}
		firstname := req.GetFirstname()
		result := "Hello" + firstname + "! "
		sendErr := stream.Send(&GreetPb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client %v", err)
			return err
		}
	}

}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error Listening:%v", err)
	}
	s := grpc.NewServer()
	GreetPb.RegisterGreetServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Error Serving the Connection:%v", err)
	}
}
