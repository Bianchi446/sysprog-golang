package main

import (
	"fmt"
)

func condition(a int) int { // Accept a as a parameter
	if r := a % 10; r != 0 { // if with short declaration
		if r > 5 {
			// if without declaration
			a -= r
		} else if r < 5 { // else if statement
			a += 10 - r
		}
	} else {
		// else statement
		a /= 10
	}
	return a
}

func main() {
	a := 47 // Example input
	result := condition(a)
	fmt.Println("Result:", result)

	tier := 1
	age := 20
	customSwitch(tier, age)

	b := 20
	looping(b)

}
