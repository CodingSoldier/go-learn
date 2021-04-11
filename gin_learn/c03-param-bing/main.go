package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string `form:"name"`
	Address  string `form:"address"`
	Birthday string `form:"birthday" time_format:"2016-01-02 01:01:01" time_utc:"1"`
}

func bindPerson(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err == nil {
		c.String(200, "%v", person)
	} else {
		c.String(200, "person bind error:%v", err)
	}
}

/**
请求参数绑定结构体
http://localhost:8080/test
{
    "name": "名字",
    "address": "地址",
    "birthday": "2020-01-01 05:05:05"
}

http://localhost:8080/test?name=名字&address=地址&birthday=2020-01-01 05:05:05

*/
func main() {
	r := gin.Default()
	r.GET("/test", bindPerson)
	r.POST("/test", bindPerson)
	r.Run()
}
