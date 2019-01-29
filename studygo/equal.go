package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c", "d"}
	if !reflect.DeepEqual(got, want) {
		fmt.Println("Equal err")
	}
}
