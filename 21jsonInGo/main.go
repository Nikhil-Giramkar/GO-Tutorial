package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	//When we wish to change the name of keys in json we use `` and give required name
	Name     string `json:"courseName"`
	Price    int    `json:"amount"`
	Platform string
	Password string   `json:"-"`              //"-" tells Json not to include this field
	Tags     []string `json:"tags,omitempty"` //omitempty specifies that if any of these tags is null then don't encode it in json
}

func main() {
	fmt.Println("Json In GoLang")
	EncodeJson()
}

func EncodeJson() {

	listOfCourses := []course{
		{"MERN Stack Dev", 299, "AWS", "abc1234", []string{"web-dev", "js"}},
		{"MEAN Stack Dev", 399, "GCP", "zxc456", []string{"web-dev", "angular", "js"}},
		{"DSA Basic", 900, "Azure", "987poi", []string{"dsa", "java"}},
		{"DSA Advance", 5612, "Uptyx", "741plm", nil},
	}

	//finalJson, err := json.Marshal(listOfCourses)
	//doesn't look pretty when printed, so use MarshalIndent

	finalJson, err := json.MarshalIndent(listOfCourses, "", "\t")
	//Syntax to make json pretty

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJson)

}
