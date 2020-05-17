package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	// 1.切片是对数组的视图 修改视图会修改原数组
	s1 := arr[:]
	fmt.Println("arr: ", arr)
	fmt.Println("s1: ", s1)
	updateSlice(s1)
	fmt.Println("s1: ", s1)
	fmt.Println("arr: ", arr)

	// 3.修改s2 s1和arr都会改变
	s2 := s1[2:]
	fmt.Println("arr: ", arr)
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	updateSlice(s2)
	fmt.Println("s2: ", s2)
	fmt.Println("s1: ", s1)
	fmt.Println("arr: ", arr)

	// 4.slice可以向后扩展 不可以向前扩展
	// 数据结构包括：
	// ptr 指向第一个元素
	// len slice的长度
	// cap ptr到数组最后一个元素的长度
	// 总结：s[i]不可以超越len(s) 向后扩展不可以超过底层数组cap(s)
	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s3 := arr2[2:6] // [2 3 4 5]
	s4 := s3[3:5]   // [5 6]
	fmt.Println(s3, s4)
	// fmt.Println(s3[4]) 报错index out of range
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3)) // s3=[2 3 4 5], len(s3)=4, cap(s3)=6
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n", s4, len(s4), cap(s4)) // s4=[5 6], len(s4)=2, cap(s4)=3
}

// 2.视图的参数不需要写长度
func updateSlice(s []int) {
	s[0] = 100
}
