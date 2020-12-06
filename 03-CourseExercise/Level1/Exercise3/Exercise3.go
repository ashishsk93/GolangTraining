package main

import (
	"fmt"
)

var x1 = 42
var y1 = "James Bond"
var z1 = true

func main() {
	s := fmt.Sprintf("X: %v\tY: %v\tZ:%v", x1, y1, z1)
	fmt.Println(s)
}
