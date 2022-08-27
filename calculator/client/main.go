package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-go-project/calculator/proto"
	"log"
)

var addr = "localhost:50002"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial%v\n", err)
	}
	defer conn.Close()
	instance := pb.NewCalcServiceClient(conn)
	doSum(instance)
}

func doSum(c pb.CalcServiceClient) {
	log.Printf("doSum is invoked")
	res, err := c.Calculator(context.Background(), &pb.CalRequest{
		A: 2,
		B: 5,
	})
	if err != nil {
		log.Fatalf("Couldn't send req %v\n", err)
	}

	log.Printf("Sum is %d", res.Res)
}
