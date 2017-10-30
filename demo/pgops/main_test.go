package main

import (
	"fmt"
	"testing"
)

type T struct {
	Age      int
	Name     string
	Children []int
}

func TestBind(t *testing.T) {
	s := new(T)
	Bind(s)
	fmt.Println("test struct is ", s)

}
