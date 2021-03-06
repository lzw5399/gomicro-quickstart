/**
 * @Author: lzw5399
 * @Date: 2020/8/8 22:54
 * @Desc: 模拟微服务中的产品服务productService
 */
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"gomicro-quickstart/product_service/model"
	"net/http"
)

func main() {
	// 添加consul地址
	cr := consul.NewRegistry(registry.Addrs("47.100.220.174:8500"))

	// 使用gin作为路由
	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.POST("list", func(c *gin.Context) {
			var req ProdRequest
			if err := c.Bind(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"data": "模型绑定失败",
				})
				c.Abort()
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"data": model.NewProductList(req.Size),
			})
		})
	}

	server := web.NewService(
		web.Name("ProductService"),                          // 当前微服务服务名
		web.Registry(cr),                                    // 注册到consul
		web.Address(":8001"),                                // 端口
		web.Metadata(map[string]string{"protocol": "http"}), // 元信息
		web.Handler(router)) // 路由

	_ = server.Init()

	_ = server.Run()
}

type ProdRequest struct {
	Size int `json:"size"`
}
