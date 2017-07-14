package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu   sync.Mutex
		x, y int
		wg   sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		x = 1                              // A1
		fmt.Println("f1 y:", y, "; x:", x) // A2
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		y = 1                              // B1
		fmt.Println("f2 y:", y, "; x:", x) // A2
	}()
	wg.Wait()
}
