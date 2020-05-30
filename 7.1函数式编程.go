package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func adder() func(int) int {
	// 自由变量
	sum := 0
	return func(v int) int {
		// v局部变量
		sum += v
		return sum
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 为斐波那契生成器时间reader接口
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func fibonacci2() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	// 如果p过小的话结果会异常
	return strings.NewReader(s).Read(p)
}

func main() {
	// a内部的sum是同一个 计算结果的时候会累加
	//a := adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(a(i))
	//}

	// 斐波那契数列
	//f := fibonacci()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

	f2 := fibonacci2()
	printFileContents(f2)
}
