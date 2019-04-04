package opts

import (
	"flag"
	"fmt"
)

func init() {
	flag.Parse()
}

var lang = flag.String("lang", "golang", "the lang of the program")
var age = flag.Int64("age", 18, "the age of the user")
var safemod = flag.Bool("safemod", true, "whether using safemode")

func getOpts() {
	//flag.Parse()
	fmt.Println("lang is: ", *lang)
	fmt.Println("age is: ", *age)
	fmt.Println("safemod is: ", *safemod)
}
