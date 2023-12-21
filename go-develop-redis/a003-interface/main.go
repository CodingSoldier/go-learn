package main

import "fmt"

type Car interface {
	Drive()
}

type TrafficTool interface {
	Drive()
	// TrafficTool加上Run()
	// Truck结构体没实现Run()，则Truck实例无法转为TrafficTool
	Run()
}

type Truck struct {
	Model string
}

/**
定义 func (t Truck) Drive() {} 后，
go编译的时候会自动加上结构体指针的Drive()
func (t *Truck) Drive() {}
所以可以用 结构体 以及 结构体指针 初始化变量
var a Car = &Truck{}
var a Car = Truck{}

但是如果定义了 func (t *Truck) Drive() {}
go编译器不会自动加上结构体的Drive()
没加：func (t Truck) Drive() {}
所以只能只用 结构体指针 初始化变量
var a Car = &Truck{}
不能使用 结构体 初始化变量
var a Car = Truck{}
*/
func (t *Truck) Drive() {
	fmt.Println(t.Model)
}

// 空接口
type K interface {
}

// 函数的入参是空接口，则函数的参数可以是任何类型
func Print01(interface{}) {
}

func main() {
	var a Car = &Truck{}
	fmt.Println(a)
	//var c Car = Truck{}
	//fmt.Println(c)
	//
	//// 使用类型断言进行弱转换，把Car类型转为Truck类型
	//t := c.(Truck)
	//fmt.Println(t.Model)
	//
	//// Truck实体也实现TrafficTool接口的所有方法
	//// 所以Truck实体也能转换为TrafficTool类型
	//tt := c.(TrafficTool)
	//fmt.Println(tt)
	//switch c.(type) {
	//case TrafficTool:
	//	fmt.Println("tt是TrafficTool类型")
	//}

	Print01("sdfd")
	Print01(123)

}
