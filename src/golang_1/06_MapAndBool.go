package golang_1

import "fmt"

func GetMap() {
	// 1. 直接声明和初始化
	var map01 map[string]int = map[string]int{
		"Chinese": 10,
		"eng":     20,
	}
	fmt.Println(map01)

	// 2. 初始化和赋值
	map02 := map[string]int{"quan": 22}
	fmt.Println(map02)

	// 3. make
	map03 := make(map[string]int)
	map03["hilda"] = 18
	fmt.Println(map03)
}

func IsExist() {
	scores := make(map[string]int)
	math, ok := scores["chinese"]
	if ok {
		fmt.Println(math)
	} else {
		fmt.Println("not exist")
	}

	// 先赋值，后判断
	if math, ok := scores["math"]; ok {
		fmt.Printf("math 的值是: %d", math)
	} else {
		fmt.Println("math 不存在")
	}
}
