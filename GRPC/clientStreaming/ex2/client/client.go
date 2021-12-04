package main

import (
	"context"
	"ex2/CalcAvg"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to port 50051: %v", err)
	}
	defer cc.Close()
	c := CalcAvg.NewCalculateServiceClient(cc)
	doClientStreaming(c)

}

func doClientStreaming(c CalcAvg.CalculateServiceClient) {
	fmt.Printf("Ill do Client Streaming")
	requests := []*CalcAvg.CalcAvgRequest{
		&CalcAvg.CalcAvgRequest{
			Num: 10.3,
		},
		&CalcAvg.CalcAvgRequest{
			Num: 16.3,
		},
		&CalcAvg.CalcAvgRequest{
			Num: 1.3,
		},
		&CalcAvg.CalcAvgRequest{
			Num: 12.3,
		},
		&CalcAvg.CalcAvgRequest{
			Num: 10.3,
		},
		&CalcAvg.CalcAvgRequest{
			Num: 19.3,
		},
	}
	stream, err := c.CalcAvg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling the server function")
	}
	for _, req := range requests {
		fmt.Println("Sending Request", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error Receiving response %v", err)
	}
	fmt.Printf("Calc Avf Response:%v", res)
}
