package golang_2

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/*
	作json转换时候会用到
*/

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

func Tag01() {
	s1 := Student{
		Name: "jack",
		Age:  12,
	}

	data1, err := json.Marshal(s1)
	if err != nil {
		panic(err)
	}

	// p1没有addr，就不会打印
	fmt.Printf("%s\n", data1)

	s2 := Student{
		Name: "tom",
		Age:  12,
		Addr: "tomHouse",
	}

	data2, err := json.Marshal(s2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data2)
}

/**************************************************/

/*
	定制化输出
*/

type Boss struct {
	Name   string `label:"Name is: "`
	Age    int    `label:"Age is: "`
	Gender string `label:"Gender is: " default:"unknown"`
}

func Tag02() {
	boss := Boss{
		Name: "BigMom",
		Age:  56,
	}

	Print(boss)
}

func Print(obj interface{}) error {
	// 获取value
	v := reflect.ValueOf(obj)

	// 解析字段
	for i := 0; i < v.NumField(); i++ {
		// 取Tag
		field := v.Type().Field(i)
		tag := field.Tag

		// 解析label和default
		label := tag.Get("label")
		defaultValue := tag.Get("default")

		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			value = defaultValue
		}

		fmt.Println(label + value)
	}

	return nil
}
