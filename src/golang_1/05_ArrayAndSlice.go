package golang_1

import (
	"fmt"
)

func Slice() {
	myarr := [...]int{1, 2, 3}
	fmt.Printf("%d 的类型是: %T", myarr[0:2], myarr[0:2])
}

func GetSlice() {
	arr := [10]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}

	s1 := arr[0:8]
	s2 := arr[0:2:5]

	fmt.Println(len(s1), cap(s1))
	fmt.Println(s1)
	fmt.Println(len(s2), cap(s2))
	fmt.Println(s2)
}

func GetSlice02() {
	// 声明，然后手动赋值
	var numList []int
	arr := [5]int{1, 2, 3}
	numList = arr[0:2]
	fmt.Println(numList)
}

func SliceTest() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:8]
	fmt.Printf("myslice为 %d, 其长度为: %d, 容量为：%d\n", myslice, len(myslice), cap(myslice))

	myslice = myslice[:(cap(myslice) + 1)]
	fmt.Println(myslice)
	fmt.Printf("myslice的第四个元素为: %d", myslice[3])
}
