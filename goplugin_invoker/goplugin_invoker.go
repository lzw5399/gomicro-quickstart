/**
 * @Author: lzw5399
 * @Date: 2020/8/10 21:37
 * @Desc:
 */
package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
	"log"
)

func main() {
	// 1. 注册consul地址
	cr := consul.NewRegistry(registry.Addrs("47.100.220.174:8500"))

	// 2. 实例化selector
	mySelector := selector.NewSelector(
		selector.Registry(cr),                     // 传入上面的consul
		selector.SetStrategy(selector.RoundRobin), // 指定获取实例的算法
	)

	// 3. 请求服务
	resp, err := callByGoPlugin(mySelector)
	if err != nil{
		log.Fatal("request API failed", err)
	}

	fmt.Printf("[服务调用结果]: %v", resp)
}

func callByGoPlugin(s selector.Selector) (map[string]interface{}, error) {
	// 1. 调用`go-plugins/client/http`包的函数获取它们提供的httpClient
	gopluginClient := http.NewClient(
		client.Selector(s),                     // 传入上面的selector
		client.ContentType("application/json"), // 指定contentType
	)

	// 2. 新建请求对象，传入: (1)服务名 (2)endpoint (3)请求参数
	req := gopluginClient.NewRequest("ProductService", "/v1/list", map[string]string{})

	// 3. 新建响应对象，并call请求，获取响应
	var resp map[string]interface{}
	err := gopluginClient.Call(context.Background(), req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
