package main

import "fmt"

/*
寻找最长不含有重复字符的子串
*/
func main() {
	result := lengthOfLongestSubstring("一二三而一")
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	// 记录每个字符最后出现的位置
	lastOccurred := make(map[byte]int)
	// 有效子串的起始
	start := 0
	// 最大长度
	maxLength := 0
	// 遍历字符串
	for i, ch := range []byte(s) {
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
