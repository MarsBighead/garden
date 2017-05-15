package main

import (
	"garden/model"
	"net/http"
)

func route() {
	http.HandleFunc("/home", model.Home)
	http.HandleFunc("/json", model.ProtocolJSON)
	http.HandleFunc("/test/protocol", model.ProtocalHTTP)
	http.HandleFunc("/tpl", model.HomeTemplate)
	// http.HandleFunc("/", model.Home)
}
