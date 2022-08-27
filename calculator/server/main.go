package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-go-project/calculator/proto"

	//pb "grpc-go-project/greet/proto"
	"log"
	"net"
)

type CalServer struct {
	pb.CalcServiceServer
}

var addr = "0.0.0.0:50002"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	log.Printf("Listening on %s", addr)
	s := grpc.NewServer()
	pb.RegisterCalcServiceServer(s, &CalServer{})
	//pb.RegisterGreetServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v\n", err)
	}
}

func (s *CalServer) Calculator(ctx context.Context, in *pb.CalRequest) (*pb.CalResp, error) {
	log.Printf("Sum fun was invoked with %v\n", in)
	return &pb.CalResp{
		Res: in.A + in.B,
	}, nil
}
