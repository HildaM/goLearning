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

func (person *Person) incre_age() {
	person.age++
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

/***************************************************/

type company struct {
	companyName string
	companyAddr string
}

type staff struct {
	name     string
	age      int
	gender   string
	position string
	company
	// 注意！这个company没有任何修饰！
}

func Struct02() {
	ByteDance := company{
		companyName: "ByteDance",
		companyAddr: "Shenzhen",
	}

	Hilda := staff{
		name:     "Quan",
		age:      22,
		gender:   "man",
		position: "Zhaoqing",
		company:  ByteDance,
	}

	fmt.Println(Hilda)

}

func Struct03() {
	// 1. 直接实例化
	myCom := company{
		companyName: "Tencent",
		companyAddr: "shenzhen",
	}
	fmt.Println(myCom)

	// 2. 使用new
	myCom2 := new(company)
	// 此时输出为0
	fmt.Println(myCom2)

	myCom2.companyName = "miHoyo"
	myCom2.companyAddr = "shanghai"
	fmt.Println(myCom2)

	// 3. 使用&
	var myCom3 *company = &company{}
	fmt.Println(myCom3)

	myCom3.companyName = "Microsoft"
	myCom3.companyAddr = "suzhou"
	fmt.Println(myCom3)
}
