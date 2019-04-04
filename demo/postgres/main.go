package main

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"log"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=test sslmode=disable password=postgres")

	if err != nil {
		log.Fatal("connect err ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Ping err ", err)
	}

	std := &student{}
	/*t := reflect.TypeOf(x)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
	}
	*/
	s := reflect.ValueOf(std).Elem()
	//v := reflect.ValueOf(&x).Elem()
	s.Field(0).SetString("Paul")
	s.Field(1).SetInt(24)
	fmt.Println("X is ", std)
}

type student struct {
	Name string
	ID   int
}

func bind(v interface{}) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		s := reflect.ValueOf(v).Elem()
		s.Field(0).SetInt(123)
		sliceValue := reflect.ValueOf([]int{1, 2, 3})
		s.FieldByName("Children").Set(sliceValue)
	}

}

//General query sample for PostgreSQL
func General(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM test_b")
	if err != nil {
		return err
	}
	columns, _ := rows.Columns()
	fmt.Println(columns)
	scanArgs := make([]interface{}, len(columns))

	for rows.Next() {
		values := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		err = rows.Scan(scanArgs...)
		if err != nil {
			return err
		}
		fmt.Printf("all rows %v\n", values)
	}
	return nil
}

const (
	gophers = 10
	entries = 10000
)

// Sessions test multiple session connection with Goroutine, PostgreSQL dump: pg_dump -U postgres -O test  -t test>test.sql
func Sessions() {
	// create string to pass
	sStmt := "insert into test (gopher_id, created) values ($1, $2)"
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

type data struct {
	id   int
	name string
}

func transaction(db *sql.DB) {
	name := "Bighead"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	d := new(data)
	d.id = 1
	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	err = txn.QueryRow(`select id,name from test where id=$1 for update`, d.id).Scan(&d.id, &d.name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)
	stmt, err := txn.Prepare(`update test set name=$1 where id=$2`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(name, d.id)
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func copy(db *sql.DB) error {
	// lazily open db (doesn't truly open until first request)

	txn, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := txn.Prepare(pq.CopyIn("test", "gopher_id", "created"))
	if err != nil {
		return err
	}

	for i := 1; i <= gophers*entries; i++ {
		_, err = stmt.Exec(i, time.Now().Format("2006-01-02 15:04:05.999999999Z07:00"))
		if err != nil {
			return err
		}
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}
	// close db
	return db.Close()

}

func gopher2(sStmt string) {

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
