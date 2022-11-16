package main

import (
	"fmt"
	"net/http"
	"sync"
)

// waitGroup is advance version of time.Sleep
var wg sync.WaitGroup

func main() {

	//We create a go routine by adding keyword "go" in front of func call

	//greeter("nikhil")

	// go greeter("nikhil")
	// greeter("giramkar")
	//Go routine gives concurrency to code

	websiteList := []string{
		"https://fb.com",
		"https://google.com",
		"https://yahoo.com",
		"https://instagram.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		//go keywork makes it concurrent as "goRoutine"
		wg.Add(1)
		//You keep adding the tasks which might delay and for which main method should wait
	}

	wg.Wait()
	//this tells main method to wait until it gets Done() signal from all tasks present in waitGroup
}

// func greeter(msg string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println(msg)
// 	}
// }

func getStatusCode(endpoint string) {
	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error in Getting Status")
	} else {
		fmt.Printf("%d Status for %s\n", result.StatusCode, endpoint)
	}

	defer wg.Done()
	//signal the waitGroup that status code is received and we can exit from wait Group
}
