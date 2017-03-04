package main

import (
	"fmt"
	"garden/bank"
	"time"
)

func main() {
	// Alice:
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
	}()
	// Bob:
	go bank.Deposit(100)
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
