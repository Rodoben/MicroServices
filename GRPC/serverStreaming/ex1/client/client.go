package main

import (
	"context"
	"ex2/greetpb"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	doServerStreaming(c)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("I'll send data to perform server streaming")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ronald",
			LastName:  "Benjamin",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetmanyTIme RPC: %v", err)
	}
	for {
		fmt.Println("Inside for loop")
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading: %v", err)
		}
		log.Printf("Response from Greeting: %v", msg.GetResult())
	}
}
