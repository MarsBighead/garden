package main

import (
	"garden"
	"garden/model"
	"garden/page"
	"log"
	"net/http"
	"os"
	"path"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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
	db, err := sqlx.Open("mysql", cfg.Databases.MySQL)
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()

	//fmt.Printf("%#v\n", env)
	dir = dir + "/.."
	r := &garden.Router{
		Environment: map[string]string{
			"HOME":     path.Clean(dir + "/../"),
			"TEMPLATE": path.Clean(dir + "/../template"),
			"DATA":     path.Clean(dir + "/../data"),
		},
	}

	log.Printf("Garden is running now")
	ps := page.NewService(r, db)
	http.HandleFunc("/home", ps.HomePage)
	http.HandleFunc("/index", ps.HomePage)
	http.HandleFunc("/list", ps.PageList)
	http.HandleFunc("/", ps.HomePage)
	log.Printf("Running Server on http://localhost:8001")
	err = http.ListenAndServe(":8001", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	select {}
}
