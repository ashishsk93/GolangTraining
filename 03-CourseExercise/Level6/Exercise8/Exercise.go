package main

import "fmt"

func test() func() {
	return func() {
		fmt.Println("I am being returned")
	}
}

func main() {
	f := test()

	f()
}
