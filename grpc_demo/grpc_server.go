/**
 * @Author: lzw5399
 * @Date: 2020/8/11 22:36
 * @Desc: gRPC server
 */
package main

import (
	"gomicro-quickstart/grpc_demo/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// 1. new一个grpc的server
	rpcServer := grpc.NewServer()

	// 2. 将刚刚我们新建的ProdService注册进去
	service.RegisterProdServiceServer(rpcServer, new(service.ProdService))

	// 3. 新建一个listener，以tcp方式监听8082端口
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}

	// 4. 运行rpcServer，传入listener
	_ = rpcServer.Serve(listener)
}
