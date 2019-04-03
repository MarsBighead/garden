package page

import (
	"garden"

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

// Person General people to visit
type Person struct {
	Name    string
	Age     int
	Emails  []string
	Company string
	Role    string
}

// OnlineUser User struct for online party
type OnlineUser struct {
	User      []*Person
	LoginTime string
}
