package main

import (
	"fmt"
	"garden/model"
	"log"
	"net/http"
	"os"
)

func testHeader(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
	fmt.Printf("Hi there, I love %v!\n", r.URL.Path[1:])
	for k, v := range r.Header {
		fmt.Printf("Key -> Value: %s ->  %s\n", k, v)
		if len(v) >= 2 {
			fmt.Printf("Values from 1 :%v\n", v[1:])
		}
	}
}
func testStatusNoContent(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	fmt.Printf("Test 204 StatusCode\n")
	//w.Write([]byte(""))
}

func say(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/", model.HttpHome)
	http.HandleFunc("/header", testHeader)
	http.HandleFunc("/t", model.HttpTemp)
	http.HandleFunc("/hello", testStatusNoContent)
	http.Handle("/handle", http.HandlerFunc(say))
	log.Printf("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	select {}
}

// checkError -Simplify error return checking
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
