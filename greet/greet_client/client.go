package main

import (
	"context"
	"fmt"
	"github.com/vinicius91/go-basic-01/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
	doServerStreaming(c)
	doClientStreaming(c)
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a client streaming RPC")

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Jessica",
				LastName:  "Jones",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Matt",
				LastName:  "Murdock",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Daniel",
				LastName:  "Rand",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Luke",
				LastName:  "Cage",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling the LongGreet: %v", err)
	}
	// We iterate the requests sending them
	for _, req := range requests {
		fmt.Printf("Sending Request: %v", req)
		err := stream.Send(req)
		time.Sleep(100 * time.Millisecond)
		if err != nil {
			log.Fatalf("Some error ocurried while send the message: %v", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet response: %v\n", res)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Start doing a streaming RPC")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Matt",
			LastName:  "Murdock",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while receiving stream of Greet: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// We've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("An error ocurried while receiving the stream message: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v \n", msg.GetResult())
	}

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Start doing unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Jessica",
			LastName:  "Jones",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.GetResult())
}
