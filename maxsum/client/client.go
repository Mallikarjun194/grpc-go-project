package main

import (
	"context"
	"fmt"
	pb "grpc-go-project/maxsum/proto"
	"io"
	"log"
	"time"
)

func doMax(c pb.SumServiceClient) {
	fmt.Println("Max sum client is invoked..!")
	stream, err := c.MaxService(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes %v\n", err)
	}
	reqs := []*pb.Request{
		{Num: 12}, {Num: 15}, {Num: 3}, {Num: 16}, {Num: 2}, {Num: 20}, {Num: 3}, {Num: 10},
	}
	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("Sending num %v\n", req)
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
			log.Printf("Max is %s\n", msg.Res)
		}
		close(waitc)
	}()
	<-waitc
}
