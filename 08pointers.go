package main

import "fmt"

func main() {

	fmt.Println("Welcome to Pointers")

	//var ptr *int

	//fmt.Println("value of pointer here :", ptr)

	mynumber := 100

	var ptr = &mynumber

	fmt.Println("here is value of pointer :", ptr)
	fmt.Println("here is the value in pointer :", *ptr)

	*ptr = *ptr * 2
	fmt.Println("the new value is :", mynumber)

}
