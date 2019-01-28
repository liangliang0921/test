package main

import (
	"fmt"
)

type Test struct {
	id string
}

func (t Test) PrintTest() {
	fmt.Println(t.id)
}

func (t *Test) PrintSTest() {
	fmt.Println(t.id)
}

func main() {
	t := Test{
		id: "123",
	}

	t.PrintTest()
	t.PrintSTest()
}
