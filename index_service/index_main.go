/**
 * @Author: lzw5399
 * @Date: 2020/8/8 22:53
 * @Desc: 模拟微服务中的主站API
 */
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"net/http"
)

func main() {
	router := gin.Default()
	// 下面准备要显示当前所有注册了的微服务名
	router.GET("/", func(c *gin.Context) {
		// data代表
		data := make([]interface{}, 0)
		c.JSON(http.StatusOK, gin.H{
			"apis": data,
		})
	})

	server := web.NewService(
		web.Address(":8000"),                                // 端口
		web.Handler(router)) // 路由

	_ = server.Run()
}
