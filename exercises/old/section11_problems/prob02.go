// ● Using a COMPOSITE LITERAL:
// 	● create a SLICE of TYPE int
// 	● assign 10 VALUES
// ● Range over the slice and print the values out.
// ● Using format printing
// 	○ print out the TYPE of the slice

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 12, 15, 18, 25, 64}
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", a)
}
