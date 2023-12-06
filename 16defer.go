package main

import (
	"fmt"
)

func main() {
	fmt.Println("Defer in Go")

	fmt.Println("Hello")
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two") // Defer used LIFO , Last in first Out
	Mydefer()

	// World , One , Two
	//Hello ; 0,1,2,3,4 ; // World , One , Two
}
func Mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("\n", i)

	}

}
