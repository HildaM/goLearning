package golang_2

import "fmt"

// 1. new
type Stu struct {
	name string
	age  int
}

func NewStu() {
	// new内建对象
	num := new(int)
	fmt.Printf("Num: %d, Num ptr: %p\n", *num, num)

	// new自定义对象
	xiaoMing := new(Stu)
	xiaoMing.name = "xiaoMing"
	fmt.Println(xiaoMing)
	fmt.Println(*xiaoMing)
}

func Make() {
	// slice
	slice := make([]int, 2, 10)
	fmt.Printf("size = %d, cap = %d\n", len(slice), cap(slice))

	// map
	map1 := make(map[string]int)
	map1["test"] = 10
	fmt.Println(map1)

	// channel
	channel := make(chan int, 10)
	fmt.Println(channel)
}
