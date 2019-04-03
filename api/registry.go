package api

import (
	"garden"

	"github.com/jmoiron/sqlx"
)

type Creator func() garden.Input

var Inputs = map[string]Creator{}

func Add(name string, creator Creator) {
	Inputs[name] = creator
}

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
