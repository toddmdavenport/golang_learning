package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%d	%#x	%b\n", a, a, a)
	b := a << 1
	fmt.Printf("%d	%#x	%b\n", b, b, b)
}
