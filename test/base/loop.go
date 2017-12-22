package base

import (
	"fmt"
	"time"
)

//GotoLoop Loop sample by goto,
//         difference in loop count, final control numbeer is the same
func GotoLoop() {
	i := 0
Loop:
	fmt.Printf("%d ", i)
	if i < 10 {
		i++
		goto Loop
	}
	fmt.Printf("\nEnd %d\n", i)
}

//NormalLoop Normal loop sample
func NormalLoop(i int) {
	fmt.Printf("Decomposition elements:%d\n", i)
	for j := 1; j <= 10; j++ {
		fmt.Printf("%d x %d = %d\n", i, j, i*j)
	}
}

func loopStrategy() {
	start := time.Now()
	var i int

	for {
		switch i {
		case 0:
			i++
		case 1:
			i = 1
		}
		time.Sleep(5 * time.Second)
		if time.Now().Sub(start) > 1e11 {
			fmt.Println(i, "Duration from", start, "to", time.Now())
			break
		}
		fmt.Println(i, "Duration from", start, "to", time.Now())

	}

}
