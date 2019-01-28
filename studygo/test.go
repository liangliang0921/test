package main

import (
	"fmt"
	//"time"
)

func Printd() {
	//time.Sleep(3 * time.Second)
	fmt.Println("End Printd")
}

func Transfer() {
	go Printd()
	fmt.Println("End Transfer")
}

func main() {
	Transfer()
	//time.Sleep(5 * time.Second)
	fmt.Println("End main")
}
