package main

import "fmt"

func main() {

	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`ashish_kumar`] = []string{`Test1`, `Test2`, `Test3`}

	delete(m, `ashish_kumar`)

	for k, v := range m {
		fmt.Println(k)
		for i, v1 := range v {
			fmt.Printf("\t%v %v\n", i, v1)
		}
	}
}
