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
	"sync"
	"time"

	_ "github.com/lib/pq"
)

const (
	gophers = 10
	entries = 10000
)

func main() {

	// create string to pass
	//var sStmt string = "insert into test (gopher_id, created) values ($1, $2)"
	var sStmt string = "insert into test (gopher_id, created) values "
	var wg sync.WaitGroup
	for i := 0; i < gophers; i++ {
		wg.Add(1)
	}
	// run the insert function using 10 go routines
	// about 0.243s
	fmt.Printf("StartTime: %v\n", time.Now())
	for i := 0; i < gophers; i++ {
		/*go func(i int, sStmt string) {
			defer wg.Done()
			gopher(i, sStmt)
			//fmt.Printf("Gopher is %d,\n%s\n", i, sStmt)
		}(i, sStmt)*/
		// spin up a gopher
		go gopher(i, sStmt, &wg)
		//gopher(i, sStmt)
	}

	// this is a simple way to keep a program open
	// the go program will close when a key is pressed
	wg.Wait()
	fmt.Printf("StopTime: %v\n", time.Now())
	//var input string
	//fmt.Scanln(&input)
}
func gopher(gopher_id int, sStmt string, wg *sync.WaitGroup) {
	defer (*wg).Done()
	// lazily open db (doesn't truly open until first request)
	db, err := sql.Open("postgres", "host=localhost user=postgres dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Gopher Id: %v || StartTime: %v\n", gopher_id, time.Now())
	base := gopher_id * entries
	var values []string
	for i := 0; i < entries; i++ {
		v := fmt.Sprintf("( %d, '%v')", i+base, time.Now().Format("2006-01-02 15:04:05.999999999Z07:00"))
		values = append(values, v)

	}
	sStmt = sStmt + strings.Join(values, ",")
	res, err := db.Exec(sStmt)
	if err != nil || res == nil {
		log.Fatal(err)
	}
	// close db
	db.Close()
	fmt.Printf("Gopher Id: %v || StopTime: %v\n", gopher_id, time.Now())

}

/*
func gopher(gopher_id int, sStmt string, wg *sync.WaitGroup) {
	defer (*wg).Done()
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

}*/