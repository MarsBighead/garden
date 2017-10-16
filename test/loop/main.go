package main

import (
	"fmt"
)

func main() {
	i := 0
Loop:
	fmt.Printf("%d\n", i)
	if i < 10 {
		i++
		goto Loop
	}
}
