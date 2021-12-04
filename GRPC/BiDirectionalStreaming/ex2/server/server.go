package main

import (
	"ex2/calculatePb"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) CalculateMaximun(stream calculatePb.CalculateMaxSevice_CalculateMaximunServer) error {
	fmt.Println("Invoked CalculateMaximum function")
	var maximum int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil

		}
		if err != nil {
			log.Fatalf("Error reading the data:%v", err)
			return err
		}
		number := req.GetNum()
		if number > maximum {
			maximum = number
			sendErr := stream.Send(&calculatePb.CalculateMaxResponse{
				Result: maximum,
			})
			if sendErr != nil {
				log.Fatalf("Error while sending data to client:%v", err)
				return err
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error Listening to the post:", err)
	}
	s := grpc.NewServer()
	calculatePb.RegisterCalculateMaxSeviceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Error while serving: %v", err)
	}
}
