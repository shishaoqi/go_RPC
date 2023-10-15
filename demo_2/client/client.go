package main

import (
	"net/rpc"
	"rpc_demo2/rpcService"
)

// 创建更形象的结构体
type HelloServiceClient struct {
	*rpc.Client
}

var HelloServiceInterface = (*HelloServiceClient)(nil)

// 结构体初始化并赋值
func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

// 把对应方法搞上
func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(rpcService.HelloServiceName+".Hello", request, reply)
}
