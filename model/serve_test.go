package model

import (
	"testing"
)

func TestMiddleHandler(t *testing.T) {
	patterns := []string{
		"/home",
		"/home.htm",
		"/home.html",
	}
	MiddleHandler(nil, "home", patterns...)

}
