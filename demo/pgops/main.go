package main

import (
	"database/sql"
	"fmt"
	"reflect"

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
	//	generalQuery(db)
	x := &X{}
	/*t := reflect.TypeOf(x)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
	}
	*/
	s := reflect.ValueOf(x).Elem()
	//v := reflect.ValueOf(&x).Elem()
	s.Field(0).SetString("Paul")
	s.Field(1).SetInt(24)
	fmt.Println("X is ", x)
}

type X struct {
	Name string
	ID   int
}

func Bind(v interface{}) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		s := reflect.ValueOf(v).Elem()
		s.Field(0).SetInt(123)
		sliceValue := reflect.ValueOf([]int{1, 2, 3})
		s.FieldByName("Children").Set(sliceValue)
	}

}

func generalQuery(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM test_b")
	if err != nil {
		log.Fatal("Fetch data err ", err)
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
		fmt.Printf("all rows %v\n", values)
	}
}
