package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func lazyHandler(w http.ResponseWriter, req *http.Request) {
	ranNum := rand.Intn(2)
	if ranNum == 0 {
		time.Sleep(6 * time.Second)
		fmt.Fprintf(w, "slow response, %d\n", ranNum)
		fmt.Printf("slow response, %d\n", ranNum)
		return
	}
	fmt.Fprintf(w, "quick response, %d\n", ranNum)
	fmt.Printf("quick response, %d\n", ranNum)
	return
}

func main() {
	http.HandleFunc("/", lazyHandler)
	http.ListenAndServe(":9200", nil)
}
