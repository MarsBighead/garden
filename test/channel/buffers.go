package channel

import "fmt"

func buffers() {
	chs := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			chs <- i
		}
		// Remember that channel should be closed after pushed values
		close(chs)
	}()
	for v := range chs {
		fmt.Println("channel value", v)
	}
}
