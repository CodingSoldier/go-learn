package main

/**
GOPATH环境变量
	unix/linux默认在 ~/go 目录
	windows默认在 C:\Users\用户名\go。使用 go env 命令行可以查看

官方推荐：所有项目和第三方库都放在一个GOPATH下
也可以将每个项目放在不同的GOPATH下

使用 go get 需要先设置代理
使用7牛云的代理：https://github.com/goproxy/goproxy.cn
打开PowerShell执行以下两句
	$env:GO111MODULE = "on"
	$env:GOPROXY = "https://goproxy.cn"
再执行go get
	go get golang.org/x/tools/cmd/godoc





*/
