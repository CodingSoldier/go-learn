package main

import (
	"fmt"
	"sync"
)

/**
100人争抢10个鸡蛋
*/
func main() {

	// 初始化channel
	eggs := make(chan int, 10)

	// 初始化10个鸡蛋
	for i := 0; i < 10; i++ {
		eggs <- i
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			select {
			case egg := <-eggs:
				fmt.Printf("People: %d get egg：%d\n", i, egg)
			default:

			}
			wg.Done()
		}(i)
	}

	wg.Wait()

}
