package main

import (
	"garden/config"
	"garden/pkg/router"
	_ "garden/pkg/router/all"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	env, err := config.GetEnvironment(dir + "/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	/*
		if env.Database != nil {
			db, err := sqlx.Open("mysql", cfg.Databases.MySQL)
			if err != nil {
				log.Fatal(err)
			}
			db.Ping()
		}
	*/
	log.Printf("Garden is running with verison %s.", env.Version)
	r := gin.Default()
	act, ok := router.Routers["/albums"]
	if !ok {
		log.Fatal("Undefined but requested input: ablums")
	}
	r.GET("/albums/:id", act().Get)
	r.POST("/albums", act().Post)

	r.Run("localhost:8080")
	/*
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
	*/
}
