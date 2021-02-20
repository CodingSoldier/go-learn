package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (ue userError) Message() string {
	return string(ue)
}

/**
系统类型error是一个接口，userError只要实现了 Error() string 方法，就是error类型
type error interface {
	Error() string
}
*/
func (ue userError) Error() string {
	return ue.Message()
}

const prefix = "/list/"

func appHandlerFile(writer http.ResponseWriter, request *http.Request) error {

	// 判断url是否以 /list/开头，不以/list/开头就返回userError
	if strings.Index(request.URL.Path, prefix) != 0 {
		msg := fmt.Sprintf("path %s must start "+"with %s", request.URL.Path, prefix)
		// 由于userError是string类型，所以使用userError(msg)创建userError对象
		return userError(msg)
	}

	path := request.URL.Path[len(prefix):]
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

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

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

			// 如果是userError则返回userError的异常信息
			if ue, ok := err.(userError); ok {
				http.Error(writer, ue.Message(), http.StatusBadRequest)
				return
			}

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
