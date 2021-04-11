package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
请求静态文件
*/
func main() {
	router := gin.Default()

	/**
	进入 E:\go-learn\gin_learn\b02-router-static
	执行 go build
	运行 b02-router-static.exe
	发送请求 http://localhost:8080/asserts/a.html
	得到a.html文本

	在idea中运行会提示请求404
	*/
	router.Static("/asserts", "./asserts")

	/**
	同样是要打包才能访问到静态文件
	*/
	router.StaticFS("/static", http.Dir("static"))

	/**
	同样是要打包才能访问到静态文件
	*/
	router.StaticFile("/a.jpg", "./resource/a.jpg")

	router.Run()

}
