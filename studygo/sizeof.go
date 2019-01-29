package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type A struct {
		a bool
		//b float64
		b int32
		c int16
	}
	type B struct {
		a bool
		c int16
		b float64
	}
	emea := A{}
	emeb := B{}
	emec := bool(true)
	emed := int16(11)
	emef := float64(11)

	fmt.Println(unsafe.Sizeof(emea))
	fmt.Println(unsafe.Sizeof(emeb))
	fmt.Println(unsafe.Sizeof(emec))
	fmt.Println(unsafe.Sizeof(emed))
	fmt.Println(unsafe.Sizeof(emef))
	fmt.Println(unsafe.Alignof(emea))
	fmt.Println(unsafe.Offsetof(emea.a))
	fmt.Println(unsafe.Offsetof(emea.b))
}
