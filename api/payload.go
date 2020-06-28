package api

import (
	"garden"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	Router      *garden.Router
	DB          *sqlx.DB
	Environment map[string]string
}

// NewService initialize page service
func NewService(r *garden.Router, db *sqlx.DB) *Service {
	return &Service{
		Router:      r,
		Environment: r.Environment,
		DB:          db,
	}
}

// RESTfulJSON For JSON API handle
func (s *Service) RESTfulJSON(w http.ResponseWriter, r *http.Request) {
	filename := s.Environment["DATA"] + "/ad-mock.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(body)
}
