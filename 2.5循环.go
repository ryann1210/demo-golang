package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0
	// 1.for同样不需要括号
	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println(convertToBin(13)) // 十进制转二进制 1101

	readFile("2.5循环.txt")

	forever()
}

func convertToBin(n int) string {
	result := ""
	// 2.可以省略初始条件
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func readFile(filename string) {
	// 3.os.Open读取文件
	if file, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		// 4.构建扫描器
		scanner := bufio.NewScanner(file)
		// 5.省略初始条件递增条件 只有结束条件 类似java的while
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

func forever() {
	// 6.死循环
	for {
		fmt.Println("abc")
		// 7.跳出循环
		break
	}
}
