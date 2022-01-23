package main

import (
	"context"
	"fmt"
	"sync"
)

/**
协程生命周期
	官方定义：协程的创建等全部生命周期的管理
	定义生命周期作用：便于协程的回收
		go语言中协程数量有最高限制，达到限制后，新创建的协程需要等待之前的协程运行完才能运行
	声明周期：
		协程创建
		协程回收
		协程中断
	实现：使用context实现中断
*/
func main() {
	// 初始化一个context
	parent := context.Background()
	// 生成一个取消的context
	ctx, cancel := context.WithCancel(parent)
	runTimes := 0
	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine Done")
				return
			default:
				fmt.Printf("Goroutine Running Times: %d\n", runTimes)
				runTimes++
			}
			if runTimes > 5 {
				cancel()
				wg.Done()
			}
		}
	}(ctx)
	wg.Wait()

}
