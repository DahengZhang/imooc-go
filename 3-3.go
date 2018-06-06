package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("%v len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s1 := arr[:5]
	//fmt.Println(s1)
	//fmt.Println(cap(s1))
	//s2 := append(s1, 10) // 会改变原数组
	//s3 := append(s2, 11)
	//s4 := append(s3, 12)
	//s5 := append(s4, 13)
	//s6 := append(s5, 14) // 超出原数组或切片长度时，会自动分配更长的空间存储
	//fmt.Println(s6)
	//fmt.Println(arr)
	var s []int // 创建切片
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	s0 := []int{1, 3, 5, 7}
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	s2 := make([]int, 16) // 空的切片，len为16
	s3 := make([]int, 10, 32) // 空的切片，len为16, cap为32
	printSlice(s2)
	printSlice(s3)

	// 切片拷贝, 直接覆盖对应项
	copy(s1, s0)
	printSlice(s1)
	copy(s2, s1)
	printSlice(s2)

	// 切片删除某一项
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	front := s2[:1]
	s2 = s2[1:]
	printSlice(front)
	printSlice(s2)

	back := s2[len(s2) -1:]
	s2 = s2[:len(s2) -1]
	printSlice(back)
	printSlice(s2)
}
