package bind

import (
	"testing"
	"time"
)

type T struct {
	Age      int
	Name     string
	Children []int
	Time     *time.Time
}

type To struct {
	Age      int
	Name     string
	Children []int
	Time     time.Time
}

func TestBind(t *testing.T) {
	t0 := new(T)
	Bind(t0)
	t1 := new(To)
	Bind(t1)
}
