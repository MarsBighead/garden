package main

import (
	"fmt"
	"garden/config"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
)

// Config  configure file for application
type Config struct {
	Application string `toml:"application"`
	Databases   struct {
		MySQL string `toml:"mysql"`
	} `toml:"databases"`
	Directory string
}

func main() {
	dir, err := currentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	db, err := config.ConnectMySQL(dir)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	fmt.Printf("db right or not?\n")
	route()
	err = http.ListenAndServe(":8001", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Printf("Running Server on http://localhost:8001")
	select {}
}

//Databases Start an test server
func Databases() {
	/*	model.TruncateTable("chr", db)
		model.InsertVal(db)
		model.DumpLoad("chr", db)
		model.GetRows(db)
		os.Exit(1)*/
}

func currentDirectory() (dir string, err error) {
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func readConfig(dir string) (*Config, error) {
	var cfg Config

	_, err := toml.DecodeFile(dir+"/config.toml", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
