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
	"sync"
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
	var sStmt string = "insert into test (gopher_id, created) values ($1, $2)"
	var wg sync.WaitGroup
	for i := 0; i < gophers; i++ {
		wg.Add(1)
	}
	// run the insert function using 10 go routines
	fmt.Printf("StartTime: %v\n", time.Now())

	for i := 0; i < gophers; i++ {

		go func(i int, sStmt string) {
			defer wg.Done()
			gopher(i, sStmt)
			//fmt.Printf("Gopher is %d,\n%s\n", i, sStmt)
		}(i, sStmt)
		// spin up a gopher
		// ~= 6.46s
		//go gopher(i, sStmt)
		//gopher(i, sStmt)
	}

	// this is a simple way to keep a program open
	// the go program will close when a key is pressed
	wg.Wait()
	fmt.Printf("StopTime: %v\n", time.Now())
	//var input string
	//fmt.Scanln(&input)
}

func gopher(gopher_id int, sStmt string) {

	// lazily open db (doesn't truly open until first request)
	db, err := sql.Open("postgres", "host=localhost user=postgres dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Gopher Id: %v || StartTime: %v\n", gopher_id, time.Now())
	base := gopher_id * entries
	for i := 0; i < entries; i++ {

		stmt, err := db.Prepare(sStmt)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec(i+base, time.Now())
		if err != nil || res == nil {
			log.Fatal(err)
		}

		stmt.Close()

	}
	// close db
	db.Close()
	fmt.Printf("Gopher Id: %v || StopTime: %v\n", gopher_id, time.Now())

}
