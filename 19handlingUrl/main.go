package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://lco.dev:3000/learn?course=reactJs&paymentId=ergfbgb"

func main() {
	fmt.Println("Handling URL")

	result, _ := url.Parse(myURL)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	var queryParams = result.Query()
	fmt.Println(queryParams["paymentId"])

	for _, val := range queryParams {
		fmt.Println(val)
	}

	//Creating a Url

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	constructedUrl := partsOfUrl.String()

	fmt.Println(constructedUrl)

}
