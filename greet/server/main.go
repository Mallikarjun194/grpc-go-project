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

var addr = "0.0.0.0:50001"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	log.Printf("Listening on %s", addr)
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v\n", err)
	}
}
