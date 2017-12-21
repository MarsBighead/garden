package base

import (
	"bytes"
	"fmt"
	"io"
)

//StudyBits Study bits method
func StudyBits() {

	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123))
}

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte('}')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// S test string struct
type S struct {
	s string
}

func readerToBytes(reader io.Reader, size int) ([]byte, error) {
	data := make([]byte, size)
	n, err := reader.Read(data)
	if err != nil {
		return data, err
	}
	fmt.Printf("Actual return data bytes length: %d\n", n)
	return data[:n], nil
}

// ReadBytes Transfer data defined by io.Reader interface to type byte
func ReadBytes() {
	data := &S{"My name is Paul Duan"}
	l := len(data.s)
	fmt.Println("Lenghth ", l)
	// Declare a io.Reader include only one object with io.Reader(obj)
	reader := io.Reader(data)
	raw, err := readerToBytes(reader, l)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Data Block raw: %q\n", raw)
}
