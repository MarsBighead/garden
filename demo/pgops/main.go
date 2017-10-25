package main

import (
	"database/sql"
	"fmt"

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
	/*no := 1
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
	}*/
	generalQuery(db)
}

func generalQuery(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM test_b")
	if err != nil {
		log.Fatal("Fetch data err ", err)
	}
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))

	for rows.Next() {
		values := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		err = rows.Scan(scanArgs...)
		fmt.Printf("all rows %v\n", values)
	}
}
