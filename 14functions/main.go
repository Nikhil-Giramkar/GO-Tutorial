package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions in Go")

	greeter()

	fmt.Println("Addition of 3 and 4:", adder(3, 4))

	fmt.Println("Sum of unknown number of params:", superAdder(1, 3, 4, 5, 10))

	years, msg := licenseGiver(5)
	fmt.Println(msg, years)
}

func greeter() {
	fmt.Println("Namastey from Go Lang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// unknown number of parameters
func superAdder(values ...int) int {

	total := 0
	for _, val := range values {

		total += val
	}
	return total
}

//return multiple things

func licenseGiver(age int) (int, string) {
	if age >= 18 {
		return 0, "You are eligible for license"
	}

	return 18 - age, "You need these many years to be eligible"
}
