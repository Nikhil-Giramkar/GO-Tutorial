package main

import "fmt"

func main() {
	fmt.Println("Maps...")

	languages := make(map[string]string)

	languages["js"] = "JavaScript"
	languages["cpp"] = "C++"
	languages["jv"] = "Java"
	languages["py"] = "Python"

	fmt.Println(languages)

	fmt.Println("JS is ", languages["js"])

	delete(languages, "cpp")

	fmt.Println(languages)

	//looping

	for key, value := range languages {
		fmt.Printf("Key: %v, Value: %v \n", key, value)
	}

	//print only values

	for _, value := range languages {
		fmt.Printf("Value: %v \n", value)
	}

	//Similary you can print only keys as well key,_

}
