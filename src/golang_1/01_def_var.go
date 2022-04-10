package golang_1

import (
	"fmt"
)

/*
	五种创建变量的方法
*/

// 一行创建一个变量
func Def01() {
	// var <name> <type>
	var name string = "test01"
	var rate float32 = 0.22
	fmt.Printf("name: %s, token: %f", name, rate)
}

// 多个变量一起声明
func Def02() {
	var (
		name   string
		age    int
		gender string
	)
	fmt.Printf("%s, %d, %s", name, age, gender)
}

// 声明和初始化一个变量
/*
	使用 := （
	推导声明写法或者短类型声明法：编译器会自动根据右值类型推断出左值的对应类型。），
	可以声明一个变量，并对其进行（显式）初始化。

	注意！！！
		这种方法有个限制就是，只能用于函数内部
		也就是说：不能用于全局变量
*/
// name := "test"  报错！
func Def03() {
	name := "test03"
	// 等价于 var name string = "test03"
	// 等价于 var name = "test03"
	fmt.Println(name)
}

// 声明和初始化多个变量
func Def04() {
	name, age := "quan", 22
	fmt.Println(name, age)

	// 可以用于交换变量
	a, b := 100, 200

	a, b = b, a // 便捷的交换！
}

// new 函数声明一个指针变量
/*
	类似cpp
*/
func Def05() {
	// 普通变量和获取其它的指针
	var age = 100
	var ptr = &age
	fmt.Printf("val: %d, ptr: %p\n", age, ptr)

	/*
		new是一个函数，返回值为一个指针（*Type）
	*/
	ptr = new(int)
	fmt.Printf("ptr address: %p \nptr value: %d", ptr, *ptr)
}

// 匿名变量，下划线表示
/*
	不分配内存，不占用内存空间
	不需要你为命名无用的变量名而纠结
	多次声明不会有任何问题
通常我们用匿名接收必须接收，但是又不会用到的值。
*/
func Def06() {
	GetData := func() (int, int) {
		return 100, 200
	}

	a, _ := GetData()
	_, b := GetData()

	fmt.Println(a, b)
}
