package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	limit := make(chan bool, 50)
	var wg sync.WaitGroup
	wg.Add(1)
	start := time.Now()
	go func() {
		defer wg.Done()
		assign(ch, 200)
	}()
	for v := range ch {
		limit <- true
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			log.Printf("Sequence %d, running goroutines number %d\n", v, runtime.NumGoroutine())
			<-limit
		}(v)
	}
	close(limit)
	wg.Wait()
	end := time.Now()
	fmt.Printf("time cost is %v\n", end.Sub(start))
}

func assign(ch chan<- int, max int) {
	for i := 0; i < max; i++ {
		ch <- i
	}
	fmt.Printf("Length of loop is %d\n", max)
	close(ch)

}
