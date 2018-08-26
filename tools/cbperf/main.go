package main

import (
	"database/sql"
	"garden/bio"
	"garden/model"
	"log"
	"os"

	"github.com/couchbase/gocb"
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
	s := bio.Service{
		DB: db,
	}
	cluster, err := gocb.Connect("couchbase://localhost")
	if err != nil {
		log.Fatal(err)
	}
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "root",
		Password: "togerme",
	})
	bucket, err := cluster.OpenBucket("refgene", "")
	if err != nil {
		log.Fatal(err)
	}
	s.QueryAllRefGene(bucket)

}
