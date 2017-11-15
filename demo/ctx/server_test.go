package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
	"testing"
	"time"
)

var (
	wg sync.WaitGroup
)

type ResPack struct {
	r   *http.Response
	err error
}

func work(ctx context.Context) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	defer wg.Done()
	c := make(chan ResPack, 1)

	req, _ := http.NewRequest("GET", "http://localhost:9200", nil)
	go func() {
		resp, err := client.Do(req)
		pack := ResPack{r: resp, err: err}
		c <- pack
	}()

	fmt.Println("work runtime.NumGoroutine() is", runtime.NumGoroutine())
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c
		fmt.Println("Timeout!")
	case res := <-c:
		if res.err != nil {
			fmt.Println(res.err)
			return
		}
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s", out)
	}
	return
}

func TestServer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	now := time.Now()
	fmt.Println("start time", now)
	defer cancel()
	wg.Add(1)
	go work(ctx)
	wg.Wait()
	fmt.Println("end time", time.Now())
	fmt.Println("Finished")
}
