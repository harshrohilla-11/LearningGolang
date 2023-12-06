package main

import "fmt"

func main() {

	var fruitlist [5]string

	fruitlist[0] = "Mango"
	fruitlist[1] = "Apple"
	fruitlist[3] = "Melons"

	fmt.Println(fruitlist)
	//len func when provided value does not count
	fmt.Println("the fruits are :", len(fruitlist))
	fmt.Printf("type is: %T\n", fruitlist)

	// this is apparaently a slice
	var vegielist = []string{"A", "B", "C", "akgs"}

	fmt.Println(vegielist)
	fmt.Println("the vegies are :", len(vegielist))
	fmt.Printf("type is: %T\n", vegielist)

	var newlist = []int64{1, 2, 3, 4, 5}
	fmt.Println(newlist)
	fmt.Printf("type----  %T\n", newlist)

}
