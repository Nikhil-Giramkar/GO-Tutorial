package main

import "fmt"

func main() {
	fmt.Println("Defer in Go")
	//when we mark anything defer
	//it is executed just before end of the function

	//works in LIFO form
	//the line marked as defer last, will be displayed first

	defer fmt.Println("Nikhil")
	defer fmt.Println("Hello")
	defer fmt.Println("World")
	fmt.Println("Good")
	fmt.Println("Morning")

	myDefer()

	/*
		Output:

		Defer in Go
		Good
		Morning
		World
		Hello
		Nikhil
	*/

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
