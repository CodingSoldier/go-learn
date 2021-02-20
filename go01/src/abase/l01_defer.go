package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	/**
	defer：确保调用在函数结束时发生
	defer调用是一个栈，先进后出，打印顺序是 3 2 1
	*/
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	//panic("报错")
	fmt.Println("不会打印")
}

func deferParam() {
	for i := 0; i < 100; i++ {
		/**
		参数在defer语句时计算
			虽然defer是在最后调用，但参数不是，打印结果是0到30
		*/
		defer fmt.Println(i)
		if i == 30 {
			panic("报错")
		}
	}
}

/**
defer在流中的应用
*/
func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// 关闭文件流
	defer file.Close()

	writer := bufio.NewWriter(file)
	// flush缓冲区
	defer writer.Flush()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, i)
	}
}

/**
单个错误处理
*/
func errorManage(filename string) {
	// os.OpenFile方法注释中有写，如果出错，则错误是*PathError
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		// 判断error是否是一个 *os.PathError
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
	}
	defer file.Close()
}

func main() {

	//tryDefer()

	//deferParam()

	//writeFile("l01_defer_输出结果.txt")

	errorManage("l01_defer_错误处理.txt")

}
