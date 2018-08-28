package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	ch := make(chan string, 1)
	go func() {
		log.Println("build ch 1")
		time.Sleep(20 * time.Second)
		ch <- "result 1"
	}()
	timeout := make(chan bool, 1)
	fmt.Println("1. bool timeout", len(timeout))
	go func() {
		time.Sleep(1 * time.Second)
		<-timeout
	}()
	select {
	case res := <-ch:
		log.Println(res)
	case <-time.After(10 * time.Second):
		timeout <- true
		fmt.Println("2. bool timeout", len(timeout))
		log.Println("timeout 1")
	}

	fmt.Println("3. bool timeout", len(timeout))

	// If we allow a longer timeout of 3s, then the receive
	// from `c2` will succeed and we'll print the result.
	c := make(chan string, 1)
	go func() {
		log.Println("build c 2")
		time.Sleep(20 * time.Second)
		c <- "result 2"
	}()
	select {
	case res := <-c:
		log.Println(res)
	case <-time.After(30 * time.Second):
		log.Println("timeout 2")
	}
}
