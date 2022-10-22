package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Methods in Go")

	nikhil := Person{"Nikhil Giramkar", 22, "nikhil@gmail.com", true}

	fmt.Println(nikhil)
	fmt.Println()

	nikhil.GetName()
	fmt.Println()

	nikhil.NewEmail()
	fmt.Println()

	//Even when you set a new email to the user using a method, the original value is not set
	//this is because we passed only the copy of struct
	//to manipulate the actual values, we need to pass by reference, using pointers

	fmt.Println(nikhil)
	fmt.Println()

	nikhil.NewEmailPermanent()
	fmt.Println()

	fmt.Println(nikhil)
	fmt.Println()

}

type Person struct {
	Name     string
	Age      int
	Email    string
	IsActive bool
}

// methods are functions which belong to a struct
func (p Person) GetName() {
	fmt.Println("Name of person is ", p.Name)
}

// pass by value
func (p Person) NewEmail() {
	emailParts := strings.Split(p.Email, "@")
	firstPart := emailParts[0]
	newEmail := firstPart + "@dev.go"
	p.Email = newEmail
	fmt.Println("Your new email id:", p.Email)
}

// pass by refernce
func (p *Person) NewEmailPermanent() {
	emailParts := strings.Split(p.Email, "@")
	firstPart := emailParts[0]
	newEmail := firstPart + "@yahoo.com"
	p.Email = newEmail
	fmt.Println("Your new email id:", p.Email)
}
