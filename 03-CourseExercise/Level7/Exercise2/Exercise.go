package main

import "fmt"

type person struct {
	first   string
	last    string
	address string
}

func changeMe(p *person) {
	(*p).address = "Changed Address"
}

func main() {
	p := person{
		first:   "Ashish",
		last:    "Kumar",
		address: "Address",
	}

	fmt.Println(p)

	changeMe(&p)

	fmt.Println(p)
}
