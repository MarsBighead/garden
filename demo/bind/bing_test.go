package bind

import (
	"fmt"
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
	t2 := To{}
	Bind(&t2)
	fmt.Println(t2)
}
func TestSliceBind(t *testing.T) {
	var ss []string
	SliceBind(&ss)
	fmt.Println(ss)
}
