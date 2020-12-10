package main

import "fmt"

func main() {
	t := foo()
	v, x := bar()

	fmt.Println(t, v, x)
}

func foo() int {
	return 36
}

func bar() (int, string) {
	return 45, "Test"
}
