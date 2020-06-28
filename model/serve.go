package model

import (
	"garden"
	"log"
	"regexp"

	"github.com/jmoiron/sqlx"
)

var reSuffix = regexp.MustCompile(`\.(htm(l)?|php|jsp)$`)

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

// MiddleHandler handle the original input information
func (s *Service) MiddleHandler(inputs map[string]garden.Creator, name string, patterns ...string) {
	creator, ok := inputs[name]
	if !ok {
		panic("Undefined but requested input:" + name)
	}
	input := creator()
	input.AddEnv(s.Environment)

	for _, pattern := range patterns {
		log.Printf("Direct url: %s", pattern)
		if reSuffix.MatchString(pattern) {
			log.Printf("URL: %s", pattern)
			//http.Handle(pattern, input.ServeHTTP)
		} else if regexp.MustCompile(name + "$").MatchString(pattern) {
			log.Printf("Original url: %s", pattern)
		}

	}

}
