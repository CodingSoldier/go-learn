package main

import (
	"fmt"
	"sync"
	"time"
)

//func add(p *int32) {
//	// 原子相加
//	atomic.AddInt32(p, 1)
//}
//
//func main() {
//	c := int32(0)
//	for i := 0; i < 1000; i++ {
//		go add(&c)
//	}
//	time.Sleep(5 * time.Second)
//	fmt.Println(c)
//}
//
//type Person struct {
//	// 加Mutex锁
//	mu     sync.Mutex
//	salary int
//	level  int
//}
//
//func promote(p *Person) {
//	// 加锁
//	p.mu.Lock()
//	// 使用defer解锁
//	defer p.mu.Unlock()
//	p.salary++
//	fmt.Println(p.salary)
//	p.level++
//	fmt.Println(p.level)
//
//	//运行代码后再解锁，可能中间代码报错，导致未解锁
//	//p.mu.Unlock()
//
//}
//
//func main() {
//	p := Person{salary: 1000, level: 1}
//	go promote(&p)
//	go promote(&p)
//	go promote(&p)
//	time.Sleep(time.Minute)
//}

//type Person struct {
//	mu     int32
//	salary int
//	level  int
//}
//
//func (p *Person) promote() {
//	for {
//		// 使用 自旋 + CAS 加锁
//		if atomic.CompareAndSwapInt32(&p.mu, 0, 1) {
//			p.salary++
//			fmt.Println(p.salary)
//			p.level++
//			fmt.Println(p.level)
//			//  使用 自旋 + CAS 解锁
//			atomic.CompareAndSwapInt32(&p.mu, 1, 0)
//			break
//		}
//	}
//}
//
//func main() {
//	p := Person{salary: 1000, level: 1, mu: 0}
//	go p.promote()
//	go p.promote()
//	go p.promote()
//	time.Sleep(time.Minute)
//}

//// WaitGroup的使用
//type Person struct {
//	mu     sync.RWMutex
//	salary int32
//	level  int32
//}
//
//func promote(p *Person, wg *sync.WaitGroup) {
//	p.mu.Lock()
//	p.salary++
//	fmt.Printf("写salary: %d\n", p.salary)
//	p.level++
//	fmt.Printf("写level: %d\n", p.level)
//	p.mu.Unlock()
//	wg.Done()
//}
//
//func printPerson(p *Person, wg *sync.WaitGroup) {
//	p.mu.RLock()
//	defer p.mu.RUnlock()
//	fmt.Printf("读salary：%d\n", p.salary)
//	fmt.Printf("读level: %d\n", p.level)
//}
//
//func main() {
//	p := Person{salary: 1000, level: 0}
//	wg := sync.WaitGroup{}
//	wg.Add(3)
//
//	go promote(&p, &wg)
//	go promote(&p, &wg)
//	go promote(&p, &wg)
//	go printPerson(&p, &wg)
//	go printPerson(&p, &wg)
//	go printPerson(&p, &wg)
//
//	wg.Wait()
//}

//// Once 只执行一次
//type Person struct {
//	mu     sync.RWMutex
//	salary int32
//	level  int32
//}
//
//func (p *Person) promote() {
//	p.salary++
//	fmt.Printf("写salary: %d\n", p.salary)
//	p.level++
//	fmt.Printf("写level: %d\n", p.level)
//}
//
//func main() {
//	p := Person{salary: 1000, level: 0}
//	once := sync.Once{}
//
//	go once.Do(p.promote)
//	go once.Do(p.promote)
//	go once.Do(p.promote)
//
//	time.Sleep(time.Minute)
//}

// 死锁问题
type Person struct {
	mu     sync.RWMutex
	salary int32
	level  int32
}

func main() {
	mu := sync.Mutex{}
	mu.Lock()
	fmt.Println("mu上锁了")

	// 拷贝锁，此时mu是上锁状态。拷贝出来的p也是上锁状态，并且会把mu的sema队列也拷贝了
	// 总结：永远不要拷贝锁，请重新创建锁
	copyMu := mu

	//// 更加隐蔽的拷贝锁，Person包含锁成员，拷贝Person把锁也拷贝了
	//person1 := Person{salary: 100, level: 2}
	//person2 := person1

	mu.Unlock()

	// 再次上锁，会报deadlock异常
	copyMu.Lock()

	time.Sleep(time.Minute)
}
