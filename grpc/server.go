package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/dictav/go-rpc-learning/grpc/proto"
)

type fizzbuzzService struct {
}

func (fbs *fizzbuzzService) FizzBuzz(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	n := req.N
	switch {
	case n%15 == 0:
		return &pb.Response{"FizzBuzz"}, nil
	case n%3 == 0:
		return &pb.Response{"Fizz"}, nil
	case n%5 == 0:
		return &pb.Response{"Buzz"}, nil
	default:
		return &pb.Response{fmt.Sprintf("%d", n)}, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	pb.RegisterFizzBuzzServiceServer(server, new(fizzbuzzService))
	server.Serve(lis)
}
