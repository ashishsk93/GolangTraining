package main

import "fmt"

func main() {
	for i := 65; i < 91; i++ {
		fmt.Println(i - 64)
		for j := 0; j < 4; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
