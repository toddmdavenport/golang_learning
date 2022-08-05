# Section 2

## 7. Note On Windows and bash

Install gitbash or WSL if you want to use the same terminal
commands as in the videos.

## 8. Variables and Functions

Will write our first functional program in Go. Create a file called `main.go`.
This is a convention in the language. It is not needed but is considered
idiomatic

Start with a package declaration `package main` and a function declaration.
Each file must end with `.go` and the first line must be a package declaration.

[section2 main.go file](../exercises/section2/8functionsAndVariables.go) shows some basic use of
the language including printing, variable declaration, combined declaration and
assignment with the `:=` operator, and function declaration and calling a
function.

All variables must be used or the program will not run. There also must be a
main() function that returns nothing.

Functions can have multiple return values in Go.

## 9. Pointers

Pointers are useful, but may be unfamillar if you are coming from an interpreted
language.

A pointer points to a specific location in memory and gives you a means to get
what is stored in that location.

A pointer is declared using a `*` and and reference is declared using an `&`.

You can use a pointer to change a variable in memeory as shown in the
[exercise](../exercises/section2/9pointers.go) for this seciton.

You can see in the example that the memory address of `myString` is loggeed when
the function is called.

## 10. Types and Structs

### Some more info on Variable Declaration

You can declare a variable and set it at the same time like `var s = "seven"`.
Local variables can override global variables in the scope of a functions.
The compiler will allow re-use of the same names at different scopes so it is
important to avoid name collision.

### Structs

It's possible to define a variable for each item that you want to store info
about but that gets cumbersome and unwieldy quickly. In other Object Oriented
languages you would define a object to store related data. Go is not Object
Oriented. It uses a typed entity called a [struct](https://gobyexample.com/structs)
instead.

[Example](https://gobyexample.com/structs):

```go
package main

import "fmt"

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {

    p := person{name: name}
    p.age = 42
    return &p
}

func main() {

    fmt.Println(person{"Bob", 20})
    fmt.Println(person{name: "Alice", age: 30})
    fmt.Println(person{name: "Fred"})
    fmt.Println(&person{name: "Ann", age: 40})
    fmt.Println(newPerson("Jon"))
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)
}
```

[exercise](../exercises/section2/10typesAndStructs.go)

In the exercise we used capital letters for the User struct. Go is not object
oriented. Go uses naming conventions to make thing private or public. If a
function starts with a capital letter, it is available outside the package. If
it starts with a lowercase letter its only accessible inside the package. Same
goes for variables and structs.

### Structs and functions

You can easily assign a function to a struct.

```go
package main

type myStruct struct {
    FirstName string
}

//pointer to the struct
// called a 'receiver`
func (m *myStruct) printFirstName() string {
    return m.FirstName
}
```

This allows you to perform business logic in the function. It's very useful and
will be used throughout the course.
