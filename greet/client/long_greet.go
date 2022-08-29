package main

import (
	"context"
	pb "grpc-go-project/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet fun was invoked...")

	reqs := []*pb.GreetRequest{
		{FirstName: "Harry"},
		{FirstName: "Zayn-Malik"},
		{FirstName: "Ravi"},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while streaming %v\n", err)
	}
	for _, req := range reqs {
		log.Printf("Sending %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receving response %v\n", err)
	}
	log.Printf("Receving response")
	log.Printf("LongGreet is %s\n", res.Result)

}
