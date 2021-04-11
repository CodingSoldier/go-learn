package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
获取get请求参数
http://localhost:8080/test?firstName=aa&lastName=bbb
*/
func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		firstName := c.Query("firstName")
		lastName := c.DefaultQuery("lastName", "默认值")

		c.String(http.StatusOK, "%s%s", firstName, lastName)
	})
	r.Run()
}
