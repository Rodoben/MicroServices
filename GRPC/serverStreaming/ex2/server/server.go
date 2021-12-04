package main

import (
	"ex2/PnDec"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) PnDec(req *PnDec.PnDecRequest, stream PnDec.PnDecService_PnDecServer) error {
	num := req.GetNum()
	var k int32 = 2
	for num > 1 {
		if num%k == 0 {
			stream.Send(&PnDec.PnDecResponse{
				Result: k,
			})
			time.Sleep(1000 * time.Millisecond)
			num = num / k
		} else {
			k++
			fmt.Println("Divisor increased to", k)
		}

	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Coudntconnect %v", err)
	}
	s := grpc.NewServer()
	PnDec.RegisterPnDecServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving: %v", err)
	}
}
