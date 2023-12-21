package main

import "fmt"
import "github.com/CodingSoldier/go-develop-redis-lib/a01"

func main() {
	fmt.Println("hello world")
	m := a01.Man{}
	fmt.Println(m)

}

/**
结构体：
type People struct {
	name  string
	title string
}

类方法
p是receiver（接收者），表示walk()是People结构体的方法，即类方法
func (p People) walk() {
}

纯粹的函数，不是类方法，没有(receiver XX)
func onlyFunc() {
}

继承：
go没有继承，但可以使用语法糖模拟继承
Man结构体可执行
	m := Man{}
	m.walk()  // 底层是这样：m.People.walk()

看起来就像是Man结构体继承了People

Go的继承只是组合，People是匿名字段（只有类型，没有名称），通过这种语法糖模拟继承
type Man struct {
	People
}

接口
struct并不是显示实现接口，而是隐式实现接口
显示实现，使用java举例：
public class People implement Live {
}
People 显示实现 Live接口

定义Live接口，定义walk()方法
只要struct实现了Live的所有方法，那么struct就实现了Live接口，这是隐式实现
People实现了Live的所有方法，所以People实现了Live接口
type Live interface {
	walk()
}

*/

/**
Go Modules
1、Java源代码编译为class字节码文件，字节码文件是一种中间语言，java项目想引用他人的项目只需要引入他人的class文件合集（jar包）即可。
	但是go没有中间语言，要引用他人的项目，只能是引用他人项目的源代码。
2、gomod的作用是将你想引用的GO项目和git项目联系起来，引用的版本就是git的tag
3、gomod就是解决“需要用哪个git项目的哪个版本”

使用gomod
1、新建go.mod，填写module名称、go版本，如下：
module a01

go 1.17

2、在a01目录下执行 go get git项目@tag ，即可通过gomod使用git项目。例如：
go get github.com/Jeffail/tunny@0.1.3

使用本地文件替换网络下载的方式，可在go.mod追加
replace github.com/Jeffail/tunny@0.1.3 => /usr/local/lib/本地文件的地址

可以使用go vendor将git项目缓存到本地
go build -mod verdor

go modules命名，使用任意git仓库地址/用户账号/项目名
module github.com/CodingSoldier/go-develop-redis-lib

go 1.17

自己创建一个go项目推送到git让别人使用
1、创建go-develop-redis-lib仓库，编写代码
go.mod为：
module github.com/CodingSoldier/go-develop-redis-lib

go 1.17

a01.go为如下（注意不能有main.go，也不能有main包，main方法）：
package a01

type People struct {
	name  string
	title string
}

func (p People) walk() {
}

推送到github，打上tag

本项目执行：
go get github.com/CodingSoldier/go-develop-redis-lib@v1.0.3

导入依赖
import "github.com/CodingSoldier/go-develop-redis-lib/a01"
就可以使用a01的代码了：
	m := a01.Man{}


*/
