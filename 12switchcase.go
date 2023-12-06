package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// fmt.Println("Switch Case in  Go")

	rand.Seed(time.Now().UnixNano())

	const consecutiveSixthreshold = 3
	var consecutiveSixCount int
	var isfirstone = true

	dicenumber := rand.Intn(6) + 1
	fmt.Println("Dice value :\n", dicenumber)

	// use switch to handle outcomes
	switch dicenumber {
	case 1:
		if isfirstone {
			fmt.Println("Value is 1 , You can move now ")
			isfirstone = false //to indicate that the special message for the first roll of 1 has been used
		} else {
			fmt.Println("Move 1 pace forward")
		}
	case 2:
		fmt.Println("Value is 2 - move 2 paces ")
	case 3:
		fmt.Println("Value is 3 - move 3 paces ")
	case 4:
		fmt.Println("Value is 4 - move 4 paces ")
	case 5:
		fmt.Println("Value is 5 - move 5 paces ")
	case 6:
		fmt.Println("Lucky Roll ! Value is 6 - Move 6 paces and You can roll Again ")
		consecutiveSixCount++

		if consecutiveSixCount == consecutiveSixthreshold {
			fmt.Println("Voided - No more Turns")
		}
	default:
		fmt.Println("What was that !")
		consecutiveSixCount = 0 // Reset the count
	}

}
