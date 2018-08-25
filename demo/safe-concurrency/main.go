package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	max := 200
	ch := make(chan int, max)
	done := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("len vs cap", len(ch), cap(ch))
	start := time.Now()
	total := max * 2000
	go func() {
		defer wg.Done()
		producer(ch, total)
		close(ch)
		for i := 0; i < max; i++ {
			done <- true
		}
		close(done)
	}()
	var count int
	num := make(chan int, max)
	for i := 0; i < max; i++ {
		go func() {
			var sig bool
			var n int
			for !sig {
				select {
				case _, ok := <-ch:
					if ok {
						n++
						count++
						time.Sleep(1 * time.Millisecond)
						//log.Printf("Sequence %d, running goroutines number %d\n", v, runtime.NumGoroutine())
					}
				case sig = <-done:

				}
			}
			num <- n
		}()
	}
	wg.Wait()
	end := time.Now()
	close(num)
	fmt.Printf("Total %d, count %d, time cost is %v\n", total, count, end.Sub(start))
	count = 0
	for n := range num {
		count += n
	}
	fmt.Printf("Total %d, channel count %d, time cost is %v\n", total, count, end.Sub(start))
}

func producer(ch chan int, total int) {
	for i := 0; i < total; i++ {
		ch <- i
	}
	fmt.Printf("Length of loop is %d\n", total)
}
