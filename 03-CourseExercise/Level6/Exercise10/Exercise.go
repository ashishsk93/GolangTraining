package main

import "fmt"

func main() {
	f := foo()

	f()
	f()

	g := foo()
	g()
	g()
	g()
}

func foo() func() {
	x := 13
	return func() {
		x += 10
		fmt.Println(x)
	}
}
