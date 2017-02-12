package models

import (
	"fmt"
)

// S Test data struct
type S struct {
	num int
}

func (p *S) get() int {
	fmt.Println("Get value")
	return p.num
}
func (p *S) put(val int) {
	fmt.Println("Put value")
	p.num = val
}

// DataOperator interface to operate data
type DataOperator interface {
	get() int
	put(int)
}

func input(p DataOperator) {
	fmt.Println(p.get())
	p.put(1)
	fmt.Println(p.get())
}

// Testinterface Test interface usage in golang
func Testinterface() {
	var s S
	input(&s)
}
