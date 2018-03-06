package main

import (
	"fmt"
	"runtime"
)

var maxTry = 3

func main() {
	var vcs []vc
	for i := 1; i <= 5; i++ {
		v := vc{
			id:   i,
			name: fmt.Sprintf("test%d", i),
		}
		vcs = append(vcs, v)
	}
	vcChan := make(chan vc)
	//ch2 := make(chan int)
	for _, v := range vcs {
		go poll(v, vcChan)
	}
	var cnt *int
	fmt.Println("Before number", runtime.NumGoroutine())
	go monitor(vcChan)
	fmt.Println("goroutine number", runtime.NumGoroutine())

	fmt.Println("Final Goroutine numbers", runtime.NumGoroutine(), cnt)
	for {
	}
	//	go pump2(ch2)
	//go suck(ch1, ch2)
}

type vc struct {
	id   int
	name string
	try  int
}

func poll(v vc, vcChan chan vc) {
	if v.id == 2 {
		vcChan <- v
	}
}
func monitor(vcChan chan vc) {
	v := <-vcChan
	fmt.Println("monitor online! Get value from channel", v)
	v.try++
	if v.try < maxTry {
		go poll(v, vcChan)
	}
	fmt.Println("Before number", runtime.NumGoroutine())
	go monitor(vcChan)
	fmt.Println("monitor offline!", runtime.NumGoroutine())

}
