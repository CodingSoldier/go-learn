package main

import "fmt"

/**
函数式编程 VS 函数指针
函数式编程：
	函数是一等公民，参数，变量，返回值都可以是函数
	高阶函数，函数的参数也可以是函数
	闭包

*/

/**
func adder() func(int) int
	定义一个函数adder()，此函数返回值也是一个函数func(int)，
	func(int)是带有一个int类型的参数，func(int)的返回值是int
*/
func adder() func(int) int {
	sum := 0

	/**
	不仅仅是返回一个函数，还把变量sum也返回了
	可以说返回了一个闭包
	*/
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...+%d=%d\n", i, a(i))
	}
}
