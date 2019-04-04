package main

import (
	"fmt"
	"testing"
)

type Person struct {
	Age      int
	Name     string
	Children []int
}

func TestBind(t *testing.T) {
	p := new(Person)
	bind(p)
	fmt.Println("test struct is ", p)

}
