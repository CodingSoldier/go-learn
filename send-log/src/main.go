package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

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
