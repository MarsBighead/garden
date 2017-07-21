package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		sizes <- 128
	}()
	go func() {
		wg.Wait()
		close(sizes)
	}()
	time.Sleep(10000000)
	var total int64
	for size := range sizes {
		total += size
	}
	fmt.Printf("Total size %d\n", total)
}
