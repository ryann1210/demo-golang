package main

import "fmt"

func main() {
	tryDefer()
}

func tryDefer() {
	// 1 defer定义的逻辑在方法结束时候执行
	// 2 defer结构为栈 先进后出 所以最终输出3 2 1
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	panic("error occurred.")
	fmt.Println(4)
}
