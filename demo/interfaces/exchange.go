package intest

import (
	"fmt"
	"io"
)

// Exchanger interface test for func Exchange
type Exchanger interface {
	Exchange()
}

// StringPair struct type  for func  Exchange string
type StringPair struct {
	first, second string
}

// Exchange exchange the pair string values
func (p *StringPair) Exchange() {
	p.first, p.second = p.second, p.first
}

// Point struct type  for func  Exchange int
type Point [2]int

// Exchange exchange the pair int values
func (p *Point) Exchange() {
	p[0], p[1] = p[1], p[0]
}

// ExchangeThese func for invoking interface Exchanger
func ExchangeThese() {
	paul := StringPair{"Paul", "Duan"}
	header := StringPair{"Bigheader", "Mars"}
	point := Point{5, -3}
	fmt.Println("Before:", paul, header, point)
	paul.Exchange()
	header.Exchange()
	//point.Exchange()
	(&point).Exchange()
	fmt.Println("After#1:", paul, header, point)
	exchangeThese(&paul, &header, &point)
	fmt.Println("After#2:", paul, header, point)
}

func exchangeThese(exchangers ...Exchanger) {
	for _, exchanger := range exchangers {
		exchanger.Exchange()
	}
}

// Read Realize func Read defined by interface io.Reader
func (p *StringPair) Read(data []byte) (n int, err error) {
	//fmt.Printf("StringPair for func Read is %v.\n", p)
	if p.first == "" && p.second == "" {
		return 0, io.EOF
	}
	if p.first != "" {
		n = copy(data, p.first)
		p.first = p.first[n:]
	}
	if n < len(data) && p.second != "" {
		m := copy(data[n:], p.second)
		p.second = p.second[m:]
		n += m
	}
	return n, nil
}

// ToBytes Transfer data defined by io.Reader interface to type byte
func ToBytes(reader io.Reader, size int) ([]byte, error) {
	data := make([]byte, size)
	n, err := reader.Read(data)
	if err != nil {
		return data, err
	}
	fmt.Printf("Actual return data bytes length: %d\n", n)
	return data[:n], nil
}

func DataBlock() {
	data := &S{"My name is Paul Duan"}
	l := len(data.s)
	fmt.Println("Lenghth ", l)
	// Declare a io.Reader include only one object with io.Reader(obj)
	reader := io.Reader(data)
	raw, err := ToBytes(reader, l)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Data Block raw: %q\n", raw)
}

type S struct {
	s string
}

func (s S) Read(data []byte) (n int, err error) {
	if s.s == "" {
		return 0, io.EOF
	}
	n = copy(data, s.s)
	return n, nil
}
