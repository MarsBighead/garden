package model

import (
	"fmt"
	"html/template"
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

func HttpTemp(w http.ResponseWriter, r *http.Request) {
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
	t, err := template.ParseFiles(FormatPath("/template/t.htm"))

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

func FormatPath(filename string) string {
	gdPath := os.Getenv("GARDEN")
	fmt.Printf("Get env variable $GARDEN=%s\n", gdPath)
	return gdPath + filename
}
