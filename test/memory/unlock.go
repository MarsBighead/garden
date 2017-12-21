package memory

import (
	"fmt"
	"log"
	"time"
)

var x, y int

func unlock() {
	log.Println("Start test unlock variable in different goroutines...")
	go func() {
		x = 1
		fmt.Println("f1 y:", y, "; x:", x) // A2
	}()
	go func() {
		y = 1
		fmt.Println("f2 y:", y, "; x:", x) // A2
	}()
	time.Sleep(1 * time.Second)
}
