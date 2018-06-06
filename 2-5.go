package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func converToBin(n int) (r string)  { // 将整数转化为二进制
	r = ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		r = strconv.Itoa(lsb) + r
	}
	return
}

func printFile(filename string)  {
	file, err :=os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // 只有限制条件的循环
		fmt.Println(scanner.Text())
	}
}

func forever()  { // 死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	// goLang没有 while
	fmt.Println(
		converToBin(5),
		converToBin(13),
		converToBin(0)) // 0
	printFile("abc.txt")
}
