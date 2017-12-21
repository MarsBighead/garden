package base

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//UUIDToHex UUID To Hex
func UUIDToHex() {
	id := "117c786be2fc1ce33a8b456c"
	//   "9223372036854775807"
	src, err := hex.DecodeString(id)
	if err != nil {
		log.Fatal(err)
	}
	var s string
	for _, b := range src {
		ss := strconv.Itoa(int(b))
		if len(ss) <= 2 {
			ss = "0" + ss
		}
		s += ss
	}
	fmt.Printf("s result %v, length of s: %v\n", s, len(s))
	fmt.Printf("Origin hex %v\n", hex.EncodeToString(src))
	rs := strings.Split(s, "")
	var i int
	for {
		if i >= 12 {
			break
		}
		st := rs[i*3] + rs[i*3+1] + rs[i*3+2]
		fmt.Printf("Origin hex %v, TrimLeft  %v\n", st, strings.TrimLeft(st, "0"))
		x, _ := strconv.Atoi(strings.TrimLeft(st, "0"))
		fmt.Printf("number %x\n", x)
		i++
	}
}
