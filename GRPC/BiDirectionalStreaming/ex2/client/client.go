package main

import (
	"context"
	"ex2/calculatePb"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to port: %v", err)
	}
	defer cc.Close()
	c := calculatePb.NewCalculateMaxSeviceClient(cc)
	doBiDirectionalStreaming(c)
}

func doBiDirectionalStreaming(c calculatePb.CalculateMaxSeviceClient) {
	fmt.Print("Hi i'll do bidirectional straming!!!")
	stream, err := c.CalculateMaximun(context.Background())
	if err != nil {
		log.Fatalf("Error while streaming:%v", err)
	}
	waitc := make(chan struct{})

	go func() {
		number := []int32{4, 3, 2, 5, 9, 4, 32, 1, 1, 5}
		for _, r := range number {
			fmt.Println("Sending Number:", r)
			stream.Send(&calculatePb.CalculateMaxRequest{
				Num: r,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Problem while reading server stream: %v", err)
				break
			}
			max := res.GetResult()
			fmt.Println("Response: NewMaximux = ", max)
		}
		close(waitc)
	}()
	<-waitc
}
