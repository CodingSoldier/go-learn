package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func add(p *int32) {
	// 原子相加
	atomic.AddInt32(p, 1)
}

func main() {
	c := int32(0)
	for i := 0; i < 1000; i++ {
		go add(&c)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(c)
}
