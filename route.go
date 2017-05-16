package main

import (
	"garden/model"
	"net/http"
)

func route() {
	http.HandleFunc("/home", model.Home)
	http.HandleFunc("/list", model.HomeList)
	http.HandleFunc("/pbt", model.Pbt)
	http.HandleFunc("/aes", model.AES)
	http.HandleFunc("/reproto", model.RebuildPbt)
	http.HandleFunc("/api/xiaodu", model.FromXiaodu)
	http.HandleFunc("/json", model.ProtocolJSON)
	http.HandleFunc("/statistic", model.HomeStatistic)
	http.HandleFunc("/statistics", model.AdvancedStatistic)
	http.HandleFunc("/test/protocol", model.ProtocalHTTP)
	http.HandleFunc("/tpl", model.HomeTemplate)
	http.HandleFunc("/", model.Home)
}
