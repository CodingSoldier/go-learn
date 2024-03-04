package main

import (
	"fmt"
	"time"
)

//func main() {
//	// 创建缓冲区为5的字符串管道，默认缓冲区是0
//	c := make(chan string, 5)
//	// a进入缓冲区
//	c <- "a"
//	// 从缓冲区拿数据，并赋值给str
//	str := <-c
//	fmt.Printf("str=%s\n", str)
//}

// 通过channel通信的方式共享内存
func watch(c chan int) {
	if <-c == 1 {
		fmt.Println("退出watch")
	}
}

func main() {
	c := make(chan int)
	go watch(c)
	time.Sleep(time.Second)
	c <- 1
	time.Sleep(time.Second)
}
