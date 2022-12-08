package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"graproject/bootstrap"
)

func main() {

	// new 一个 Gin Engine 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	err := router.Run(":3000")
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
