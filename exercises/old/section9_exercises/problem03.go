package main

import "fmt"

func main() {
	then := 1977
	now := 2022
	for then != now {
		fmt.Println(then)
		then++
	}
}
