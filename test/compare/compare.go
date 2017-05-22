package main

import (
	"fmt"
	"garden/demo/bank"

	"time"
)

func main() {
	// Alice:
	name := "Alice"
	go func() {
		bank.Deposit(200, name)
	}()
	time.Sleep(1 * time.Second)
	// Bob:
	name = "Bob"
	go bank.Deposit(100, name)
	delayTag(1000000 * time.Second)
}

func delayTag(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)

		}

	}
}
