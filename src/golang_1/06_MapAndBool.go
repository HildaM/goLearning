package golang_1

import (
	"fmt"
	"sync"
)

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

func Map() {
	mutilMap := make(map[int][]int) // 注意：make(map[int][]int, 100, 200) 会报错

	// 添加元素
	// 缔造一个题型数组
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			mutilMap[i] = append(mutilMap[i], j)
		}
	}

	// 遍历
	for idx, arr := range mutilMap {
		fmt.Printf("第 %d 行：", idx)
		fmt.Println(arr)
	}
}

func SyncMap() {
	var syncMap sync.Map

	/*
		sync.Map 需要使用Store函数存储
		同时，Map没有限制存储的类型
	*/
	for i := 0; i < 10; i++ {
		syncMap.Store(i, i*10)
	}
	syncMap.Store("a", "A")
	syncMap.Store("b", "B")
	syncMap.Store("c", "C")

	// 获取指定的值
	fmt.Println(syncMap.Load('a'))
	fmt.Println(syncMap.Load(2))

	// 删除指定值
	syncMap.Delete('q')

	// range遍历
	syncMap.Range(func(k, v interface{}) bool {
		fmt.Printf("k: %v, v: %v\n", k, v)
		return true
	})

}
