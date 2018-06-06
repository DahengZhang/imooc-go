package main

import "fmt"

func updateSlice(arr []int)  { // 不加长度为切片
	arr[0] = 100
}

func main() { // 切片
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s := arr[2:8] // 切片 左闭右开
	fmt.Println(s)
	fmt.Println(arr[:8])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])

	updateSlice(arr[:])
	fmt.Println(s)
	fmt.Println(arr)
	updateSlice(s)
	fmt.Println(s)
	fmt.Println(arr)

	s2 := arr[3:8]
	fmt.Println(s2)
	s2 = s2[2:] // 切片之上可以建立切片
	fmt.Println(s2)
	s3 := arr[2:6]
	fmt.Println(s3)
	fmt.Println(cap(s3))
	s4 := s3[2:8]
	fmt.Println(s4)
	fmt.Println(cap(s4)) // cap为上一级切片或数组的len
	s5 := arr[:0]
	fmt.Println(s5)
	fmt.Println(len(s5))
	fmt.Println(cap(s5)) // 包含被切掉的部分，切片仅仅展示所选范围，并不对原数组剪切
}
