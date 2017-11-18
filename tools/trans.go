package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var ns = flag.Float64("ns", 1000000, "UPDATE goroutine jobs")

func main() {
	flag.Parse()
	if ns != nil {
		t := *ns * time.Nanosecond.Seconds()
		fmt.Println("nanoseconds to second", t)
	}

	start := time.Now()
	ops := 1000000
	/*for i := 0; i < ops; i++ {
		rand.Seed(time.Now().UnixNano())
		rand.Intn(3)
	}*/

	numConcurrency := 4
	concurrencyKey := make(chan int, numConcurrency)
	var wg sync.WaitGroup
	for n := 0; n < ops; n++ {
		wg.Add(1)
		concurrencyKey <- n
		go func(i int) {
			defer func() {
				wg.Done()
				<-concurrencyKey
			}()
			genRand()
		}(n)
	}
	fmt.Println("go routine numbers", runtime.NumGoroutine())
	wg.Wait()
	close(concurrencyKey)
	dur := time.Now().Sub(start)
	avg := float64(dur.Nanoseconds()) / float64(ops) / 1000
	fmt.Printf(`
##rand
  duration  %v
  ops number %v
  avg: %v macroseconds
`, dur, ops, avg)

}

func genRand() {
	rand.Seed(time.Now().UnixNano())
	rand.Intn(3)
}
