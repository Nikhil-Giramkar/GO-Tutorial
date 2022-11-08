package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Web verbs")

	PerformGet()
}

func PerformGet() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Response: ", response.Status)
	fmt.Println("Content Length: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Content:", string(content))

}
