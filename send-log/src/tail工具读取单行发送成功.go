package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

//
//// 测试用的文本文件11M大小
//var log1 string = `D:\wls\applogs\access\spring-error.log`
//
//// 测试用的文本文件400M大小
////var m400 string = `G:\runtime\log\ccapi\400M.log`
//
//// 文件一块一块的读取
//func readBlock(filePath string) {
//
//	f, err := os.Open(filePath)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	defer f.Close()
//	// 设置每次读取字节数
//	buffer := make([]byte, 1024)
//	for {
//		n, err := f.Read(buffer)
//		// 控制条件,根据实际调整
//		if err != nil && err != io.EOF {
//			log.Println(err)
//		}
//		if n == 0 {
//			break
//		}
//		// 如下代码打印出每次读取的文件块(字节数)
//		//fmt.Println(string(buffer[:n]))
//	}
//	fmt.Println("readBolck spend : ", time.Now().Sub(start1))
//}
//
//func Match() (chan<- string, <-chan string) {
//	in := make(chan string)
//	out := make(chan string)
//	go func() {
//		for {
//			select {
//			case s := <-in:
//				if strings.Contains(s, "[Error]") {
//					out <- s
//				}
//			}
//		}
//	}()
//	return in, out
//}
//
//func main() {
//	//https: //www.eet-china.com/mp/a60842.html
//
//	for {
//		f, err := os.Open("D:\\wls\\applogs\\access\\spring-error.log")
//		finfo, err := f.Stat()
//		fmt.Println("file size:", finfo.Size())
//		i := finfo.Size()
//		if err != nil {
//			fmt.Println(err)
//		}
//		if finfo.Size() < i {
//			i = 0
//		}
//		if err != nil {
//			fmt.Println(err)
//		}
//		seek, err := f.Seek(i, 0)
//		if err != nil {
//			fmt.Println(seek, err)
//		}
//		rd := bufio.NewReader(f)
//		for {
//			line, err := rd.ReadString('\n')
//			if err != nil || io.EOF == err {
//				if io.EOF != err {
//					fmt.Println(err)
//				}
//				break
//			}
//			fmt.Println("数据：" + line)
//			n := len([]byte(line))
//			i = i + int64(n)
//			time.Sleep(time.Millisecond * 100)
//		}
//		time.Sleep(time.Second)
//		f.Close()
//	}
//
//}

/**
window用户不是adminnistror，gopath是C:\Users\用户名\go
idea的gopath也要设置为C:\Users\用户名\go
进入send-log根目录
go get -u github.com/hpcloud/tail  即可下载依赖
idea会自动导入依赖
*/
func main() {
	fileName := "D:\\wls\\applogs\\access\\spring-error.log"
	config := tail.Config{
		ReOpen: false, // 重新打开
		Follow: true,  // 是否跟随
		//File.Seek(offset, whence)，设置光标的未知
		// offset,偏移量
		// whence，从哪开始：0从头，1当前，2末尾
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)

		// 发送企微信息
		//message.WxSend(line.Text)
	}
}
