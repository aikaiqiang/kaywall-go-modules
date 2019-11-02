package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	//if len(os.Args) != 2 {
	//	fmt.Println("Usage: ", os.Args[0], "server:port")
	//	log.Fatal(1)
	//}
	//service := os.Args[1]
	service := "127.0.0.1:1234"

	client, err := jsonrpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith 11-error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith 11-error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
