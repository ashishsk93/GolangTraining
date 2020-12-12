package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Name:", p.first, p.last, "Age:", p.age)
}

func main() {
	p := person{
		first: "Ashish",
		last:  "Kumar",
		age:   27,
	}

	p.speak()
}
