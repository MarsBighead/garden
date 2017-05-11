package main

import (
	"garden/model"
	"net/http"
)

func route() {
	http.HandleFunc("/", model.Home)
	//http.HandleFunc("/t", model.Test)
}
