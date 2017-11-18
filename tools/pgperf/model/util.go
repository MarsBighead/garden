package model

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

func toCSV(filename string, data interface{}) {
	fh, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()
	gocsv.MarshalFile(data, fh)
}
