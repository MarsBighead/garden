package main

import (
	"fmt"
	"garden/config"
	"garden/model"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Service struct {
}

func main() {
	dir, err := currentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	db, err := config.GetDB(*dir)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	fmt.Printf("db right or not?\n")
	Server()
	select {}

	/*	model.TruncateTable("chr", db)
		model.InsertVal(db)
		model.DumpLoad("chr", db)
		model.GetRows(db)
		os.Exit(1)*/
}

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse parameters, default none
	fmt.Println(r.Form)
	fmt.Printf("User Agent: %v\n", r.Header.Get("User-Agent"))
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	var content = "Hello!\n" + "<a href=\"\">matrix API</a>"

	//w.Write([]byte("Hello<a href=\"\">matrix API</a>"))
	fmt.Fprintf(w, content) //这个写入到w的是输出到客户端的
}

//Server Start an test server
func Server() {
	route()
	http.HandleFunc("/list", model.HomeList)
	err := http.ListenAndServe(":8001", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func currentDirectory() (*string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}
	return &dir, nil
}
