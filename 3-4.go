package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "icain",
		"age": "24"}

	m2 := make(map[string]int)

	var m3 map[int]int
	fmt.Println(m, m2, m3)

	for i, v :=range m {
		fmt.Println(i, v) // map 中是个无序数组，每次打印顺序不一
	}

	mapName := m["name"]
	fmt.Println(mapName)
	mapAg, ok := m["agse"] // 如果所选项不存在，则会返回空, ok为false
	fmt.Println(mapAg, ok)
	if mapAg, ok := m["age"]; ok == true {
		fmt.Println(mapAg)
	} else {
		fmt.Println("对应项不存在")
	}

	delete(m, "age")
	fmt.Println(m)
}
