package golang_2

import "fmt"

/*
	目前看来：只有interface类型的东西才用得上断言
*/

// 1. 第一种使用方法
func Asser01() {
	var i interface{} = 10

	t1 := i.(int) // 判断i是否为int类型
	fmt.Println(t1)

	fmt.Println("***************")

	t2 := i.(string)
	fmt.Println(t2)

	fmt.Println("***************")
}
