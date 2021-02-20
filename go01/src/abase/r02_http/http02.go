package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/**
统一异常处理
*/

/**
将 http.HandleFunc(pattern string, handler func(ResponseWriter, *Request)) handler的具体实现提取出来，并返回error，以便能统一处理error
*/
func appHandlerFile(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err

	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)

	return nil
}

/**
定义一个类型为函数，作为errWrapper这个包装函数的入参
*/
type appHandler func(writer http.ResponseWriter, request *http.Request) error

/**
定义包装函数，处理error
包赚函数的返回值是 http.HandleFunc 的第二个参数
*/
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		// recover可能出现的panic
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("【异常】recover panic %s \n", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		code := http.StatusOK
		if err != nil {
			fmt.Printf("【异常】%s \n", err.Error())
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

/**
访问：http://localhost:8888/list/file.txt
*/
func main() {
	http.HandleFunc("/", errWrapper(appHandlerFile))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
