package main

import (
	"fmt"
	"testing"
)

func TestTax2011QuickDeduction(t *testing.T) {
	fmt.Println("Year 2011:")
	rate := tax2011Rate()
	beforeTax := []float64{
		18000,
		54000,
		108000,
		420000,
		660000,
		960000,
		-1,
	}
	afterTax, qdWithTax := quickDeduction(0, beforeTax, rate)
	fmt.Printf("Atfer Tax:%v\n", afterTax)
	for k, v := range qdWithTax {
		if k > 0 {
			fmt.Printf("%v: %v\n", k/12.0, v)
		}
	}
	afterTax, qdWithoutTax := quickDeduction(1, beforeTax, rate)
	fmt.Printf("Atfer Tax:%v\n", afterTax)
	for k, v := range qdWithoutTax {
		if k > 0 {
			fmt.Printf("%v: %v\n", k/12.0, v)

		}
	}
}
