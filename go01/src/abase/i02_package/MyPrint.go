package i02_package

import "fmt"

func (node Node) MyPrint() {
	fmt.Printf("MyPrint %v \n", node.Value)
}
