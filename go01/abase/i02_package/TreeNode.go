package i02_package

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node TreeNode) print() {
	fmt.Println(node.Value)
}

func (node *TreeNode) setValue(value int) {
	// nil指针也可以调用方法，但是 nil.属性 则会报错。所以多加了一行判断
	if node == nil {
		fmt.Println("不能给nil设置值")
		return
	}
	node.Value = value
}
