package main

import (
	"fmt"
	"time"
)

var x, y int

func main() {
	go func() {
		x = 1
		fmt.Print("y:", y, " ")
	}()
	go func() {
		y = 1
		fmt.Print("x:", x, " ")
	}()
	time.Sleep(2 * time.Second)
}
