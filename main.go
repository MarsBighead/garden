package main

import (
	"garden/config"
	"garden/model"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := config.GetDBConfig()
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	model.TruncateTable("chr", db)
	model.InsertVal(db)
	model.DumpLoad("chr", db)
	model.GetRows(db)
	os.Exit(1)
}
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
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

func apiMatrix(w http.ResponseWriter, r *http.Request) {
	userFile := "mock.json"
	file, err := ioutil.ReadFile(userFile)
	fmt.Printf("User Agent: %v\n", r.Header.Get("User-Agent"))
	// fmt.Printf("Open file in func hello!\n")
	check(err)
	//fmt.Print(string(file))
	fmt.Fprintf(w, string(file)) //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/m/api/matrix", apiMatrix)
	err := http.ListenAndServe(":8001", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
