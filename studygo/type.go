package main

import (
	"fmt"
	"reflect"
)

func Interfa(in interface{}) {
	intr := reflect.ValueOf(in)
	fmt.Println(intr.Pointer)
	n := intr.NumField()
	fmt.Println(intr.NumField)
	fmt.Println(n)
	for i, n := 0, intr.NumField(); i < n; i++ {
		fmt.Println(intr.Field(i))
	}
}

func main() {
	// 类型不同不可以用 == 来判断
	//	a := int16(11)
	//	b := int32(11)
	//	if a == b {
	//		fmt.Printf("%d == %d \n", a, b)
	//	} else {
	//		fmt.Println("type different")
	//	}

	type Test struct {
		a string
		b int64
	}
	test := Test{"123", 123}
	Interfa(test)
}
