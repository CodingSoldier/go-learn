package main

import "fmt"

/**
channel，信道，用于多个协程之间通信

channel在go lang中的妙用
	传递方面：
		消息传递、任务发送、事件广播
	控制方面：
		资源争抢、并发控制、流程开关

使用channel
	输入 ->chan
	输出 <-chan
	关闭 close(chan)


*/
func main() {

	// 初始化channel
	// channel在使用前，必须进行make初始化
	// 否则会是一个nil
	ch := make(chan int)

	// 输出channel，必须先在一个goroutine中输出channel，之后才能写channel输入
	go func() {
		fmt.Println(<-ch)
	}()

	// 输入channel
	ch <- 1

	// 关闭channel
	close(ch)

	// channel关闭后，不可以再输入
	//ch <- 2

	// channel关闭之后，可以再输出
	<-ch

}
