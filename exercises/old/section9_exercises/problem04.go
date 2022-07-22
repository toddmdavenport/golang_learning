package main

import "fmt"

func main() {
	then := 1977
	now := 2022
	for {
		fmt.Println(then)
		then++
		if then == now {
			break
		}
	}
}
