package performance

import (
	"bytes"
	"strings"
)

func Builder() {
	var builder strings.Builder
	str := " test "
	for i := 0; i < 100; i++ {
		builder.WriteString(str)
	}
	//fmt.Println(builder.String())
}

func Add() {
	str := " test "
	ans := ""
	for i := 0; i < 100; i++ {
		ans += str
	}
}

func Buffer() {
	var buffer bytes.Buffer
	for i := 0; i < 100; i++ {
		buffer.WriteString(" test ")
	}
}
