package main

import (
	"ex2/CalcAvg"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) CalcAvg(stream CalcAvg.CalculateService_CalcAvgServer) error {
	var c float32 = 0.0
	var sum float32 = 0.0
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			avg := sum / c
			return stream.SendAndClose(&CalcAvg.CalcAvgResponse{
				Result: avg,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading %v", err)
		}
		num := msg.GetNum()
		c++
		sum += num

	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error Listening on port 50051: %v", err)
	}
	s := grpc.NewServer()
	CalcAvg.RegisterCalculateServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Error while Serviing: %v", err)
	}
}
