package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	lists := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&lists[i])
	}
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%v ", lists[i])
	}
	fmt.Println()
}
