package main

import (
	"fmt"
	"testing"
)

func TestTax2011QuickDeduction(t *testing.T) {
	fmt.Println("Year 2011:")
	tqd := tax2011QuickDeduction()
	for k, v := range tqd {
		if k > 0 {
			fmt.Printf("%v: %v\n", k/12, v/12)
		}
	}
}

func TestTax2018QuickDeduction(t *testing.T) {
	tqd := tax2018QuickDeduction()
	fmt.Println("Year 2018:")
	for k, v := range tqd {
		if k > 0 {
			fmt.Printf("%v: %v\n", k/12, v/12)
		}
	}
}
