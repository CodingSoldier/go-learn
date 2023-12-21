package main

import (
	"fmt"
)

func main() {

	// nil在builtin.go中，
	//nil是一个变量，是 pointer, channel, func, interface, map, or slice type 的0值
	var a *int
	// 此时，a是空指针
	fmt.Println(a == nil)

	// true
	var b map[int]int
	fmt.Println(b == nil)

	// a和b不能比较，虽然a、b都是nil，但是他们的类型不相等，由此推导出nil还具备类型
	//fmt.Println( a == b)

	////// 空结构体不是nil
	//var s struct{}
	//fmt.Println(s == nil)

	// i是空接口，i是nil
	// ii是空指针，ii是nil
	// i = ii 之后，i不是nil
	var i interface{}
	fmt.Println(i == nil)
	var ii *int
	fmt.Println(ii == nil)
	i = ii
	fmt.Println(i == nil)

}
