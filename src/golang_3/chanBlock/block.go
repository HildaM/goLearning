package main

import (
	"fmt"
	"time"
)

/*
	注意：go 开启的协程必须在之前。
	如果在开启协程之前，chan有任何操作的话，都会被阻塞！直接panic
*/

func main() {
	//ch := make(chan int)
	ch1 := make(chan int, 1)

	// 先发送再接收
	//go recv(ch)
	//send(ch)

	// 先接收再发送
	//go send(ch)
	//recv(ch)

	// 在有缓冲的通道下，这样操作就可以！
	send(ch1)
	recv(ch1)

	time.Sleep(1e9)
}

func recv(ch chan int) {
	ret := <-ch
	fmt.Println("接收成功：", ret)
}

func send(ch chan int) {
	ch <- 100
	ch <- 101
	fmt.Println("发送成功！")
}
