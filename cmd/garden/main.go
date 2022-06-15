package main

import (
	"context"
	"garden/config"
	"garden/pkg/router"
	_ "garden/pkg/router/all"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := config.NewConfig(ctx)
	env, err := cfg.NewEnvironment()
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
	urls := []string{
		"/albums",
		"/advertiser",
	}
	for _, uri := range urls {
		tackeAction(r, env, uri)
	}
	log.Printf("Running Server on http://localhost:8001")
	r.Run("localhost:8001")
	/*
		ps := page.NewService(r, db)
		http.HandleFunc("/home", ps.HomePage)
		http.HandleFunc("/index", ps.HomePage)
		http.HandleFunc("/list", ps.PageList)
		http.HandleFunc("/", ps.HomePage)
		err = http.ListenAndServe(":8001", nil) //设置监听的端口
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
		select {}
	*/
}

func tackeAction(r *gin.Engine, env *config.Environment, uri string) {
	act, ok := router.Routers[uri]
	if !ok {
		log.Fatalf("Undefined but requested input: %s\n", uri)
	}
	router := act(env)
	r.GET(uri+"/:id", router.Get)
	r.POST(uri, router.Post)
}
