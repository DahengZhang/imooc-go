package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int , op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("Something was wrong! %s", op)
	}
}

func apply(op func(int, int) int, a, b int) int  {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args " + // 换行粘贴
		"(%d %d)\n", opName, a, b)
	return op(a, b)
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) (s int) { // 可变参数列表
	s = 0
	for i := range numbers {
		s += numbers[i]
	}
	return
}

func main() {
	if result, err := eval(3, 4, "+"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(div(3, 4))
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7))
}
