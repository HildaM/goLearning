package golang_1

import "fmt"

func forLoop() {
	// 1.
	a := 1
	for a <= 5 {
		fmt.Println(a)
		a++
	}
	fmt.Println()

	// 2.
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 3. 无限循环
	//for {
	//
	//}
	//for ;; {
	//
	//}

	// 4. for-range

}
