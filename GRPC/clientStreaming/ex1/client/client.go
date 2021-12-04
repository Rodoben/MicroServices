package main

import (
	"context"
	"ex1/longGreet"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while dialing: %v", err)
	}
	defer cc.Close()
	c := longGreet.NewLongGreetServiceClient(cc)
	doClientStreaming(c)
}

func doClientStreaming(c longGreet.LongGreetServiceClient) {
	fmt.Println("I am doing Client Stream")
	reqs := []*longGreet.LongGreetRequest{
		&longGreet.LongGreetRequest{
			FirstName: "Ronald",
		},
		&longGreet.LongGreetRequest{
			FirstName: "Mansi",
		},
		&longGreet.LongGreetRequest{
			FirstName: "Lucy",
		},
		&longGreet.LongGreetRequest{
			FirstName: "Stephane",
		},
		&longGreet.LongGreetRequest{
			FirstName: "Komal",
		},
		&longGreet.LongGreetRequest{
			FirstName: "Marsh",
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet %v", err)
	}
	for _, r := range reqs {
		fmt.Println("Sending req: \n", r)
		stream.Send(r)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet %v", err)
	}
	fmt.Printf("LongGreet Response:%v", res)
}
