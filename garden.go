package main

import (
	"garden/config"
	"garden/models"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := config.GetDBConfig()
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	models.TruncateTable("chr", db)
	models.InsertVal(db)
	models.GetRows(db)
	os.Exit(1)
}
