package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var fruits [4]string
	fruits[0] = "apple"
	fruits[2] = "banana"
	fruits[3] = "papaya"
	fmt.Println("Fruits:", fruits)
	fmt.Println("Fruits len", len(fruits))

	var veggies = [5]string{"beans", "cabbage", "tomato"}
	fmt.Println("Veggies: ", veggies)

}
