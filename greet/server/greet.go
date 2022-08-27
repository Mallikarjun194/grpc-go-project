package main

import (
	"context"
	pb "grpc-go-project/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResp, error) {
	log.Printf("Greeting fun was invoked with %v\n", in)

	return &pb.GreetResp{
		Result: "Hello " + in.FirstName,
	}, nil
}
