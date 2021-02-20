package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n",
		s, len(s), cap(s))
}

func main() {

	/**
	定义一个slice： var s []int
		s的初始值为nil，go语言没有null，而是用nil
	nil可以直接使用append，不会报空指针异常
	*/
	//var s []int
	//for i := 0; i < 100; i++ {
	//	printSlice(s)
	//	s = append(s, 2*i+1)
	//}

	// 创建slice，并赋初值
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	// 创建slice并指定len，cap
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("拷贝slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("删除元素")
	// 将s2下标为4后面的元素添加到s2下标0到下标2的slice中去，达到删除下标3元素的目的
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("弹出前面的元素")
	// 取出第0个元素
	front := s2[0]
	// s2 = s2下标1后面的元素
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("弹出末尾元素")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

}
