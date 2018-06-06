package main

import (
	"io/ioutil"
	"fmt"
)

func grade(score int) (g string) {
	g = ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g ="C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Something was wrong"))
	}
	return
}

func main() { // 条件语句
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil { // contents 与 err 作用在 if 语句内
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}
	result := grade(90)
	fmt.Println(result)
}
