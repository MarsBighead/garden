package main

import (
	"fmt"
	"testing"
)

func TestTax2011QuickDeduction(t *testing.T) {
	fmt.Println("Year 2011:")

	afterTax, qdWithTax := tax2011QuickDeduction(0)
	fmt.Printf("Atfer Tax:%v\n", afterTax)
	for k, v := range qdWithTax {
		if k > 0 {
			fmt.Printf("%v: %v\n", k/12.0, v)
		}
	}
	afterTax, qdWithoutTax := tax2011QuickDeduction(1)
	fmt.Printf("Atfer Tax:%v\n", afterTax)
	for k, v := range qdWithoutTax {
		if k > 0 {
			fmt.Printf("%v: %v\n", k/12.0, v)

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
