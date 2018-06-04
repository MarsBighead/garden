package main

import (
	"fmt"
	"github.com/couchbase/gocb"
)

//https://developer.couchbase.com/documentation/server/5.1/sdk/go/sdk-xattr-example.html
func main() {
	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "password",
	})

	bucket, _ := cluster.OpenBucket("travel-sample", "")

	// Add key-value pairs to hotel_10138, representing traveller-Ids and associated discount percentages
	bucket.MutateIn("hotel_10138", 0, 0).
		UpsertEx("discounts.jsmith123", "20", gocb.SubdocFlagXattr|gocb.SubdocFlagCreatePath).
		UpsertEx("discounts.pjones356", "30", gocb.SubdocFlagXattr|gocb.SubdocFlagCreatePath).
		// The following lines, "insert" and "remove", simply demonstrate insertion and
		// removal of the same path and value
		InsertEx("discounts.jbrown789", "25", gocb.SubdocFlagXattr|gocb.SubdocFlagCreatePath).
		RemoveEx("discounts.jbrown789", gocb.SubdocFlagXattr).
		Execute()

	// Add key - value pairs to hotel_10142, again representing traveller - Ids and associated discount percentages
	bucket.MutateIn("hotel_10142", 0, 0).
		UpsertEx("discounts.jsmith123", "15", gocb.SubdocFlagXattr|gocb.SubdocFlagCreatePath).
		UpsertEx("discounts.pjones356", "10", gocb.SubdocFlagXattr|gocb.SubdocFlagCreatePath).
		Execute()

	q := gocb.NewN1qlQuery("SELECT id, META().id AS docID FROM `travel-sample`")
	rows, _ := bucket.ExecuteN1qlQuery(q, nil)

	var row struct {
		DocID string `json:"docID"`
	}
	for rows.Next(&row) {
		res, err := bucket.LookupIn(row.DocID).GetEx("discounts.jsmith123", gocb.SubdocFlagXattr).Execute()
		if err == nil {
			var discount string
			res.ContentByIndex(0, &discount)
			fmt.Printf("%s - %s\n", discount, row.DocID)
		}
	}
}
