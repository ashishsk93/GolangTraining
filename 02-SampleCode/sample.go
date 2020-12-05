package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is a test statement")

	foo()

	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Print(i)
			fmt.Print(" is divisible by three")
		} else {
			fmt.Print(i)
			fmt.Print(" is not divisible by three")
		}
		fmt.Println()
	}

	exit()
}

func foo() {
	fmt.Println("I am inside foo")
}

func exit() {
	fmt.Println("About to exit the program!!")
}
