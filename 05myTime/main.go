package main

import (
	"fmt"
	"time"
)

func main() {
	presentDate := time.Now()

	fmt.Println(presentDate)

	fmt.Println(presentDate.Format("02-01-2006 Monday"))

	createdDate := time.Date(2000, time.December, 29, 11, 45, 02, 2, time.UTC)

	fmt.Println("Created Date: ", createdDate)
	fmt.Println(createdDate.Format("02-Jan-2006 Monday"))
}
