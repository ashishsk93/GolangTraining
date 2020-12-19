package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	v := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			c := v
			c++
			fmt.Println(c)
			runtime.Gosched()
			v = c
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("At end", v)
}
