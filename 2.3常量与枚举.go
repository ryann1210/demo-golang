package main

import (
	"fmt"
	"math"
)

func main() {
	// 常量
	consts()
	// 枚举
	emuns()
}

func consts() {
	// 1. 常量的数值可以作为各种类型使用
	const fileName = "123.txt"
	const a, b = 3, 4
	var d int = a
	var c = math.Sqrt(a*a + b*b)
	fmt.Println(a, b, c, d, fileName)
	// 2. 常量不必全部大写 首字母大写的变量另有含义
	// 3. 常量也可以集中定义
	const (
		name = "ryan"
		age  = 18
	)
	fmt.Println(name, age)
}

func emuns() {
	// 1.通过const定义枚举类型
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)
	// 2. 通过iota定义枚举类型
	// iota表示一组自增值 _表示跳过某个自增值
	const (
		beijing = iota
		_
		_
		shanghai
		guangzhou
		shenzhen
	)
	fmt.Println(beijing, shanghai, guangzhou, shenzhen)
	// 3. b kb mb gb tb pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
