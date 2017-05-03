package main

import "fmt"

func main() {
	// fmt.Println("vim-go")
	var n int
	fmt.Scan(&n)
	dicts := make(map[string]string, n)
	for i := 0; i < n; i++ {
		var s, phone string
		fmt.Scan(&s)
		fmt.Scan(&phone)
		dicts[s] = phone
	}
	for i := 0; i < 100000; i++ {
		var s string
		fmt.Scanln(&s)
		if s == "" {
			break
		}
		if v, ok := dicts[s]; ok {
			fmt.Printf("%s=%s\n", s, v)
		} else {
			fmt.Printf("Not found\n")
		}
	}
}
