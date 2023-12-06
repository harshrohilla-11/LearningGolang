package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to functions in Goland")
	Namaste() // functon written place dont matter the call matters for execution

	bread() // calling a function

	justadding := adder(3, 5)
	fmt.Println(justadding)

	proresult := proadder()
	fmt.Println(proresult)

}
func Namaste() {
	fmt.Println("Namsate from Golang")

}
func bread() {

	fmt.Println("lets have breakfast")

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo

}

func proadder(values ...int) int {
	total := 0

	for _, val := range values {
		fmt.Scanf("%d", &val)
		total += val
	}
	return total
}
