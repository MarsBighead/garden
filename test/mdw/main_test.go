package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"testing"
)

func TestMiddleware(t *testing.T) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	req, _ := http.NewRequest("GET", "http://localhost:9201", nil)
	fmt.Println("work runtime.NumGoroutine() is", runtime.NumGoroutine())
	req.Header.Set("X-Request-ID", "68")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("work runtime.NumGoroutine() is", runtime.NumGoroutine())
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println("work runtime.NumGoroutine() is", runtime.NumGoroutine())

}
