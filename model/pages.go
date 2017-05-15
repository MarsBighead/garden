package model

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

//HomeList list all handle in the top path
func HomeList(w http.ResponseWriter, r *http.Request) {
	log.Print("Running http handle model.HomeList!")
	t, err := template.ParseFiles(GetCurrentDir() + "/template/list.htm")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

// Home  home/index web page
func Home(w http.ResponseWriter, r *http.Request) {
	log.Print("Running http handle model.Home!")
	t, err := template.ParseFiles(GetCurrentDir() + "/template/home.htm")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// HomeTemplate  build web page with template
func HomeTemplate(w http.ResponseWriter, r *http.Request) {
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

	log.Print("Running http handle model.HomeTemplate!")
	t, err := template.ParseFiles(GetCurrentDir() + "/template/tpl.htm")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, onlineUser)
	if err != nil {
		log.Fatal(err)
	}
}

// ProtocalHTTP method for test http protocal output in dashboard
func ProtocalHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse parameters, default none
	fmt.Println(r.Form)
	fmt.Println("User-Agent:", r.Header.Get("User-Agent"))
	fmt.Println("HTTP scheme:", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	var content = "Hello!\n" + "<a href=\"\">matrix API</a>"

	w.Write([]byte(content)) //这个写入到w的是输出到客户端的
}
