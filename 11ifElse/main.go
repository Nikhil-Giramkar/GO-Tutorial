package main

import "fmt"

func main() {
	fmt.Println("If else....")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Less than 10"
	} else if loginCount > 10 {
		result = "More than 10"
	} else {
		result = "Exactly 10"
	}

	fmt.Println(result)

	//We can also assign values on the go in if block
	//Generally when fetch some values from API

	if myAge := 22; myAge > 25 {
		fmt.Println("Yacha lagna karun taka")
	} else {
		fmt.Println("Kamvaychi akkal yeu de")
	}
}
