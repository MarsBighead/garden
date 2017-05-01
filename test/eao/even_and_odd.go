package main

import (
	"fmt"
	"strings"
)

type R struct {
	even string
	odd  string
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t int
	var s string
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&s)
		result := getEvenAndOdd(s)
		fmt.Printf("%s %s\n", result.even, result.odd)
	}
}

func getEvenAndOdd(s string) (data *R) {
	var even, odd string
	list := strings.Split(s, "")
	for i := 0; i < len(list); i++ {
		if i%2 == 0 {
			even += list[i]
		} else {
			odd += list[i]
		}
	}
	return &R{
		even: even,
		odd:  odd,
	}
}
