//Big Sky - 使える FizzBuzz
//http://mattn.kaoriya.net/software/lang/go/20120809174459.htm

package main

import (
	"net/rpc"
)

func main() {
	client, _ := rpc.Dial("tcp", "localhost:8001")
	var ret string
	for i := 1; i <= 100; i++ {
		client.Call("FizzBuzz.Serve", i, &ret)
		println(ret)
	}
}
