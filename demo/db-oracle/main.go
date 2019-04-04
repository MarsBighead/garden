package main

import (
	"database/sql"
	"fmt"

	_ "gopkg.in/goracle.v2"
)

func main() {
	db, err := sql.Open("oci8", "usgmtr/vmware@10.158.13.105/XE")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

}
