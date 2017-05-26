package main

import (
	"garden/bio"
	"garden/model"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg, err := model.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Open("mysql", cfg.Databases.MySQL)
	if err != nil {
		log.Fatal(err)
	}
	service := Service{
		DB:     db,
		Config: cfg,
	}
	log.Printf("Garden is running now")
	service.route()
	err = http.ListenAndServe(":8001", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Printf("Running Server on http://localhost:8001")
	select {}
}

// Service garden http service
type Service struct {
	DB     *sqlx.DB
	Config *model.Config
}

func (s *Service) route() {
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
	http.HandleFunc("/hg38/refgene/modes", bio.Hg38RefgeneModes)
	http.HandleFunc("/hg38/refgene/gene", bio.Hg38Refgene)
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
