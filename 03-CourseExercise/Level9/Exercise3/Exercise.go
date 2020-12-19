package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	v := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			c := v
			c++
			fmt.Println(c)
			runtime.Gosched()
			v = c
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("At end", v)
}
