package main

import (
	"fmt"
	"runtime"
	"time"
)

var c = make(chan int, 6)
var count = 0
var done = make(chan int, 3)

func main() {
	//Consumer
	start := time.Now()
	for i := 0; i < 3; i++ {
		go consumer(i)
	}
	// Producer
	for i := 0; i < 100; i++ {
		c <- i
	}
	/** here **/
	close(c)
	//Hold consumer channel
	fmt.Println("cnt goroutine", runtime.NumGoroutine())
	for i := 0; i < 3; i++ {
		fmt.Println(<-done)
	}
	close(done)
	fmt.Println("Main goroutine time cost:", time.Now().Sub(start))
	//time.Sleep(2 * time.Second)
	fmt.Println("cnt is", count)
}
func consumer(index int) {
	start := time.Now()
	for target := range c {
		fmt.Printf("Consumer %d; Producer:%d\n", index, target)
		index = target
		count++
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Consumer goroutine time cost", time.Now().Sub(start))
	done <- index
}
