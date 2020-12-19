package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user

// Len is part of sort.Interface.
func (b byAge) Len() int {
	return len(b)
}

// Swap is part of sort.Interface.
func (b byAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (b byAge) Less(i, j int) bool {
	return (b[i].Age < b[j].Age)
}

type byName []user

// Len is part of sort.Interface.
func (b byName) Len() int {
	return len(b)
}

// Swap is part of sort.Interface.
func (b byName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (b byName) Less(i, j int) bool {
	return (b[i].Last < b[j].Last)
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	b := byAge(users)

	sort.Sort(b)

	fmt.Println(users)

	b1 := byName(users)

	sort.Sort(b1)

	for i, v := range users {
		fmt.Println("Person #", i)
		fmt.Println("\tFirst", v.First, "Last:", v.Last)
		sort.Strings(v.Sayings)
		for id, val := range v.Sayings {
			fmt.Println("\t\t", fmt.Sprintf("%v", id)+".", val)
		}
	}
}
