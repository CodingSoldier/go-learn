# go-learn
golang学习

安装golang环境
    1、https://studygolang.com/dl
    2、下载window版本（msi文件）安装
    3、idea安装go插件（点击Marketplace安装）
    4、在go-learn目录下新建go01项目

gland导入 go-learn
    1、Project Structrue -> go01配置为Resource Root
    2、运行 j01_interface01.go，run kid 为 Package
    3、Packege path 为 go01/src/j01_interface


新建go01项目
1、新建go01项目
go01
go.mod
go1.go
2、Setting -> GOPATH
Project GOPATH 不添加任何目录
Use GOPATH that's defined 不勾选
Module GOPATH 不添加任何目录
3、运行 go1.go

新建go02-gin
1、新建 go02-gin
2、进入 go02-gin
go mod init go02-gin
3、下载依赖
go get -u github.com/gin-gonic/gin
4、新建go02gin.go
输入gin，idea会自动导入依赖
5、运行go02gin.go
6、http://localhost:8000/ok 能得到响应
