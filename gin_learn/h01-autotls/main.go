package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

/**
下载依赖，运行 go get github.com/gin-gonic/autotls@0.0.2

必须在公网环境运行
*/
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	autotls.Run(r, "www.itpp.tk")
}
