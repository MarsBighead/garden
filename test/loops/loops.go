package main

import "fmt"

func main() {
	var i int

	fmt.Scan(&i)
	fmt.Printf("%d\n", i)
	for j := 1; j <= 10; j++ {
		fmt.Printf("%d x %d = %d\n", i, j, i*j)
	}
}
