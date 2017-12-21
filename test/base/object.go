package base

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
	if initialAge < 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		initialAge = 0
	}
	p.age = initialAge
	return p
}

func (p person) amIOld() {
	//Do some computation in here and print out the correct statement to the console
	if p.age < 13 {
		fmt.Println("You are young.")
	} else if p.age >= 18 {
		fmt.Println("You are old.")
	} else if p.age >= 13 && p.age < 18 {
		fmt.Println("You are a teenager.")
	}
}

func (p person) yearPasses() person {
	//Increment the age of the person in here
	p.age++
	return p
}

//GoObject include Class and instance
func GoObject() {
	enhance := 3
	ages := []int{2, 17, 10, 19}

	for _, age := range ages {
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < enhance; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}
