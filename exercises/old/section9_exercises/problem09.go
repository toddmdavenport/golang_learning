package main

import "fmt"

var favsport string

func main() {
	favsport = "Basketball"
	switch favsport {
	case "Basketball":
		fmt.Println("Basketball")
	case "Baseball":
		fmt.Println("Baseball")
	case "Skiing":
		fmt.Println("Skiing")
	case "Football":
		fmt.Println("Football")
	}
}
