package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

/**
访问：http://localhost:8888/list/file.txt
*/
func main() {
	http.HandleFunc("/list/",
		func(writer http.ResponseWriter, request *http.Request) {
			path := request.URL.Path[len("/list/"):]
			file, err := os.Open(path)
			if err != nil {
				// 访问 http://localhost:8888/list/file.txt1
				// 报错，返回错误信息给前端
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return

			}
			defer file.Close()

			all, err := ioutil.ReadAll(file)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			writer.Write(all)
		})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
