package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	var (
		fileName = "D:/test.txt"
		content  = "Hello golang"
		file     *os.File
		err      error
	)

	// 使用追加模式打开文件
	file, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	if err != nil {
		panic(errors.New("打开文件失败"))
	}
	defer file.Close() // 养成好习惯，打开资源后紧跟着defer关闭

	// 写文件
	writer := bufio.NewWriter(file)
	n, err := writer.Write([]byte(content))
	if err != nil {
		panic(errors.New("写入文件失败"))
	}
	fmt.Println("写入文件成功，文件长度为 = ", n)
	writer.Flush() // 刷新

}
