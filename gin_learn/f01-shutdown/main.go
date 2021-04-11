package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
优雅停机
进入 E:\go-learn\gin_learn\f01-shutdown
执行 go run main.go
发送请求 http://127.0.0.1:8085
按下ctrl+c，5秒后服务停止
*/
func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		// 模拟一个处理时间很长的请求
		time.Sleep(5 * time.Second)
		fmt.Println("Welcome Gin Server")
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8085",
		Handler: router,
	}

	// 启动一个监听
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 将shutdown信号存入channel中
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server...")

	// 等待一段时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 真正执行关闭鼓舞
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("服务退出")
	router.Run()
}
