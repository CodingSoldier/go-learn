package main

import (
	"testing"
)

/**
运行测试用例
1、使用idea运行
2、cd src\abase\m01_test
	go test .   // 使用命令行运行测试用例

查看代码覆盖率
	go tool cover  // 查看命令介绍

*/
func TestTriangle(t *testing.T) {
	// go语言使用表格驱动测试
	// 定义一个struct数组作为一组测试数据
	// a, b作为calcTriangele(a, b int)的参数。c作为calcTriangele(a, b int)的返回结果
	tests := []struct {
		a, b, c int
	}{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	// 遍历表格数据，并调用被测方法
	for _, tt := range tests {
		if actual := calcTriangele(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); got %d; expected %d",
				tt.a, tt.b, actual, tt.c)
		}
	}
}
