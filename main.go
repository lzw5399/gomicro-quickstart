package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.String(http.StatusOK, "user api")
	})

	server := web.NewService(
		web.Address(":8081"),                                // 端口
		web.Metadata(map[string]string{"protocol": "http"}), // 元信息
		web.Handler(r)) // 路由
		
	_ = server.Run()
}
