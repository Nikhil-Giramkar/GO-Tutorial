package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Go")

	content := "This content will go in file."

	file, err := os.Create("./myFile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("File is written, length of content:", length)
	defer file.Close()

	readFile("./myFile.txt")
}

func readFile(filePath string) {
	contentInBytes, err := ioutil.ReadFile(filePath)

	checkNilError(err)

	fmt.Println("Text in file:", string(contentInBytes))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
