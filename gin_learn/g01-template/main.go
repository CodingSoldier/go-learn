package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
模板渲染
*/
func main() {
	router := gin.Default()
	/**
	项目的根目录是 gin_learn
	可直接在idea运行main方法
	*/
	router.LoadHTMLGlob("g01-template/template/*")

	/**
	进入 E:\go-learn\gin_learn\g01-template\template
	执行 go run main.go
	*/
	//router.LoadHTMLGlob("template/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "标题",
		})
	})

	router.Run()
}
