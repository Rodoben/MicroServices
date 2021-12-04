package main

import (
	"context"
	"ex2/sumpb"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Coudnt connect to server %v", err)
	}
	defer cc.Close()
	c := sumpb.NewSumServiceClient(cc)
	doUnary(c)
}

func doUnary(c sumpb.SumServiceClient) {
	fmt.Println("Starting to do unary RPC")
	req := &sumpb.SumRequest{
		Numbers: &sumpb.Numbers{
			Num1: 10.00,
			Num2: 20.00,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC %v", err)
	}
	log.Printf("Response from Greet:%v", res.Result)

}
