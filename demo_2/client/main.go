package main

import (
	"fmt"
	"log"
)

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	// 输出服务端的响应
	fmt.Println(reply)
}
