package main

import (
	"Deadlines/CalculateSI"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while dialing to the port:50051: %v", err)
	}

	defer cc.Close()
	c := CalculateSI.NewCalculateSIServiceClient(cc)
	doUnarySICalc(c)
	doUnarySICalcWithDeadline(c, 5*time.Second) // should complete
	doUnarySICalcWithDeadline(c, 1*time.Second) // should timeout

}

func doUnarySICalc(c CalculateSI.CalculateSIServiceClient) {
	fmt.Println("Hi i am client and will do unary operation on simple intrest")

	res, err := c.CalculateSiAndTotAmount(context.Background(), &CalculateSI.CalculateSIRequest{Principle: -1000.0, Rate: 10.5, Time: 13})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println("Response Message:", respErr.Message())
			fmt.Println("Response Code:", respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("Sent a negative number")
			} else {
				log.Fatalf("Big error sending negative number")
				return
			}
		}
	}
	fmt.Println("SI : ", res.GetSi(), "Amount : ", res.GetTot())
}

func doUnarySICalcWithDeadline(c CalculateSI.CalculateSIServiceClient, timeout time.Duration) {
	fmt.Println("Hi i am client and will do unary operation on simple intrest with deadline:", timeout, "Seconds")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := c.CalculateSiAndTotAmountWithDeadline(ctx, &CalculateSI.CalculateSIRequest{Principle: 1000.0, Rate: 10.5, Time: 13})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println("Response Message:", respErr.Message())
			fmt.Println("Response Code:", respErr.Code())

			if respErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded")
			} else {
				fmt.Printf("unexpected error: %v", respErr)
			}

			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("Sent a negative number")
			} else {
				log.Fatalf("Big error sending negative number")
				return
			}
		}
	}
	fmt.Println("SI : ", res.GetSi(), "Amount : ", res.GetTot())
}
