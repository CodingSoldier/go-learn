package main

import "github.com/gin-gonic/gin"

// 发送请求 http://localhost:8080/ping
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 监听0.0.0.0:8080
	r.Run()
}
