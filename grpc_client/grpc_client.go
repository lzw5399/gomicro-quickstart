/**
 * @Author: lzw5399
 * @Date: 2020/8/11 23:09
 * @Desc: grpc服务端
 */
package main

import (
	"context"
	"fmt"
	"gomicro-quickstart/grpc_client/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	tls, err := credentials.NewClientTLSFromFile("grpc_client/keys/server.crt", "")

	if err != nil {
		log.Fatal("客户端获取证书失败: ", err)
	}

	// 1. 新建连接，端口是服务端开放的8082端口
	// 并且添加grpc.WithInsecure()，不然没有证书会报错
	conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(tls))
	if err != nil {
		log.Fatal(err)
	}

	// 退出时关闭链接
	defer conn.Close()

	// 2. 调用Product.pb.go中的NewProdServiceClient方法
	productServiceClient := service.NewProdServiceClient(conn)

	// 3. 直接像调用本地方法一样调用GetProductStock方法
	resp, err := productServiceClient.GetProductStock(context.Background(), &service.ProductRequest{ProdId: 233})
	if err != nil {
		log.Fatal("调用gRPC方法错误: ", err)
	}

	fmt.Println("调用gRPC方法成功，ProdStock = ", resp.ProdStock)
}
