package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// solution1 : sync.Mutex
	// var mu sync.Mutex

	// solution2: atomic
	var count int64
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				//mu.Lock()
				atomic.AddInt64(&count, 1)
				//mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
