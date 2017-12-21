package base

import (
	"fmt"
	"strings"
)

// R Even and odd result structure
type R struct {
	even string
	odd  string
}

// EvenAndOdd Topic from hackerrank
func EvenAndOdd() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t int
	var s string
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&s)
		result := evenAndOdd(s)
		fmt.Printf("%s %s\n", result.even, result.odd)
	}
}

func evenAndOdd(s string) (data *R) {
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
