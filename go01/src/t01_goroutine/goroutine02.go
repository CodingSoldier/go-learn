package main

import (
	"fmt"
	"sync"
)

/**
多协程
	官方定义：一段时间内协程的并行
	实际应用：某个任务使用多个协程同时进行处理
*/
func main() {

	/**
	多协程等待
	*/
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			fmt.Printf("Goroutine Test %d\n", num)
			wg.Done()
		}(i)
	}

	wg.Wait()

}
