package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println(div1(13, 3)) // 4 1

	// _表示不需要接受该参数
	q, _ := div2(13, 3)
	fmt.Println(q) // 4 1

	fmt.Println(div3(13, 3)) // 4 1

	if result, err := math1(13, 2, "asdf"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(
		math2(func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))
}

// 1.允许多个返回值
func div1(a, b int) (int, int) {
	return a / b, a % b
}

// 2.多返回值另一种写法 可以给返回值变量名
func div2(a, b int) (q, r int) {
	return a / b, a % b
}

// 3.方法内操作返回变量 最后直接return
func div3(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func math1(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("操作符错误：" + op)
	}
}

// 类似java函数式接口 把实现逻辑交到调用方
func math2(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()    // 反射获取指针
	opName := runtime.FuncForPC(pointer).Name() // 获取方法名称
	fmt.Printf("Calling func %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	result := 0
	for _, v := range numbers {
		result += v
	}
	return result
}
