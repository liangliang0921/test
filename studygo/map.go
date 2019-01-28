package main

import (
	"fmt"
)

func main() {
	m1 := make(map[string]string)
	m1["1234"] = "1312313"

	var m2 map[string]string
	m2 = make(map[string]string)
	m2["13213"] = "eqweq"
	fmt.Println(m2)
	info, ok := m2["13213"]
	if !ok {
		fmt.Println(info)
	}
	fmt.Println(m1)
}
