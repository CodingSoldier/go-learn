package httpimpl

import (
	"net/http"
	"net/http/httputil"
	"time"
)

/**
一个真实的实现类
	结构仅仅是具备Get(url string)方法，与Interface01再无任何关系

也可以说实现者不需要实现接口，而是实现接口的方法即可，
即实现者在行为表现上想一只鸭子，虽然它不是传统观念中真正意义上的鸭子
*/
type HttpClient struct {
	UserAgent string
	TimeOut   time.Duration
}

func (httpClient HttpClient) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}
