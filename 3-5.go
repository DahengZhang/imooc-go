package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	last0ccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := last0ccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i -start + 1
		}
		last0ccurred[ch] = i
	}
	return maxLength
}

func myLengthOfNonRepeatingSubStr(s string) int { // 存在BUG，逻辑出错
	maxNoRP := make(map[byte]int)
	maxLength := 0
	for i, ch := range []byte(s) {
		if _, ok := maxNoRP[ch]; ok {
			delete(maxNoRP, ch)
		}
		maxNoRP[ch] = i
		if len(maxNoRP) >= maxLength {
			maxLength = len(maxNoRP)
		}
	}
	return maxLength
}

func lengthOfNonRepeatingSubStrIn(s string) int {
	last0ccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := last0ccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i -start + 1
		}
		last0ccurred[ch] = i
	}
	return maxLength
}

func myLengthOfNonRepeatingSubStrIn(s string) int { // 存在BUG，逻辑出错
	maxNoRP := make(map[rune]int)
	maxLength := 0
	for i, ch := range []rune(s) {
		if _, ok := maxNoRP[ch]; ok {
			delete(maxNoRP, ch)
		}
		maxNoRP[ch] = i
		if len(maxNoRP) >= maxLength {
			maxLength = len(maxNoRP)
		}
	}
	return maxLength
}

func main() {
	fmt.Println(0, lengthOfNonRepeatingSubStr("abcabndvd"))
	//fmt.Println(1, myLengthOfNonRepeatingSubStr("abcabndvd")) // 存在BUG，逻辑出错
	//fmt.Println(1, myLengthOfNonRepeatingSubStr("abcab张大亨vd")) // 存在BUG，逻辑出错
	fmt.Println(2, lengthOfNonRepeatingSubStrIn("abcab张大亨vd"))
	//fmt.Println(2, myLengthOfNonRepeatingSubStrIn("abcab张大亨vd")) // 存在BUG，逻辑出错
}
