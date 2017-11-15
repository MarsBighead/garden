package main

// schema we can use along with some select statements
// create table test ( gopher_id int, created timestamp );
// select * from test order by created asc limit 1;
// select * from test order by created desc limit 1;
// select count(created) from test;

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	gophers = 10
	entries = 10000
)

type TestB struct {
	ID   int    `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=test sslmode=disable")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()
	// create string to pass
	//var sStmt string = "insert into test (gopher_id, created) values ($1, $2)"
	var testB []TestB
	query := db.Find(&testB)
	if query.Error != nil {
		panic(query.Error)
	}
	fmt.Println(testB)
}
