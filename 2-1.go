package main

import "fmt"

// 包内变量，作用于当前文件，不可省略 var 关键字
var  (
	aa = 1
	bb = 2
)

func variabelZeroValue() { // 不定义初值
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variabelInitialValue() { // 定义初值
	var a, b int = 5, 15
	var s string = "icain"
	fmt.Printf("%d %d %q\n", a, b, s)
}

func variabelTypeDeduction() { // 自动识别变量类型与省略 var 关键字
	a, b, s := 5, 15, "icain"
	fmt.Printf("%d %d %q\n", a, b, s)
}

func main()  { // 主体函数
	fmt.Println("hello world")
	variabelZeroValue()
	variabelInitialValue()
	variabelTypeDeduction()
}
