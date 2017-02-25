package test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func GetMultiUri(strs []string) {
	start := time.Now()
	s := strs[:]
	for i := 0; i < 10000; i++ {
		fmt.Printf("Nums: %v\n", i)
		s = append(s, strs[0], strs[1], strs[2])
	}
	fmt.Printf("New strs: %v\n", len(s))
	//os.Exit(0)
	ch := make(chan string)
	for _, url := range s {
		go fetch(url, ch) // start a goroutine
	}
	for range s {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes := io.LimitReader(resp.Body, 0)
	buf, err := ioutil.ReadAll(nbytes)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%s\n%v\n", secs, url, string(buf))
}
