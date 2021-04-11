package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

/**
使用日志中间件
*/
func main() {
	// 将日志打印到gin.log
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.New()
	//r.Use(gin.Logger())

	/**
	gin.Logger()log中间件
	gin.Recovery() 全局捕获异常，但是发生panic()没返回信息
	*/
	//r.Use(gin.Logger(), gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "默认值")

		fmt.Println("执行############")
		panic("报错")

		c.String(200, "%s", name)
	})

	r.Run()
}
