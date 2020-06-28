package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	day := flag.String("d", "2019-10-21", "Start time with short date")
	flag.Parse()
	fmt.Printf("Transformat Day=%s to Unix time:\n\n", *day)
	t, err := time.Parse(time.RFC3339, *day+"T00:00:00Z")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.Unix())
}
