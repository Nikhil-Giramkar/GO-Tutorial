package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Sunday"}

	//Method 1
	for i := 0; i < len(days); i++ {
		fmt.Println("Hello, this is ", days[i])
	}

	//Method 2
	for day := range days {
		fmt.Println(days[day])
	}

	//Retrieve index and value
	for idx, day := range days {
		fmt.Printf("Index:%v, Day:%v\n", idx, day)
	}

	//Using for as while loop
	counter := 1

	for counter < 5 {
		fmt.Println("Count:", counter)
		counter++
	}

	counter = 1

	//Break loop
	for counter < 10 {
		if counter == 3 {
			break
		}
		fmt.Println("Count:", counter)
		counter++
	}

	counter = 1

	//Contiinue
	for counter < 7 {
		if counter == 3 {
			counter++
			continue
		}
		if counter == 5 {
			goto nik
		}
		fmt.Println("Count:", counter)
		counter++
	}

nik:
	fmt.Println("Label nikhil")

}
