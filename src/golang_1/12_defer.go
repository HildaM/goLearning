package golang_1

import "fmt"

func Defer() {
	defer fmt.Println("类似于join")
	fmt.Println("先走一步")
}

func SnapShot() {
	name := "go"
	defer fmt.Println(name)

	name = "java"
	fmt.Println(name)
}

func SnapShotEX() {
	fmt.Println()

	name := "go"
	defer func() {
		fmt.Println(name)
	}()

	name = "cpp"
	fmt.Println(name)
}
