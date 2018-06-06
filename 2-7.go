package main

import "fmt"

func swapByValue(a, b int) {
	b, a = a, b
}

func swapByRef(a, b *int) {
	*b, *a = *a, *b
}

func swapByReturn(a, b int) (int, int) {
	return b, a
}

func main() { // 值传递，引用传递
	a, b := 1, 2
	// 值传递
	swapByValue(a, b)
	fmt.Println(a, b) // 值传递结果
	swapByRef(&a, &b)
	fmt.Println(a, b) // 引用传递结果
	a, b = swapByReturn(a, b)
	fmt.Println(a, b)
}
