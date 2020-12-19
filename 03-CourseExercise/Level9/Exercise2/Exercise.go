package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p *person) speak() {
	fmt.Println("My Name is:", p.first, p.last)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "Ashish",
		last:  "Kumar",
	}

	//saySomething(p)
	saySomething(&p)

	fmt.Println("Exiting")
}
