package inembed

import (
	"fmt"
	"unicode"
)

type LowerCaser interface {
	LowerCase()
}

type UpperCaser interface {
	UpperCase()
}

type LowerUpperCaser interface {
	//LowerCase()
	LowerCaser
	//UpperCase()
	UpperCaser
}

type FixCaser interface {
	FixCase()
}

type Part struct {
	Num  int
	Name string
}

func (part *Part) FixCase() {
	part.Name = fixCase(part.Name)
}

//
type StringPair struct {
	First, Second string
}

func (p *StringPair) FixCase() {
	p.First = fixCase(p.First)
	p.Second = fixCase(p.Second)
}

func fixCase(s string) string {
	var chars []rune
	upper := true
	for _, char := range s {
		if upper {
			char = unicode.ToUpper(char)
		} else {
			char = unicode.ToLower(char)
		}
		chars = append(chars, char)
		upper = unicode.IsSpace(char) || unicode.Is(unicode.Hyphen, char)
	}
	return string(chars)
}

func TR() {
	toastRack := Part{8427, " TOAST RACK"}
	toastRack.FixCase()
	lobelia := StringPair{" LOBELIA", "SACKVILLE-BAGGINS"}
	(&lobelia).FixCase()
	fmt.Println(lobelia, toastRack)
}
