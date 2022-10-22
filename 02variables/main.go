package main

import "fmt"

const LoginToken string = "wqrwegaasdv"

//any variable with first letter capital becomes Public

func main() {
	fmt.Println("Vraiables in GO")

	var username string = "Nikhil Giramkar"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isHuman bool = true
	fmt.Println(isHuman)
	fmt.Printf("Variable is of type: %T \n", isHuman)

	var smallint uint8 = 255
	fmt.Println(smallint)
	fmt.Printf("Variable is of type: %T \n", smallint)

	var smallFloat float32 = 255.324252235134124152
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//Default values

	var defaultInt int
	fmt.Println(defaultInt)

	//Implicit style
	var anotherVariable = "This is a string"
	fmt.Println(anotherVariable)
	//anotherVariable = 5  - throws error

	//No var style
	numberOfPeople := 22334455
	fmt.Println(numberOfPeople)
	// we can use a variable with := and without var only within any method, not for global variables

	fmt.Println("Global Public constant variable value: ", LoginToken)
}
