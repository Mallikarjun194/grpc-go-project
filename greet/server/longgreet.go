package main

import (
	"fmt"
	pb "grpc-go-project/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Print("Welcome to client-stream longGeet func was invoked...")
	res := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResp{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading %v\n", err)
		}
		log.Printf("Receving request %s\n", req)
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
