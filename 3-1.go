package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100 // 数组为值类型传递，拷贝之前数组，此处修改并不会影响原数组
	for _, v := range arr {
		fmt.Println(v)
	}
}

func printArrayRef(arr *[5]int) {
	arr[0] = 100 // 指针传递直接传递元数据，此处修改会影响原数组
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [2]int{1, 2}
	arr3 := [...]int{3, 4, 6, 7, 9} //长度不同的数组会被认成不同的数据类型
	var grid [4][5]int // 4行5列
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	fmt.Println("======")
	for i, v := range arr2{
		fmt.Println(i, v)
	}
	fmt.Println("======")
	for _, v := range arr3{
		fmt.Println(v)
	}
	fmt.Println("======")
	printArray(arr3)
	fmt.Println(arr3)
	printArrayRef(&arr3) // 地址传递
	fmt.Println(arr3)
}
