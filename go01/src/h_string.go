package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "Yes我爱慕课网！"

	//for _, b := range []byte(s){
	//	// 使用UTF-8编码打印
	//	fmt.Printf("%X", b)
	//}

	//fmt.Println()
	//for i, ch := range s{
	//	// ch是int32类型，是一个4字节的整数
	//	// ch就是一个rune
	//
	//	// 转为Unicode编码
	//	fmt.Printf("(%d, %X)", i, ch)
	//}

	//// []byte(字符串)获取字节
	//bytes := []byte(s)
	//for len(bytes) > 0 {
	//	// byte转utf8
	//	ch, size := utf8.DecodeRune(bytes)
	//	bytes = bytes[size:]
	//	fmt.Printf("%c ", ch)
	//}

	// 获取字节长度
	fmt.Println(len(s))
	// 获取字符数量
	fmt.Println(utf8.RuneCountInString(s))

	//// 将字符串s转为rune数组，遍历出来的ch是字符
	//for i, ch := range []rune(s){
	//	fmt.Printf("(%d, %c)", i, ch)
	//}

	/**
	运行配置
		Run kind：File
		Run after build
	*/
	str := `sfadf
		"带双引号"
122334`
	fmt.Println(str)

}
