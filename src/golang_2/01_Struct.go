package golang_2

import "fmt"

/*
	golang中没有类，只有结构体
*/

// 声明结构体
type Person struct {
	// name,gender string 属性相同可以合并
	name   string
	age    int8
	gender string
	bro    *Person // 指针
}

// 绑定方法
func (person Person) FmtBro() {
	fmt.Println(person.bro)
}

func Struct01() {
	Hilda := Person{
		name:   "Hilda",
		age:    22,
		gender: "boy",
		bro:    nil,
	}

	quan := Person{
		name:   "Hilda",
		age:    22,
		gender: "boy",
		bro:    &Hilda,
	}

	fmt.Println(Hilda, quan)
	quan.FmtBro()
}
