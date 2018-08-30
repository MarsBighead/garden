package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"garden/bio"
	"garden/model"
	"io/ioutil"
	"log"
	"os"
	"time"

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
	body, err := ioutil.ReadFile(dir + "/config.json")
	if err != nil {
		log.Fatal(err)
	}
	cb := new(couchbase)
	err = json.Unmarshal(body, cb)
	if err != nil {
		log.Fatal(err)
	}
	/*b, err := json.MarshalIndent(cb, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("config.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	}*/

	cluster, err := gocb.Connect(fmt.Sprintf("%s://%s", cb.Driver, cb.Host))
	if err != nil {
		log.Fatal(err)
	}
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: cb.Username,
		Password: cb.Password,
	})
	bucket, err := cluster.OpenBucket(cb.Bucket, "")
	if err != nil {
		log.Fatal(err)
	}
	err = bucket.Manager(cb.Username, cb.Password).Flush()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	s.WriteBucket(bucket, 1000)

}

type couchbase struct {
	Driver   string `yaml:"driver"    json:"driver"`
	Host     string `yaml:"host"      json:"host"`
	Username string `yaml:"username"  json:"username"`
	Password string `yaml:"password"  json:"password"`
	Bucket   string `yaml:"bucket"    json:"bucket"`
	Switch   string `yaml:"switch"    json:"switch"`
	Port     string `yaml:"-"         json:"-"`
}
