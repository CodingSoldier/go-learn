package main

import "fmt"

//
//func main()  {
//	variableZeroValue()
//}
//
//func variableZeroValue(){
//
//	///**
//	//定义变量，变量名在前，类型在后
//	//变量有默认值
//	//*/
//	//var a int  // 默认值是0
//	//var s string  // 默认值是""
//	//fmt.Printf("%d %q", a, s)
//
//
//	//// 变量赋值，同一个类型的多个变量赋值
//	//var a, b int = 3, 4
//	//var s string = "abc"
//	//fmt.Println(a, b, s)
//
//
//	//// 不声明类型，定义变量可以写在一行
//	//var a, b, c, s = 3, 4, true, "def"
//	//fmt.Println(a, b, c, s)
//
//
//	// 使用 := 声明变量，替代 var
//	a, b, c, s := 3, 4, true, "def"
//	// 再次赋值不能用 := 需要用 =
//	b = 5
//	fmt.Println(a, b, c, s)
//
//}

/**
在函数外定义变量是包内变量，不是全局变量
*/
var a = 3
var b = true
var s = "ssss"

//c := 1234  // 函数外不能使用 := 定义变量

// 定义多个变量可以写在括号中
var (
	aa = 22
	bb = true
)

func main() {
	variableZeroValue()
}

func variableZeroValue() {

	fmt.Println(a, b, s)
	fmt.Println(aa, bb)

}
