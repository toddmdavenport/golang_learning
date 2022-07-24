package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	var whatToSay string
	var i int
	whatToSay = "Goodbye, cruel world!"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)
	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("The function returned:", whatWasSaid)
	fmt.Println("The other thing that was said was:", theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
