package main

import "fmt"

func main() {
	g := true == false
	h := 32 <= 23
	i := 1 >= 0
	j := true != false
	k := 0 < 100
	l := 100 > 0
	fmt.Println(g, h, i, j, k, l)
}
