package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("In go routine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("In go routine 2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done with main")
}
