package main

import "fmt"

func main() {
	a := []string{`Shaken, not stirred`, `Martinis`, `Women`}
	b := []string{`James Bond`, `Literature`, `Computer Science`}
	c := []string{`Being evil`, `Ice cream`, `Sunsets`}
	x := map[string][]string{
		`bond_james`:      a,
		`moneypenny_miss`: b,
		`no_dr`:           c,
	}

	for k, v := range x {
		fmt.Println(k)
		for i, v1 := range v {
			fmt.Printf("\t%v %v\n", i, v1)
		}
	}
}
