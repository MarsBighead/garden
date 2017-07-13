package main

import (
	"fmt"
	"time"
)

func main() {
	sizes := make(chan int64)
	go func() {
		sizes <- 128
		close(sizes)
	}()
	time.Sleep(10000000)
	var total int64
	for size := range sizes {
		total += size
	}
	fmt.Printf("Total size %d\n", total)
}
