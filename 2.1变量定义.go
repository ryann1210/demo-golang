package main

import "fmt"

func variable() {
	// 1.初始化
	// 变量名在前，变量类型在后
	// 默认初始值int为0，string为""
	// Printf为格式化打印，%q会打印出空字符串
	// 变量定义了必须要用，不然会报红
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b) // 0 ""

	// 2.初始化赋值
	// 初始化可以赋值，可以多个变量同时赋值
	var c, d int = 3, 4
	var e string = "hello"
	fmt.Println(c, d, e)

	// 3. 类型推断
	// 初始化赋值可以不用写类型
	var f, g, h = 3, true, "def"
	fmt.Println(f, g, h)

	// 4. 初始化赋值的简化写法 :=
	i, j, k := 3, false, "hello"
	fmt.Println(i, j, k)
}

// 5. 方法外函数
// 方法外变量初始化赋值不可以用:=
// 不是全局变量，只在该package下可用
var aa = 3
var bb = false
var cc = "kkk"

// 多个变量可以集中定义，该方式在函数内也可以
var (
	dd = 3
	ee = false
	ff = "ggg"
)

func main() {
	variable()
	fmt.Println(aa, bb, cc)
	fmt.Println(dd, ee, ff)
}
