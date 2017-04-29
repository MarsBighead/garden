package main

import (
	"fmt"
)

type person struct {
	age int
}

func main() {
	/*
		var n int
		var persons []person
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			n, _ = strconv.Atoi(scanner.Text())
			if n >= 1 && n <= 4 {
				break
			}
		}
		fmt.Printf("Numbers is %d\n", n)
		for scanner.Scan() {
			age, _ := strconv.Atoi(scanner.Text())
			var p person
			p = p.NewPerson(age)

			n--
			if n <= 0 {
				break
			}
		}*/
	var T, age int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
	if initialAge < 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		p.age = 0
	}
	return p
}

func (p person) amIOld() {
	//Do some computation in here and print out the correct statement to the console
	if p.age < 13 {
		fmt.Println("You are young.")
	} else if p.age >= 18 {
		fmt.Println("You are old.")
	} else if p.age >= 13 && p.age < 18 {
		fmt.Println("You are teenager.")
	}
}

func (p person) yearPasses() person {
	//Increment the age of the person in here
	p.age++
	return p
}
