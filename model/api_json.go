package model

import (
	"io/ioutil"
	"log"
	"net/http"
)

// ProtocolJSON For JSON API handle
func ProtocolJSON(w http.ResponseWriter, r *http.Request) {
	jsonFile := GetCurrentDir() + "/data/mock.json"
	body, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(body)
}
