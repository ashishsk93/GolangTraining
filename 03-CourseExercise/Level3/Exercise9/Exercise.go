package main

import "fmt"

func main() {
	favSport := "Cricket"
	switch favSport {
	case "Football":
		fmt.Println("Football")
	case "Cricket":
		fmt.Println("Cricket")
	case "Baseball":
		fmt.Println("Baseball")
	case "Basketball":
		fmt.Println("Basketball")
	default:
		fmt.Println("Neither")
	}
}
