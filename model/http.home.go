package model

import (
	"html/template"
	"net/http"
)

func HttpHome(w http.ResponseWriter, r *http.Request) {
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
	t, err := template.ParseFiles(FormatPath("/template/home.htm"))

	checkError(err)

	err = t.Execute(w, onlineUser)
	checkError(err)
}
