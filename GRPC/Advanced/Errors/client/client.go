package main

import (
	"Errors/SquareRoot"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to port 50051: %v", err)
	}
	defer cc.Close()

	c := SquareRoot.NewSquareRootServiceClient(cc)

	doUnary(c, 10.23)
	doUnary(c, -20.23)
}

func doUnary(c SquareRoot.SquareRootServiceClient, n float32) {
	fmt.Println("I am clinet and will do unary call")
	res, err := c.SquareRoot(context.Background(), &SquareRoot.SquareRootRequest{Num: n})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We send a Negative number")
			} else {
				log.Fatalf("Big error calling SquareRoot")
				return
			}

		}

	}
	fmt.Println("Result of Square root %v:%v", res.GetResult())
}
