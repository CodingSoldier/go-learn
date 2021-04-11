gin学习项目

创建工程
    1、新建project —> Go Modules -> 填写project location -> finish
    2、下载gin依赖
        项目根目录执行
        go get github.com/gin-gonic/gin@v1.4.0
        go.mod多了依赖，依赖和版本之间没有@
        github.com/gin-gonic/gin v1.4.0
    3、新建start目录，新建main.go，输入如下内容。
        运行main()，发送请求 http://localhost:8080/ping
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}







