package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	threadNum := 4
	length := 200
	start := time.Now()
	chChan(threadNum, length)
	sp1 := time.Now()
	fmt.Println("type 1 duration", sp1.Sub(start))
	wgChan(threadNum, length)
	sp2 := time.Now()
	fmt.Println("type 2 duration", sp2.Sub(sp1))
}

func chChan(threadNum int, length int) {
	var maxG int
	done := make(chan bool)
	defer close(done)
	for n := 0; n < length; {
		var i int
		for i = 0; i < threadNum; i++ {
			n++
			if n >= length {
				break
			}
			go func() {
				done <- true
			}()

		}

		// wait until both are done
		for c := 0; c < i; c++ {
			if c == 0 {
				if runtime.NumGoroutine() > maxG {
					maxG = runtime.NumGoroutine()
				}
			}
			<-done
		}
	}
	fmt.Println("max goroutine is", maxG)
}
func wgChan(threadNum int, length int) {
	var maxG int
	threadKey := make(chan int, threadNum)
	var wg sync.WaitGroup
	for n := 0; n < length; n++ {
		wg.Add(1)
		threadKey <- n
		go func(i int) {
			defer func() {
				wg.Done()
				<-threadKey
			}()
			if runtime.NumGoroutine() > maxG {
				maxG = runtime.NumGoroutine()
			}
			//fmt.Println(time.Now(), "Goruntine number", runtime.NumGoroutine())
		}(n)
	}
	wg.Wait()
	close(threadKey)
	fmt.Println("max goroutine is", maxG)
}
