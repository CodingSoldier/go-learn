package a01

type People struct {
	name  string
	title string
}

/**
类方法
p是receiver（接收者），表示walk()是People结构体的方法，即类方法
*/
func (p People) walk() {

}

// 纯粹的函数，不是类方法，没有(receiver XX)
func onlyFunc() {

}

/**
继承：
go没有继承，但可以使用语法糖模拟继承
Man结构体可执行
	m := Man{}
	m.walk()  // 底层是这样：m.People.walk()

看起来就像是Man结构体继承了People

Go的继承只是组合，People是匿名字段（只有类型，没有名称），通过这种语法糖模拟继承
*/
type Man struct {
	People
}

/**
接口
struct并不是显示实现接口，而是隐式实现接口
显示实现，使用java举例：
public class People implement Live {

}
People 显示实现 Live接口

定义Live接口，定义walk()方法
只要struct实现了Live的所有方法，那么struct就实现了Live接口，这是隐式实现
People实现了Live的所有方法，所以People实现了Live接口
*/
type Live interface {
	walk()
}
