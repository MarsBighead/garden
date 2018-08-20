package main

import (
	"database/sql"
	"garden/bio"
	"garden/model"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := model.Parse(dir + "/config.toml")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", cfg.Databases.MySQL)
	if err != nil {
		log.Fatal(err)
	}
	service := bio.Service{
		DB: db,
	}
	log.Printf("Restful refGene  API is running now.")
	http.HandleFunc("/hg38/refgene/mode", service.Hg38RefgeneMode)
	http.HandleFunc("/hg38/refgene/gene", service.Hg38Refgene)
	log.Printf("Running Server on http://localhost:8001/hg38/refgene/mode?chrom=chrX")
	err = http.ListenAndServe(":8001", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	select {}
}
