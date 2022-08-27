package main

import (
	"google.golang.org/grpc"
	pb "grpc-go-project/greet/proto"
	"log"
	"net"
)

type Server struct {
	pb.GreetServiceServer
}

var addr string = "127.0.0.1:5001"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	log.Fatal("Listening on:", addr)
	s := grpc.NewServer()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v\n", err)
	}
}
