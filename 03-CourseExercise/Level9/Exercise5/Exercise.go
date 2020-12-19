package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var v int64
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&v, 1)
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&v))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("At end", v)
}
