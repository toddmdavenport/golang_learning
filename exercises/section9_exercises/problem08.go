package main

import "fmt"

func main() {
	x := "Todd"
	switch {
	case x == "Todd":
		fmt.Println("It's capitalized")
	case x == "todd":
		fmt.Println("Not captalized")
	}
}
