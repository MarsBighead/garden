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
	var max int
	for v := range ch {
		limit <- true
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			cnt := runtime.NumGoroutine()
			if cnt > max {
				max = cnt
			}
			log.Printf("Sequence %d, running goroutines number %d\n", v, cnt)
			<-limit
		}(v)
	}
	close(limit)
	wg.Wait()
	end := time.Now()
	fmt.Printf("time cost is %v, max %d\n", end.Sub(start), max)
}

func assign(ch chan<- int, max int) {
	for i := 0; i < max; i++ {
		ch <- i
	}
	fmt.Printf("Length of loop is %d\n", max)
	close(ch)

}
