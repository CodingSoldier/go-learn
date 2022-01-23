package main

import (
	"fmt"
	"time"
)

///**
//channel实现协程与协程之间的通信
//
// */
//func chanDemo(){
//	// 创建一个channel
//	channel := make(chan int)
//
//	// 在这里往channel发送数据会发生死锁，因为往channel发送数据之后要有协程接收数据
//	//channel <- 1
//
//	go func() {
//		for {
//			n := <-channel
//			fmt.Println(n)
//		}
//	}()
//
//	channel <- 1
//	channel <- 4
//
//	time.Sleep(time.Millisecond)
//}

func worker(id int, c chan int) {
	/**
	n := range c
	遍历取出通道中的值
	*/
	for n := range c {
		fmt.Printf("Worker %d received %c \n", id, n)
	}
}

/**
chan<- int
返回一个只能接收int数据的channel
*/
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
}

// 创建带有缓冲区的Channel
func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
}

// 关闭channel
func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
}

/**
go的并发模式是CSP模型

go强调不要通过共享内存通信，而是通过通信共享内存
*/
func main() {

	//chanDemo()

	//bufferedChannel()

	channelClose()

	time.Sleep(time.Millisecond)

}
