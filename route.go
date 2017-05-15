package main

import (
	"garden/model"
	"net/http"
)

func route() {
	http.HandleFunc("/", model.Home)
	http.HandleFunc("/home", model.Home)
	http.HandleFunc("/json", model.ProtocolJSON)
	http.HandleFunc("/test/protocol", model.Home)
}
