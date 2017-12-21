package base

import "fmt"

//Maps Study map/hash method in golang
func Maps(n int) {
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
