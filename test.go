package main

import (
	"fmt"
	"time"
)

func main() {
	ti := int64(time.Now().UnixNano())
	fmt.Println(ti)
	fmt.Println(time.Now().Unix())
}
