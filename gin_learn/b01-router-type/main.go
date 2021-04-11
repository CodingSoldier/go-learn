package main

import "github.com/gin-gonic/gin"

/**
请求类型
*/
func main() {
	r := gin.Default()

	/**
	接收请求 http://localhost:8080/get
	返回string类型数据 get
	*/
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "get")
	})

	/**
	接收 http://localhost:8080/post 请求
	*/
	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post")
	})

	/**
	delete类型的请求 http://localhost:8080/delete
	*/
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "delete")
	})

	/**
	r.Any()能处理任何类型的请求
	*/
	r.Any("/any", func(c *gin.Context) {
		c.String(200, "any")
	})

	r.Run()
}
