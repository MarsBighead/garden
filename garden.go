package main

import (
	"fmt"
	"garden/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := config.GetDBConfig()
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	stmtIns, err := db.Prepare("INSERT INTO chr  VALUES( ?, ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT * FROM chr ")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()
	// Insert square numbers for 0-24 in the database
	for i := 6; i < 1025; i++ {
		_, err = stmtIns.Exec(i, (i * i), "xxx") // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	fmt.Printf("Test insert and select successfully!\n")
}
