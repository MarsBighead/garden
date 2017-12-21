package base

import "fmt"

//StudyComplex study data complex
func StudyComplex() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Printf("%v\n", x)
	fmt.Println(x * y)       // "(-5+10i)"
	fmt.Println(real(x * y)) // "-5"
	fmt.Println(imag(x * y)) // "10"
}
