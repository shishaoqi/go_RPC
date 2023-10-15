package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"rpc_demo2/rpcService"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	fmt.Println("server Hello func receive: ", request)
	*reply = "rpc server reply: " + request
	return nil
}

func main() {
	rpcService.RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listenTCP:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accep error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
