package main

import (
	"fmt"
	"math"
)

type square struct {
	side int
}

type circle struct {
	radius int
}

func (s square) area() {
	fmt.Println("Area:", s.side*s.side)
}

func (c circle) area() {
	fmt.Println("Area:", math.Pi*float64(c.radius)*float64(c.radius))
}

type shape interface {
	area()
}

func info(s shape) {
	s.area()
}

func main() {
	s := square{
		side: 5,
	}

	c := circle{
		radius: 7,
	}

	info(c)
	info(s)
}
