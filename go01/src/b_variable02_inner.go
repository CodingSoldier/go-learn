/**
内荐变量类型
boolean, string等

(u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
加u表示有符号类型，不加u表示无符号类型
int不规定长度，由操作系统决定
uintptr 是指针

byte

rune 就是go语言的char类型，是32位（4字节）

float32, float64

complex64, complex128 是复数，有实部虚部
复数基础：i=根号下-1
复数例子：3 + 4i ，3是实部，4是虚部
详情：b_variable02_inner_复数.go.jpg

*/
package main

import "fmt"

//func main()  {
//
//	//// 复数绝对值
//	//c := 3 + 4i
//	//fmt.Println(cmplx.Abs(c))
//
//	// 强制类型转换 math.Sqrt()返回float64类型
//	var a, b int = 3, 4
//	var c int
//	c = int(math.Sqrt(float64(a*a + b*b)))
//	fmt.Println(c)
//}

//// 常量当然可以定义为在包内部
//const filename = "abc.txt"
//
//func main()  {
//
//	/**
//	const定义常量，没写类型，则常量可以是任意类型
//	a*a + b*b 会被自动识别为float64类型
//
//	golang的常量不需要大写
//	 */
//	const a, b = 3, 4
//	//a = 5  常量不允许修改
//	var c int
//	c = int(math.Sqrt(a*a + b*b))
//	fmt.Println(c)
//	fmt.Println(filename)
//}

//定义枚举类型
func enums() {

	//const (
	//	cpp = 0
	//	java = 1
	//	python = 2
	//	javascript = 3
	//)
	//fmt.Println(cpp, java, python, javascript)

	//使用iota迭代自增
	//const (
	//	cpp = iota
	//	java
	//	python
	//	javascript
	//)
	//fmt.Println(cpp, java, python, javascript)

	// 使用 _ 跳过一个迭代值
	const (
		cpp = iota
		_
		python
		javascript
	)
	fmt.Println(cpp, python, javascript)

}

func main() {
	enums()
}
