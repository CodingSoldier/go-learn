package main

import "fmt"

// s []int 函数参数是一个切片
// 切片是数组的一个视图，是指针传递
func updateSlice(s []int){
	s[0] = 100
}

func main(){

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	//// arr[2:6]半闭半开区间，即包含下标2的值，不包含下标6的值
	//fmt.Println("arr[2:6] = ", arr[2:6])


	//// 使用切片
	//updateSlice(arr[:])
	//fmt.Println(arr)

	//// 对slice再次slice
	//s := arr[0:5]
	//fmt.Println(s)
	//s = s[3:5]
	//fmt.Println(s)


	//s := arr[2:6]
	//fmt.Println(s)
	//// slice仍然能取到slice之后的值
	//// slice可以向后扩展，但不能向前扩展
	//// f_array_slice取值.jpg
	//s = s[3:5]
	//fmt.Println(s)
	//// s[:]仍然取到可见部分，不会取到不可见部分
	//fmt.Println(s[:])


	// f_array_slice取值02.jpg
	// 切片s的capacity比length大
	s := arr[2:6]
	fmt.Printf("s=%v, len(s)=%d, cap(s)=%d \n", s, len(s), cap(s))




}
