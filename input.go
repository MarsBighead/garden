package garden

import (
	"net/http"
)

// Input items for garden server
type Input interface {
	Category() string
	Description() string
	AddEnv(map[string]string)
	http.Handler
}
