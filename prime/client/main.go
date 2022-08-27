package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-go-project/prime/proto"
	"io"
	"log"
)

var addr = "localhost:50003"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect%v\n", err)
	}
	instance := pb.NewPrimeServiceClient(conn)
	doPrime(instance)

}

func doPrime(c pb.PrimeServiceClient) {
	log.Printf("doPrime is invoked")
	req := &pb.PrimeReq{
		Num: 150,
	}

	stream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Prime %v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v\n", err)
		}
		log.Printf("Prime is %d\n", msg.Res)
	}
}
