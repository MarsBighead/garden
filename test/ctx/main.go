package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	now := time.Now()
	fmt.Println("start time", now)
	defer cancel()
	wg.Add(1)
	go tctx(ctx)
	wg.Wait()
	fmt.Println("end time", time.Now())
}

func tctx(ctx context.Context) {
	defer wg.Done()
	c := make(chan int, 1)
	go func() {
		fmt.Println("Test ctx at", time.Now())
		time.Sleep(4 * time.Second)
		c <- 1
		fmt.Println("Test end ctx at", time.Now())
	}()
	select {
	case <-ctx.Done():
		fmt.Println("Timeout!")
	case n := <-c:
		fmt.Println("channel flag is", n)
	}
	return
}
