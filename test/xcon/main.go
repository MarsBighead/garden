package main

import (
	"fmt"
	"time"
)

func dgucoproducer(c chan int, max int) {
	for i := 0; i < max; i++ {
		c <- i
	}
	//close(c)
}

func dgucoconsumer(c chan int) {
	ok := true
	value := 0
	for ok {
		fmt.Println("Wait receive")
		if value, ok = <-c; ok {
			fmt.Println(value)
		}
		if ok == false {
			fmt.Println("*******Break********")
		}
	}
}

func main() {
	c := make(chan int)
	defer close(c)
	go dgucoproducer(c, 10)
	go dgucoconsumer(c)
	time.Sleep(time.Second * 5)
	fmt.Println("Done")
}
