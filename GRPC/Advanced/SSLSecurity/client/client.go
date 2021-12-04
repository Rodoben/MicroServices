package main

import (
	"SSLSecurity/MultiplicationPb"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error Dialing to port:%v", err)
	}

	defer cc.Close()

	c := MultiplicationPb.NewMultiplicationServiceClient(cc)
	doServerStreaming(c, 1*time.Second)
	doServerStreaming(c, 11*time.Second)
}

func doServerStreaming(c MultiplicationPb.MultiplicationServiceClient, timeout time.Duration) {
	var n int32 = 56
	fmt.Println("I'll send data to perform server streaming")
	req := &MultiplicationPb.MultiplicationRequest{
		Num: n,
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resStream, err := c.MultiplicationTableWithDeadline(ctx, req)
	if err != nil {
		log.Fatalf("Error While Calling the server code")
	}
	respErr, ok := status.FromError(err)
	if ok {
		fmt.Println("Response Message:", respErr.Message())
		fmt.Println("Response Code:", respErr.Code())
		if respErr.Code() == codes.InvalidArgument {
			fmt.Println("We send negative number")
		} else {
			for i := 1; ; i++ {

				msg, err := resStream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Printf("Error Reading the response %v", err)
					break
				}
				fmt.Println("Response:", i, "*", n, "=", msg.GetResult())
			}
		}
	}

}
