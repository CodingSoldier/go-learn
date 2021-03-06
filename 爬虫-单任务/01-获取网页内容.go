package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	/**
	获取珍爱网征婚页内容
	*/
	request, err := http.NewRequest(http.MethodGet, "https://www.zhenai.com/zhenghun", nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")

	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// 判断状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("resp.StatusCode ==", resp.StatusCode)
	}

	// 打印内容
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}
