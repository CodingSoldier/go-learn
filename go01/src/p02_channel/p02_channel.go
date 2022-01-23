package main

import (
	"fmt"
	"sync"
)

/**
使用channel等待goroutine的结束
*/

/**
定义一个worker
in 是一个channel
done是goroutine结束时的回调
*/
type worker struct {
	in   chan int
	done func()
}

/**
执行worker函数体
取出channel中的数据
*/
func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c \n", id, n)
		// 表示worker执行完毕
		w.done()
	}
}

/**
创建worker
*/
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// 等待worker执行完毕
	wg.Wait()
}

func main() {

	chanDemo()

}
