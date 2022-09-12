package main

import (
	pb "grpc-go-project/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetManyTimes func was invoked...")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//log.Fatalf("Error while reading client stream %v\n", err)
			return nil
		}
		res := "Hello" + req.FirstName + "!"
		err = stream.Send(&pb.GreetResp{
			Result: res,
		})
		if err != nil {
			log.Fatalf("Error while sending data %v\n", err)
		}
	}
	return nil
}
