package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}
	wg.Wait()
}
