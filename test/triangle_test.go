package main

import (
	"math"
	"testing"
)

// 文件名必须_test结尾 方法名必须Test开头
// 这样才会有可执行方法箭头
// 也可以在Terminal输入go test .
func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	for _, test := range tests {
		if actual := calcTriangle(test.a, test.b); actual != test.c {
			t.Errorf("calcTriangle(%d, %d);got %d;expected %d", test.a, test.b, actual, test.c)
		}
	}
}

func calcTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}
