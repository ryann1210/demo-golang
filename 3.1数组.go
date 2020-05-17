package main

import "fmt"

func main() {
	// 1.数组初始化及赋值
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2, arr3)
	var grid [4][5]bool
	fmt.Println(grid)

	// 2.遍历数组 用range 用i遍历也可以
	// 3.如果只用v的话 i用_占位
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// 4.数组是值类型
	fmt.Println(arr1)
	printArr(arr1)
	fmt.Println(arr1)

	// 5.可以通过指针传递进行操作
	fmt.Println(arr1)
	printArr2(&arr1)
	fmt.Println(arr1)
}

func printArr(arr [5]int) {
	fmt.Println(arr)
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArr2(arr *[5]int) {
	fmt.Println(arr)
	arr[0] = 100
	fmt.Println(arr)
}
