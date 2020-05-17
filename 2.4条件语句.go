package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "2.4条件语句.txt"

	// 1.函数可以有多个返回值
	contents, err := ioutil.ReadFile(filename)
	// 2.nil有点像java的null
	if err != nil {
		fmt.Println(err)
	} else {
		// 3.contents为bytes数组，用Printf可以打印出内容
		fmt.Printf("%s\n", contents)
	}

	// 4.可以在if里写执行和判断的逻辑 这样写contents和err只能在if内访问
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(eval(3, 4, "+"))

	fmt.Println(grade(12))
}

func eval(a, b int, op string) int {
	var result int
	// 5.swtich会自动break 除非使用fallthrough
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		// 7.相当于报错
		panic("不支持的操作符")
	}
	return result
}

func grade(score int) string {
	var result string
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		result = "不及格"
	case score <= 100:
		result = "及格"
	}
	return result
}
