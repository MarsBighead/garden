package channel

import (
	"fmt"
	"sync"
	"time"
)

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

func controller() {
	start := time.Now()
	limit := 3
	origin := limit
	ch := make(chan string, limit)
	var gen int
	for {
		if gen == 0 {
			go func() {
				for i := 0; i < limit; i++ {
					ch <- fmt.Sprintf("id %d", i)
				}
			}()
			gen++
		}

		if limit > origin {
			tmp := make(chan string, limit)
			extChanCap(ch, tmp)
			ch = tmp
			origin = limit
			//break
		}
		go func(ch chan string) {
			var control int
			fmt.Printf("Channel capacity %d, generate flag: %v, re-cosumer: ", cap(ch), gen)
			for s := range ch {
				fmt.Printf("%s ", s)
				control++
				ch <- s
				if control >= limit {
					break
				}
			}
			fmt.Println()
		}(ch)

		time.Sleep(10 * time.Second)
		if time.Now().Sub(start) > 3e10 && time.Now().Sub(start) <= 5e10 {
			fmt.Println(limit, "Continue from", start, "to", time.Now())
			limit++
			continue
		}
		if time.Now().Sub(start) > 5e10 {
			break
		}
		fmt.Println(limit, "Duration from", start, "to", time.Now())

	}

}

//extChanCap extend capacity of channel
func extChanCap(ch, tmp chan string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("cap: ch vs tmp %d:%d\n", cap(ch), cap(tmp))
		for i := 0; i < cap(ch); i++ {
			tmp <- <-ch
		}
		close(ch)
		tmp <- fmt.Sprintf("Id %d", cap(tmp)-1)
	}()
	wg.Wait()
}
