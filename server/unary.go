package main

import (
	"context"

	pb "github.com/MilkiasTonji/gRPC-go/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}

//start from here tomorrow morning....
