package main

import (
	"fmt"
	"reflect"
	"runtime"
)


// go语言可以返回多个值，(q, r int)是返回值，返回值是int类型
func div(a, b int) (q, r int){
	return a / b, a%b
}

// 函数返回值有两个，类型是int、error
func eval(a, b int, op string) (int, error){
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// _ 标识不要此返回值
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("不支持的运算符：%s", op)
	}
}


/**
go语言中函数是一等公民
apply第一个参数是一个函数op func(float64, float64) float64
	函数名是op，参数列表是(float64, float64)，返回值是float64类型
a, b float64是第二、第三个参数
 */
func apply(op func(float64, float64) float64, a, b float64) float64{
	// 获取函数名
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("函数名是：%s ,参数是：(%f, %f)\n", opName, a, b)
	return op(a, b)
}

// 函数可变参数列表
func sum(numbers ...int) int{
	s := 0
	for i := range numbers{
		s += numbers[i]
	}
	return s
}


// swap01(a, b int) 值传递
func swap01(a, b int){
	a, b = b, a
}

// swap02(a, b *int)指针传递
func swap02(a, b *int){
	// 使用*取指针地址
	// 与C++指针最大的不同在于go语言指针不支持运算
	*a, *b = *b, *a
}

func main(){

	//// 函数返回多个值
	//q, r := div(3, 1)
	//fmt.Println(q, r)


	//// 调用函数
	//if i, err := eval(1, 2, "+"); err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println(i)
	//}


	//// 函数参数也是函数
	//r := apply(func(a float64, b float64) float64 {
	//	return math.Pow(a, b)
	//}, 3, 4)
	//fmt.Println(r)


	// 可变参数
	//i := sum(1, 2, 3, 4, 5)
	//fmt.Println(i)



	a1, b1 := 1, 2
	swap01(a1, b1)
	fmt.Println(a1, b1)

	a2, b2 := 3, 4
	// 传递指参数使用&
	swap02(&a2, &b2)
	fmt.Println(a2, b2)




}
