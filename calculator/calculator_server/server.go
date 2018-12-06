package main

import (
	"context"
	"fmt"
	"github.com/vinicius91/go-basic-01/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {

}

func (*server) Sum(ctx context.Context, req *calculatorpb.OperationRequest) (*calculatorpb.OperationResponse, error) {
	fmt.Printf("Sum function invoked with %v\n", req)
	firstNumber := req.GetNumbers().GetFirstNumber()
	secondNumber := req.GetNumbers().GetSecondNumber()

	result := firstNumber + secondNumber

	res := &calculatorpb.OperationResponse{
		Result:result,
	}

	return res, nil

}

func (*server) Subtract(ctx context.Context, req *calculatorpb.OperationRequest) (*calculatorpb.OperationResponse, error) {
	fmt.Printf("Subtract function invoked with %v\n", req)
	firstNumber := req.GetNumbers().GetFirstNumber()
	secondNumber := req.GetNumbers().GetSecondNumber()

	result := firstNumber - secondNumber

	res := &calculatorpb.OperationResponse{
		Result:result,
	}

	return res, nil

}

func (*server) Multiply(ctx context.Context, req *calculatorpb.OperationRequest) (*calculatorpb.OperationResponse, error) {
	fmt.Printf("Multiply function invoked with %v\n", req)
	firstNumber := req.GetNumbers().GetFirstNumber()
	secondNumber := req.GetNumbers().GetSecondNumber()

	result := firstNumber * secondNumber

	res := &calculatorpb.OperationResponse{
		Result:result,
	}

	return res, nil

}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Starting Prime Number Decomposition %v\n", req)
	number := req.GetNumber()
	divisor := int32(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor:divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v\n", divisor)
		}
	}

	return nil
}

func main() {
	fmt.Println("Server Started listening on port 50051")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
