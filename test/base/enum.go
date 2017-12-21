package base

import "fmt"

//LanguageType Language Type
type LanguageType uint8

const (
	//STRICTLYTYPE STRICTLY TYPE
	STRICTLYTYPE = iota + 1
	// LOOSELYTYPE LOOSELY TYPE
	LOOSELYTYPE
)

type language struct {
	name   string
	typing int
}

//StudyEnum enum type method in golang
func StudyEnum() {
	lang := language{"Go", STRICTLYTYPE}
	fmt.Printf("%#v\n", lang)
}
