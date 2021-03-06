package main

import (
	"fmt"
	"time"
)

/**
生产者消费者
*/

var infos = make(chan int, 10)

func producer(index int) {
	infos <- index
	fmt.Printf("Producer: %d, Sent: %d\n", index, index)
}

func consumer(index int) {
	fmt.Printf("Comsumer: %d, Get: %d\n", index, <-infos)
}

func main() {

	// 生产者与消费者相同
	//for i := 0; i < 5; i++ {
	//	go producer(i)
	//}
	//for i := 0; i < 5; i++ {
	//	go consumer(i)
	//}

	// 生产者大于消费者，只能消费5个
	// 生产者能生产15次，chan缓冲大小 + 消费者消费数量 = 15
	for i := 0; i < 50; i++ {
		go producer(i)
	}
	for i := 0; i < 5; i++ {
		go consumer(i)
	}

	// 消费者大于生产者，也只是消费了5个
	//for i := 0; i < 5; i++ {
	//	go producer(i)
	//}
	//for i := 0; i < 50; i++ {
	//	go consumer(i)
	//}

	// 睡眠，观察阻塞
	time.Sleep(10 * time.Second)
}
