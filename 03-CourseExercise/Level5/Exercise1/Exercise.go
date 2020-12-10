package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {
	p1 := person{
		firstName: "Ashish",
		lastName:  "Kumar",
		iceCream: []string{
			"Vanilla",
			"Butterscotch",
		},
	}

	p2 := person{
		firstName: "Anand",
		lastName:  "Kumar",
		iceCream: []string{
			"Chocolate",
			"Mango",
		},
	}

	for _, v := range p1.iceCream {
		fmt.Println(v)
	}

	for _, v := range p2.iceCream {
		fmt.Println(v)
	}
}
