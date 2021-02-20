package main

import (
	"fmt"
)

/**

panic
	停止当前函数执行
	一直向上返回执行每一层defer
	如果没有遇见recover，程序退出

recover
	仅在defer调用中使用
	获取panic的值
	如果无法处理，可重新panic

尽量使用error而不是用panic
意料之中的：使用error。如：文件打不开
意料之外的：使用panic。如：数组越界

*/

func tryRecover() {
	defer func() {
		r := recover()
		if r == nil {
			fmt.Println("没发生错误")
			return
		}
		if err, ok := r.(error); ok {
			fmt.Println("recover error，让程序继续执行。错误为: ", err)
		} else {
			fmt.Println(err)
			panic("不是error，无法处理")
		}
	}()

	//panic(errors.New("制造一个error"))

	// runtime error
	b := 0
	a := 5 / b
	fmt.Println(a / b)

}

func main() {

	tryRecover()

}
