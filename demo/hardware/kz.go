package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"
)

var MaxConcurrency = 4
var JobNum = 10000000000

type Workload struct {
	seq   int
	sleep time.Duration
}

func jobCallback(workload Workload) {
	log.Printf("handle job %d\n", workload.seq)
	time.Sleep(workload.sleep)
}

func kz() {
	concurrencyCh := make(chan bool, MaxConcurrency)
	jobChan := make(chan Workload, 100000)

	// job producer
	go func() {
		for i := 0; i < JobNum; i++ {
			workload := Workload{
				seq:   i,
				sleep: time.Second * time.Duration(rand.Intn(10)),
			}
			jobChan <- workload
		}
		log.Printf("sent all jobs")
	}()

	// a goroutine that prints the goroutine number
	go func() {
		for {
			select {
			case <-time.Tick(time.Second * 1):
				fmt.Printf("goroutine num: %d\n", runtime.NumGoroutine())
			}
		}
	}()

	for {
		select {
		case concurrencyCh <- true:
			select {
			case workload := <-jobChan:
				go func() {
					defer func() {
						<-concurrencyCh
					}()
					jobCallback(workload)
				}()
			}
		}
	}
}
