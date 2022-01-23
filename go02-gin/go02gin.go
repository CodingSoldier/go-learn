package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ok", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world!")
	})
	fmt.Printf("start server (port:%s)", "8000")
	r.Run(":8000")
}
