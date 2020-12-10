package main

import "fmt"

type vehicle struct {
	doors  int
	colors string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors:  4,
			colors: "black",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors:  4,
			colors: "red",
		},
		luxury: true,
	}

	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Println(t1.colors)
	fmt.Println(s1.colors)
}
