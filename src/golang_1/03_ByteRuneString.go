package golang_1

import "fmt"

func ByteAndRune() {
	var a byte = 65
	var b uint8 = 66
	fmt.Printf("a = %c, b = %c", a, b)
}

func String() {
	// string的本质是一个byte数组
	myStr01 := "hello"
	myStr02 := [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("mystr01: %s\n", myStr01)
	fmt.Printf("mystr02: %s", myStr02)
}
