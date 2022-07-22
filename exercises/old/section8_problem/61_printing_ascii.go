//look at the ascii docs. from decimal 33 to 122, print the decimal and then
// string.

package main

import "fmt"

func main() {
	// i := 33
	// for {
	// 	fmt.Printf("decimal: %d\t ascii: %c\n", i, i)
	// 	i++
	// 	if i == 123 {
	// 		break
	// 	}
	// }

	// for i := 33; i <= 122; i++ {
	// 	fmt.Printf("decimal: %v\t ascii: %c\n", i, i)
	// }

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
}
