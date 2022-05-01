package golang_1

import (
	"container/list"
	"fmt"
)

func List() {
	// list 声明
	var myList list.List
	//myList2 := list.New()

	// 插入元素
	myList.PushBack("pushBack")
	myList.PushFront("pushFront")
	myList.PushBack(11)

	// 删除元素
	element := myList.PushBack(11)
	myList.Remove(element)

	// 访问元素：使用迭代器访问
	// 顺序访问
	for iter := myList.Front(); iter != nil; iter = iter.Next() {
		fmt.Printf("%v ", iter.Value)
	}
	// 倒序访问
	for iter := myList.Back(); iter != nil; iter = iter.Prev() {
		fmt.Printf("%v ", iter.Value)
	}

}
