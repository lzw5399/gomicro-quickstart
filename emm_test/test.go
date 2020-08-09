/**
 * @Author: lzw5399
 * @Date: 2020/8/8 23:30
 * @Desc: 测试consul服务发现
 */
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	// 1.连接到consul
	cr := consul.NewRegistry(registry.Addrs("47.100.220.174:8500"))

	for {
		// 2.根据service name获取对应的微服务列表
		services, err := cr.GetService("productService")
		if err != nil {
			log.Fatal("cannot get service list")
		}

		// 3.使用random随机获取其中一个实例
		next := selector.RoundRobin(services)
		svc, err := next()
		if err != nil {
			log.Fatal("cannot get service")
		}

		fmt.Println("[测试输出]:", svc.Address)
		time.Sleep(time.Second * 1)
	}
}
