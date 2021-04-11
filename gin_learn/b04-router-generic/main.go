package main

import "github.com/gin-gonic/gin"

/**
范绑定
处理以 http://localhost:8080/user/ 开头的请求

例如：
http://localhost:8080/user/xxx

*/
func main() {
	engine := gin.Default()
	engine.GET("/user/*action", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "hello world",
		})
	})

	engine.Run()
}
