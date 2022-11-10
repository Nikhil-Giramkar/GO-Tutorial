package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web verbs")

	PerformGet()

	PerformPostJsonRequest()

	PerformPostFormRequest()
}

func PerformGet() {

	fmt.Println("--------------GET REQUEST----------------------")
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

func PerformPostJsonRequest() {
	fmt.Println("--------------POST REQUEST----------------------")

	myUrl := "http://localhost:8000/post"

	//fake data
	requestBody := strings.NewReader(`
		{
			"name":"Tony Stark",
			"heroName": "Iron Man",
			"Language":"Java"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status: ", response.Status)
	fmt.Println("Content Length: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Content: ", string(content))

}

func PerformPostFormRequest() {
	fmt.Println("--------------POST FORM REQUEST----------------------")

	myUrl := "http://localhost:8000/postform"

	//form data

	formData := url.Values{}

	formData.Add("company", "microsoft")
	formData.Add("employee", "nikhil")
	formData.Add("address", "Ghar")

	response, err := http.PostForm(myUrl, formData)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status: ", response.Status)
	fmt.Println("Content Length:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Content: ", string(content))

}
