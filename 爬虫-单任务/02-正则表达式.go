package main

import (
	"fmt"
	"regexp"
)

const text = `
test bac
abc@qq.com adfs
sss
email = 123@gmail.com.cn
`

func main() {
	// 正则表达式
	// ``里面的内容不会被转译。使用""这\.要写成\\.
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`)
	// 查找正则匹配所有结果
	match := re.FindAllString(text, -1)
	fmt.Println(match)
	fmt.Println()
	// 查找子匹配
	submatch := re.FindAllStringSubmatch(text, -1)
	for _, m := range submatch {
		fmt.Println(m)
	}
}
