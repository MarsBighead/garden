package main

import (
	"database/sql"
	"fmt"
	"os"

	"log"

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

type data struct {
	id   int
	name string
}
