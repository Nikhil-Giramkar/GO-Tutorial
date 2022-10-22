package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case....")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Printf("Value of dice: %d\n", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1")
	case 2:
		fmt.Println("Move by 2 spots")
	case 3:
		fmt.Println("Move by 3 spots")
	case 4:
		fmt.Println("Move by 4 spots")
		fallthrough
	case 5:
		fmt.Println("Move by 5 spots")
	case 6:
		fmt.Println("Move by 6 spots")
	default:
		fmt.Println("Invalid number")
	}

	//No break statement
	//fallthrough statement used to execute current and next case
}
