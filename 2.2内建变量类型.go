package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	// 1. 验证欧拉公式
	euler()
	// 2. 强制类型转换
	triangle()
}

func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

func triangle() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) // 没有float64 int两次强转会报错，不会隐式转换
	fmt.Println(c)                         // 5
}
