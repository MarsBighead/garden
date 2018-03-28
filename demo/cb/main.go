package main

import (
	"encoding/json"
	"fmt"

	"github.com/couchbase/gocb"
)

type hotel struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}

func main() {
	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "duanp1",
	})
	bucket, _ := cluster.OpenBucket("travel-sample", "")
	query := gocb.NewN1qlQuery("SELECT name,phone FROM `travel-sample` WHERE type=$1 and directions IS NOT MISSING limit 10;")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{"hotel"})
	var hs []*hotel
	row := new(hotel)
	for rows.Next(row) {
		fmt.Printf("Name: %v, Phone: %v\n", row.Name, row.Phone)
		hs = append(hs, row)
		row = new(hotel)
	}
	b, _ := json.MarshalIndent(hs, "", "    ")
	fmt.Println(string(b))

}
