package models

import (
	"database/sql"
	"fmt"
	"io"
	"math/rand"
	"strconv"

	"bytes"

	"github.com/go-sql-driver/mysql"
)

func GetRows(db *sql.DB) {

	// Prepare statement for reading data
	rows, err := db.Query("SELECT * FROM chr limit 1,5;")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	//fmt.Printf("Get data from table chr:\n%v\n", rows)
	var i int32
	n, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	//fmt.Printf("rows.Columns result: %s\n", n[0])
	for i, col := range n {
		if i >= len(n)-1 {
			fmt.Printf("%s\n", col)
		} else {
			fmt.Printf("%s\t", col)
		}
	}
	fmt.Printf("num\trs_id\tmaf\ttest\n")
	for rows.Next() {

		var rs string
		var maf string
		var test string
		err = rows.Scan(&rs, &maf, &test)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		fmt.Printf("%d\t%s\t%s\t%s\n", i, rs, maf, test)
		i++
	}
}

// InsertVal Insert value into database
func InsertVal(db *sql.DB) {
	rowIns, err := db.Prepare("INSERT INTO chr  VALUES( ?, ?, ? )") // ? = placeholder
	// Insert square numbers for 0-24 in the database
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rowIns.Close() // Close the statement when we leave main() / the program terminates
	for i := 6; i < 12; i++ {
		_, err := rowIns.Exec(i, (i * i), "xxx") // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
}

// TruncateTable Truncate data in the table
func TruncateTable(table string, db *sql.DB) {

	rowsTruncate, err := db.Prepare("Truncate " + table) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rowsTruncate.Close() // Close the statement when we leave main() / the program terminates
	_, err = rowsTruncate.Exec()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("Truncate table chr successfully!\n")
}

// DumpLoad Truncate data in the table
func DumpLoad(table string, db *sql.DB) {
	var content string
	for i := 11; i < 15; i++ {
		str := fmt.Sprintf("%d\t%d\t%s\n", i, i*i, "xxxxx")
		content += str
	}
	fmt.Printf("DumpLoad:\n%s\n", content)
	//buf := bytes.NewBufferString(content)
	readerName := "r" + strconv.Itoa(rand.Int())
	mysql.RegisterReaderHandler(readerName, func() io.Reader { return bytes.NewBufferString(content) })
	defer mysql.DeregisterReaderHandler(readerName)
	cmd := "LOAD DATA LOCAL INFILE 'Reader::" + readerName + "' " +
		"IGNORE INTO TABLE chr "
	_, err := db.Exec(cmd)
	if err != nil {
		panic(err.Error())
	}
}
