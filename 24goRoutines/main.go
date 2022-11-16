package main

import (
	"fmt"
	"time"
)

func main() {

	//We create a go routine by adding keyword "go" in front of func call

	//greeter("nikhil")

	go greeter("nikhil")

	greeter("giramkar")

	//Go routine gives concurrency to code

}

func greeter(msg string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(msg)
	}
}
