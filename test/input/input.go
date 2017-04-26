package main

import "fmt"

func main() {
	var ai int
	var ad float64
	var as string
	i := 4
	d := 4.0
	s := "Hackerrank "
	fmt.Scanln(&ai)
	fmt.Scanln(&ad)
	fmt.Scanln(&as)
	fmt.Printf("%d\n", i+ai)
	fmt.Printf("%.1f\n", d+ad)
	fmt.Printf("%s%s\n", s, as)
}
