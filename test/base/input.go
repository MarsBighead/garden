package base

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func input() {
	var i uint64 = 4
	d := 4.0
	s := "HackerRank "
	var ai uint64
	var ad float64
	var as string
	scanner := bufio.NewScanner(os.Stdin)
	/*fmt.Scanln(&ai)
	fmt.Scanln(&ad)
	fmt.Scanln(&as)*/
	n := 0
	for scanner.Scan() {
		switch n {
		case 0:
			ai, _ = strconv.ParseUint(scanner.Text(), 10, 0)
		case 1:
			ad, _ = strconv.ParseFloat(scanner.Text(), 64)
		case 2:
			as = scanner.Text()
		}
		n++
		if n >= 3 {
			break
		}
	}

	fmt.Printf("%d\n", i+ai)
	fmt.Printf("%.1f\n", d+ad)
	fmt.Printf("%s%s\n", s, as)
}
