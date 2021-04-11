package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `form:"name" binding:"required"`
	Age  int    `form:"age" binding:"required,gt=10"`
}

func bindPerson(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.String(200, "person bind error:%v", err)
		c.Abort()
		return
	}

	c.String(200, "%v", person)
}

/**
请求参数认证
官方文档
https://pkg.go.dev/github.com/go-playground/validator/v10

http://localhost:8080/test
{
    "name": "名字",
    "age": 1
}
*/
func main() {
	r := gin.Default()
	r.GET("/test", bindPerson)
	r.POST("/test", bindPerson)
	r.Run()
}
