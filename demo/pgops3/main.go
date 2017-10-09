package main

// schema we can use along with some select statements
// create table test ( gopher_id int, created timestamp );
// select * from test order by created asc limit 1;
// select * from test order by created desc limit 1;
// select count(created) from test;

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const (
	gophers = 10
	entries = 10000
)

// Dump table strcuture
//  pg_dump -U postgres -O test  -t test>test.sql
func main() {

	// create string to pass
	fmt.Printf("StartTime: %v\n", time.Now())
	var sStmt string = "insert into test (gopher_id, created) values "
	gopher(sStmt)
	//var input string
	//fmt.Scanln(&input)
	fmt.Printf("StopTime: %v\n", time.Now())
}

func gopher(sStmt string) {

	// lazily open db (doesn't truly open until first request)
	db, err := sql.Open("postgres", "host=localhost user=postgres dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	var values []string
	for i := 1; i <= gophers*entries; i++ {
		v := fmt.Sprintf("( %d, '%v')", i, time.Now().Format("2006-01-02 15:04:05.999999999Z07:00"))
		values = append(values, v)
	}
	sStmt = sStmt + strings.Join(values, ",")
	res, err := db.Exec(sStmt)
	if err != nil || res == nil {
		log.Fatal(err)
	}
	// close db
	db.Close()

}
