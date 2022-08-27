package main

import (
	"google.golang.org/grpc"
	pb "grpc-go-project/prime/proto"
	"log"
	"net"
)

var addr = "localhost:50003"

type Server struct {
	pb.PrimeServiceServer
}

func main() {
	list, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	log.Printf("Listening on %s", addr)
	s := grpc.NewServer()
	pb.RegisterPrimeServiceServer(s, &Server{})
	if err = s.Serve(list); err != nil {
		log.Fatalf("Failed to server %v\n", err)
	}
}

func (s *Server) Prime(in *pb.PrimeReq, stream pb.PrimeService_PrimeServer) error {
	log.Printf("Prime func was invoked... with %v\n", in)
	var k uint32
	k = 2
	for in.Num > 1 {
		if in.Num%k == 0 {
			stream.Send(&pb.PrimeResp{
				Res: k,
			})
			in.Num = in.Num / k
		} else {
			k += 1
		}
	}
	return nil

}

//func Result(N uint32) {
//	var k uint32
//	k = 2
//	for N > 1 {
//		if N%k == 0 {
//			fmt.Println(k)
//			N = N / k
//		} else {
//			k += 1
//		}
//	}
//}
