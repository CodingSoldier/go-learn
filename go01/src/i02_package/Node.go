package i02_package

import "fmt"

/**

需要设置项目的路径，貌似不需要
setting -> Languages & Frameworks -> Go -> GOPATH -> Project GOPATH -> go01

为结构定义的方法必须在同一个包内，方法可以在同一个包内的不同文件中

go语言没有继承，如果扩充类型
	定义别名 + 使用组合
*/
type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	// nil指针也可以调用方法，但是 nil.属性 则会报错。所以多加了一行判断
	if node == nil {
		fmt.Println("不能给nil设置值")
		return
	}
	node.Value = value
}
