package main

import "fmt"

func main() {
	for x := 33; x < 123; x++ {
		fmt.Printf("%d - %#U\n", x, x)
	}
}
