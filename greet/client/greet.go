package main

import (
	"context"
	pb "grpc-go-project/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet is invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Zayn-Malik",
	})
	if err != nil {
		log.Fatalf("Couldn't greet %v\n", err)
	}

	log.Printf("Greeting is %s ", res.Result)
}
