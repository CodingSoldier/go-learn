package i02_package

import "fmt"

/**
go语言没有继承，如果扩充类型？
	使用：定义别名 + 使用组合
*/
type CustomNode struct {
	Node *Node
}

func (customNode *CustomNode) CustomPrint() {
	fmt.Println("自定义结构别名+组合实现扩充原有结构")
	customNode.Node.Print()
}
