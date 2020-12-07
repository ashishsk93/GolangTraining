package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println(i, "is Divisible by 4")
		} else if i%4 == 1 {
			fmt.Println(i, "gives a remainder of 1 when divided by 4")
		} else if i%4 == 2 {
			fmt.Println(i, "gives a remainder of 2 when divided by 4")
		} else {
			fmt.Println(i, "gives a remainder of 3 when divided by 4")
		}
	}
}
