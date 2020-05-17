package main

import (
	"fmt"
)

func main() {
	// 初始化
	m1 := map[string]string{"name": "ryan", "sex": "male"} // 空map
	m2 := make(map[string]string)                          // 空map
	var m3 map[string]string                               // nil
	fmt.Println(m1, m2, m3)

	// range方式遍历
	// 这里如果用m2m3遍历也不会出错
	// map无序所以顺序不一定
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 通过key取值 如果key不存在返回空字符串
	name := m1["name"]
	fmt.Println(name)
	// 接受两个参数判断是否存在
	if sex, exist := m1["sex"]; exist {
		fmt.Println(sex)
	} else {
		fmt.Println("Key not exist.")
	}

	// 删除
	delete(m1, "sex")
	fmt.Println(m1)
}
