package main

import (
	"fmt"
	"time"
)

/**
协程英文单词：Coroutine
	轻量级的“线程”
	线程是抢占式的，抢占式在切换线程的时候需要保存更多的状态；协程是非抢占式，由协程主动交出控制权
	协程是编译器/解释器/虚拟机层面的多任务
	多个协程可能是在一个或者多个线程上运行

*/
func main() {
	var a [10]int

	/**
	设置 i<100
	执行 go run -race o01_goroutine.go
	会提示 panic: runtime error: index out of range [98] with length 10

	*/
	for i := 0; i < 10; i++ {
		// go后面加一个立即执行的函数，就是开协程
		go func(i int) {
			// for死循环
			for {
				// 打印是一个IO操作，会发生协程切换
				fmt.Println("打印 i=", i)

				// 协程永远不交出控制权
				//a[i]++

				// 手动交出控制权，一般不用
				//runtime.Gosched()
			}
		}(i)
	}

	// 睡眠避免程序立即退出
	time.Sleep(time.Minute)

	fmt.Println(a)
}

/**
goroutine定义：
	任何函数只需要加上go就能送给调度器运行，一个协程可能会在一个线程中运行，多个协程也可能在一个线程中运行。o02_goroutine.jpg
	不需要在定义时区分是否是异步函数（python需要）
	调度器在合适的点进行切换
	使用-race来检测数据访问冲突

goroutine可能的切换点（只是个参考）
	I/O,select
	channel
	等待锁
	函数调用（有时）
	runtime.Gosched()

*/
