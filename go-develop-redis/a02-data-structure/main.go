package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/**
	查看基本数据类型的大小，
	int类型是8字节，由于我电脑是64位的，所以是8字节；如果是32位系统，int则是4字节
	*/
	fmt.Println(unsafe.Sizeof(int(0)))
	fmt.Println(unsafe.Sizeof(int64(0)))
	// int32是4字节
	fmt.Println(unsafe.Sizeof(int32(0)))

	// 指针在64位系统也是8字节
	i := int(0)
	p := &i
	fmt.Println(unsafe.Sizeof(p))
}
