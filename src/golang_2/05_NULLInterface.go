package golang_2

import (
	"fmt"
)

type empty_iface interface {
}

func GetInfo() {
	var i interface{}
	fmt.Printf("type: %T, value: %v", i, i)
}

// 1. 充当 “java的Object类”，承载任何值
func ObjectOfAll() {
	var i interface{}

	i = 1
	fmt.Println(i)

	i = "Object of all"
	fmt.Println(i)

	i = func(int) int {
		return 10
	}(10)
	fmt.Println(i)
}

// 2. 让函数接收任何值
func printAll(iface interface{}) {
	fmt.Println(iface)
}

// 接收任意个类型的值
func receiveAll(ifaces ...interface{}) {
	for _, iface := range ifaces {
		fmt.Println(iface)
	}
}

func GetPrint() {
	a := 10
	printAll(a)

	b := "Object of all"
	printAll(b)

	c := func(int) {
		type innerC struct {
			name string
			age  int
		}
	}
	printAll(c)

	receiveAll("***************", c, b, a)
}

// 3. 定义一个包容万象的切片
func ReceiveAll() {
	any := make([]interface{}, 5)
	any[0] = 10
	any[1] = "hello Any"
	any[2] = true
	any[3] = []int{1, 2, 3, 4}

	for _, value := range any {
		fmt.Println(value)
	}
}

/////////////////////////////////

// a. “江河不能承载大海”
func smallCantHoldBig() {
	//var a int = 1
	//var i interface{} = a
	//var b string = i
	//
	//fmt.Println(b)
}

// b. "不能回头“
func GoForward() {
	//sli := []int{1, 2, 3, 4, 5}
	//
	//var i interface{} = sli
	//
	//g := i[1:3]
	//
	//fmt.Println(g)
}

// c. ”空接口就是一张面具，被它修饰后类型就被隐藏起来了“
func Hide() {
	a := 10
	b := true
	c := []string{"hello", "world"}
	myfunc(a, b, c)
}

func myfunc(i ...interface{}) {
	for _, value := range i {
		switch value.(type) {
		case int:
			fmt.Println(value, "类型为int")
		case bool:
			fmt.Println(value, "类型为bool")
		case []string:
			fmt.Println(value, "类型为string数组")
		}
	}
}
