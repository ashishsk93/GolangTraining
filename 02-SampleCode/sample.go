package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is a test statement")

	foo()

	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, "is divisible by three")
		} else {
			fmt.Println(i, "is not divisible by three")
		}
	}

	exit()
}

func foo() {
	fmt.Println("I am inside foo")
}

func exit() {
	fmt.Println("About to exit the program!!")
}
