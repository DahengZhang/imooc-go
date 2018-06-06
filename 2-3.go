package main

import (
	"math"
	"fmt"
)

func consts() { // 常量
	const filename = "abc.txt"
	const a, b = 3, 4 // 常量可以不用指定数据类型，赋整数值后也可以当做 float 类型使用
	var c int
	c = int(math.Sqrt((a * a + b * b)))
	fmt.Print(filename, c)
}

func enums() { // 枚举
	const (
		cpp = 0
		java = 1
		golang = 2
	)
	const (
		cppi = iota
		_
		javai
		golangi
	)
	const (
		b = 1 << (10 * iota)
		_
		kb
		tb
	)
	fmt.Println(cpp, java, golang, cppi, javai, golangi, b, kb, tb )
}

func main() {
	consts()
	enums()
}
