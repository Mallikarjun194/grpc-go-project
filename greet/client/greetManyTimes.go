package main

import (
	"context"
	pb "grpc-go-project/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("doGreetManyTimes was invoked...")
	req := &pb.GreetRequest{
		FirstName: "zayn-malik",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes %v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v\n", err)
		}
		log.Printf("GreetManyTimes is %s\n", msg.Result)
	}
}
