package main

import "fmt"

func main() {
	fmt.Println("Struct....")

	//No inheritance, super, parent concepts in golang

	nikhil := User{"Nikhil Giramkar", "nikhil@gmail.com", true, 22}

	fmt.Println(nikhil)
	fmt.Printf("Detailed information: %+v", nikhil)
	fmt.Printf("\nName: %v \nEmail: %v \nIs Verified: %v \nAge: %v", nikhil.Name, nikhil.Email, nikhil.Status, nikhil.Age)

}

//Struct name  and its properties in capital as they will be exported out and used further
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
