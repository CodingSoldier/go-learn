package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

/**
获取post请求参数
http://localhost:8080/test

*/
func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		/**
		获取form-data数据
		*/
		firstName := c.PostForm("firstName")
		lastName := c.DefaultPostForm("lastName", "lastName默认值")

		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}

		// 读取了body之后，需要将参数再放回body，后面才能继续获取请求参数
		//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		//firstName := c.PostForm("firstName")
		//lastName := c.DefaultPostForm("lastName", "lastName默认值")
		fmt.Println(bodyBytes)
		//c.String(http.StatusOK, "%s", string(bodyBytes))

		c.String(http.StatusOK, "%s, %s", firstName, lastName)
	})
	r.Run()
}
