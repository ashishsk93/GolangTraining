package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	s1 := foo(n...)
	fmt.Println(s1)
	s2 := bar(n)
	fmt.Println(s2)
}

func foo(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func bar(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}
