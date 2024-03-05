package main

// 使用select非阻塞运行channel
//func main() {
//	c1 := make(chan int, 5)
//	c2 := make(chan int)
//
//	// 使用select非阻塞运行channel
//	select {
//	// 如果从c1取出数据
//	case <-c1:
//		fmt.Println("c1")
//	// 如果向c2写入了数据
//	case c2 <- 22:
//		fmt.Println("c2")
//	default:
//		fmt.Println("默认")
//	}
//
//	time.Sleep(time.Second)
//}

type MyMutex chan struct {
}

func NewMyMutex() MyMutex {
	ch := make(chan struct{}, 1)
	return ch
}

func (m *MyMutex) Lock() {
	*m <- struct{}{}
}

func (m *MyMutex) Unlock() {
	<-*m
}
