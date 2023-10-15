package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc_demo1/rpcService"
)

func main() {
	rpc.RegisterName("HelloService", new(rpcService.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)
}
