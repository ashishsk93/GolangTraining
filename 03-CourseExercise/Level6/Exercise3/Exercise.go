package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Runs first")
}

func foo() {
	fmt.Println("Runs at the end.")
}
