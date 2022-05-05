package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// define chan
	var c1 chan int // 定义一个传输int的通道
	c2 := make(chan int)
	fmt.Println(c1, c2)

	// 通信操作符 <-
	// 注意：操作chan的时候要使用goroutine
	go sendData(c2)
	go getData(c2)

	time.Sleep(1e8) // 避免主程序过快结束
}

func sendData(ch chan int) {
	// 流向通道
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func getData(ch chan int) {
	var num int
	// 流出通道
	for {
		num = <-ch
		fmt.Print(strconv.Itoa(num) + " ")
	}
}
