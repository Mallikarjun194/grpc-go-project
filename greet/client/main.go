package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-go-project/greet/proto"
	"log"
)

var addr = "localhost:50001"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect%v\n", err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
}
