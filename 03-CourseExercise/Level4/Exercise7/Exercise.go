package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not Stired"}
	y := []string{"Miss", "Moneypenny", "Hellooooo, James"}
	z := [][]string{x, y}

	for i, v := range z {
		fmt.Println(i)
		for j, v1 := range v {
			fmt.Printf("\t%v %v\n", j, v1)
		}
	}
}
