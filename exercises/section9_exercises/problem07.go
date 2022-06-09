package main

import "fmt"

func main() {
	if true == true {
		fmt.Println("It's true!")
	} else if true == true {
		fmt.Println("This won't print because the above will always be true")
	} else {
		fmt.Println("This will never print :( ")
	}
}
