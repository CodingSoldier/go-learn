package main

import (
	"github.com/gin-gonic/gin"
)

/**
自定义中间件，校验IP是否在白名单中
*/
func IpAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.1",
		}
		flag := false
		clientIP := c.ClientIP()
		for _, host := range ipList {
			if clientIP == host {
				flag = true
				break
			}
		}

		if !flag {
			c.String(401, "ip没权限")
			c.Abort()
		}
	}
}

/**
发送请求
http://127.0.0.1:8080/test?name=11111111
*/
func main() {
	r := gin.New()

	r.Use(IpAuthMiddleware())

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "成功")
	})

	r.Run()
}
