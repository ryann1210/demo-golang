package main

import (
	"fmt"
	"reflect"
)

func main() {
	a, b := 1, 2
	fmt.Println(reflect.TypeOf(a))   // int
	fmt.Println(&a)                  // 十六进制内存地址 0xc00000a0b8
	fmt.Println(reflect.TypeOf(&a))  // *int 指针类型
	fmt.Println(reflect.TypeOf(*&a)) // 指针前面加* 得到指针对应的值

	fmt.Println("方法外 起始值", a, b) // 1 2
	valSwap(a, b)
	fmt.Println("方法外 值传递结束", a, b) // 1 2

	a, b = 1, 2
	fmt.Println("方法外 起始值", a, b) // 1 2
	refSwap(&a, &b)
	fmt.Println("方法外 指针传递结束", a, b) // 2 1
}

// 值传递对原本的a b不影响
func valSwap(a, b int) {
	fmt.Println("方法内 值传递开始: ", a, b) // 1 2
	a, b = b, a
	fmt.Println("方法内 值传递结束: ", a, b) // 2 1
}

// 指针传递修改了对应的值
func refSwap(a, b *int) {
	fmt.Println("方法内 指针传递开始: ", a, b) // 0xc00000a0b8 0xc00000a0d0
	*a, *b = *b, *a
	fmt.Println("方法内 指针传递结束: ", a, b) // 0xc00000a0b8 0xc00000a0d0
}
