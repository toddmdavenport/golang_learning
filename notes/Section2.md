# Section 2

## 7. Note On Windows and bash

Install gitbash or WSL if you want to use the same terminal
commands as in the videos.

## 8. Variables and Functions

Will write our first functional program in Go. Create a file called `main.go`.
This is a convention in the language. It is not needed but is considered
idomatic

Start with a package declaration `package main` and a function declaration.
Each file must end with `.go` and the first line must be a package declaration.

[section2 main.go file](../exercises/section2/8functionsAndVariables.go) shows some basic use of
the language including printing, variable declaration, combimed declaration and
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

Some more info on Variable Declaration

you can declare a variable and set it at the same time like `var s = "seven"`
Local variables can override global variables in the scope of a functions.
The compiler will allow re-use of the same names at different scopes so it is
important to avoid name collision.

[exercise](../exercises/section2/10typesAndStructs.go)
