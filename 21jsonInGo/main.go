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

type student struct {
	Name     string   `json:"studName"`
	RollNo   int      `json:"studRoll"`
	Address  string   `json:"studAddress"`
	Subjects []string `json:"subjects"`
}

func main() {
	fmt.Println("Json In GoLang")
	EncodeJson()

	DecodeJson()
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

func DecodeJson() {
	receivedJson := []byte(`
	{
		"studName": "Nikhil G",
		"studRoll": 17,
		"studAddress": "TX, USA",
		"subjects": ["Maths","Science","History"]
	}`)

	var myStudent student

	isJsonValid := json.Valid(receivedJson)

	if isJsonValid {
		fmt.Println("We received a vaild json")

		json.Unmarshal(receivedJson, &myStudent)

		fmt.Printf("%#v \n", myStudent)
	} else {
		fmt.Println("JSON not Valid")
	}

	//In some cases, we do not want to fit Json in any struct format
	//just extract key value pairs from json and perform further operations
	//Then use map, to extract key,value

	var myRandomJson map[string]interface{}
	//The above syntax denotes that myRandomJson is a mpa with keys of type string, but values can be anything hence interface{}

	json.Unmarshal(receivedJson, &myRandomJson)
	fmt.Printf("%#v \n", myStudent)

	for k, v := range myRandomJson {
		fmt.Printf("%v --> %v , Type of val: %T \n", k, v, v)
	}

}
