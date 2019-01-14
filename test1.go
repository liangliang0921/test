package main

import (
	"fmt"
)

type TestStruct struct {
	a string
	b string
}

func main() {
	testStruct := &TestStruct{}
	var s []string
	s[0] = "123"
	s[1] = "456"
	for _, option := range s {
		option(testStruct)
	}
	fmt.Println(testStruct)
}
