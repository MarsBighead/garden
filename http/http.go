package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Person struct {
	Name    string
	Age     int
	Emails  []string
	Company string
	Role    string
}

type OnlineUser struct {
	User      []*Person
	LoginTime string
}

func handler(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/", handler)
	http.HandleFunc("/t", tempHandler)
	http.HandleFunc("/hello", testStatusNoContent)
	http.Handle("/handle", http.HandlerFunc(say))
	log.Printf("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	select {}
}

func tempHandler(w http.ResponseWriter, r *http.Request) {
	dumx := Person{
		Name:    "zoro",
		Age:     27,
		Emails:  []string{"dg@gmail.com", "dk@hotmail.com"},
		Company: "Omron",
		Role:    "SE"}

	chxd := Person{
		Name:   "chx",
		Age:    26,
		Emails: []string{"test@gmail.com", "d@hotmail.com"}}

	onlineUser := OnlineUser{User: []*Person{&dumx, &chxd}}

	//t := template.New("Person template")
	//t, err := t.Parse(templ)
	t, err := template.ParseFiles("template/t.htm")
	checkError(err)

	err = t.Execute(w, onlineUser)
	checkError(err)
}

// checkError -Simplify error return checking
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
