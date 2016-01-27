package main

import (
	pb "github.com/dictav/go-rpc-learning/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func fizzbuzz(n int32) (string, error) {

	conn, err := grpc.Dial("127.0.0.1:8001", grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()
	client := pb.NewFizzBuzzServiceClient(conn)

	res, err := client.FizzBuzz(context.Background(), &pb.Request{Answer: n})
	if err != nil {
		return "", err
	}
	return res.Answer, nil
}

func main() {
	for i := 1; i <= 100; i++ {
		ret, err := fizzbuzz(int32(i))
		if err != nil {
			panic(err.Error())
		}

		println(ret)
	}
}
