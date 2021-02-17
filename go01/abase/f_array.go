package main

import "fmt"

/**
函数参数传递数组是值传递
[10]int 和 [20]int 是不同类型
调用func f(arr [10]int)会拷贝数组
 */
func printArray(arr [5]int){
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 函数参数是指针传递
// go语言一般不使用数组作为参数，而是使用slice
func printArray02(arr *[5]int){
	arr[0] = 100 // 函数内数组不用使用*
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main(){

	// 定义数组的方法
	// arr1是长度为5的数组
	var arr1 [5]int
	// arr2长度为3并且赋初值
	arr2 := [3]int{1, 3, 5}
	// arr3让编译器自动识别长度，并赋初值
	arr3 := [...]int{2, 4, 6, 8, 10}
	// 4行5列数组
	var grid[4][5]int

	fmt.Println(arr1, arr2, arr3, grid)


	//// 使用 range arr 遍历数组，可同时得到下标与值
	//for i, v := range arr2 {
	//	fmt.Printf("下标：%d，值：%d \n", i, v)
	//}


	//// 数组是值传递
	//printArray(arr3)
	////printArray(arr2)  // 编译不通过，[3]int和[5]int不是同一种类型
	//fmt.Println(arr3)


	// 传递指针
	printArray02(&arr3)
	fmt.Println(arr3)



}
