/**
 * @Author: lzw5399
 * @Date: 2020/8/8 23:30
 * @Desc: 测试consul服务发现.
 */
package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	// 1.连接到consul
	cr := consul.NewRegistry(registry.Addrs("47.100.220.174:8500"))

	// 2.根据service name获取对应的微服务列表
	services, err := cr.GetService("productService")
	if err != nil {
		log.Fatal("cannot get service list")
	}

	// 3.使用random随机获取其中一个实例
	next := selector.Random(services)
	svc, err := next()
	if err != nil {
		log.Fatal("cannot get service")
	}

	fmt.Println("[测试输出]:", svc.Address)

	resp, err := RequestApi(http.MethodGet, svc.Address, "/v1/list", nil)
	if err != nil {
		log.Fatal("request api failed")
	}
	fmt.Println("[请求API结果]:", resp)
}

// 简单封装一个请求api的方法
func RequestApi(method string, host string, path string, body io.Reader) (string, error) {
	if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		host = "http://" + host
	}
	req, _ := http.NewRequest(method, host+path, body)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(buff), nil
}
