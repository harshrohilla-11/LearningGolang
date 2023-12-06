package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcopme to If Else in Golang")

	// fmt.Println("Enter your login Count :")
	var logincounter int64
	// fmt.Scanln(&logincounter)

	fmt.Printf("Enter your value :")
	fmt.Scanf("%d", &logincounter)

	if logincounter < 10 {

		fmt.Println("NOt Regular User")

	} else if logincounter >= 20 {
		fmt.Println("Regular User")

	} else {
		fmt.Println("Modrate usecase")

	}
	// can initialize a variable (her num) in if itself
	if num := 20; num < 11 {
		fmt.Println("Smaller than 11")

	} else {
		fmt.Println("Greater than 11")
	}

}
