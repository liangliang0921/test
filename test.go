package main

import (
	"fmt"
	"time"
)

func main() {
	ti := int64(time.Now().UnixNano())
	fmt.Println(ti)
	fmt.Println(time.Now().Unix())

	var str string
	if str == "" {
		fmt.Println("str 为空")
	}

	if len(str) == 0 {
		fmt.Println("str len 为 0")
	}

	str2 := "你好sda😸213"
	str3 := []rune(str2)
	fmt.Println(len(str3))

	s1, s2 := 1, 2
	fmt.Println(s1, s2)
}
