package main

import (
	"context"
	pb "grpc-go-project/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Printf("doGreetManyTimes was invoked...")
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes %v\n", err)
	}
	reqs := []*pb.GreetRequest{
		{FirstName: "Harry"},
		{FirstName: "Zayn-Malik"},
		{FirstName: "Ravi"},
	}
	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("Sending req %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	go func() {

		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v\n", err)
				break
			}
			log.Printf("Everyone is %s\n", msg.Result)
		}
		close(waitc)
	}()
	<-waitc
}
