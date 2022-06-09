package main

import "fmt"

func main() {
	if true == true {
		fmt.Println("It's true!")
		if true == false {
			fmt.Println("This will never print :( ")
		}
	}
}
