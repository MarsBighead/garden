package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 100000)
	t := time.Tick(time.Second)

	go func() {
		for {
			select {
			case <-t:
				watching()
			}
		}
	}()

	for i := 0; i < 10000000; i++ {
		c <- true
		go worker(i, c)
	}

	fmt.Println("Done")
}

func watching() {
	fmt.Printf("Num Goroutine: %d\n", runtime.NumGoroutine())
}

func worker(i int, c chan bool) {
	//fmt.Println("worker", i)
	time.Sleep(100 * time.Microsecond)
	<-c
}
