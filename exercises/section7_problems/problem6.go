package main

import "fmt"

const (
	a = 2021 + iota
	b
	c
	d
)

func main() {
	a := 2021
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
