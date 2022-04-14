package golang_2

import (
	"fmt"
	"reflect"
)

func Reflect() {
	var age interface{} = 25

	fmt.Printf("原接口遍历类型：%T, 值为：%v\n", age, age)

	// 1. 第一定律
	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", t)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)

	// 2. 第二定律
	i := v.Interface()
	fmt.Printf("从反射对象到接口变量：新对象的类型为 %T 值为 %v \n", i, i)
}

func Settable() {
	var name = "Hilda golang"

	v1 := reflect.ValueOf(&name)
	fmt.Println("可写性为：", v1.CanSet())

	v2 := v1.Elem()
	fmt.Println("可写性为：", v2.CanSet())
	fmt.Println(v2)

	// 修改后，原本的name也会发生变化
	v2.SetString(v2.String() + "test")
	fmt.Println(v2, name)
}
