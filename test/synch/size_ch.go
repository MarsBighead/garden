package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sizes := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	files := []int{1, 2, 3, 4, 5, 120}
	go func() {
		defer wg.Done()
		var sum int
		for _, file := range files {
			sum += file
		}
		fmt.Println("sum", sum)
		sizes <- sum
	}()
	go func() {
		wg.Wait()
		close(sizes)
	}()
	time.Sleep(10000000)
	var total int
	for size := range sizes {
		total += size
	}
	fmt.Printf("Total size %d\n", total)
}
