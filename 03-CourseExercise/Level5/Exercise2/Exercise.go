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

	m := map[string]person{
		p1.firstName: p1,
		p2.firstName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, v1 := range v.iceCream {
			fmt.Println(i, v1)
		}
	}
}
