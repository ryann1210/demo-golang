package main

import (
	"fmt"
)

func main() {
	tryRecover()
}

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred: ", err)
		} else {
			panic(fmt.Sprintf("i dont know what to do: %v", r))
		}
	}()

	// 可以捕获异常处理
	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	// 不是error类型 捕获后继续panic
	panic(123)
}
