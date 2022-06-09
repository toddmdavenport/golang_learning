package main

import "fmt"

func main() {
	count := 3
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\n", i)
		for j := 0; j <= 2; j++ {
			fmt.Printf("\t%#U\n", i)
			count++
		}
	}
}
