package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	// 添加元素时如果超越cap 系统会重新分配更大的底层数组
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(arr) // [0 1 2 3 4 5 6 10]
	fmt.Println(s1)  // [2 3 4 5]
	fmt.Println(s2)  // [5 6]
	fmt.Println(s3)  // [5 6 10]
	fmt.Println(s4)  // [5 6 10 11]
	fmt.Println(s5)  // [5 6 10 11 12]

	var s6 []int // s6 == nil
	for i := 0; i < 10; i++ {
		s6 = append(s6, 2*i+1)
		fmt.Printf("s6=%v, len(s6)=%d, cap(s6)=%d\n", s6, len(s6), cap(s6))
	}
	/*
	   s6=[1], len(s6)=1, cap(s6)=1
	   s6=[1 3], len(s6)=2, cap(s6)=2
	   s6=[1 3 5], len(s6)=3, cap(s6)=4
	   s6=[1 3 5 7], len(s6)=4, cap(s6)=4
	   s6=[1 3 5 7 9], len(s6)=5, cap(s6)=8
	   s6=[1 3 5 7 9 11], len(s6)=6, cap(s6)=8
	   s6=[1 3 5 7 9 11 13], len(s6)=7, cap(s6)=8
	   s6=[1 3 5 7 9 11 13 15], len(s6)=8, cap(s6)=8
	   s6=[1 3 5 7 9 11 13 15 17], len(s6)=9, cap(s6)=16
	   s6=[1 3 5 7 9 11 13 15 17 19], len(s6)=10, cap(s6)=16
	*/

	s7 := []int{2, 4, 6, 8}
	s8 := make([]int, 16)
	s9 := make([]int, 16, 32)
	fmt.Printf("s7=%v, len(s7)=%d, cap(s7)=%d\n", s7, len(s7), cap(s7))
	fmt.Printf("s8=%v, len(s8)=%d, cap(s8)=%d\n", s8, len(s8), cap(s8))
	fmt.Printf("s9=%v, len(s9)=%d, cap(s9)=%d\n", s9, len(s9), cap(s9))

	copy(s8, s7)
	fmt.Printf("s8=%v, len(s8)=%d, cap(s8)=%d\n", s8, len(s8), cap(s8))

	s8 = append(s8[:3], s8[4:]...)
	fmt.Printf("s8=%v, len(s8)=%d, cap(s8)=%d\n", s8, len(s8), cap(s8))

	s8 = s8[1:]
	fmt.Printf("s8=%v, len(s8)=%d, cap(s8)=%d\n", s8, len(s8), cap(s8))

	s8 = s8[:len(s8)-1]
	fmt.Printf("s8=%v, len(s8)=%d, cap(s8)=%d\n", s8, len(s8), cap(s8))
}
