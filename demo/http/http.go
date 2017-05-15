package main

import (
	"fmt"
	"garden/model"
	"log"
	"net/http"
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
func status204(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	fmt.Printf("Test 204 StatusCode\n")
	//w.Write([]byte(""))
}

func say(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/", model.Home)
	http.HandleFunc("/header", testHeader)
	http.HandleFunc("/tpl", model.HomeTemplate)
	http.HandleFunc("/204", status204)
	http.Handle("/handle", http.HandlerFunc(say))
	log.Printf("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	select {}
}
