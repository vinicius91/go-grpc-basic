package main

import (
	"context"
	"fmt"
	"github.com/vinicius91/go-basic-01/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Starting Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("An error ocurried while tring to connect %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	sum(c, 15, 12)
	sum(c, 15, 24)
	subtract(c, 16, 10)
	subtract(c, 16, 25)
	multiply(c, 9, 9)
	multiply(c, 99.57, 81.92)
	decomposePrime(c, 12)
	decomposePrime(c, 1465456442)

}

func sum(c calculatorpb.CalculatorServiceClient, f float32, s float32 ) {
	fmt.Println("Starting the sum")

	req := &calculatorpb.OperationRequest{
		Numbers: &calculatorpb.Numbers{
			FirstNumber:f,
			SecondNumber:s,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("An error has ocurried while trying to sum %v\n", err)
	}

	log.Printf("The result is %v\n", res.Result)
}

func subtract(c calculatorpb.CalculatorServiceClient, f float32, s float32 ) {
	fmt.Println("Starting the subtract")

	req := &calculatorpb.OperationRequest{
		Numbers: &calculatorpb.Numbers{
			FirstNumber:f,
			SecondNumber:s,
		},
	}

	res, err := c.Subtract(context.Background(), req)
	if err != nil {
		log.Fatalf("An error has ocurried while trying to sum %v\n", err)
	}

	log.Printf("The result is %v\n", res.Result)
}


func multiply(c calculatorpb.CalculatorServiceClient, f float32, s float32 ) {
	fmt.Println("Starting the multiply")

	req := &calculatorpb.OperationRequest{
		Numbers: &calculatorpb.Numbers{
			FirstNumber:f,
			SecondNumber:s,
		},
	}

	res, err := c.Multiply(context.Background(), req)
	if err != nil {
		log.Fatalf("An error has ocurried while trying to sum %v\n", err)
	}

	log.Printf("The result is %v\n", res.Result)
}

func decomposePrime(c calculatorpb.CalculatorServiceClient, n int32) {
	fmt.Println("Starting the multiply")

	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number:n,
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while sending the requisition to PrimenumberDecomposition: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Printf("End of PrimeNumberDecomposition")
			break
		}
		if err != nil {
			log.Fatalf("An error ocurried while receiving the stream of PrimeNumberDecomposition: %v", err)
		}
		log.Printf("Divisor: %v", msg.GetPrimeFactor())
	}

}