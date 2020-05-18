package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱天安门!" // 字符串初始化是utf8编码

	fmt.Println(len(s))                    // 19 中文占3字节
	fmt.Println(utf8.RuneCountInString(s)) // 9 以rune为单位返回字符串长度

	for i, ch := range s { // ch就是rune 类型为int32 固定4字节
		fmt.Printf("(%d %X)", i, ch) // (0 59)(1 65)(2 73)(3 6211)(6 7231)(9 5929)(12 5B89)(15 95E8)(18 21)
	}
	fmt.Println()
	for i, ch := range []rune(s) { // 转成rune数组再输出
		fmt.Printf("(%d %c)", i, ch) // (0 Y)(1 e)(2 s)(3 我)(4 爱)(5 天)(6 安)(7 门)(8 !)
	}
	fmt.Println()

	bytes := []byte(s)
	fmt.Printf("%x\n", []byte(s)) // 596573e68891e788b1e5a4a9e5ae89e997a821
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch) // Y e s 我 爱 天 安 门 !
	}
	fmt.Println()

	// 完善上个章节例题对中文的支持
	// strings包里有很多对字符串的操作
	result := lengthOfLongestSubstringRune("一二三而一")
	fmt.Println(result)
}

func lengthOfLongestSubstringRune(s string) int {
	// 记录每个字符最后出现的位置
	lastOccurred := make(map[rune]int)
	// 有效子串的起始
	start := 0
	// 最大长度
	maxLength := 0
	// 遍历字符串
	for i, ch := range []rune(s) {
		// 当前字符已出现过 并且在有效子串内
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			// 有效子串的起始值+1
			start = lastI + 1
		}
		// 重新计算最大长度
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		// 更新当前字符最后出现的位置
		lastOccurred[ch] = i
	}
	return maxLength
}
