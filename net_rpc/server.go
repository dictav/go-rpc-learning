//Big Sky - 使える FizzBuzz
//http://mattn.kaoriya.net/software/lang/go/20120809174459.htm
package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type FizzBuzz int

func (fb *FizzBuzz) Serve(n int, r *string) error {
	switch {
	case n%15 == 0:
		*r = "FizzBuzz"
	case n%3 == 0:
		*r = "Fizz"
	case n%5 == 0:
		*r = "Buzz"
	default:
		*r = fmt.Sprintf("%d", n)
	}
	return nil
}

func main() {
	fb := new(FizzBuzz)
	rpc.Register(fb)

	server, _ := net.Listen("tcp", ":8001")
	for {
		client, err := server.Accept()
		if err != nil {
			println(err.Error())
			continue
		}
		rpc.ServeConn(client)
	}
}
