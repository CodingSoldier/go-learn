package main

import (
	"fmt"
)

type K struct {
}

type F struct {
	num1 K
	num2 int32
}

func main() {
	///**
	//查看基本数据类型的大小，
	//int类型是8字节，由于我电脑是64位的，所以是8字节；如果是32位系统，int则是4字节
	//*/
	//fmt.Println(unsafe.Sizeof(int(0)))
	//fmt.Println(unsafe.Sizeof(int64(0)))
	//// int32是4字节
	//fmt.Println(unsafe.Sizeof(int32(0)))
	//
	//// 指针在64位系统也是8字节
	//i := int(0)
	//p := &i
	//fmt.Println(unsafe.Sizeof(p))

	//a := K{}
	//b := int(0)
	//c := K{}
	//// 独立空结构体大小也是0
	//fmt.Println(unsafe.Sizeof(a))
	//// 独立空结构体的指针/地址是一样的
	//fmt.Printf("%p\n", &a)
	//fmt.Printf("%p\n", &b)
	//fmt.Printf("%p\n", &c)
	//fmt.Println("##############")
	//f := F{}
	//// 非独立空结构体大小不是0
	//fmt.Println(unsafe.Sizeof(f))
	//// 非独立空结构体的指针地址不是空结构体指针地址常量
	//fmt.Printf("%p\n", &f)
	//fmt.Println(unsafe.Sizeof(f.num1))
	//fmt.Printf("%p\n", &f.num1)
	///**
	//空结构体：
	//	没有任何属性的结构体称为空结构体
	//	空结构体的地址均相同（前提：空结构体本身不是其他结构体的属性成员）
	//	空结构体是为了节省内存
	//		使用map加空结构体模拟hashset
	//		m := map[string]struct{}{}
	//		m["a"] = struct{}{}
	//
	//		使用chanel发送空结构体，即只发送信号，但没有内容
	//		a:make(chan struct{})
	//*/

	//// 字符串的大小都是16字节，即64位，跟系统是多少位有关
	//fmt.Println(unsafe.Sizeof("慕课网"))
	//fmt.Println(unsafe.Sizeof("慕课网Moody老师"))
	///**
	//type stringStruct struct {
	//	str unsafe.Pointer
	//	len int
	//}
	//字符串本质是个结构体
	//Data指针指向底层Byte数组，字符串底层是数组
	//*/

	// len表示字符串字节数，单个中文3字节+3
	a := "中文abc"
	fmt.Println(len(a))

}
