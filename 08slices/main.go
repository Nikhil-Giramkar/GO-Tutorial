package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices Tutorial")

	//Arrays need to mention size, here don't need to mentionn size
	var fruitList = []string{"Apple", "Orange"}
	fmt.Printf("Type of fruitList, %T  \n", fruitList)

	fruitList = append(fruitList, "Banana", "Papaya")
	fmt.Println("Fruits: ", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("Fruits", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println("Fruits", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 111
	highScores[1] = 999
	highScores[2] = 333
	highScores[3] = 777
	//highScores[4] = 555

	highScores = append(highScores, 444, 555, 666)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	var courses = []string{"java", "perl", "goLang", "Python", "C++", "Aws", "Azure"}
	fmt.Println(courses)

	//Remove from secific index
	fmt.Println("Enter index to delete")
	//reader := bufio.NewReader(os.Stdin)

	var index int
	//fmt.Scanf("%d", index)

	index = 3
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
