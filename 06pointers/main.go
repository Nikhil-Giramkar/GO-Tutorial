package main

import "fmt"

func main() {
	fmt.Println("Pointers....")

	var myPtr *int
	fmt.Println("Default Value of myPtr:", myPtr)

	myNum := 23

	var ptr = &myNum
	fmt.Println("Value in pointer:", ptr)
	fmt.Println("Value in pointer", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Value of myNum:", *ptr)
	fmt.Println("Address of myNum:", ptr)

}
