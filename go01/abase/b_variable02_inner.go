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

import (
	"fmt"
	"math/cmplx"
)

func main()  {

	// 复数绝对值
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

}


