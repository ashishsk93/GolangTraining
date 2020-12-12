package main

import "fmt"

func foo(f func()) {
	f()
}

func main() {
	foo(func() {
		fmt.Println("Hello how are you?")
	})
}
