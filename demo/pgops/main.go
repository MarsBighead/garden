package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"log"
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
	no := 1
	rows, err := db.Query("SELECT student_name,age FROM student WHERE no >=$1", no)
	if err != nil {
		log.Fatal("Fetch data err ", err)
	} else {
		for rows.Next() {
			var age int
			var studentName string
			err = rows.Scan(&studentName, &age)
			fmt.Printf("name=%s, id=%d\n", studentName, age)
		}
	}
}
