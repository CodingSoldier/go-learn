package main

import "fmt"

/**
go仅支持封装，不支持继承和多态
*/

/**
定义一个结构体
*/
type treeNode struct {
	value       int
	left, right *treeNode
}

/**
go语言没有构造函数，使用工厂方法替代构造函数
结构体创建在堆上还是栈上？答：不需要知道。
*/
func createTreeNode(value int) *treeNode {
	/**
	对于C++来说，本函数体内的treeNode是局部变量，对象在栈中创建
	go语言很智能，如果函数体内的局部变量只在函数体中使用，则在栈中创建对象
	如果函数中的局部变量的地址作为函数返回值，则在堆上创建对象
	*/
	return &treeNode{value: value}
}

/**
为结构体定义方法，使用方式 treeNode实例.print()
显式定义和命名方法接受者
*/
func (node treeNode) print() {
	fmt.Println(node.value)
}

/**
结构体方法默认也是值传递，不是引用传递，使用指针则是指针传递
只有使用指针才可以改变结构体内容
*/
func (node *treeNode) setValue(value int) {
	// nil指针也可以调用方法，但是 nil.属性 则会报错。所以多加了一行判断
	if node == nil {
		fmt.Println("不能给nil设置值")
		return
	}
	node.value = value
}

func main() {
	// 定义/创建treeNode的几种方式
	var root treeNode

	//root = treeNode{value: 3}
	//root.left = &treeNode{}
	//root.right = &treeNode{5, nil, nil}
	//root.right.left = new(treeNode)
	//
	//// 使用工厂方法创建对象
	//root.left.right = createTreeNode(2)
	//fmt.Println(root.left.right)
	//
	//// 创建treeNode数组
	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, nil},
	//}
	//fmt.Println(nodes)

	// 调用结构体方法
	root.print()
	root.setValue(11)
	root.print()

	// root.left是nil，也可以调用方法
	root.left.setValue(22)

	/**
	  值接受者 VS 指针接受者
	  	1、要改变内容必须使用指针接受者
	  	2、结构过大也考虑使用指针接受者
	  	3、一致性：如有指针接受者，最好都是指针接受者
	*/

}
