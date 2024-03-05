package main

<<<<<<< HEAD
// 使用select非阻塞运行channel
//func main() {
//	c1 := make(chan int, 5)
//	c2 := make(chan int)
//
//	// 使用select非阻塞运行channel
//	select {
//	// 如果从c1取出数据
//	case <-c1:
//		fmt.Println("c1")
//	// 如果向c2写入了数据
//	case c2 <- 22:
//		fmt.Println("c2")
//	default:
//		fmt.Println("默认")
//	}
//
//	time.Sleep(time.Second)
//}

type MyMutex chan struct {
}

func NewMyMutex() MyMutex {
	ch := make(chan struct{}, 1)
	return ch
}

func (m *MyMutex) Lock() {
	*m <- struct{}{}
}

func (m *MyMutex) Unlock() {
	<-*m
=======
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
>>>>>>> origin/main
}
