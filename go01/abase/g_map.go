package main

import (
	"fmt"
)

func main() {

	// 创建map并赋值
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// make创建空map
	m2 := make(map[string]int)

	// var定义map
	var m3 map[string]int

	fmt.Println(m2, m3)

	courseName := m["course"]
	fmt.Println("获取map的值", courseName)

	notkey := m["notkey"]
	fmt.Println("获取map的值", notkey)

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key不存在")
	}

	fmt.Println("删除map元素")
	delete(m, "name")
	name, ok := m["name"]
	fmt.Println("name, ok", name, ok)

	/**
	map使用哈希表，必须可以比较相等
	除了slice，map，function的内建类型都可以作为key
	Struct类型不包括上述字段，也可以作为key
	*/

}
