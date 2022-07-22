package main

import (
	"fmt"
)

type hamburger int

var x hamburger
var y int

func main() {
	//from previous exercise
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	// new code
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
