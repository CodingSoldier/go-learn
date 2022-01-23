package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	/**
	`json:"name"`是定义序列化key的别名
	在go中叫做tag
	*/
	ServerName string `json:"name"`
	ServerIp   string
	ServerPort int
}

// 序列化结构体
func SerializeStruct() {
	server := new(Server)
	server.ServerName = "名字"
	server.ServerIp = "192.168.1.1"
	server.ServerPort = 8080

	// 序列化成json字节数组
	b, err := json.Marshal(server)
	if err != nil {
		fmt.Println("序列化错误：", err.Error())
		return
	}
	fmt.Println("序列化后转为string", string(b))
}

// 序列化map
func SerializeMap() {
	server := make(map[string]interface{})

	server["ServerName"] = "名字"
	server["ServerIp"] = "192.168.0.0"
	server["ServerPort"] = 9090

	// 序列化成json字节数组
	b, err := json.Marshal(server)
	if err != nil {
		fmt.Println("序列化错误：", err.Error())
		return
	}
	fmt.Println("序列化后转为string", string(b))
}

// 反序列化
func UnSerializeStruct() {
	// key的名称要和struct的tag相同
	jsonString := `{"name":"名字","ServerIp":"192.168.1.1","ServerPort":8080}`
	server := new(Server)
	err := json.Unmarshal([]byte(jsonString), &server)
	if err != nil {
		fmt.Println("反序列化异常：", err)
	}
	fmt.Println("反序列化结果", server)
}

// 反序列化
func UnSerializeMap() {
	// key的名称要和struct的tag相同
	jsonString := `{"name":"名字","ServerIp":"192.168.1.1","ServerPort":8080}`
	map01 := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &map01)
	if err != nil {
		fmt.Println("反序列化异常：", err)
	}
	fmt.Println("反序列化结果", map01)
}

func main() {

	// 序列化
	//SerializeStruct()

	// 序列化
	//SerializeMap()

	// 反序列化
	//UnSerializeStruct()

	// 反序列化
	UnSerializeMap()

}
