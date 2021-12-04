package main

import (
	"context"
	"ex2/PnDec"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("Hi I am Client!!")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Coudn't connect: %v", err)
	}
	defer cc.Close()
	c := PnDec.NewPnDecServiceClient(cc)

	doServerStreaming(c)
}

func doServerStreaming(c PnDec.PnDecServiceClient) {
	fmt.Println("I am Client and Perform PrimeNumberDecomposition")
	req := &PnDec.PnDecRequest{
		Num: 142,
	}
	resStream, err := c.PnDec(context.Background(), req)
	if err != nil {
		log.Fatalf("Errror while calling Server PNDEC %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error Reading: %v", err)
		}
		log.Printf("Response From PrimeNumberDecomposition: %v", msg.GetResult())
	}
}
