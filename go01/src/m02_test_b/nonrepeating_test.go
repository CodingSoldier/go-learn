package main

import "testing"

/**
性能测试

命令行执行方式
cd src\abase\m02_test_b
go test -bench .

go test -bench . -cpuprofile cpq.out

下载graphviz
	http://www.graphviz.org/download/
	下载zip包，EXE文件被360认为是病毒
	path添加路径 C:\mysoftware\Graphviz\bin

go tool pprof cpq.out
web
使用浏览器打开svg文件


*/
func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d ", actual, s, ans)
		}
	}
}
