package main

import "log"

type myStruct struct {
	FirstName string
}

// create a function that returns
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	//	log.Println("myVar is set to", myVar.FirstName)
	// log.Println("myVar2 is set to", myVar2.FirstName)
	//can call the function instead of accessing the struct element directly
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
