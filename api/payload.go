package api

import (
	"io/ioutil"
	"log"
	"net/http"
)

// RESTfulJSON For JSON API handle
func (s *Service) RESTfulJSON(w http.ResponseWriter, r *http.Request) {
	filename := s.Environment["DATA"] + "/ad-mock.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(body)
}
