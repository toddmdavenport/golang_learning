package main

import (
	"fmt"
)

type hamburger int

var x hamburger

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
