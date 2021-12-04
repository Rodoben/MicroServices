package main

import (
	"context"
	"ex1/GreetPb"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error Dialing:%v", err)
	}
	defer cc.Close()
	c := GreetPb.NewGreetServiceClient(cc)
	doBidirectionStreming(c)
}

func doBidirectionStreming(c GreetPb.GreetServiceClient) {
	fmt.Println("Doing Bi-Directional Streaming")
	// we create a stream by invoking the server
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error While creating stream:%v", err)
		return
	}
	requests := []*GreetPb.GreetEveryoneRequest{
		&GreetPb.GreetEveryoneRequest{
			Firstname: "Ronald",
		},
		&GreetPb.GreetEveryoneRequest{
			Firstname: "Alex",
		},
		&GreetPb.GreetEveryoneRequest{
			Firstname: "Lucy",
		},
		&GreetPb.GreetEveryoneRequest{
			Firstname: "Abram",
		},
		&GreetPb.GreetEveryoneRequest{
			Firstname: "Binod",
		},
		&GreetPb.GreetEveryoneRequest{
			Firstname: "Mike",
		},
	}

	waitc := make(chan struct{})
	//we send a bunch of message to the client(go routine)
	go func() {
		for _, req := range requests {
			fmt.Println("Sending Request", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()

	}()

	// we receive a bunch of message from the client (go routine)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receving: %v", err)
				break
			}
			fmt.Println("Received:", res.GetResult())
		}
		close(waitc)

	}()

	// block until everything is done
	<-waitc
}
