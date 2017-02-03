package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
    fmt.Printf( "Hi there, I love %v!\n", r.URL.Path[1:])
    fmt.Printf( "Header %v!\n", r.Header.Get("Cache-Control"))
    fmt.Printf( "Header %v!\n", r.Header.Get("X-Forwarded-For"))
    for k, v := range r.Header {
        fmt.Printf("Key -> Value: %s ->  %s\n", k, v)
        if len(v)>=2 {
            fmt.Printf("Values from 1 :%v\n",  v[1:])
        }
    }
}
func hello(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello"))
}

func say(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello"))
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/hello", hello)
    http.Handle("/handle",http.HandlerFunc(say));
    http.ListenAndServe(":8080", nil)
    select{};
}
