package main

import "fmt"

type LanguageType uint8

const (
	STRICTLY_TYPE = iota + 1
	LOOSELY_TYPE
)

type Language struct {
	name   string
	typing int
}

func main() {
	fmt.Println("vim-go")
	lang := Language{"Go", STRICTLY_TYPE}
	fmt.Printf("%#v\n", lang)
}
