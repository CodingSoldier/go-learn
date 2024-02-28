package main

import (
	"fmt"
	"math"
	"time"
)

func do(i int, ch chan struct{}) {
	fmt.Println(i)
	time.Sleep(time.Second)
	<-ch
}

func main() {
	// chanel缓冲区是3000，缓冲区满了之后无法放入struct
	c := make(chan struct{}, 3000)
	for i := 0; i < math.MaxInt32; i++ {
		c <- struct{}{}
		go do(i, c)
	}
	time.Sleep(time.Hour)
}
