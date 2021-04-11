package main

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
)

/**
go module使用

go mod 命令	说明
download	download modules to local cache(下载依赖包)
edit	edit go.mod from tools or scripts（编辑go.mod)
graph	print module requirement graph (打印模块依赖图)
verify	initialize new module in current directory（在当前目录初始化mod）
tidy	add missing and remove unused modules(拉取缺少的模块，移除不用的模块)
vendor	make vendored copy of dependencies(将依赖复制到vendor下)
verify	verify dependencies have expected content (验证依赖是否正确）
why	explain why packages or modules are needed(解释为什么需要依赖)


go mod init 工程名  // 初始化工程
1、idea新建project，选择go modules。
	不能在empty project中建，必须单独作为一个project

	// 使用命令行新建go module的方式
		新建一个目录module02
		在module02目录下执行 go mod init module02
		便会生成go.mod文件
		go.sum是对go.mod的补充
		go.mod、go.sum都要放到代码仓库

2、mudule项目会生成一个go.mod文件。go.mod配置了module的名称
3、E:\go-learn\module01> 目录下执行 go get github.com/hashicorp/golang-lru 下载依赖
	依赖放置在GOPATH的\pkg\mod\cache\download目录下
	C:\Users\Administrator\go\pkg\mod\cache\download\github.com
4、在main函数中可直接使用依赖库的方法 lru.New(128)

go.mod定义了
1、module的名字
2、go版本
3、依赖

在module01目录下执行
go build
便会生成 module01.exe

go mod graph  // 打印模块依赖图

go mod download  // 下载依赖
删除 C:\Users\Administrator\go\pkg\mod\github.com 下的 hashicorp
再次执行
go mod download
idea的setting -> Go -> GOPATH 必须配置 GOPATH 为 C:\Users\Administrator\go

go mod tidy     // 拉取缺少的模块，移除不用的模块
注释掉main函数中的所有代码，仅保留
	fmt.Println("---->")
go.mod中依旧有hashicorp依赖
执行
go mod tidy
go.mod中的hashicorp依赖被删除

如果 lru.New(128) 变红，则执行
go get github.com/hashicorp/golang-lru

go mod verify   // 验证依赖是否正确
修改go.mod的golang-lru 版本为 v0.5.5
执行
go mod verify
提示版本未找到

go mod why    // 解释为什么需要依赖
go mod why -m github.com/hashicorp/golang-lru


go mod vendor   // 将依赖复制到vendor下，一般在DevOps中使用
执行
go mod vendor
module01目录下多了一个vendor目录，vendor目录存放了依赖
删除GOPATH下的hashicorp依赖
执行
go build
依旧能执行成功


go.mod
排除依赖
exclude github.com/hashicorp/golang-lru v0.5.3
执行，注意是使用@
go get github.com/hashicorp/golang-lru@v0.5.3
提示依赖被排除


删除vendor后执行
go list -m all   // 列出所有依赖


*/
func main() {

	cache, _ := lru.New(128)
	for i := 0; i < 256; i++ {
		cache.Add(1, nil)
	}
	fmt.Println("---->", cache.Len())

	fmt.Println("---->")

}
