package main

import (
	"github.com/sirupsen/logrus"
	pb "grpc-go-project/maxsum/proto"
	"io"
	"log"
)

func (s *Server) MaxService(stream pb.SumService_MaxServiceServer) error {
	logrus.Infof("MaxService func was invoked...!!")
	var maxnum int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//log.Fatalf("Error while reading client stream %v\n", err)
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}
		//slice := make([]int32, 0)
		//slice = append(slice, req.Num)
		if num := req.Num; num > maxnum {
			maxnum = num
			err = stream.Send(&pb.Resp{
				Res: maxnum,
			})
			if err != nil {
				log.Fatalf("Error while sending data %v\n", err)
			}
		}
	}
	return nil
}

//func sum(slice []int32) int32 {
//	s := int32(0)
//	for _, v := range slice {
//		s += v
//	}
//	return s
//
//}
