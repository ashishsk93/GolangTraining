package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Anonymous Function called")
	}

	f()
}
