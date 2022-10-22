package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name: ")

	//comma ok or error syntax

	input, _ := reader.ReadString('\n')

	//Go does not have try catch so, if all ok, value stored in "input"
	//If error occurs, it is ignored due to _
	// if we wish to store error msg, we can use
	//input, err := reader.ReadString('\n')

	fmt.Printf("Hello Mr. %s, Happy to see you!", input)

}
