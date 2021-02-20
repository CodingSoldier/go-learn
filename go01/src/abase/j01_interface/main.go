package main

import (
	"abase/j01_interface/httpimpl"
	"fmt"
)

/**
定义接口
*/
type Interface01 interface {
	Get(url string) string
}

/**
定义实现，实现只需要具备Get(url string)方法即可
与java不同，实现结构的代码没出现Interface01这几个字符
这是一个假的实现
*/
type Implement01 struct {
	Contents string
}

func (r Implement01) Get(url string) string {
	return r.Contents
}

/**
使用接口与实现
*/
func downloadBaidu(i Interface01) string {
	return i.Get("https://www.baidu.com")
}

/**
duck typing 鸭子类型
	描述事物的外部行为而非内部结构

严格来说go属于结构化类型系统，类似duck typing

*/
func main() {

	//r := Implement01{"伪造一个网页内容"}
	//str := downloadBaidu(r)
	//fmt.Println(str)

	//r2 := httpimpl.HttpClient{}
	//// downloadBaidu参数是值类型，也可以传指针
	//resp := downloadBaidu(&r2)
	//fmt.Println(resp)

	var inter Interface01
	r3 := Implement01{"001"}
	inter = &r3
	// 判断接口实现者的类型
	if impl01, ok := inter.(*Implement01); ok {
		fmt.Println(impl01.Contents)
	}

	inter = &httpimpl.HttpClient{}
	// 判断接口实现者的类型
	if httpimpl, ok := inter.(*httpimpl.HttpClient); ok {
		fmt.Println(httpimpl.TimeOut)
	}

	// // 使用值类型与使用指针类型，inspect比较
	//inter = httpimpl.HttpClient{}
	//if httpimpl, ok := inter.(httpimpl.HttpClient); ok {
	//	fmt.Println(httpimpl.TimeOut)
	//}

	// 通过switch判断接口实现者类型
	inspect(inter)

	/**
	接口变量包含：
		接口实现者的类型（type）
		接口实现者的值（或者指针）
	*/

}

// 通过switch判断接口实现者类型
func inspect(inter Interface01) {
	/**
	接口变量自带指针
	接口变量同样采用值传递，几乎不需要使用接口的指针
	指针接受者实现只能以指针方式使用，值接受者都可以
	*/
	fmt.Printf("%T %v\n", inter, inter)

	switch v := inter.(type) {
	case *Implement01:
		fmt.Println(v.Contents)
	case *httpimpl.HttpClient:
		fmt.Println(v.TimeOut)
	}
}
