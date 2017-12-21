package base

import (
	"fmt"
)

//OutputList Out list element one by one
func OutputList(a []string) {
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Printf("%v\n", a[i])
	}
}

//InputList build list element via inputing one by one
func InputList(n int) []string {
	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	return a
}
