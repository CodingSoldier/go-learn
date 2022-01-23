package main

import "go01/src/i02_package"

/**
entry.go作为程序入口，包名是main，包含main方法
*/
func main() {

	// 类型是 包名.结构名
	var root i02_package.Node
	root.SetValue(111)
	root.Print()

	// 结构的方法可以在同一个包的不同文件中
	root.MyPrint()

	// 使用扩充结构
	customRoot := i02_package.CustomNode{&root}
	customRoot.CustomPrint()
}
