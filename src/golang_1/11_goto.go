package golang_1

import "fmt"

func Goto() {
	var name string
	i := 1
	for {
		if i > 5 {
			goto flag
		}
		fmt.Println(i)
		i++
	}
	// var name = "test"
flag:
	fmt.Println(name)
}
