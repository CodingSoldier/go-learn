package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

/**
获取post请求参数
http://localhost:8080/test
{
    "name": "名字"
}
*/
func main() {
	r := gin.Default()

	r.POST("/test", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		c.String(http.StatusOK, "%s", string(bodyBytes))
	})

	r.Run()
}
