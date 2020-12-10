package main

import "fmt"

func main() {
	t := struct {
		first string
		last  string
	}{
		first: "Ashish",
		last:  "Kumar",
	}

	fmt.Println(t)
}
