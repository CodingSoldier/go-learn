package main

import (
	"fmt"
	"time"
)

/**
进程：
	一种系统运行行动

线程：
	CPU运算调度的最小单元
	作用：同时运算多个任务
	线程观摩：
		ps -M [pid]   查看pid对应的线程

协程（Coroutine）是轻量级的线程
协程的定位：用户控制的函数
协程的优势：
	1、协程的内存消耗更小
		一个线程可以包含多个协程
		线程大概会申请8MB内存，协程大概申请2KB内存
	2、协程切换更快
		线程申请内存需要经过内核，协程申请内存不需要经过内核，协程申请内存少了一道手续
	3、线程的切换由内核控制，协程由用户控制

Goroutine实际上也是一种协程
Goroutine对协程的优化：
	去掉冗余的协程声明周期管理
		协程创建
		协程完成
		协程重用
	降低额外的延迟和开销
	降低加锁/解锁的频率
*/

// 协程例子
func action() {
	fmt.Println("test goroutine")
}

func main() {
	go action()

	time.Sleep(1 * time.Second)
}
