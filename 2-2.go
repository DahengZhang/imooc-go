package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

func euler() { // 验证复数
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
}

func triangle() { // 强制类型转换
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64((a * a + b * b))))
	fmt.Print(c)
}

func main()  {
	euler()
	triangle()
}