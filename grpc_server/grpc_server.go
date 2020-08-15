/**
 * @Author: lzw5399
 * @Date: 2020/8/11 22:36
 * @Desc: gRPC server
 */
package main

import (
	"fmt"
	"gomicro-quickstart/grpc_server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net/http"
)

func main() {
	// 1. 引用证书
	tls, err := credentials.NewServerTLSFromFile("grpc_server/keys/server.crt", "grpc_server/keys/server_no_password.key")
	if err != nil {
		log.Fatal("服务端获取证书失败: ", err)
	}

	// 2. new一个grpc的server，并且加入证书
	rpcServer := grpc.NewServer(grpc.Creds(tls))

	// 3. 将刚刚我们新建的ProdService注册进去
	service.RegisterProdServiceServer(rpcServer, new(service.ProdService))

	// 4. 新建一个listener，以tcp方式监听8082端口
	//listener, err := net.Listen("tcp", ":8082")
	//if err != nil {
	//	log.Fatal("服务监听端口失败", err)
	//}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request)
		rpcServer.ServeHTTP(writer, request)
	})
	// 5. 运行rpcServer，传入listener
	// _ = rpcServer.Serve(listener)
	httpServer := http.Server{
		Addr:    ":8082",
		Handler: mux,
	}

	httpServer.ListenAndServeTLS("grpc_server/keys/server.crt", "grpc_server/keys/server_no_password.key")
}
