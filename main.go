package main

import (
	"garden/bio"
	"garden/model"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Printf("Garden is running now")
	route()
	err := http.ListenAndServe(":8001", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Printf("Running Server on http://localhost:8001")
	select {}
}

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
	http.HandleFunc("/hg38/refgene/modes", bio.PayloadHg38Modes)
	http.HandleFunc("/hg38/refgene/gene", bio.PayloadHg38RefGene)
	http.HandleFunc("/", model.Home)
}

//Databases Start an test server
func Databases() {
	/*	model.TruncateTable("chr", db)
		model.InsertVal(db)
		model.DumpLoad("chr", db)
		model.GetRows(db)
		os.Exit(1)*/
}
