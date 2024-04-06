package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var body [4]byte
	addr := conn.RemoteAddr()
	for {
		_, err := conn.Read(body[:])
		if err != nil {
			break
		}
		fmt.Printf("收到%s的消息：%s\n", addr, string(body[:]))
		_, err = conn.Write(body[:])
		if err != nil {
			break
		}
		fmt.Printf("发送给%s，消息：%s\n", addr, string(body[:]))
	}
	fmt.Printf("与%s断开连接", addr)
}
