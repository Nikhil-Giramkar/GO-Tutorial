package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web requests in Go")
	response, err := http.Get(url)

	checkNilError(err)

	fmt.Printf("Response type: %T \n", response)

	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)

	checkNilError(err)

	fmt.Println("Response obtained:")
	fmt.Println(string(dataBytes))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
