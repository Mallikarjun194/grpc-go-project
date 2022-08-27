package main

import (
	"fmt"
	pb "grpc-go-project/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes func was invoked... with %v\n", in)
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s number %d", in.FirstName, i)

		stream.Send(&pb.GreetResp{
			Result: res,
		})
	}
	return nil
}
