package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/**
定义一个类型是 func() int
在go语言中函数是一等公民，所以可以定义类型为函数类型
*/
type intGen func() int

/**
为函数实现接口

	函数返回值是一个接口
*/
func fibonacci02() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

/**
type Reader interface {
	Read(p []byte) (n int, err error)
}
io.Reader接口定义了一个方法，使用intGen实现此方法

Read的调用者是(g intGen)，即函数的调用者也可以是一个函数，函数是一等公民
*/
func (g intGen) Read(p []byte) (n int, err error) {
	// 调用闭包返回值
	next := g()

	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

/**
函数参数是io.Reader接口
*/
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	intGenF := fibonacci02()
	printFileContents(intGenF)
}
