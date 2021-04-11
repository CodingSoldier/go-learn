package main

import "github.com/gin-gonic/gin"

/**
在url中使用参数
http://localhost:8080/abc/123
*/
func main() {
	engine := gin.Default()
	engine.GET("/:name/:id", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": context.Param("name"),
			"id":   context.Param("id"),
		})
	})

	engine.Run()
}
