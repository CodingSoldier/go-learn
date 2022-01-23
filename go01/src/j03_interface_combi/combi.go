package main

import "fmt"

type Reader interface {
	Read(path string) string
}

type Writer interface {
	Write(path string) bool
}

/**
接口组合
*/
type ReaderWriter interface {
	Reader
	Writer
}

type ReaderWriterImpl struct {
}

/**
实现者方法最好统一，要么全部是指针接受者，要么全部是值接受者
*/
func (rw *ReaderWriterImpl) Read(path string) string {
	return "返回内容"
}

func (rw *ReaderWriterImpl) Write(path string) bool {
	return true
}

func main() {
	rw := ReaderWriterImpl{}
	read := rw.Read("www")
	fmt.Println(read)

	write := rw.Write("www")
	fmt.Println(write)

}
